package mapper

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

const (
	FlareChainID = 14
	FlareAssetID = "foMCFvzKECiGVJmmkAEHm9Vt43hYjuxreiNX5PfqfecaVsZBT"

	SongbirdChainID = 19
	SongbirdAssetID = "1S3PSi4VsVpD8iK2vdykuajxVeuCV2xhjPSkQ4K88mqWGozMP"

	OpCall         = "CALL"
	OpFee          = "FEE"
	OpCreate       = "CREATE"
	OpCreate2      = "CREATE2"
	OpSelfDestruct = "SELFDESTRUCT"
	OpCallCode     = "CALLCODE"
	OpDelegateCall = "DELEGATECALL"
	OpStaticCall   = "STATICCALL"
	OpDestruct     = "DESTRUCT"
	OpImport       = "IMPORT"
	OpExport       = "EXPORT"

	StatusSuccess = "SUCCESS"
	StatusFailure = "FAILURE"
)

var (
	StageBootstrap = &types.SyncStatus{
		Synced: types.Bool(false),
		Stage:  types.String("BOOTSTRAP"),
	}

	StageSynced = &types.SyncStatus{
		Synced: types.Bool(true),
		Stage:  types.String("SYNCED"),
	}

	FlareCurrency = &types.Currency{
		Symbol:   "FLR",
		Decimals: 18, //nolint:gomnd
	}

	SongbirdCurrency = &types.Currency{
		Symbol:   "SGB",
		Decimals: 18, //nolint:gomnd
	}

	ActiveCurrency *types.Currency

	OperationStatuses = []*types.OperationStatus{
		{
			Status:     StatusSuccess,
			Successful: true,
		},
		{
			Status:     StatusFailure,
			Successful: false,
		},
	}

	OperationTypes = []string{
		OpFee,
		OpCall,
		OpCreate,
		OpCreate2,
		OpSelfDestruct,
		OpCallCode,
		OpDelegateCall,
		OpStaticCall,
		OpDestruct,
		OpImport,
		OpExport,
	}

	CallMethods = []string{
		"eth_getTransactionReceipt",
	}
)

func CallType(t string) bool {
	callTypes := []string{
		OpCall,
		OpCallCode,
		OpDelegateCall,
		OpStaticCall,
	}

	for _, callType := range callTypes {
		if callType == t {
			return true
		}
	}

	return false
}

func CreateType(t string) bool {
	createTypes := []string{
		OpCreate,
		OpCreate2,
	}

	for _, createType := range createTypes {
		if createType == t {
			return true
		}
	}

	return false
}
