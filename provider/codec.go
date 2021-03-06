package provider

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/bandprotocol/bandchain/chain/x/zoracle"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewCodec() *codec.Codec {
	var cdc = codec.New()
	sdk.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	zoracle.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	staking.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	params.RegisterCodec(cdc)
	return cdc
}
