package protocol_v1

import (
	"encoding/binary"
	"fmt"

	p "github.com/trist725/mgsu/network/protocol"
)

const (
	PH = 0
)

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
	data := p.allocator.Alloc(iMsg.Size() + 4)

	// 拷贝消息id
	binary.BigEndian.PutUint32(data[:4], iMsg.MessageID())

	// 拷贝消息
	if _, err := iMsg.MarshalTo(data[4:]); err != nil {
		return nil, fmt.Errorf("marshal [%v] fail, %v", iMsg, err)
	}

	err := p.encryptor.Encrypt(data)
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

	id := binary.BigEndian.Uint32(data[:4])

	iMsg, err := p.Produce(id)
	if err != nil {
		return nil, fmt.Errorf("produce message fail, %v", err)
	}

	if err := iMsg.Unmarshal(data[4:]); err != nil {
		return nil, fmt.Errorf("unmarshal [%v] fail, %v", iMsg, err)
	}

	return iMsg, nil
}
