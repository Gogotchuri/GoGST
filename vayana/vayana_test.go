package vayana

import (
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	"github.com/gogotchuri/GoGST/types/EInvTypes"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"testing"
	"time"
)

const TOKEN = "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiZmI4OTM0YjktODkzZS00NmRjLThjNmEtMTU2NGUwOTBhOGYwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoQGtlcm5lbC5maW5hbmNlIiwiaXNzIjoidi10aGVvIiwibmFtZSI6IklsaWEgR29nb3RjaHVyaSIsIm9yZ3MiOlsie1wib2lkXCI6XCJkYTNhOTgzOS1jMDcyLTQ2YjItYWM0Zi02MDFmNzJkMzk2NzRcIixcInByaW1cIjp0cnVlLFwiYWRtXCI6dHJ1ZSxcInNlcnZcIjpbXCJ2c1wiLFwiZ3NwXCIsXCJlYXBpXCJdfSJdLCJleHAiOjE2NzY2NDg3OTYsImlhdCI6MTY3NjY0NTE5Nn0.FwryhaC4EET8EfaFh9LRvLAt0v9yTssLpHPNGOEzGQNvvHS3EXsOfSP7XaQ4-88QFf1IPVwhFsLEaNTnMDIbJsCRE5GZl57m_jz-Ge0XXkQ2evdz3F4dYiQ0GgGXVFAN5mVLXHKVsMZYz9YG7uV2bJg_wd2t5M1HuzOGxXMEczKcwo_VfCmZ-W5SLc0fWwP_GtXthPK3TNeBrkvBtCKuDGvzSEQGeH4oUSUSlm0AnnQpWPmH4EOua9ZfKesUxVwiOcJzvPeONaVXwIQQdhJ-GmYn718w6sXDkW6cDVk9nMVytEL_JTCG4OW98Y1H8cXAXcZgAGXPjy6o8KeZdmGkP_lfYyO7FRwn2Fscnpz4zjuvumw0MqJ2tum9_CnU9EFALWpBVd3P5wSs3wSgp-LW-sRLcXDBtd4iLB14kLP7FoDLb9Q9j_hv3o8btIhxzzxfWJJ8IKzibhIfoL42r1zf5LKVEBRVjBsI2xBe5bquSusKMbYCJ4DIRaaxPQlhABMnG61mXVC2B2fdTsi81SzbKHowy5g6wSQ3VJCeT8DNEtyuRjz5QFMqGfsNCv_vyYIdhxZtFrs24HtM5LY9QSFW2Aww57APvT0C2JAXnhty9vIQADbiHksdHWVMOLdHtNYAjWI3kP-RcqXJ54vcsxFXxxvECuXkFfbQIgvZ-LYy2aY"
const IsProduction = true
const OrgID = "da3a9839-c072-46b2-ac4f-601f72d39674"

const GSTIN = "33AAHCE3087B1ZZ"
const TUser = "Easy_invoicing"
const TPass = "Easy@2022"
const Password = "Strawhats16!"

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

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(IsProduction, OrgID)
	err := client.Authenticate("tech@kernel.finance", Password)
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
	resp, err := gspC.GetGSTINDetails(GSTIN)
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
