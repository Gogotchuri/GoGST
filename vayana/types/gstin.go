package vayanaTypes

type GSTINDetailsResponse struct {
	Gstin     string      `json:"Gstin"`
	TradeName interface{} `json:"TradeName"`
	LegalName string      `json:"LegalName"`
	StateCode string      `json:"StateCode"`
	AddrBnm   interface{} `json:"AddrBnm"`
	AddrBno   interface{} `json:"AddrBno"`
	AddrFlno  interface{} `json:"AddrFlno"`
	AddrSt    interface{} `json:"AddrSt"`
	AddrLoc   interface{} `json:"AddrLoc"`
	TxpType   interface{} `json:"TxpType"`
	BlkStatus interface{} `json:"BlkStatus"`
	Status    string      `json:"Status"`
	AddrPncd  int         `json:"AddrPncd"`
	DtReg     interface{} `json:"DtReg"`
	DtDReg    interface{} `json:"DtDReg"`
}
