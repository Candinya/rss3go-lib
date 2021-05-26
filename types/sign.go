package types

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nyawork/rss3go_lib/utils"
)

func (o *RSS3) SetSign(key *ecdsa.PrivateKey) {

	o.Profile.SetSign(key)

	ethAddr := crypto.PubkeyToAddress(key.PublicKey).String()

	o.Id = ethAddr

	for index, _ := range o.Links {
		l := &o.Links[index]
		l.SetSign(key)
	}

	for index, _ := range o.Items {
		i := &o.Items[index]
		i.Id = fmt.Sprintf("%s-item-%d", ethAddr, index)
		//if index != 0 {
		//	i.Upstream = fmt.Sprintf("%s-item-%d", ethAddr, index-1)
		//} else {
		//	i.Upstream = ""
		//}
		i.SetSign(key)
	}

	// Sign self

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS34S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}

func (o *RSS3Profile4F) SetSign(key *ecdsa.PrivateKey) {

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS3Profile4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}

func (o *RSS3Links4F) SetSign(key *ecdsa.PrivateKey) {

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS3Links4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}

func (o *RSS3Item4F) SetSign(key *ecdsa.PrivateKey) {

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS3Item4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}

func (o *RSS3Items4F) SetSign(key *ecdsa.PrivateKey) {

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS3Items4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}

func (o *RSS3List4F) SetSign(key *ecdsa.PrivateKey) {

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return
	}

	var objCheckSign RSS3List4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return
	}

	sign := utils.SignMsg(key, jsonBytes4S)

	signHex := hex.EncodeToString(sign)

	o.Signature = signHex

}
