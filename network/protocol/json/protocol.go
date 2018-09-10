package protocol_json

import (
	"fmt"

	p "gitee.com/nggs/network/protocol"
	jsoniter "github.com/json-iterator/go"
)

const (
	PH = 0
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type IProtocol interface {
	IMessageFactoryManager
	Encode(i IMessage) ([]byte, error)
	Decode(data []byte) (IMessage, error)
}

type protocol struct {
	IMessageFactoryManager
	allocator p.IAllocator
	encryptor p.IEncryptor
	decryptor p.IDecryptor
}

func New(allocator p.IAllocator, encryptor p.IEncryptor, decryptor p.IDecryptor) IProtocol {
	proto := &protocol{
		IMessageFactoryManager: newMessageFactoryManager(),
		allocator:              &p.NonAllocator{},
		encryptor:              &p.NonEncryptor{},
		decryptor:              &p.NonDecryptor{},
	}
	if allocator != nil {
		proto.allocator = allocator
	}
	if encryptor != nil {
		proto.encryptor = encryptor
	}
	if decryptor != nil {
		proto.decryptor = decryptor
	}
	return proto
}

func (p *protocol) Encode(iMsg IMessage) ([]byte, error) {
	data, err := json.Marshal(map[string]IMessage{iMsg.MessageID(): iMsg})
	if err != nil {
		return nil, fmt.Errorf("marshar [%v] fail, %v", iMsg, err)
	}

	err = p.encryptor.Encrypt(data)
	if err != nil {
		return nil, fmt.Errorf("encrypt [%v] fail, %v", iMsg, err)
	}

	return data, err
}

func (p *protocol) Decode(data []byte) (IMessage, error) {
	err := p.decryptor.Decrypt(data)
	if err != nil {
		return nil, fmt.Errorf("decrypt message fail, %v", err)
	}

	iter := jsoniter.ParseBytes(jsoniter.ConfigFastest, data)

	id := iter.ReadObject()
	iMsg, err := p.Produce(id)
	if err != nil {
		return nil, fmt.Errorf("produce message fail, %v", err)
	}
	iter.ReadVal(iMsg)
	err = iter.Error

	return iMsg, err
}
