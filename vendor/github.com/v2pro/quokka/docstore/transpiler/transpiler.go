package transpiler

import (
	"github.com/robertkrimen/otto/parser"
	"github.com/robertkrimen/otto/ast"
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/compiler"
	"strconv"
	"fmt"
	"bytes"
	"github.com/robertkrimen/otto/file"
	"github.com/v2pro/plz/countlog"
)

func Compile(input string) (func(doc interface{}, req interface{}) interface{}, error) {
	output, err := translate(input)
	if err != nil {
		return nil, err
	}
	plugin, err := compiler.DynamicCompile(output)
	if err != nil {
		return nil, err
	}
	fn, err := plugin.Lookup("Fn")
	if err != nil {
		return nil, err
	}
	return fn.(func(doc interface{}, req interface{}) interface{}), nil
}

func translate(input string) (string, error) {
	funcLiteral, err := parser.ParseFunction("", input)
	if err != nil {
		return "", err
	}
	output := []byte(`
package main
import "github.com/v2pro/quokka/docstore/runtime"
import "github.com/v2pro/plz/countlog"
import "fmt"

var Object = runtime.ObjectClass
var Date = runtime.DateClass
var Math = runtime.MathClass

func Fn(doc interface{}, req interface{}) (ret interface{}) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			jsErr, _ := recovered.(*runtime.JavascriptError)
			if jsErr != nil {
				ret = recovered
				if _, isThrown := jsErr.Recovered.(runtime.Thrown); isThrown {
					ret = jsErr.Recovered
				}
			} else {
				countlog.Fatal("event!runtime.panic", "err", recovered, "stacktrace", countlog.ProvideStacktrace)
				ret = fmt.Errorf("unknown error: %v", recovered)
			}
		}
	}()
	return handle([]interface{}{doc, req})
}

func log_trace(event string, properties ...interface{}) {
	runtime.Log(countlog.LevelTrace, event, properties...)
}
`)
	translators := map[string]*translator{}
	for _, funcStmtObj := range funcLiteral.Body.(*ast.BlockStatement).List {
		funcStmt := funcStmtObj.(*ast.FunctionStatement)
		funcName := funcStmt.Function.Name.Name
		tl := &translator{
			input: input[funcStmt.Function.Body.Idx0()-14: funcStmt.Function.Body.Idx1()-16],
		}
		paramNames := []string{}
		for _, param := range funcStmt.Function.ParameterList.List {
			paramNames = append(paramNames, param.Name)
		}
		tl.translate(funcName, paramNames, bytes.Count(output, []byte{'\n'}))
		if tl.err != nil {
			return "", tl.err
		}
		output = append(output, tl.output...)
		translators[funcName] = tl
	}
	for funcName, tl := range translators {
		output = append(output, `var `...)
		output = append(output, funcName...)
		output = append(output, `_JS_SOURCE=`...)
		output = append(output, '`')
		output = append(output, tl.input...)
		output = append(output, '`', '\n')
		output = append(output, `var `...)
		output = append(output, funcName...)
		output = append(output, `_GO_SOURCE=`...)
		output = append(output, '`')
		output = append(output, tl.goSrc...)
		output = append(output, '`', '\n')
	}
	return string(output), nil
}

type translator struct {
	input          string
	output         []byte
	jsOffsets      []int
	goOffsets      []int
	goSrc          []byte
	err            error
	bodyStartIdx   file.Idx
}

func (tl *translator) reportError(node ast.Node, errmsg string) {
	if tl.err != nil {
		return
	}
	tl.err = errors.New(errmsg)
}

func (tl *translator) translate(funcName string, paramNames []string, absoluteLineNo int) {
	funcLiteral, err := parser.ParseFunction("", tl.input)
	if err != nil {
		tl.err = err
		return
	}
	tl.bodyStartIdx = funcLiteral.Body.Idx0() + 2
	tl.translateStatement(funcLiteral.Body)
	tl.goSrc = tl.output
	tl.output = []byte(`func `)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, `(args []interface{}) interface{} {`...)
	tl.output = append(tl.output, `
	defer func() {
		recovered := recover()
		if recovered != nil {
			runtime.ReportError("`...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, `", `...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, "_JS_SOURCE, []int{"...)
	for i, offset := range tl.jsOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, "}, "...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, "_GO_SOURCE, []int{"...)
	for i, offset := range tl.goOffsets {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, strconv.Itoa(offset)...)
	}
	tl.output = append(tl.output, `}, `...)
	tl.output = append(tl.output, strconv.Itoa(absoluteLineNo)...)
	tl.output = append(tl.output, `, recovered)
		}
	}()
`...)
	for i, paramName := range paramNames {
		tl.output = append(tl.output, paramName...)
		tl.output = append(tl.output, " := args["...)
		tl.output = append(tl.output, strconv.Itoa(i)...)
		tl.output = append(tl.output, `]
		runtime.BlackHole(`...)
		tl.output = append(tl.output, paramName...)
		tl.output = append(tl.output, ")\n"...)

	}
	tl.output = append(tl.output, tl.goSrc...)
	tl.output = append(tl.output, "\nreturn nil\n}\n"...)
}

func (tl *translator) translateStatement(stmt ast.Statement) {
	switch typedStmt := stmt.(type) {
	case *ast.BlockStatement:
		tl.translateBlockStatement(typedStmt)
	case *ast.ExpressionStatement:
		tl.translateExpression(typedStmt.Expression)
	case *ast.ReturnStatement:
		tl.translateReturnStatement(typedStmt)
	case *ast.ThrowStatement:
		tl.translateThrowStatement(typedStmt)
	case *ast.IfStatement:
		tl.translateIfStatement(typedStmt)
	case *ast.VariableStatement:
		tl.translateVariableStatement(typedStmt)
	default:
		tl.reportError(stmt, "can not handle "+reflect.TypeOf(stmt).String())
	}
}
func (tl *translator) translateReturnStatement(stmt *ast.ReturnStatement) {
	tl.output = append(tl.output, "return "...)
	tl.translateExpression(stmt.Argument)
}

func (tl *translator) translateThrowStatement(stmt *ast.ThrowStatement) {
	tl.output = append(tl.output, "panic(runtime.Thrown{"...)
	tl.translateExpression(stmt.Argument)
	tl.output = append(tl.output, "})"...)
}

func (tl *translator) translateBlockStatement(stmt *ast.BlockStatement) {
	for _, child := range stmt.List {
		jsOffset := int(child.Idx0() - tl.bodyStartIdx)
		if jsOffset < 0 {
			countlog.Error("event!transpiler.failed to calculate js offset",
				"bodyStartIdx", tl.bodyStartIdx,
					"blockIdx0", stmt.Idx0(), "blockIdx1", stmt.Idx1(),
					"idx0", child.Idx0(), "idx1", child.Idx1(), "type", reflect.TypeOf(child).String())
			tl.err = errors.New("jsOffset calc failed")
		}
		tl.jsOffsets = append(tl.jsOffsets, jsOffset)
		tl.goOffsets = append(tl.goOffsets, len(tl.output))
		tl.translateStatement(child)
		tl.output = append(tl.output, '\n')
	}
}

func (tl *translator) translateExpression(expr ast.Expression) {
	switch typedExpr := expr.(type) {
	case *ast.AssignExpression:
		tl.translateAssignExpression(typedExpr)
	case *ast.StringLiteral:
		tl.translateStringLiteral(typedExpr)
	case *ast.Identifier:
		tl.output = append(tl.output, typedExpr.Name...)
	case *ast.BracketExpression:
		tl.translateBracketExpression(typedExpr)
	case *ast.DotExpression:
		tl.translateDotExpression(typedExpr)
	case *ast.ObjectLiteral:
		tl.translateObjectLiteral(typedExpr)
	case *ast.ArrayLiteral:
		tl.translateArrayLiteral(typedExpr)
	case *ast.NumberLiteral:
		tl.output = append(tl.output, fmt.Sprintf("%v", typedExpr.Value)...)
	case *ast.NullLiteral:
		tl.output = append(tl.output, "nil"...)
	case *ast.CallExpression:
		tl.translateCallExpression(typedExpr)
	case *ast.BinaryExpression:
		tl.translateBinaryExpression(typedExpr.Operator.String(), typedExpr.Left, typedExpr.Right)
	case *ast.UnaryExpression:
		tl.translateUnaryExpression(typedExpr)
	case *ast.VariableExpression:
		tl.translateVariableExpression(typedExpr)
	default:
		tl.reportError(typedExpr, "can not handle "+reflect.TypeOf(typedExpr).String())
	}
}

func (tl *translator) translateDotExpression(expr *ast.DotExpression) {
	tl.output = append(tl.output, `runtime.AsObj(`...)
	tl.translateExpression(expr.Left)
	tl.output = append(tl.output, `).Get("`...)
	tl.output = append(tl.output, expr.Identifier.Name...)
	tl.output = append(tl.output, `")`...)
}

func (tl *translator) translateBracketExpression(expr *ast.BracketExpression) {
	tl.output = append(tl.output, `runtime.AsObj(`...)
	tl.translateExpression(expr.Left)
	tl.output = append(tl.output, ").Get("...)
	tl.translateExpression(expr.Member)
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateAssignExpression(expr *ast.AssignExpression) {
	operator := expr.Operator.String()
	switch leftExpr := expr.Left.(type) {
	case *ast.BracketExpression:
		tl.output = append(tl.output, `runtime.AsObj(`...)
		tl.translateExpression(leftExpr.Left)
		tl.output = append(tl.output, ").Set("...)
		tl.translateExpression(leftExpr.Member)
		tl.output = append(tl.output, ", "...)
		if operator == "=" {
			tl.translateExpression(expr.Right)
		} else {
			tl.translateBinaryExpression(operator, leftExpr, expr.Right)
		}
		tl.output = append(tl.output, ')')
	case *ast.DotExpression:
		tl.output = append(tl.output, `runtime.AsObj(`...)
		tl.translateExpression(leftExpr.Left)
		tl.output = append(tl.output, `).Set("`...)
		tl.output = append(tl.output, leftExpr.Identifier.Name...)
		tl.output = append(tl.output, `", `...)
		if operator == "=" {
			tl.translateExpression(expr.Right)
		} else {
			tl.translateBinaryExpression(operator, leftExpr, expr.Right)
		}
		tl.output = append(tl.output, ')')
	case *ast.Identifier:
		tl.output = append(tl.output, leftExpr.Name...)
		tl.output = append(tl.output, '=')
		if operator == "=" {
			tl.translateExpression(expr.Right)
		} else {
			tl.translateBinaryExpression(operator, leftExpr, expr.Right)
		}
	default:
		tl.reportError(expr.Left, "can not handle "+reflect.TypeOf(expr.Left).String())
	}
}

func (tl *translator) translateStringLiteral(expr *ast.StringLiteral) {
	tl.output = append(tl.output, '"')
	tl.output = append(tl.output, expr.Value...)
	tl.output = append(tl.output, '"')
}

func (tl *translator) translateObjectLiteral(expr *ast.ObjectLiteral) {
	tl.output = append(tl.output, `runtime.NewObject(`...)
	for i, prop := range expr.Value {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.output = append(tl.output, '"')
		tl.output = append(tl.output, prop.Key...)
		tl.output = append(tl.output, `", `...)
		tl.translateExpression(prop.Value)
	}
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateArrayLiteral(expr *ast.ArrayLiteral) {
	tl.output = append(tl.output, `runtime.NewList(`...)
	for i, prop := range expr.Value {
		if i > 0 {
			tl.output = append(tl.output, ", "...)
		}
		tl.translateExpression(prop)
	}
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateCallExpression(expr *ast.CallExpression) {
	tl.output = append(tl.output, "runtime.Call("...)
	tl.translateExpression(expr.Callee)
	for _, arg := range expr.ArgumentList {
		tl.output = append(tl.output, ", "...)
		tl.translateExpression(arg)
	}
	tl.output = append(tl.output, ')')
}

var binaryOperators = map[string]string{
	"+":   "Add",
	"-":   "Subtract",
	"*":   "Multiply",
	"/":   "Divide",
	">":   "GT",
	">=":  "GE",
	"<":   "LT",
	"<=":  "LE",
	"==":  "EQ",
	"===": "EQ",
}

func (tl *translator) translateBinaryExpression(operator string, left ast.Expression, right ast.Expression) {
	funcName := binaryOperators[operator]
	if funcName == "" {
		tl.reportError(left, "do not support operator "+operator)
		return
	}
	tl.output = append(tl.output, "runtime."...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, '(')
	tl.translateExpression(left)
	tl.output = append(tl.output, ", "...)
	tl.translateExpression(right)
	tl.output = append(tl.output, ')')
}

var unaryOperators = map[string]string{
	"-": "NegativeOf",
}

func (tl *translator) translateUnaryExpression(expr *ast.UnaryExpression) {
	operator := expr.Operator.String()
	funcName := unaryOperators[operator]
	if funcName == "" {
		tl.reportError(expr, "do not support operator "+operator)
		return
	}
	tl.output = append(tl.output, "runtime."...)
	tl.output = append(tl.output, funcName...)
	tl.output = append(tl.output, '(')
	tl.translateExpression(expr.Operand)
	tl.output = append(tl.output, ')')
}

func (tl *translator) translateIfStatement(stmt *ast.IfStatement) {
	tl.output = append(tl.output, "if (runtime.AsBool("...)
	tl.translateExpression(stmt.Test)
	tl.output = append(tl.output, ")) {\n"...)
	tl.translateStatement(stmt.Consequent)
	tl.output = append(tl.output, "}\n"...)
}

func (tl *translator) translateVariableStatement(stmt *ast.VariableStatement) {
	for _, expr := range stmt.List{
		tl.translateExpression(expr)
		tl.output = append(tl.output, '\n')
	}
}

func (tl *translator) translateVariableExpression(expr *ast.VariableExpression) {
	tl.output = append(tl.output, expr.Name...)
	tl.output = append(tl.output, ":="...)
	tl.translateExpression(expr.Initializer)
}