package msg

import (
	protocol "github.com/trist725/mgsu/network/protocol/json"
	jsoniter "github.com/json-iterator/go"
)

var Protocol = protocol.New(nil, nil, nil)

var json = jsoniter.ConfigCompatibleWithStandardLibrary
