package mapper

import (
	"github.com/flare-foundation/coreth/plugin/evm"
	"github.com/flare-foundation/flare/codec"
	"github.com/flare-foundation/flare/codec/linearcodec"
	"github.com/flare-foundation/flare/utils/wrappers"
	"github.com/flare-foundation/flare/vms/secp256k1fx"
)

const (
	preApricotCodecVersion uint16 = 0
	codecRegistrationSkip  int    = 3
)

var (
	codecManager codec.Manager
)

func init() {
	codecManager = codec.NewDefaultManager()

	errs := wrappers.Errs{}
	preApricotCodec := initPreApricotCodec(&errs)
	errs.Add(
		codecManager.RegisterCodec(preApricotCodecVersion, preApricotCodec),
	)

	if errs.Errored() {
		panic(errs.Err)
	}
}

func initPreApricotCodec(errs *wrappers.Errs) linearcodec.Codec {
	c := linearcodec.NewDefault()

	errs.Add(
		c.RegisterType(&evm.UnsignedImportTx{}),
		c.RegisterType(&evm.UnsignedExportTx{}),
	)

	c.SkipRegistrations(codecRegistrationSkip)

	errs.Add(
		c.RegisterType(&secp256k1fx.TransferInput{}),
		c.RegisterType(&secp256k1fx.MintOutput{}),
		c.RegisterType(&secp256k1fx.TransferOutput{}),
		c.RegisterType(&secp256k1fx.MintOperation{}),
		c.RegisterType(&secp256k1fx.Credential{}),
		c.RegisterType(&secp256k1fx.Input{}),
		c.RegisterType(&secp256k1fx.OutputOwners{}),
	)

	return c
}
