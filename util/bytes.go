package util

import "encoding/binary"

func Int32ToByteArr(clientID int32, littleEndian bool) []byte {
	id := make([]byte, 4)
	if littleEndian {
		binary.LittleEndian.PutUint32(id, uint32(clientID))
	} else {
		binary.BigEndian.PutUint32(id, uint32(clientID))
	}
	return id
}
