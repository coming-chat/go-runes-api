package btc

type RuneTransferTransaction struct {
	CommitId    string `json:"commit_id"`
	CommitPsbt  string `json:"commit_psbt"`
	NetworkFee  int64  `json:"network_fee"`
	SatpointFee int64  `json:"satpoint_fee"`
	ServiceFee  int64  `json:"service_fee"`
	CommitFee   int64  `json:"commit_fee"`
	CommitVsize string `json:"commit_vsize"`
	// CommitCustom *types.Brc20CommitCustom `json:"commit_custom"`
}