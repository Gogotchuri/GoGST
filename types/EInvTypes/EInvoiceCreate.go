package EInvTypes

//import (
//	"github.com/gogotchuri/go-validator"
//	"github.com/mcuadros/go-defaults"
//)

type Transporter struct {
	TaxScheme string `json:"TaxSch" default:"GST" validate:"oneof=GST,required"` //TODO: validation default
	//Type of Supply: B2B-Business to Business, SEZWP - SEZ with payment, SEZWOP - SEZ without payment, EXPWP - Export with Payment, EXPWOP - Export without payment,DEXP - Deemed Export
	SupplyType string `json:"SupTyp" validate:"default=B2B,oneof='B2B' 'EZWP' 'EZWOP' 'XPWP' 'XPWOP' 'EXP',required"`
	//whether the tax liability is payable under reverse charge
	RegRev         string `json:"RegRev" validate:"default=N,oneof='Y' 'N'"`
	ECommerceGstin string `json:"EcmGstin"`
	//indicates the supply is intra state but chargeable to IGST
	IgstOnIntra string `json:"IgstOnIntra" validate:"default=N,oneof='Y' 'N'"`
}

type DocumentDetails struct {
	Type       string `json:"Typ" validate:"default=INV,oneof='INV' 'CRN' 'DBN',required"`
	DocumentNo string `json:"No"  validate:"min=1,max=16,alphanumeric,required,startsnotwith=0,startsnotwith=/,startsnotwith=-"`
	Date       string `json:"Dt"  validate:"date_format=02/01/2006,required"`
}

type Address struct {
	Address1  string `json:"Addr1" validate:"min=1,max=100,required"`
	Address2  string `json:"Addr2" validate:"min=3,max=100"`
	Location  string `json:"Loc"   validate:"min=3,max=50,required"`
	Pin       uint   `json:"Pin"   validate:"min=100000,max=999999,required"`
	StateCode string `json:"Stcd"  validate:"min=1,max=2,numeric,required"`
}

type Company struct {
	GSTIN     string `json:"Gstin" validate:"min=15,max=15,alphanumeric,india_gstin,required"`
	LegalName string `json:"LglNm" validate:"min=1,max=100,required"`
	TradeName string `json:"TrdNm" validate:"min=1,max=100"`
	Address
	Phone string `json:"Ph"    validate:"min=6,max=12,numeric"`
	Email string `json:"Em"    validate:"min=6,max=100,email"`
}

type Seller struct {
	Company
}

type Buyer struct {
	Company
	// State code of Place of supply. If POS lies outside the country, the code shall be 96.
	// Intra-state tax, POS shall be different from the seller’s state code.
	// Inter-state tax (CGST/SGST), POS should be the same as the seller’s state code.
	PlaceOfSupply string `json:"Pos" validate:"min=1,max=2,numeric,required"`
}

type DispatchDetails struct {
	Name string `json:"Nm"    validate:"min=3,max=100,required"`
	Address
}

type ShipToDetails struct {
	Gstin     string `json:"Gstin" validate:"min=15,max=15,alphanumeric,india_gstin"`
	LegalName string `json:"LglNm" validate:"min=3,max=100,required"`
	TradeName string `json:"TrdNm" validate:"min=3,max=100"`
	Address
}

type ItemBase struct {
	SerialNo           string  `json:"SlNo" validate:"min=1,max=6,numeric,required"`
	ProductDescription string  `json:"PrdDesc" validate:"min=3,max=300"`
	IsService          string  `json:"IsServc" validate:"oneof='Y' 'N',required"`
	HSNCode            string  `json:"HsnCd" validate:"min=4,max=8,numeric,required"`
	Barcode            string  `json:"Barcde" validate:"min=3,max=30"`
	Unit               string  `json:"Unit" validate:"min=3,max=8,alpha"` //TODO: unit validation
	UnitPrice          float64 `json:"UnitPrice" validate:"min=0,max=9999999999,required"`
	Quantity           float64 `json:"Qty" validate:"min=0,max=9999999999"`
	NetAmount          float64 `json:"TotAmt" validate:"min=0,max=9999999999,required"`
	Discount           float64 `json:"Discount" validate:"min=0,max=9999999999"`
	TaxableAmount      float64 `json:"AssAmt" validate:"min=0,max=9999999999,required"`
	IGSTRate           float64 `json:"GstRt" validate:"min=0,max=100,required"`
	IGSTAmount         float64 `json:"IgstAmt" validate:"min=0,max=9999999999"`
	CGSTAmount         float64 `json:"CgstAmt" validate:"min=0,max=9999999999"`
	SGSTAmount         float64 `json:"SgstAmt" validate:"min=0,max=9999999999"`
	AdditionalAmount   float64 `json:"OthChrg" validate:"min=0,max=9999999999"`
	TotalAmount        float64 `json:"TotItemVal" validate:"min=0,max=9999999999,required"`
}

type Item struct {
	ItemBase
	FreeQuantity           float64           `json:"FreeQty" validate:"min=0,max=9999999999"`
	PreTaxAmount           float64           `json:"PreTaxVal" validate:"min=0,max=9999999999"`
	CESRate                float64           `json:"CesRt" validate:"min=0,max=100"`
	CESAmount              float64           `json:"CesAmt" validate:"min=0,max=9999999999"`
	CESNonAdvalAmount      float64           `json:"CesNonAdvlAmt" validate:"min=0,max=9999999999"`
	StateCesRate           float64           `json:"StateCesRt" validate:"min=0,max=100"`
	StateCesAmount         float64           `json:"StateCesAmt" validate:"min=0,max=9999999999"`
	StateCesNonAdvalAmount float64           `json:"StateCesNonAdvlAmt" validate:"min=0,max=9999999999"`
	OrdLineReference       string            `json:"OrdLineRef" validate:"min=1,max=50"`
	CountryOfOrigin        string            `json:"OrgCntry" validate:"min=2,max=2,alpha"`
	ProductSerialNo        string            `json:"PrdSlNo" validate:"min=1,max=20"`
	BatchDetails           *ItemBatchDetails `json:"BchDtls"`
	ItemAttributes         []ItemAttribute   `json:"AttribDtls"`
}

type ItemBatchDetails struct {
	Name           string `json:"Nm"    validate:"min=3,max=20,required"`
	ExpirationDate string `json:"ExpDt" validate:"date_format=02/01/2006"`
	WarrantyDate   string `json:"WrDt"  validate:"date_format=02/01/2006"`
}

type ItemAttribute struct {
	Name  string `json:"Nm"  validate:"min=1,max=100"`
	Value string `json:"Val" validate:"min=1,max=100"`
}

type DocumentValues struct {
	TaxableValue     float64 `json:"AssVal" validate:"min=0,max=9999999999,required"`
	CGSTValue        float64 `json:"CgstVal" validate:"min=0,max=9999999999"`
	SGSTValue        float64 `json:"SgstVal" validate:"min=0,max=9999999999"`
	IGSTValue        float64 `json:"IgstVal" validate:"min=0,max=9999999999"`
	Discount         float64 `json:"Discount" validate:"min=0,max=9999999999"`
	AdditionalAmount float64 `json:"OthChrg" validate:"min=0,max=9999999999"`
	TotalAmount      float64 `json:"TotInvVal" validate:"min=0,max=9999999999,required"`

	RoundedOffAmount           float64 `json:"RndOffAmt" validate:"min=-99,max=99"`
	CESSValue                  float64 `json:"CesVal" validate:"min=0,max=9999999999"`
	StateCESSValue             float64 `json:"StCesVal" validate:"min=0,max=9999999999"`
	TotalAmountForeignCurrency float64 `json:"TotInvValFc" validate:"min=0,max=9999999999"`
}

type PaymentDetails struct {
	Name     string `json:"Nm"`
	AccDet   string `json:"AccDet"`
	Mode     string `json:"Mode"`
	FinInsBr string `json:"FinInsBr"`
	PayTerm  string `json:"PayTerm"`
	PayInstr string `json:"PayInstr"`
	CrTrn    string `json:"CrTrn"`
	DirDr    string `json:"DirDr"`
	CrDay    int    `json:"CrDay"`
	PaidAmt  int    `json:"PaidAmt"`
	PaymtDue int    `json:"PaymtDue"`
}

type Period struct {
	StartDate string `json:"InvStDt" validate:"date_format=02/01/2006,required"`
	EndDate   string `json:"InvEndDt" validate:"date_format=02/01/2006,required"`
}

type PrecedingDocumentDetails struct {
	InvoiceNo        string `json:"InvNo" validate:"min=1,max=16,alphanumeric,required"`
	InvoiceDate      string `json:"InvDt" validate:"date_format=02/01/2006,required"`
	OtherReferenceNo string `json:"OthRefNo" validate:"min=1,max=20"`
}

type ContractDetails struct {
	ReceiptAdvRefr string `json:"RecAdvRefr" validate:"min=1,max=20"`
	ReceiptAdvDt   string `json:"RecAdvDt" validate:"date_format=02/01/2006"`
	TendRefr       string `json:"TendRefr" validate:"min=1,max=20"`
	ContrRefr      string `json:"ContrRefr" validate:"min=1,max=20"`
	ExtRefr        string `json:"ExtRefr" validate:"min=1,max=20"`
	ProjRefr       string `json:"ProjRefr" validate:"min=1,max=20"`
	PORefr         string `json:"PORefr" validate:"min=1,max=16"`
	PORefDt        string `json:"PORefDt" validate:"date_format=02/01/2006"`
}

type ReferenceDetails struct {
	Remark                   string                     `json:"InvRm" validate:"min=3,max=100"`
	DocumentPeriod           *Period                    `json:"DocPerdDtls"`
	PrecedingDocumentDetails []PrecedingDocumentDetails `json:"PrecDocDtls"`
	ContractDetails          []ContractDetails          `json:"ContrDtls"`
}

type AdditionalDocumentDetails struct {
	Url            string `json:"Url" validate:"min=3,max=100,url"`
	DocB64         string `json:"Docs" validate:"min=3,max=100,base64"`
	AdditionalInfo string `json:"Info" validate:"min=3,max=1000"`
}

type ExportDetails struct {
	ShipBNo string  `json:"ShipBNo"`
	ShipBDt string  `json:"ShipBDt"`
	Port    string  `json:"Port"`
	RefClm  string  `json:"RefClm"`
	ForCur  string  `json:"ForCur"`
	CntCode string  `json:"CntCode"`
	ExpDuty float64 `json:"ExpDuty"`
}

type EWBDetails struct {
	TransId    string `json:"TransId" validate:"india_transin"`
	TransName  string `json:"TransName" validate:"min=3,max=100"`
	Distance   int    `json:"Distance" validate:"min=0,max=9999999999"`
	TransMode  string `json:"TransMode" validate:"min=1,max=1,oneof=1 2 3 4"`
	TransDocNo string `json:"TransDocNo" validate:"min=1,max=15,alphanumeric"`
	TransDocDt string `json:"TransDocDt" validate:"date_format=02/01/2006"`
	VehNo      string `json:"VehNo" validate:"min=4,max=20,alphanumeric"`
	VehType    string `json:"VehType" validate:"min=1,max=1,oneof=R O"`
}

type EInvoiceCreate struct {
	Version            string          `json:"Version" validate:"default=1.1,numeric,required,min=1,max=6"`
	ReferenceNumber    string          `json:"Irn" validate:"isdefault"`
	TransporterDetails Transporter     `json:"TranDtls" validate:"required"`
	DocumentDetails    DocumentDetails `json:"DocDtls" validate:"required"`
	SellerDetails      Seller          `json:"SellerDtls" validate:"required"`
	BuyerDetails       Buyer           `json:"BuyerDtls" validate:"required"`
	//if transaction type is Bill From - Dispatch From
	DispatchDtls *DispatchDetails `json:"DispDtls"`
	//if Transaction Type is Bill To - Ship To | if "Combination of Both" both are provided
	ShipToDetails             *ShipToDetails              `json:"ShipDtls"`
	ItemList                  []Item                      `json:"ItemList" validate:"min=1,max=1000,required,unique=SerialNo"`
	DocumentValues            DocumentValues              `json:"ValDtls" validate:"required"`
	PaymentDetails            *PaymentDetails             `json:"PayDtls"`
	ReferenceDetails          *ReferenceDetails           `json:"RefDtls"`
	AdditionalDocumentDetails []AdditionalDocumentDetails `json:"AddlDocDtls"`
	ExportDetails             *ExportDetails              `json:"ExpDtls"`
	EWBDetails                *EWBDetails                 `json:"EwbDtls"`
}

func (e *EInvoiceCreate) Validate() error {
	//TODO set default values
	//TODO validate
	return nil
}
