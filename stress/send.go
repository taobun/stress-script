package main

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/tendermint/tendermint/crypto/secp256k1"

	"github.com/taobun/stress-script/provider"
)

func sendCoin(from secp256k1.PrivKeySecp256k1, to []sdk.AccAddress, val int64, gas uint64) (sdk.TxResponse, error) {
	p, err := provider.NewBandProvider(nodeURI, from, chainID)
	if err != nil {
		panic(err)
	}
	msgs := make([]sdk.Msg, 0)
	for _, t := range to {
		msgs = append(msgs, bank.MsgSend{
			FromAddress: p.Sender(),
			ToAddress:   t,
			Amount:      sdk.NewCoins(sdk.NewCoin("uband", sdk.NewInt(val))),
		})
	}
	return p.SendTransaction(msgs, 0, gas, "", "", flags.BroadcastBlock)
}
