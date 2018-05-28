// Generated from AntlrThrift.g4 by ANTLR 4.7.

package thrift // AntlrThrift

import "github.com/antlr/antlr4/runtime/Go/antlr"

// AntlrThriftListener is a complete listener for a parse tree produced by AntlrThriftParser.
type AntlrThriftListener interface {
	antlr.ParseTreeListener

	// EnterDocument is called when entering the document production.
	EnterDocument(c *DocumentContext)

	// EnterHeader is called when entering the header production.
	EnterHeader(c *HeaderContext)

	// EnterInclude is called when entering the include production.
	EnterInclude(c *IncludeContext)

	// EnterCppInclude is called when entering the cppInclude production.
	EnterCppInclude(c *CppIncludeContext)

	// EnterNamespace is called when entering the namespace production.
	EnterNamespace(c *NamespaceContext)

	// EnterStandardNamespace is called when entering the standardNamespace production.
	EnterStandardNamespace(c *StandardNamespaceContext)

	// EnterNamespaceScope is called when entering the namespaceScope production.
	EnterNamespaceScope(c *NamespaceScopeContext)

	// EnterPhpNamespace is called when entering the phpNamespace production.
	EnterPhpNamespace(c *PhpNamespaceContext)

	// EnterXsdNamespace is called when entering the xsdNamespace production.
	EnterXsdNamespace(c *XsdNamespaceContext)

	// EnterDefinition is called when entering the definition production.
	EnterDefinition(c *DefinitionContext)

	// EnterConstDef is called when entering the constDef production.
	EnterConstDef(c *ConstDefContext)

	// EnterConstValue is called when entering the constValue production.
	EnterConstValue(c *ConstValueContext)

	// EnterConstList is called when entering the constList production.
	EnterConstList(c *ConstListContext)

	// EnterConstMap is called when entering the constMap production.
	EnterConstMap(c *ConstMapContext)

	// EnterConstMapEntry is called when entering the constMapEntry production.
	EnterConstMapEntry(c *ConstMapEntryContext)

	// EnterTypedef is called when entering the typedef production.
	EnterTypedef(c *TypedefContext)

	// EnterEnumDef is called when entering the enumDef production.
	EnterEnumDef(c *EnumDefContext)

	// EnterEnumMember is called when entering the enumMember production.
	EnterEnumMember(c *EnumMemberContext)

	// EnterSenum is called when entering the senum production.
	EnterSenum(c *SenumContext)

	// EnterStructDef is called when entering the structDef production.
	EnterStructDef(c *StructDefContext)

	// EnterUnionDef is called when entering the unionDef production.
	EnterUnionDef(c *UnionDefContext)

	// EnterExceptionDef is called when entering the exceptionDef production.
	EnterExceptionDef(c *ExceptionDefContext)

	// EnterServiceDef is called when entering the serviceDef production.
	EnterServiceDef(c *ServiceDefContext)

	// EnterFunction is called when entering the function production.
	EnterFunction(c *FunctionContext)

	// EnterFieldList is called when entering the fieldList production.
	EnterFieldList(c *FieldListContext)

	// EnterField is called when entering the field production.
	EnterField(c *FieldContext)

	// EnterRequiredness is called when entering the requiredness production.
	EnterRequiredness(c *RequirednessContext)

	// EnterThrowsList is called when entering the throwsList production.
	EnterThrowsList(c *ThrowsListContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// EnterBaseType is called when entering the baseType production.
	EnterBaseType(c *BaseTypeContext)

	// EnterContainerType is called when entering the containerType production.
	EnterContainerType(c *ContainerTypeContext)

	// EnterMapType is called when entering the mapType production.
	EnterMapType(c *MapTypeContext)

	// EnterListType is called when entering the listType production.
	EnterListType(c *ListTypeContext)

	// EnterSetType is called when entering the setType production.
	EnterSetType(c *SetTypeContext)

	// EnterCppType is called when entering the cppType production.
	EnterCppType(c *CppTypeContext)

	// EnterAnnotationList is called when entering the annotationList production.
	EnterAnnotationList(c *AnnotationListContext)

	// EnterAnnotation is called when entering the annotation production.
	EnterAnnotation(c *AnnotationContext)

	// EnterSeparator is called when entering the separator production.
	EnterSeparator(c *SeparatorContext)

	// ExitDocument is called when exiting the document production.
	ExitDocument(c *DocumentContext)

	// ExitHeader is called when exiting the header production.
	ExitHeader(c *HeaderContext)

	// ExitInclude is called when exiting the include production.
	ExitInclude(c *IncludeContext)

	// ExitCppInclude is called when exiting the cppInclude production.
	ExitCppInclude(c *CppIncludeContext)

	// ExitNamespace is called when exiting the namespace production.
	ExitNamespace(c *NamespaceContext)

	// ExitStandardNamespace is called when exiting the standardNamespace production.
	ExitStandardNamespace(c *StandardNamespaceContext)

	// ExitNamespaceScope is called when exiting the namespaceScope production.
	ExitNamespaceScope(c *NamespaceScopeContext)

	// ExitPhpNamespace is called when exiting the phpNamespace production.
	ExitPhpNamespace(c *PhpNamespaceContext)

	// ExitXsdNamespace is called when exiting the xsdNamespace production.
	ExitXsdNamespace(c *XsdNamespaceContext)

	// ExitDefinition is called when exiting the definition production.
	ExitDefinition(c *DefinitionContext)

	// ExitConstDef is called when exiting the constDef production.
	ExitConstDef(c *ConstDefContext)

	// ExitConstValue is called when exiting the constValue production.
	ExitConstValue(c *ConstValueContext)

	// ExitConstList is called when exiting the constList production.
	ExitConstList(c *ConstListContext)

	// ExitConstMap is called when exiting the constMap production.
	ExitConstMap(c *ConstMapContext)

	// ExitConstMapEntry is called when exiting the constMapEntry production.
	ExitConstMapEntry(c *ConstMapEntryContext)

	// ExitTypedef is called when exiting the typedef production.
	ExitTypedef(c *TypedefContext)

	// ExitEnumDef is called when exiting the enumDef production.
	ExitEnumDef(c *EnumDefContext)

	// ExitEnumMember is called when exiting the enumMember production.
	ExitEnumMember(c *EnumMemberContext)

	// ExitSenum is called when exiting the senum production.
	ExitSenum(c *SenumContext)

	// ExitStructDef is called when exiting the structDef production.
	ExitStructDef(c *StructDefContext)

	// ExitUnionDef is called when exiting the unionDef production.
	ExitUnionDef(c *UnionDefContext)

	// ExitExceptionDef is called when exiting the exceptionDef production.
	ExitExceptionDef(c *ExceptionDefContext)

	// ExitServiceDef is called when exiting the serviceDef production.
	ExitServiceDef(c *ServiceDefContext)

	// ExitFunction is called when exiting the function production.
	ExitFunction(c *FunctionContext)

	// ExitFieldList is called when exiting the fieldList production.
	ExitFieldList(c *FieldListContext)

	// ExitField is called when exiting the field production.
	ExitField(c *FieldContext)

	// ExitRequiredness is called when exiting the requiredness production.
	ExitRequiredness(c *RequirednessContext)

	// ExitThrowsList is called when exiting the throwsList production.
	ExitThrowsList(c *ThrowsListContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)

	// ExitBaseType is called when exiting the baseType production.
	ExitBaseType(c *BaseTypeContext)

	// ExitContainerType is called when exiting the containerType production.
	ExitContainerType(c *ContainerTypeContext)

	// ExitMapType is called when exiting the mapType production.
	ExitMapType(c *MapTypeContext)

	// ExitListType is called when exiting the listType production.
	ExitListType(c *ListTypeContext)

	// ExitSetType is called when exiting the setType production.
	ExitSetType(c *SetTypeContext)

	// ExitCppType is called when exiting the cppType production.
	ExitCppType(c *CppTypeContext)

	// ExitAnnotationList is called when exiting the annotationList production.
	ExitAnnotationList(c *AnnotationListContext)

	// ExitAnnotation is called when exiting the annotation production.
	ExitAnnotation(c *AnnotationContext)

	// ExitSeparator is called when exiting the separator production.
	ExitSeparator(c *SeparatorContext)
}
