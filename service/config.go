package service

import (
	"math/big"

	"github.com/coinbase/rosetta-sdk-go/types"
	ethtypes "github.com/flare-foundation/coreth/core/types"
)

// Config holds the service configuration
type Config struct {
	Mode             string
	ChainID          *big.Int
	NetworkID        *types.NetworkIdentifier
	GenesisBlockHash string
	AvaxAssetID      string
}

const (
	ModeOffline = "offline"
	ModeOnline  = "online"
)

// IsOfflineMode returns true if running in offline mode
func (c Config) IsOfflineMode() bool {
	return c.Mode == ModeOffline
}

// IsOnlineMode returns true if running in online mode
func (c Config) IsOnlineMode() bool {
	return c.Mode == ModeOnline
}

// Signer returns an eth signer object for a given chain
func (c Config) Signer() ethtypes.Signer {
	return ethtypes.LatestSignerForChainID(c.ChainID)
}
