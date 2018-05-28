package docstore

import (
	"sync"
	"github.com/v2pro/quokka/docstore/runtime"
	"github.com/v2pro/quokka/docstore/transpiler"
	"github.com/v2pro/plz/countlog"
	"fmt"
	"encoding/base64"
	"crypto/md5"
)

var entityTypes = map[string]*entityTypeDef{}
var entityTypesMutex = &sync.Mutex{}

type CommandHandler func(doc interface{}, request interface{}) (resp interface{})

type commandDef struct {
	requestSchema  *runtime.Schema
	responseSchema *runtime.Schema
	handlerVersion string
	handler        CommandHandler
}

type entityTypeDef struct {
	schema           *runtime.Schema
	commandDefs      map[string]*commandDef
	commandDefsMutex *sync.Mutex
}

func (typ *entityTypeDef) Command(commandType string, jsHandler string, thriftIDL string) *entityTypeDef {
	handler, err := transpiler.Compile(jsHandler)
	if err != nil {
		countlog.Error("event!command_handler.failed to transpile javascript",
			"err", err, "jsHandler", jsHandler)
		return typ
	}
	schemas, err := runtime.ThriftSchemas(thriftIDL)
	if err != nil {
		panic(err)
	}
	typ.AddCommand(commandType, handler, calcJsHandlerVersion(jsHandler), schemas["Request"], schemas["Response"])
	return typ
}

func calcJsHandlerVersion(jsHandler string) string {
	hash := md5.New()
	hash.Write([]byte(jsHandler))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}

func (typ *entityTypeDef) AddCommand(commandType string, handler CommandHandler, handlerVersion string,
	requestSchema *runtime.Schema, responseSchema *runtime.Schema) *entityTypeDef {
	typ.commandDefsMutex.Lock()
	defer typ.commandDefsMutex.Unlock()
	typ.commandDefs[commandType] = &commandDef{
		requestSchema:  requestSchema,
		responseSchema: responseSchema,
		handler:        handler,
		handlerVersion: handlerVersion,
	}
	return typ
}

func (typ *entityTypeDef) getCommandDef(commandType string) *commandDef {
	typ.commandDefsMutex.Lock()
	defer typ.commandDefsMutex.Unlock()
	return typ.commandDefs[commandType]
}

func Entity(entityType string, thriftIDL string) *entityTypeDef {
	schemas, err := runtime.ThriftSchemas(thriftIDL)
	if err != nil {
		panic(fmt.Sprintf("%s entity schema invalid: %s", entityType, err.Error()))
	}
	return AddEntity(entityType, schemas["Entity"])
}

func AddEntity(entityType string, schema *runtime.Schema) *entityTypeDef {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	if entityTypes[entityType] != nil {
		panic("already has store for " + entityType)
	}
	typ := &entityTypeDef{
		schema:           schema,
		commandDefs:      map[string]*commandDef{},
		commandDefsMutex: &sync.Mutex{},
	}
	entityTypes[entityType] = typ
	return typ
}

func getEntityType(entityType string) *entityTypeDef {
	entityTypesMutex.Lock()
	defer entityTypesMutex.Unlock()
	return entityTypes[entityType]
}
