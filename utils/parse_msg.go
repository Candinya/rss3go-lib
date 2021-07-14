package utils

import (
	"strconv"
)

// Refer to https://web3js.readthedocs.io/en/v1.4.0/web3-eth-accounts.html#sign
// https://github.com/ChainSafe/web3.js/blob/9b7540d283c3c72978ebdfacc0dea2d0f3f36a05/packages/web3-eth-accounts/src/index.js#L299-L307

func ParseMsgToWeb3JSSignFormat(message string) []byte {

	// Is not hex
	src := []byte(message)

	preamble := []byte("\x19Ethereum Signed Message:\n" + strconv.Itoa(len(src)))
	wrapped := append(preamble, src...)

	return wrapped
}
