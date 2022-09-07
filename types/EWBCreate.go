package types

const CreatePath = "/generate/"

type EWBItem struct {
	ProductName   string  `json:"productName"`
	ProductDesc   string  `json:"productDesc"`
	HsnCode       int     `json:"hsnCode"`
	Quantity      float64 `json:"quantity"`
	QtyUnit       string  `json:"qtyUnit"`
	CgstRate      float64 `json:"cgstRate"`
	SgstRate      float64 `json:"sgstRate"`
	IgstRate      float64 `json:"igstRate"`
	TaxableAmount float64 `json:"taxableAmount"`
	//Ignore following 2 fields
	CessRate     float64 `json:"cessRate"`     //Not needed for GST
	CessNonadvol float64 `json:"cessNonadvol"` //Not needed for GST
}

type EWBBase struct {
	// We can use default values for these fields
	SupplyType    string `json:"supplyType"`
	SubSupplyType string `json:"subSupplyType"`
	SubSupplyDesc string `json:"subSupplyDesc"`
	// Doc Fields
	DocType string `json:"docType"`
	DocDate string `json:"docDate"`
	DocNo   string `json:"docNo"`
	// Company Info
	FromGstin             string `json:"fromGstin"`
	FromTrdName           string `json:"fromTrdName"`
	FromAddr1             string `json:"fromAddr1"`
	FromAddr2             string `json:"fromAddr2"`
	FromPlace             string `json:"fromPlace"`
	FromPincode           int    `json:"fromPincode"`
	SupplierFromStateCode int    `json:"actFromStateCode"`
	FromStateCode         int    `json:"fromStateCode"`
	//Customer Info
	ToGstin        string `json:"toGstin"`
	ToTrdName      string `json:"toTrdName"`
	ToAddr1        string `json:"toAddr1"`
	ToAddr2        string `json:"toAddr2"`
	ToPlace        string `json:"toPlace"`
	ToPincode      int    `json:"toPincode"`
	ActToStateCode int    `json:"actToStateCode"`
	ToStateCode    int    `json:"toStateCode"`
	//autofilled
	TransactionType int `json:"transactionType"`
	//amounts
	// Amounts
	TotalValue        float64 `json:"totalValue"`
	OtherValue        float64 `json:"otherValue"`
	CgstValue         float64 `json:"cgstValue"`
	SgstValue         float64 `json:"sgstValue"`
	IgstValue         float64 `json:"igstValue"`
	TotalInvoiceValue float64 `json:"totInvValue"`

	//Ignore following 2 fields
	CessValue         float64 `json:"cessValue"`
	CessNonAdvolValue float64 `json:"cessNonAdvolValue"`

	//Transport Details
	TransporterId   string `json:"transporterId"`
	TransporterName string `json:"transporterName"`
	VehicleType     string `json:"vehicleType"`
	//Items
	ItemList []EWBItem `json:"itemList"`
}

type EWBCreateRequest struct {
	EWBBase
	TransDocNo    string `json:"transDocNo"`
	TransMode     string `json:"transMode"`
	TransDistance string `json:"transDistance"`
	TransDocDate  string `json:"transDocDate"`
	VehicleNo     string `json:"vehicleNo"`
}

type EWBCreateResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Alert   string `json:"alert"`
	Info    string `json:"info"` //Set when status is 0

	Uuid         string `json:"uuid"`
	ValidUpto    string `json:"validUpto"`
	EwayBillNo   string `json:"ewayBillNo"`
	EwayBillDate string `json:"ewayBillDate"`
}

type EWBGetResponse struct {
	EWBBase

	EwayBillNo         string           `json:"ewbNo"`
	EwayBillDate       string           `json:"ewayBillDate"`
	UserGstin          string           `json:"userGstin"`
	NoValidDays        int              `json:"noValidDays"`
	Distance           int              `json:"actualDist"`
	RejectStatus       string           `json:"rejectStatus"`
	Status             string           `json:"status"`
	ValidUpto          string           `json:"validUpto"`
	VehicleListDetails []VehicleDetails `json:"VehiclListDetails"`
}

type VehicleDetails struct {
	UpdMode          string `json:"updMode"`
	VehicleNo        string `json:"vehicleNo"`
	FromPlace        string `json:"fromPlace"`
	FromState        int    `json:"fromState"`
	TripshtNo        int    `json:"tripshtNo"`
	UserGSTINTransin string `json:"userGSTINTransin"`
	EnteredDate      string `json:"enteredDate"`
	TransMode        string `json:"transMode"`
	TransDocNo       string `json:"transDocNo"`
	VehicleType      string `json:"vehicleType"`
	TransDocDate     string `json:"transDocDate"`
	GroupNo          string `json:"groupNo"`
}
