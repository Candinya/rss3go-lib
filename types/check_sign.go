package types

import (
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/nyawork/rss3go_lib/utils"
	"strings"
)

func getSigner(msgHashBytes []byte, signHex string) (string, error) {

	signHex = utils.FixSign(signHex)

	signatureReceived, err := hex.DecodeString(signHex)
	if err != nil {
		return "", err
	}

	recoveredPubBytes, err := crypto.Ecrecover(msgHashBytes, signatureReceived)
	if err != nil {
		return "", err
	}

	recoveredPub, err := crypto.UnmarshalPubkey(recoveredPubBytes)
	if err != nil {
		return "", err
	}

	signer := crypto.PubkeyToAddress(*recoveredPub).String()

	return signer, nil
}

func (o *RSS3) CheckSign() (bool, error) {

	if o.Profile != nil {
		if success, err := o.Profile.CheckSign(o.Id); err != nil {
			return false, err
		} else if !success {
			return false, nil
		}
	}

	if o.Links != nil {
		for index, _ := range o.Links {
			l := &o.Links[index]
			if success, err := l.CheckSign(o.Id); err != nil {
				return false, err
			} else if !success {
				return false, nil
			}
		}
	}

	if o.Items != nil {
		for index, _ := range o.Items {
			i := &o.Items[index]
			if success, err := i.CheckSign(); err != nil {
				return false, err
			} else if !success {
				return false, nil
			}
		}
	}

	// Check self

	signHex := o.Signature

	if signHex == "" {
		return false, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS34S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == o.Id, err

}

func (o *RSS3Profile4F) CheckSign(Id RSS3ID) (bool, error) {

	signHex := o.Signature

	if signHex == "" {
		return true, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS3Profile4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == Id, err

}

func (o *RSS3Links4F) CheckSign(Id RSS3ID) (bool, error) {

	signHex := o.Signature

	if signHex == "" {
		return false, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS3Links4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == Id, err

}

func (o *RSS3Item4F) CheckSign() (bool, error) {

	signHex := o.Signature

	if signHex == "" {
		return false, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS3Item4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == strings.Split(o.Id, "-")[0], err

}

func (o *RSS3Items4F) CheckSign() (bool, error) {

	signHex := o.Signature

	if signHex == "" {
		return false, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS3Items4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == strings.Split(o.Id, "-")[0], err

}

func (o *RSS3List4F) CheckSign() (bool, error) {

	signHex := o.Signature

	if signHex == "" {
		return false, nil
	}

	jsonBytes, err := json.Marshal(o)
	if err != nil {
		return false, err
	}

	var objCheckSign RSS3List4S

	err = json.Unmarshal(jsonBytes, &objCheckSign)
	if err != nil {
		return false, err
	}

	jsonBytes4S, err := json.Marshal(objCheckSign)
	if err != nil {
		return false, err
	}

	msgHashBytes := crypto.Keccak256Hash(jsonBytes4S).Bytes()

	signer, err := getSigner(msgHashBytes, signHex)

	return signer == strings.Split(o.Id, "-")[0], err

}
