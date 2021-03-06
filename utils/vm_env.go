package utils

import (
	"github.com/ethereum/eth-go/ethchain"
	"github.com/ethereum/eth-go/ethstate"
	"math/big"
)

type VMEnv struct {
	state *ethstate.State
	block *ethchain.Block

	transactor []byte
	value      *big.Int
}

func NewEnv(state *ethstate.State, block *ethchain.Block, transactor []byte, value *big.Int) *VMEnv {
	return &VMEnv{
		state:      state,
		block:      block,
		transactor: transactor,
		value:      value,
	}
}

func (self *VMEnv) Origin() []byte         { return self.transactor }
func (self *VMEnv) BlockNumber() *big.Int  { return self.block.Number }
func (self *VMEnv) PrevHash() []byte       { return self.block.PrevHash }
func (self *VMEnv) Coinbase() []byte       { return self.block.Coinbase }
func (self *VMEnv) Time() int64            { return self.block.Time }
func (self *VMEnv) Difficulty() *big.Int   { return self.block.Difficulty }
func (self *VMEnv) Value() *big.Int        { return self.value }
func (self *VMEnv) State() *ethstate.State { return self.state }
