package types

type EWBItem struct {
	ProductName   string  `json:"productName"`
	ProductDesc   string  `json:"productDesc"`
	HsnCode       int     `json:"hsnCode"`
	Quantity      float64 `json:"quantity"`
	QtyUnit       string  `json:"qtyUnit"`
	CgstRate      float64 `json:"cgstRate"`
	SgstRate      float64 `json:"sgstRate"`
	IgstRate      float64 `json:"igstRate"`
	CessRate      float64 `json:"cessRate"`
	CessNonadvol  float64 `json:"cessNonadvol"`
	TaxableAmount float64 `json:"taxableAmount"`
}

type EWBCreateRequest struct {
	SupplyType        string    `json:"supplyType"`
	SubSupplyType     string    `json:"subSupplyType"`
	SubSupplyDesc     string    `json:"subSupplyDesc"`
	DocType           string    `json:"docType"`
	DocNo             string    `json:"docNo"`
	DocDate           string    `json:"docDate"`
	FromGstin         string    `json:"fromGstin"`
	FromTrdName       string    `json:"fromTrdName"`
	FromAddr1         string    `json:"fromAddr1"`
	FromAddr2         string    `json:"fromAddr2"`
	FromPlace         string    `json:"fromPlace"`
	FromPincode       int       `json:"fromPincode"`
	ActFromStateCode  int       `json:"actFromStateCode"`
	FromStateCode     int       `json:"fromStateCode"`
	ToGstin           string    `json:"toGstin"`
	ToTrdName         string    `json:"toTrdName"`
	ToAddr1           string    `json:"toAddr1"`
	ToAddr2           string    `json:"toAddr2"`
	ToPlace           string    `json:"toPlace"`
	ToPincode         int       `json:"toPincode"`
	ActToStateCode    int       `json:"actToStateCode"`
	ToStateCode       int       `json:"toStateCode"`
	TransactionType   int       `json:"transactionType"`
	OtherValue        float64   `json:"otherValue"`
	TotalValue        float64   `json:"totalValue"`
	CgstValue         float64   `json:"cgstValue"`
	SgstValue         float64   `json:"sgstValue"`
	IgstValue         float64   `json:"igstValue"`
	CessValue         float64   `json:"cessValue"`
	CessNonAdvolValue float64   `json:"cessNonAdvolValue"`
	TotalInvoiceValue float64   `json:"totInvValue"`
	TransporterId     string    `json:"transporterId"`
	TransporterName   string    `json:"transporterName"`
	TransDocNo        string    `json:"transDocNo"`
	TransMode         string    `json:"transMode"`
	TransDistance     int       `json:"transDistance"`
	TransDocDate      string    `json:"transDocDate"`
	VehicleNo         string    `json:"vehicleNo"`
	VehicleType       string    `json:"vehicleType"`
	ItemList          []EWBItem `json:"itemList"`
}
