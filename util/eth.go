package util

import (
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignVerify 验证以太坊签名
// addrHex: 钱包地址Hex
// signedHex: 用私钥对原始数据签名后的密文Hex
// message: 原始数据
func SignVerify(addrHex, signedHex, message string) bool {
	signed, err := hexutil.Decode(signedHex)
	if err != nil {
		log.Println(err)
		return false
	}
	if signed[64] != 1 && signed[64] != 0 && signed[64] != 27 && signed[64] != 28 {
		log.Println("signedHex Decode error.")
		return false
	}

	if signed[64] != 1 && signed[64] != 0 {
		signed[64] -= 27
	}

	pubKey, errS := crypto.SigToPub(crypto.Keccak256([]byte(message)), signed)
	if err != nil {
		log.Println(errS)
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*pubKey)
	return addrHex == recoveredAddr.Hex()
}
