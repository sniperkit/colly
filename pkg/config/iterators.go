package config

import (
	jsoniter "github.com/json-iterator/go"
	thrifter "github.com/thrift-iterator/go"
	"github.com/thrift-iterator/go/general"
	"github.com/thrift-iterator/go/protocol"
	"github.com/thrift-iterator/go/raw"
)

var (
	// Jsoniter
	json = jsoniter.ConfigCompatibleWithStandardLibrary
	// Thriftier
	thriftConfig           *thrifter.Config
	thriftMessageHeader    *protocol.MessageHeader
	thriftMessageGeneral   *general.Message
	thriftMessageArguments *raw.Struct
)
