package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
)

func SignMsg(key *ecdsa.PrivateKey, msg []byte) []byte {

	msgHashBytes := crypto.Keccak256Hash(msg).Bytes()
	signature, err := crypto.Sign(msgHashBytes, key)
	if err != nil {
		return nil
	}
	return signature

}
