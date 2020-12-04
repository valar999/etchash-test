package main

import (
	"fmt"
	"math/big"
	"strconv"
	"github.com/valar999/etchash-test/etchash"
	"github.com/ethereum/go-ethereum/common"
)

type EtcBlock struct {
	diff int64
	hashNoNonce string
	nonce string
	mixDigest string
	number uint64
}

func (b EtcBlock) Difficulty() *big.Int {
	return big.NewInt(b.diff)
}

func (b EtcBlock) HashNoNonce() common.Hash {
	return common.HexToHash(b.hashNoNonce)
}

func (b EtcBlock) Nonce() uint64 {
	value, _ := strconv.ParseUint(b.nonce, 16, 64)
	return value
}

func (b EtcBlock) MixDigest() common.Hash {
	return common.HexToHash(b.mixDigest)
}

func (b EtcBlock) NumberU64() uint64 {
	return b.number
}

func main() {
	var ecip1099FBlockClassic uint64 = 11700000
	var hasher = etchash.New(&ecip1099FBlockClassic)
	block := EtcBlock{
		diff: 0,
		hashNoNonce: "cf1a5330162c4e138aa3a37ad270040d29c73f3d41b1748a2c912d8944e3bb42",
		nonce: "002930008bd54abd",
		mixDigest: "e79f0f63030bf691445c2b9d0266b24a9619e355194067f2ad2c73a8e0a26c65",
		number: 11736639,
	}
	result := hasher.Verify(block)
	fmt.Printf("%#v\n%v\n", block, result)
}
