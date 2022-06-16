package types

//Supply type
const (
	SupplyType_Inward  = "I"
	SupplyType_Outward = "O"
)

//Sub-Supply type
const (
	SubSupplyType_Supply            = 1
	SubSupplyType_Import            = 2
	SubSupplyType_Export            = 3
	SubSupplyType_JobWork           = 4
	SubSupplyType_ForOwnUse         = 5
	SubSupplyType_JobWorkReturns    = 6
	SubSupplyType_SalesReturn       = 7
	SubSupplyType_Others            = 8
	SubSupplyType_SKD_CKD_Lots      = 9
	SubSupplyType_LineSales         = 10
	SubSupplyType_RecipientNotKnown = 11
	SubSupplyType_ExhibitionOrFairs = 12
)

//Document type
const (
	DocumentType_TaxInvoice      = "INV"
	DocumentType_BillOfSupply    = "BIL"
	DocumentType_BillOfEntry     = "BOE"
	DocumentType_DeliveryChallan = "CHL"
	DocumentType_Other           = "OTH"
)

//Consignment Status
const (
	ConsignmentStatus_InMovement = "M"
	ConsignmentStatus_InTransit  = "T"
)

//Transportation mode
const (
	TransportationMode_Road      = 1
	TransportationMode_Rail      = 2
	TransportationMode_Air       = 3
	TransportationMode_Ship      = 4
	TransportationMode_InTransit = 5
)

//Unit
const (
	Unit_BAGS              = "BAG"
	Unit_BALE              = "BAL"
	Unit_BUNDLES           = "BDL"
	Unit_BUCKLES           = "BKL"
	Unit_BILLION_OF_UNITS  = "BOU"
	Unit_BOX               = "BOX"
	Unit_BOTTLES           = "BTL"
	Unit_BUNCHES           = "BUN"
	Unit_CANS              = "CAN"
	Unit_CUBIC_METERS      = "CBM"
	Unit_CUBIC_CENTIMETERS = "CCM"
	Unit_CENTI_METERS      = "CMS"
	Unit_CARTONS           = "CTN"
	Unit_DOZENS            = "DOZ"
	Unit_DRUMS             = "DRM"
	Unit_GREAT_GROSS       = "GGK"
	Unit_GRAMMES           = "GMS"
	Unit_GROSS             = "GRS"
	Unit_GROSS_YARDS       = "GYD"
	Unit_KILOGRAMS         = "KGS"
	Unit_KILOLITRE         = "KLR"
	Unit_KILOMETRE         = "KME"
	Unit_LITRES            = "LTR"
	Unit_METERS            = "MTR"
	Unit_MILILITRE         = "MLT"
	Unit_METRIC_TON        = "MTS"
	Unit_NUMBERS           = "NOS"
	Unit_OTHERS            = "OTH"
	Unit_PACKS             = "PAC"
	Unit_PIECES            = "PCS"
	Unit_PAIRS             = "PRS"
	Unit_QUINTAL           = "QTL"
	Unit_ROLLS             = "ROL"
	Unit_SETS              = "SET"
	Unit_SQUARE_FEET       = "SQF"
	Unit_SQUARE_METERS     = "SQM"
	Unit_SQUARE_YARDS      = "SQY"
	Unit_TABLETS           = "TBS"
	Unit_TEN_GROSS         = "TGM"
	Unit_THOUSANDS         = "THD"
	Unit_TONNES            = "TON"
	Unit_TUBES             = "TUB"
	Unit_US_GALLONS        = "UGS"
	Unit_UNITS             = "UNT"
	Unit_YARDS             = "YDS"
)

// Update and consolidation reasons
const (
	UpdateReason_DueToBreakDown     = 1
	UpdateReason_DueToTransshipment = 2
	UpdateReason_OthersToSpecify    = 3
	UpdateReason_FirstTime          = 4
)

//Generation mode
const (
	GenerationMode_API            = "API"
	GenerationMode_BulkUpload     = "Exc"
	GenerationMode_SMSFacility    = "SMS"
	GenerationMode_MobileAPP      = "APP"
	GenerationMode_WebBasedSystem = "WEB"
)

//EWB status
const (
	EWBStatus_Active    = "ACT"
	EWBStatus_Cancelled = "CNL"
	EWBStatus_Discarded = "DIS"
)

//Cancel Reason
const (
	CancelReason_Duplicate        = 1
	CancelReason_OrderCancelled   = 2
	CancelReason_DataEntryMistake = 3
	CancelReason_Others           = 4
)

//Vehicle Type
const (
	VehicleType_Regular       = "R"
	VehicleType_OverDimential = "O"
)

// Reason for extension
const (
	ReasonForExtension_NaturalCalamity      = 1
	ReasonForExtension_LawAndOrderSituation = 2
	ReasonForExtension_Transshipment        = 4
	ReasonForExtension_Accident             = 5
	ReasonForExtension_Others               = 99
)

//Transaction Type
const (
	TransactionType_Regular              = 1
	TransactionType_BillToShipTo         = 2
	TransactionType_BillFromDispatchFrom = 3
	TransactionType_BillFromTo           = 4
)
