package util

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// SignVerify 验证以太坊签名 eth_sign
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

// SignVerify191 验证EIP191签名  personal_sign
func SignVerify191(addrHex, signedHex, message string) bool {
	signed, err := hexutil.Decode(signedHex)
	if err != nil {
		log.Println(err)
		return false
	}

	if signed[64] < 27 {
		if !(signed[64] == 0 || signed[64] == 1) {
			log.Println("Invalid last byte")
			return false
		}
	} else {
		signed[64] -= 27 // shift byte?
	}

	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))

	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), signed)
	if err != nil {
		log.Printf("Ecrecover failed:%s\n", err.Error())
		return false
	}
	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		log.Printf("UnmarshalPubkey failed:%s\n", err.Error())
		return false
	}

	recoveredAddr := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()

	return strings.EqualFold(addrHex, recoveredAddr)
}

// PersonalSign EIP191签名
func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}

	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}
