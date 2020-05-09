package v3

import (
	"encoding/binary"
	"fmt"

	p "github.com/trist725/mgsu/network/protocol"
)

const (
	PH               = 0
	SizeofMessageID  = 2
	SizeofMessageSeq = 4
)

var (
	ByteOrder = binary.LittleEndian // 小端
)

type IProtocol interface {
	IMessageFactoryManager
	Alloc(size int) []byte
	Free(buffer []byte)
	Encrypt(data []byte) (err error)
	Decrypt(data []byte) (err error)
	EncodeTo(iMsg IMessage, seq MessageSeq, data []byte) (err error)
	Encode(iMsg IMessage, seq MessageSeq) (data []byte, err error)
	Decode(data []byte) (iMsg IMessage, seq MessageSeq, err error)
}

type protocol struct {
	IMessageFactoryManager
	allocator p.IAllocator
	encryptor p.IEncryptor
	decryptor p.IDecryptor
}

func New(allocator p.IAllocator, encryptor p.IEncryptor, decryptor p.IDecryptor) IProtocol {
	proto := &protocol{
		IMessageFactoryManager: NewMessageFactoryManager(),
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

func (p *protocol) Alloc(size int) []byte {
	return p.allocator.Alloc(size)
}

func (p *protocol) Free(buffer []byte) {
	p.allocator.Free(buffer)
}

func (p *protocol) Encrypt(data []byte) (err error) {
	return p.encryptor.Encrypt(data)
}

func (p *protocol) Decrypt(data []byte) (err error) {
	return p.decryptor.Decrypt(data)
}

func (p *protocol) EncodeTo(iMsg IMessage, seq MessageSeq, data []byte) (err error) {
	// 拷贝消息id
	ByteOrder.PutUint16(data[:SizeofMessageID], iMsg.MessageID())

	// 拷贝消息序列
	ByteOrder.PutUint32(data[SizeofMessageID:SizeofMessageID+SizeofMessageSeq], seq)

	// 拷贝消息
	_, err = iMsg.MarshalTo(data[SizeofMessageID+SizeofMessageSeq:])
	if err != nil {
		return fmt.Errorf("marshal [%v] fail, %v", iMsg, err)
	}

	err = p.encryptor.Encrypt(data)
	if err != nil {
		return fmt.Errorf("encrypt [%v] fail, %v", iMsg, err)
	}

	return
}

func (p *protocol) Encode(iMsg IMessage, seq MessageSeq) (data []byte, err error) {
	data = p.allocator.Alloc(SizeofMessageID + SizeofMessageSeq + iMsg.Size())

	// 拷贝消息id
	ByteOrder.PutUint16(data[:SizeofMessageID], iMsg.MessageID())

	// 拷贝消息序列
	ByteOrder.PutUint32(data[SizeofMessageID:SizeofMessageID+SizeofMessageSeq], seq)

	// 拷贝消息
	_, err = iMsg.MarshalTo(data[SizeofMessageID+SizeofMessageSeq:])
	if err != nil {
		return nil, fmt.Errorf("marshal %s fail, %v", iMsg.JsonString(), err)
	}

	err = p.encryptor.Encrypt(data)
	if err != nil {
		return nil, fmt.Errorf("encrypt message fail, id=%d, %s", iMsg.MessageID(), err)
	}

	return
}

func (p *protocol) Decode(data []byte) (iMsg IMessage, seq MessageSeq, err error) {
	err = p.decryptor.Decrypt(data)
	if err != nil {
		return nil, 0, fmt.Errorf("decrypt message fail, %s", err)
	}

	id := ByteOrder.Uint16(data[:SizeofMessageID])

	iMsg, err = p.Produce(id)
	if err != nil {
		return nil, 0, fmt.Errorf("produce message fail, %s", err)
	}

	seq = ByteOrder.Uint32(data[SizeofMessageID : SizeofMessageID+SizeofMessageSeq])

	err = iMsg.Unmarshal(data[SizeofMessageID+SizeofMessageSeq:])
	if err != nil {
		return nil, 0, fmt.Errorf("unmarshal message fail, id=%d, %s", id, err)
	}

	return
}
