package vayana

import (
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"testing"
	"time"
)

const TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiYmIyZGUyYmEtNWQ0NC00MDQ4LWJiNDktYTNiMTQ3NDhlMjk5IiwibW9iIjoiKzkxLTU5NzczMDM1OSIsImVtbCI6InRlY2grdmFuYXlhQGtlcm5lbC5maW5hbmNlIiwiaXNzIjoidi10aGVvIiwibmFtZSI6IklsaWEgR29nb3RjaHVyaSIsIm9yZ3MiOlsie1wib2lkXCI6XCI1ZGJlMTNmOC1jNjBiLTQ4YTYtODcwNS1kNzM0YjhlMTM0ZTVcIixcInByaW1cIjp0cnVlLFwiYWRtXCI6dHJ1ZSxcInNlcnZcIjpbXCJ2c1wiLFwiZ3NwXCIsXCJlYXBpXCIsXCJzYWhpZ3N0XCJdfSJdLCJleHAiOjE2Njk3Mjc1NjQsImlhdCI6MTY2OTcwNTk2NH0.2CPTxNmiqUWJCyAzKnrl70XdVFszfbl93UCIUN1mVtWcwIbKGPaC1CxycVjjpHT0_m3urhq6KQCQ-mln7xnV6V_RCO5YsSDjvvZp99pa6SSL1xIguK2pRqLZTn1GTWik1wRNY7ejcCwD1XBNPr5AUokGgnCsuEK75MTy5B3mhTcVTriGpUHNDyLcktHMBBPAgENoGFQQFPgZSGcBIUotSSWGIO9qE6y1X2o0jNc0uZjYK55wmR6mCBg6N4UgsVLypKuuoWjCRAbObVAzkq8FrM-o0rliBW_cOj-_Gs5wsuyssxrknPtFIe5ifupQAexXui1IDZd6b6MW3LInj1MgnQ"
const IsProduction = false
const OrgID = "5dbe13f8-c60b-48a6-8705-d734b8e134e5"

func getGSPClient() GoGST.GSPClient {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	gspC, _ := client.CreateGSPClient("29AAACW4202F1ZM", "test_dlr228", "test_dlr228")
	return gspC
}

func getEInvoicesClient() GoGST.GSPEInvoiceClient {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	client.SetActiveToken(TOKEN)
	eInvoicesC, _ := client.CreateGSPEInvoicesClient("29AAACW4202F1ZM", "test_dlr228", "test_dlr228")
	return eInvoicesC
}

func TestClient_Ping(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	err := client.Authenticate("", "")
	if err != nil {
		t.Error(err)
	}
	err = client.Logout()
	if err != nil {
		t.Error("error logging out", err.Error())
	}
}

func TestClient_GetGSTINDetails(t *testing.T) {
	gspC := getGSPClient()
	resp, err := gspC.GetGSTINDetails("29AAACW4202F1ZM")
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
			DocNo:                 "70001-4534",
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
		GSTIN:     "29AAACW4202F1ZM",
		LegalName: "Legal Company",
		Address: EInvTypes.Address{
			Address1:  "Varketili",
			Address2:  "Parder 231f",
			Location:  "Tbilisi",
			Pin:       516101,
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
				GSTIN:     "29AEKPV7203E1Z9",
				Address: EInvTypes.Address{
					Address1:  "Pradhan Mantri Awas Yojana",
					Location:  "Beml Nagar",
					Pin:       516101,
					StateCode: "37",
				},
			},
			PlaceOfSupply: "37",
		},
		ItemList: []EInvTypes.Item{
			{
				ItemBase: EInvTypes.ItemBase{
					SerialNo:           "1",
					ProductDescription: "Wheat",
					HSNCode:            "100101",
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
