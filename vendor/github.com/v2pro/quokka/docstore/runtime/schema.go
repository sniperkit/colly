package runtime

import (
	"reflect"
	"errors"
	"github.com/v2pro/quokka/docstore/thrift"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"fmt"
	"github.com/v2pro/plz/countlog"
)

var thriftTypes = map[string]*Schema{
	"string": mustBeString,
	"i64": mustBeInteger,
	"float64": mustBeNumber,
}

type Schema struct {
	Fields    map[string]*Schema
	Element   *Schema
	Validator func(value interface{}) error
}

func Validate(obj interface{}, schema *Schema) (err error) {
	defer func() {
		recovered := recover()
		if recovered != nil {
			err, _ = recovered.(error)
			if err == nil {
				err = errors.New(fmt.Sprintf("%v", recovered))
			}
		}
	}()
	if schema == nil {
		return nil
	}
	if schema.Validator != nil {
		return schema.Validator(obj)
	} else if schema.Fields != nil {
		dobj, _ := obj.(*DObject)
		if dobj == nil {
			return fmt.Errorf("%v is not object", obj)
		}
		dobj.Schema = schema
		dobj.validate()
	} else if schema.Element != nil {
		dlist, _ := obj.(*DList)
		if dlist == nil {
			return fmt.Errorf("%v is not list", obj)
		}
		dlist.Schema = schema
		dlist.validate()
	}
	return nil
}

func (obj *DList) validate() {
	if obj.Schema == nil {
		return
	}
	if obj.Schema.Element == nil {
		return
	}
	for key, value := range obj.data {
		obj.validateElem(key, value)
	}
}

func (obj *DList) validateElem(key int, value interface{}) {
	if obj.Schema == nil {
		return
	}
	elemSchema := obj.Schema.Element
	if elemSchema == nil {
		return
	}
	err := Validate(value, elemSchema)
	if err != nil {
		panic(fmt.Errorf("[%v]->%v", key, err.Error()))
	}
}

func (obj *DObject) validate() {
	if obj.Schema == nil {
		return
	}
	for key, value := range obj.data {
		obj.validateKey(key, value)
	}
}

func (obj *DObject) validateKey(key string, value interface{}) {
	if obj.Schema == nil {
		return
	}
	fieldSchema := obj.Schema.Fields[key]
	if fieldSchema == nil {
		fieldSchema = obj.Schema.Element
	}
	if fieldSchema == nil {
		return
	}
	err := Validate(value, fieldSchema)
	if err != nil {
		panic(fmt.Errorf("%s->%v", key, err.Error()))
	}
}

var stringType = reflect.TypeOf((*string)(nil)).Elem()
var mustBeString = &Schema{
	Validator: func(value interface{}) error {
		if reflect.TypeOf(value) != stringType {
			return errors.New("must be string")
		}
		return nil
	},
}

var intType = reflect.TypeOf((*int)(nil)).Elem()
var float64Type = reflect.TypeOf((*float64)(nil)).Elem()
var mustBeInteger = &Schema{
	Validator: func(value interface{}) error {
		if reflect.TypeOf(value) != intType {
			return errors.New("must be int")
		}
		return nil
	},
}
var mustBeNumber = &Schema{
	Validator: func(value interface{}) error {
		if reflect.TypeOf(value) != intType && reflect.TypeOf(value) != float64Type {
			return errors.New("must be number")
		}
		return nil
	},
}

func ThriftSchemas(thriftIDL string) (map[string]*Schema, error) {
	charStream := antlr.NewInputStream(thriftIDL)
	lexer := thrift.NewAntlrThriftLexer(charStream)
	tokenStream := antlr.NewCommonTokenStream(lexer, 0)
	parser := thrift.NewAntlrThriftParser(tokenStream)
	schemas := map[string]*Schema{}
	for charStream.Index() < len(thriftIDL) {
		structDef := parser.StructDef()
		structName := structDef.GetChild(1).GetPayload().(*antlr.CommonToken).GetText()
		schema := &Schema{
			Fields: map[string]*Schema{},
		}
		for i := 3; i < structDef.GetChildCount(); i++ {
			field, _ := structDef.GetChild(i).(*thrift.FieldContext)
			if field == nil {
				break
			}
			_, isFieldType := field.GetChild(3).(*thrift.FieldTypeContext)
			idx := 3
			if !isFieldType {
				idx = 2
			}
			fieldName := field.GetChild(idx + 1).GetPayload().(*antlr.CommonToken).GetText()
			fieldTypeNode := getFieldTypeNode(field.GetChild(idx))
			fieldTypeSchema, err := fieldTypeToSchema(schemas, fieldTypeNode)
			if err != nil {
				return nil, err
			}
			schema.Fields[fieldName] = fieldTypeSchema
		}
		schemas[structName] = schema
	}
	return schemas, nil
}

func getFieldTypeNode(node antlr.Tree) interface{} {
	switch node.(type) {
	case *thrift.ListTypeContext:
		return node.GetPayload()
	case *thrift.MapTypeContext:
		return node.GetPayload()
	}
	if node.GetChild(0) == nil {
		return node.GetPayload()
	}
	return getFieldTypeNode(node.GetChild(0))
}

func fieldTypeToSchema(schemas map[string]*Schema, fieldTypeNode interface{}) (*Schema, error) {
	switch typedNode := fieldTypeNode.(type) {
	case *antlr.CommonToken:
		typeSchema := thriftTypes[typedNode.GetText()]
		if typeSchema == nil {
			typeSchema = schemas[typedNode.GetText()]
		}
		if typeSchema == nil {
			countlog.Warn("event!schema.no schema defined for type", "type", typedNode.GetText())
		}
		return typeSchema, nil
	case *antlr.BaseParserRuleContext:
		containerType := typedNode.GetChild(0).GetPayload().(*antlr.CommonToken).GetText()
		switch containerType {
		case "list":
			elem, err := fieldTypeToSchema(schemas, getFieldTypeNode(typedNode.GetChild(2)))
			if err != nil {
				return nil, err
			}
			return &Schema{Element: elem}, nil
		case "map":
			elem, err := fieldTypeToSchema(schemas, getFieldTypeNode(typedNode.GetChild(4)))
			if err != nil {
				return nil, err
			}
			return &Schema{Fields: map[string]*Schema{}, Element: elem}, nil
		default:
			return nil, fmt.Errorf("does not support container type: %s", containerType)
		}
	default:
		return nil, fmt.Errorf("unexpected node type: %s", reflect.TypeOf(typedNode).String())
	}
}
