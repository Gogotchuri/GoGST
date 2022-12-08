package EInvTypes

import (
	"fmt"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/gogotchuri/go-validator"
	"github.com/gogotchuri/go-validator/translations/en"
	"github.com/mcuadros/go-defaults"
	"strings"
)

type TransportationDetails struct {
	TaxScheme string `json:"TaxSch" default:"GST" validate:"oneof=GST"`
	//Type of Supply: B2B-Business to Business, SEZWP - SEZ with payment, SEZWOP - SEZ without payment, EXPWP - Export with Payment, EXPWOP - Export without payment,DEXP - Deemed Export
	SupplyType string `json:"SupTyp" default:"B2B" validate:"oneof='B2B' 'EZWP' 'EZWOP' 'XPWP' 'XPWOP' 'EXP'"`
	//whether the tax liability is payable under reverse charge
	RegRev         string `json:"RegRev" default:"N" validate:"oneof='Y' 'N'"`
	ECommerceGstin string `json:"EcmGstin,omitempty"`
	//indicates the supply is intra state but chargeable to IGST
	IgstOnIntra string `json:"IgstOnIntra" default:"N" validate:"oneof='Y' 'N'"`
}

type DocumentDetails struct {
	Type       string `json:"Typ" default:"INV" validate:"oneof='INV' 'CRN' 'DBN'"`
	DocumentNo string `json:"No"  validate:"alphanum,min=1,max=16,startsnotwith=0,startsnotwith=/,startsnotwith=-"`
	Date       string `json:"Dt"  validate:"date_format=02/01/2006"`
}

type Address struct {
	Address1  string `json:"Addr1" validate:"min=1,max=100"`
	Address2  string `json:"Addr2,omitempty" validate:"omitempty,min=3,max=100"`
	Location  string `json:"Loc"   validate:"min=3,max=50"`
	Pin       uint   `json:"Pin"   validate:"min=100000,max=999999"`
	StateCode string `json:"Stcd"  validate:"min=1,max=2,numeric"`
}

type Company struct {
	GSTIN     string `json:"Gstin" validate:"india_transin"` //TODO: validate gstin after production is ready
	LegalName string `json:"LglNm" validate:"min=1,max=100"`
	TradeName string `json:"TrdNm,omitempty" validate:"omitempty,min=1,max=100"`
	Address
	Phone string `json:"Ph,omitempty"    validate:"omitempty,min=6,max=12,numeric"`
	Email string `json:"Em,omitempty"    validate:"omitempty,min=6,max=100,email"`
}

type Seller struct {
	Company
}

type Buyer struct {
	Company
	// State code of Place of supply. If POS lies outside the country, the code shall be 96.
	// Intra-state tax, POS shall be different from the seller’s state code.
	// Inter-state tax (CGST/SGST), POS should be the same as the seller’s state code.
	PlaceOfSupply string `json:"Pos" validate:"min=1,max=2,numeric"`
}

type DispatchDetails struct {
	Name string `json:"Nm"    validate:"min=3,max=100"`
	Address
}

type ShipToDetails struct {
	Gstin     string `json:"Gstin" validate:"omitempty,india_gstin"`
	LegalName string `json:"LglNm" validate:"min=3,max=100"`
	TradeName string `json:"TrdNm" validate:"omitempty,min=3,max=100"`
	Address
}

type ItemBase struct {
	SerialNo           string  `json:"SlNo" validate:"min=1,max=6,numeric"`
	ProductDescription string  `json:"PrdDesc" validate:"omitempty,min=3,max=300"`
	IsService          string  `json:"IsServc" default:"N" validate:"required,oneof='Y' 'N'"`
	HSNCode            string  `json:"HsnCd" validate:"min=6,max=8,numeric"`
	Barcode            string  `json:"Barcde,omitempty" validate:"omitempty,min=3,max=30"`
	Unit               string  `json:"Unit" validate:"omitempty,min=3,max=8,alpha"` //TODO: unit validation
	UnitPrice          float64 `json:"UnitPrice" validate:"min=0,max=9999999999"`
	Quantity           float64 `json:"Qty" validate:"min=0,max=9999999999"`
	NetAmount          float64 `json:"TotAmt" validate:"min=0,max=9999999999"`
	Discount           float64 `json:"Discount" validate:"min=0,max=9999999999"`
	TaxableAmount      float64 `json:"AssAmt" validate:"min=0,max=9999999999"`
	IGSTRate           float64 `json:"GstRt" validate:"min=0,max=100"`
	IGSTAmount         float64 `json:"IgstAmt" validate:"min=0,max=9999999999"`
	CGSTAmount         float64 `json:"CgstAmt" validate:"min=0,max=9999999999"`
	SGSTAmount         float64 `json:"SgstAmt" validate:"min=0,max=9999999999"`
	AdditionalAmount   float64 `json:"OthChrg" validate:"min=0,max=9999999999"`
	TotalAmount        float64 `json:"TotItemVal" validate:"min=0,max=9999999999"`
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
	OrdLineReference       string            `json:"OrdLineRef,omitempty" validate:"omitempty,min=1,max=50"`
	CountryOfOrigin        string            `json:"OrgCntry,omitempty" default:"IN" validate:"required,min=2,max=2,alpha"`
	ProductSerialNo        string            `json:"PrdSlNo,omitempty" validate:"omitempty,min=1,max=20"`
	BatchDetails           *ItemBatchDetails `json:"BchDtls"`
	ItemAttributes         []ItemAttribute   `json:"AttribDtls"`
}

type ItemBatchDetails struct {
	Name           string `json:"Nm"    validate:"min=3,max=20"`
	ExpirationDate string `json:"ExpDt" validate:"omitempty,date_format=02/01/2006"`
	WarrantyDate   string `json:"WrDt"  validate:"omitempty,date_format=02/01/2006"`
}

type ItemAttribute struct {
	Name  string `json:"Nm"  validate:"min=1,max=100"`
	Value string `json:"Val" validate:"min=1,max=100"`
}

type DocumentValues struct {
	TaxableValue     float64 `json:"AssVal" validate:"min=0,max=9999999999"`
	CGSTValue        float64 `json:"CgstVal" validate:"min=0,max=9999999999"`
	SGSTValue        float64 `json:"SgstVal" validate:"min=0,max=9999999999"`
	IGSTValue        float64 `json:"IgstVal" validate:"min=0,max=9999999999"`
	Discount         float64 `json:"Discount" validate:"min=0,max=9999999999"`
	AdditionalAmount float64 `json:"OthChrg" validate:"min=0,max=9999999999"`
	TotalAmount      float64 `json:"TotInvVal" validate:"min=0,max=9999999999"`

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
	StartDate string `json:"InvStDt" validate:"date_format=02/01/2006"`
	EndDate   string `json:"InvEndDt" validate:"date_format=02/01/2006"`
}

type PrecedingDocumentDetails struct {
	InvoiceNo        string `json:"InvNo" validate:"min=1,max=16,alphanum"`
	InvoiceDate      string `json:"InvDt" validate:"date_format=02/01/2006"`
	OtherReferenceNo string `json:"OthRefNo" validate:"omitempty,min=1,max=20"`
}

type ContractDetails struct {
	ReceiptAdvRefr string `json:"RecAdvRefr" validate:"omitempty,min=1,max=20"`
	ReceiptAdvDt   string `json:"RecAdvDt" validate:"omitempty,date_format=02/01/2006"`
	TendRefr       string `json:"TendRefr" validate:"omitempty,min=1,max=20"`
	ContrRefr      string `json:"ContrRefr" validate:"omitempty,min=1,max=20"`
	ExtRefr        string `json:"ExtRefr" validate:"omitempty,min=1,max=20"`
	ProjRefr       string `json:"ProjRefr" validate:"omitempty,min=1,max=20"`
	PORefr         string `json:"PORefr" validate:"omitempty,min=1,max=16"`
	PORefDt        string `json:"PORefDt" validate:"omitempty,date_format=02/01/2006"`
}

type ReferenceDetails struct {
	Remark                   string                     `json:"InvRm" validate:"omitempty,min=3,max=100"`
	DocumentPeriod           *Period                    `json:"DocPerdDtls"`
	PrecedingDocumentDetails []PrecedingDocumentDetails `json:"PrecDocDtls"`
	ContractDetails          []ContractDetails          `json:"ContrDtls"`
}

type AdditionalDocumentDetails struct {
	Url            string `json:"Url" validate:"omitempty,min=3,max=100,url"`
	DocB64         string `json:"Docs" validate:"omitempty,min=3,max=100,base64"`
	AdditionalInfo string `json:"Info" validate:"omitempty,min=3,max=1000"`
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
	TransId    string `json:"TransId" validate:"omitempty,india_transin"`
	TransName  string `json:"TransName" validate:"omitempty,min=3,max=100"`
	Distance   int    `json:"Distance" validate:"omitempty,min=0,max=9999999999"`
	TransMode  string `json:"TransMode" validate:"omitempty,min=1,max=1,oneof=1 2 3 4"`
	TransDocNo string `json:"TransDocNo" validate:"omitempty,min=1,max=15,alphanum"`
	TransDocDt string `json:"TransDocDt" validate:"omitempty,date_format=02/01/2006"`
	VehNo      string `json:"VehNo" validate:"omitempty,min=4,max=20,alphanum"`
	VehType    string `json:"VehType" validate:"omitempty,min=1,max=1,oneof=R O"`
}

type EInvoiceCreate struct {
	Version               string                `json:"Version" default:"1.1" validate:"numeric,min=1,max=6"`
	ReferenceNumber       string                `json:"Irn" default:"" validate:"isdefault"`
	TransportationDetails TransportationDetails `json:"TranDtls" validate:"required"`
	DocumentDetails       DocumentDetails       `json:"DocDtls" validate:"required"`
	SellerDetails         Seller                `json:"SellerDtls" validate:"required"`
	BuyerDetails          Buyer                 `json:"BuyerDtls" validate:"required"`
	//if transaction type is Bill From - Dispatch From
	DispatchDtls *DispatchDetails `json:"DispDtls"`
	//if Transaction Type is Bill To - Ship To | if "Combination of Both" both are provided
	ShipToDetails             *ShipToDetails              `json:"omitempty,ShipDtls"`
	ItemList                  []Item                      `json:"ItemList" validate:"dive,min=1,max=1000,unique=SerialNo"`
	DocumentValues            DocumentValues              `json:"ValDtls" validate:"required"`
	PaymentDetails            *PaymentDetails             `json:"omitempty,PayDtls"`
	ReferenceDetails          *ReferenceDetails           `json:"omitempty,RefDtls"`
	AdditionalDocumentDetails []AdditionalDocumentDetails `json:"omitempty,AddlDocDtls"`
	ExportDetails             *ExportDetails              `json:"omitempty,ExpDtls"`
	EWBDetails                *EWBDetails                 `json:"omitempty,EwbDtls"`
}

func (e *EInvoiceCreate) Validate(validate *validator.Validate) ValidationErrors {
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")
	err := en.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return FromError(fmt.Errorf("error while registering default translations: %v", err))
	}
	//set default values
	defaults.SetDefaults(e)
	//validate
	err = validate.Struct(e)
	if err == nil {
		return nil
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return FromError(fmt.Errorf("error while convering validation erros"))
	}

	if len(errs) == 0 {
		return nil
	}
	vErrs := make([]error, len(errs))
	for i, e := range errs {
		vErrs[i] = fmt.Errorf("%s : %s", e.Namespace()[15:], e.Translate(trans))
	}
	return FromErrors(vErrs)
}

type ValidationErrors []error

func (v ValidationErrors) Errors() []error {
	return v
}

func (v ValidationErrors) Error() string {
	strs := make([]string, len(v))
	for i, err := range v {
		strs[i] = err.Error()
	}
	return strings.Join(strs, ", ")
}

func FromError(err error) ValidationErrors {
	if err == nil {
		return nil
	}
	if errs, ok := err.(ValidationErrors); ok {
		return errs
	}
	return []error{err}
}
func FromErrors(errs []error) ValidationErrors {
	if len(errs) == 0 {
		return nil
	}
	return errs
}
