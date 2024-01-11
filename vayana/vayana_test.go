package vayana

import (
	"testing"
	"time"

	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
)

const TOKEN = "."
const IsProduction = true
const OrgID = ""

const GSTIN = "29AAAPI3182M000"
const TUser = ""
const TPass = "Info21einv#Done"
const Password = ""

func getGSPClient() GoGST.GSPClient {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	gspC, _ := client.CreateGSPClient(GSTIN, TUser, TPass)
	return gspC
}

func getEInvoicesClient() GoGST.GSPEInvoiceClient {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	eInvoicesC, _ := client.CreateGSPEInvoicesClient(GSTIN, TUser, TPass)
	return eInvoicesC
}

func TestClient_Ping(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_AuthenticatedPing(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	err := client.AuthenticatedPing()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	err := client.Authenticate(TUser, Password)
	if err != nil {
		t.Fatal(err)
	}
	err = client.Logout()
	if err != nil {
		t.Error("error logging out", err.Error())
	}
}

func TestClient_GetGSTINDetails(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	gspC, err := client.CreateGSTNClient(GSTIN)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := gspC.GetTaxPayerDetails("29AAACW4202F1ZM")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestClient_GenerateEWB(t *testing.T) {
	gspC := getGSPClient()
	resp, err := gspC.CreateEWaybill(types.EWBCreateRequest{
		EWBBase: types.EWBBase{
			SupplyType:            "O",
			SubSupplyType:         "1",
			DocType:               "INV",
			DocNo:                 "70001-4534AA",
			DocDate:               "15/05/2022",
			FromGstin:             "29AAACW4202F1ZM",
			FromTrdName:           "welton",
			FromAddr1:             "2ND CROSS NO 59  19  A",
			FromAddr2:             "GROUND FLOOR OSBORNE ROAD",
			FromPlace:             "FRAZER TOWN",
			FromPincode:           560090,
			FromStateCode:         29,
			SupplierFromStateCode: 29,
			ToGstin:               "29AEKPV7203E1Z9",
			ToTrdName:             "sthuthya",
			ToAddr1:               "Shree Nilaya",
			ToAddr2:               "Dasarahosahalli",
			ToPlace:               "Beml Nagar",
			ToPincode:             516101,
			ToStateCode:           37,
			ActToStateCode:        37,
			TransactionType:       4,
			OtherValue:            100,
			TotalValue:            56099,
			CgstValue:             0,
			SgstValue:             0,
			IgstValue:             300,
			CessValue:             400,
			CessNonAdvolValue:     400,
			TotalInvoiceValue:     68358,
			VehicleType:           "R",
			TransporterId:         "29AAACW6874H1ZS",
			ItemList: []types.EWBItem{
				{
					ProductName:   "Wheat",
					ProductDesc:   "Wheat",
					HsnCode:       1001,
					Quantity:      4,
					QtyUnit:       "BOX",
					CgstRate:      0,
					SgstRate:      0,
					IgstRate:      3,
					TaxableAmount: 56098.89,
					CessRate:      3,
					CessNonadvol:  0,
				},
			},
		},
		TransMode:     "1",
		VehicleNo:     "PVC1234",
		TransDistance: "100",
		TransDocDate:  "15/05/2022",
	})
	if err != nil {
		t.Error(err, err == vayanaTypes.ErrorTokenExpired)
		return
	}
	t.Logf("%+v", resp)
	ewb, err := gspC.GetEWayBill(resp.EwayBillNo)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ewb)
}

func TestGspClient_CancelEWaybill(t *testing.T) {
	gspC := getGSPClient()
	res, err := gspC.CancelEWaybill(types.EWBCancelRequest{
		CancelRemark:     "Other",
		CancelReasonCode: 3,
		EwayBillNo:       141002145189,
	})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
}

func TestGspClient_GetEWayBillsByDate(t *testing.T) {
	gspC := getGSPClient()
	res, err := gspC.GetEWayBillsByDate(time.Now())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
}

func TestGSPEInvoicesClient_CreateEInvoice(t *testing.T) {
	einvClient := getEInvoicesClient()
	res, err := einvClient.CreateEInvoice(getSampleInvoiceCreateRequest())
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", res)
}

func getSampleInvoiceCreateRequest() EInvTypes.EInvoiceCreate {
	seller := EInvTypes.Company{
		GSTIN:     GSTIN,
		LegalName: "Legal Company",
		Address: EInvTypes.Address{
			Address1:  "Varketili",
			Address2:  "Parder 231f",
			Location:  "Tbilisi",
			Pin:       560090,
			StateCode: "29",
		},
	}
	return EInvTypes.EInvoiceCreate{
		DocumentDetails: EInvTypes.DocumentDetails{
			DocumentNo: "DOC102335",
			Date:       "30/11/2022",
		},
		SellerDetails: EInvTypes.Seller{
			Company: seller,
		},
		BuyerDetails: EInvTypes.Buyer{
			Company: EInvTypes.Company{
				LegalName: "Legal Company Buyer",
				GSTIN:     "29AAACW6288M1ZH",
				Address: EInvTypes.Address{
					Address1:  "Pradhan Mantri Awas Yojana",
					Location:  "Beml Nagar",
					Pin:       560090,
					StateCode: "29",
				},
			},
			PlaceOfSupply: "37",
		},
		ItemList: []EInvTypes.Item{
			{
				ItemBase: EInvTypes.ItemBase{
					SerialNo:           "1",
					ProductDescription: "Wheat",
					HSNCode:            "01011020",
					Unit:               "PCS",
					UnitPrice:          10,
					Quantity:           10,
					NetAmount:          100,
					TaxableAmount:      100,
					IGSTRate:           3,
					IGSTAmount:         3,
					TotalAmount:        103,
				},
			},
		},
		DocumentValues: EInvTypes.DocumentValues{
			TaxableValue: 100,
			IGSTValue:    3,
			TotalAmount:  103,
		},
	}
}
