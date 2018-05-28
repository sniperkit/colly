// Generated from AntlrThrift.g4 by ANTLR 4.7.

package thrift // AntlrThrift

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseAntlrThriftListener is a complete listener for a parse tree produced by AntlrThriftParser.
type BaseAntlrThriftListener struct{}

var _ AntlrThriftListener = &BaseAntlrThriftListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseAntlrThriftListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseAntlrThriftListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseAntlrThriftListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseAntlrThriftListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterDocument is called when production document is entered.
func (s *BaseAntlrThriftListener) EnterDocument(ctx *DocumentContext) {}

// ExitDocument is called when production document is exited.
func (s *BaseAntlrThriftListener) ExitDocument(ctx *DocumentContext) {}

// EnterHeader is called when production header is entered.
func (s *BaseAntlrThriftListener) EnterHeader(ctx *HeaderContext) {}

// ExitHeader is called when production header is exited.
func (s *BaseAntlrThriftListener) ExitHeader(ctx *HeaderContext) {}

// EnterInclude is called when production include is entered.
func (s *BaseAntlrThriftListener) EnterInclude(ctx *IncludeContext) {}

// ExitInclude is called when production include is exited.
func (s *BaseAntlrThriftListener) ExitInclude(ctx *IncludeContext) {}

// EnterCppInclude is called when production cppInclude is entered.
func (s *BaseAntlrThriftListener) EnterCppInclude(ctx *CppIncludeContext) {}

// ExitCppInclude is called when production cppInclude is exited.
func (s *BaseAntlrThriftListener) ExitCppInclude(ctx *CppIncludeContext) {}

// EnterNamespace is called when production namespace is entered.
func (s *BaseAntlrThriftListener) EnterNamespace(ctx *NamespaceContext) {}

// ExitNamespace is called when production namespace is exited.
func (s *BaseAntlrThriftListener) ExitNamespace(ctx *NamespaceContext) {}

// EnterStandardNamespace is called when production standardNamespace is entered.
func (s *BaseAntlrThriftListener) EnterStandardNamespace(ctx *StandardNamespaceContext) {}

// ExitStandardNamespace is called when production standardNamespace is exited.
func (s *BaseAntlrThriftListener) ExitStandardNamespace(ctx *StandardNamespaceContext) {}

// EnterNamespaceScope is called when production namespaceScope is entered.
func (s *BaseAntlrThriftListener) EnterNamespaceScope(ctx *NamespaceScopeContext) {}

// ExitNamespaceScope is called when production namespaceScope is exited.
func (s *BaseAntlrThriftListener) ExitNamespaceScope(ctx *NamespaceScopeContext) {}

// EnterPhpNamespace is called when production phpNamespace is entered.
func (s *BaseAntlrThriftListener) EnterPhpNamespace(ctx *PhpNamespaceContext) {}

// ExitPhpNamespace is called when production phpNamespace is exited.
func (s *BaseAntlrThriftListener) ExitPhpNamespace(ctx *PhpNamespaceContext) {}

// EnterXsdNamespace is called when production xsdNamespace is entered.
func (s *BaseAntlrThriftListener) EnterXsdNamespace(ctx *XsdNamespaceContext) {}

// ExitXsdNamespace is called when production xsdNamespace is exited.
func (s *BaseAntlrThriftListener) ExitXsdNamespace(ctx *XsdNamespaceContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseAntlrThriftListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseAntlrThriftListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterConstDef is called when production constDef is entered.
func (s *BaseAntlrThriftListener) EnterConstDef(ctx *ConstDefContext) {}

// ExitConstDef is called when production constDef is exited.
func (s *BaseAntlrThriftListener) ExitConstDef(ctx *ConstDefContext) {}

// EnterConstValue is called when production constValue is entered.
func (s *BaseAntlrThriftListener) EnterConstValue(ctx *ConstValueContext) {}

// ExitConstValue is called when production constValue is exited.
func (s *BaseAntlrThriftListener) ExitConstValue(ctx *ConstValueContext) {}

// EnterConstList is called when production constList is entered.
func (s *BaseAntlrThriftListener) EnterConstList(ctx *ConstListContext) {}

// ExitConstList is called when production constList is exited.
func (s *BaseAntlrThriftListener) ExitConstList(ctx *ConstListContext) {}

// EnterConstMap is called when production constMap is entered.
func (s *BaseAntlrThriftListener) EnterConstMap(ctx *ConstMapContext) {}

// ExitConstMap is called when production constMap is exited.
func (s *BaseAntlrThriftListener) ExitConstMap(ctx *ConstMapContext) {}

// EnterConstMapEntry is called when production constMapEntry is entered.
func (s *BaseAntlrThriftListener) EnterConstMapEntry(ctx *ConstMapEntryContext) {}

// ExitConstMapEntry is called when production constMapEntry is exited.
func (s *BaseAntlrThriftListener) ExitConstMapEntry(ctx *ConstMapEntryContext) {}

// EnterTypedef is called when production typedef is entered.
func (s *BaseAntlrThriftListener) EnterTypedef(ctx *TypedefContext) {}

// ExitTypedef is called when production typedef is exited.
func (s *BaseAntlrThriftListener) ExitTypedef(ctx *TypedefContext) {}

// EnterEnumDef is called when production enumDef is entered.
func (s *BaseAntlrThriftListener) EnterEnumDef(ctx *EnumDefContext) {}

// ExitEnumDef is called when production enumDef is exited.
func (s *BaseAntlrThriftListener) ExitEnumDef(ctx *EnumDefContext) {}

// EnterEnumMember is called when production enumMember is entered.
func (s *BaseAntlrThriftListener) EnterEnumMember(ctx *EnumMemberContext) {}

// ExitEnumMember is called when production enumMember is exited.
func (s *BaseAntlrThriftListener) ExitEnumMember(ctx *EnumMemberContext) {}

// EnterSenum is called when production senum is entered.
func (s *BaseAntlrThriftListener) EnterSenum(ctx *SenumContext) {}

// ExitSenum is called when production senum is exited.
func (s *BaseAntlrThriftListener) ExitSenum(ctx *SenumContext) {}

// EnterStructDef is called when production structDef is entered.
func (s *BaseAntlrThriftListener) EnterStructDef(ctx *StructDefContext) {}

// ExitStructDef is called when production structDef is exited.
func (s *BaseAntlrThriftListener) ExitStructDef(ctx *StructDefContext) {}

// EnterUnionDef is called when production unionDef is entered.
func (s *BaseAntlrThriftListener) EnterUnionDef(ctx *UnionDefContext) {}

// ExitUnionDef is called when production unionDef is exited.
func (s *BaseAntlrThriftListener) ExitUnionDef(ctx *UnionDefContext) {}

// EnterExceptionDef is called when production exceptionDef is entered.
func (s *BaseAntlrThriftListener) EnterExceptionDef(ctx *ExceptionDefContext) {}

// ExitExceptionDef is called when production exceptionDef is exited.
func (s *BaseAntlrThriftListener) ExitExceptionDef(ctx *ExceptionDefContext) {}

// EnterServiceDef is called when production serviceDef is entered.
func (s *BaseAntlrThriftListener) EnterServiceDef(ctx *ServiceDefContext) {}

// ExitServiceDef is called when production serviceDef is exited.
func (s *BaseAntlrThriftListener) ExitServiceDef(ctx *ServiceDefContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseAntlrThriftListener) EnterFunction(ctx *FunctionContext) {}

// ExitFunction is called when production function is exited.
func (s *BaseAntlrThriftListener) ExitFunction(ctx *FunctionContext) {}

// EnterFieldList is called when production fieldList is entered.
func (s *BaseAntlrThriftListener) EnterFieldList(ctx *FieldListContext) {}

// ExitFieldList is called when production fieldList is exited.
func (s *BaseAntlrThriftListener) ExitFieldList(ctx *FieldListContext) {}

// EnterField is called when production field is entered.
func (s *BaseAntlrThriftListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseAntlrThriftListener) ExitField(ctx *FieldContext) {}

// EnterRequiredness is called when production requiredness is entered.
func (s *BaseAntlrThriftListener) EnterRequiredness(ctx *RequirednessContext) {}

// ExitRequiredness is called when production requiredness is exited.
func (s *BaseAntlrThriftListener) ExitRequiredness(ctx *RequirednessContext) {}

// EnterThrowsList is called when production throwsList is entered.
func (s *BaseAntlrThriftListener) EnterThrowsList(ctx *ThrowsListContext) {}

// ExitThrowsList is called when production throwsList is exited.
func (s *BaseAntlrThriftListener) ExitThrowsList(ctx *ThrowsListContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseAntlrThriftListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseAntlrThriftListener) ExitFieldType(ctx *FieldTypeContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseAntlrThriftListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseAntlrThriftListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterContainerType is called when production containerType is entered.
func (s *BaseAntlrThriftListener) EnterContainerType(ctx *ContainerTypeContext) {}

// ExitContainerType is called when production containerType is exited.
func (s *BaseAntlrThriftListener) ExitContainerType(ctx *ContainerTypeContext) {}

// EnterMapType is called when production mapType is entered.
func (s *BaseAntlrThriftListener) EnterMapType(ctx *MapTypeContext) {}

// ExitMapType is called when production mapType is exited.
func (s *BaseAntlrThriftListener) ExitMapType(ctx *MapTypeContext) {}

// EnterListType is called when production listType is entered.
func (s *BaseAntlrThriftListener) EnterListType(ctx *ListTypeContext) {}

// ExitListType is called when production listType is exited.
func (s *BaseAntlrThriftListener) ExitListType(ctx *ListTypeContext) {}

// EnterSetType is called when production setType is entered.
func (s *BaseAntlrThriftListener) EnterSetType(ctx *SetTypeContext) {}

// ExitSetType is called when production setType is exited.
func (s *BaseAntlrThriftListener) ExitSetType(ctx *SetTypeContext) {}

// EnterCppType is called when production cppType is entered.
func (s *BaseAntlrThriftListener) EnterCppType(ctx *CppTypeContext) {}

// ExitCppType is called when production cppType is exited.
func (s *BaseAntlrThriftListener) ExitCppType(ctx *CppTypeContext) {}

// EnterAnnotationList is called when production annotationList is entered.
func (s *BaseAntlrThriftListener) EnterAnnotationList(ctx *AnnotationListContext) {}

// ExitAnnotationList is called when production annotationList is exited.
func (s *BaseAntlrThriftListener) ExitAnnotationList(ctx *AnnotationListContext) {}

// EnterAnnotation is called when production annotation is entered.
func (s *BaseAntlrThriftListener) EnterAnnotation(ctx *AnnotationContext) {}

// ExitAnnotation is called when production annotation is exited.
func (s *BaseAntlrThriftListener) ExitAnnotation(ctx *AnnotationContext) {}

// EnterSeparator is called when production separator is entered.
func (s *BaseAntlrThriftListener) EnterSeparator(ctx *SeparatorContext) {}

// ExitSeparator is called when production separator is exited.
func (s *BaseAntlrThriftListener) ExitSeparator(ctx *SeparatorContext) {}
