package vayana

import (
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"testing"
	"time"
)

func getGSPClient() GoGST.GSPClient {
	client, _ := NewDefaultClient(true, "da3a9839-c072-46b2-ac4f-601f72d39674")
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiZmI4OTM0YjktODkzZS00NmRjLThjNmEtMTU2NGUwOTBhOGYwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoQGtlcm5lbC5maW5hbmNlIiwiaXNzIjoidi10aGVvIiwibmFtZSI6IklsaWEgR29nb3RjaHVyaSIsIm9yZ3MiOlsie1wib2lkXCI6XCJkYTNhOTgzOS1jMDcyLTQ2YjItYWM0Zi02MDFmNzJkMzk2NzRcIixcInByaW1cIjp0cnVlLFwiYWRtXCI6dHJ1ZSxcInNlcnZcIjpbXCJ2c1wiLFwiZ3NwXCIsXCJlYXBpXCJdfSJdLCJleHAiOjE2NjMyNjEyNDQsImlhdCI6MTY2MzIzOTY0NH0.g5g9CLOYU5QyjR_uPQzlNypRKlPtyTrkagH-xFqM9WyMYQOfvV10KYPAAEv_MoRsjmx4pRBNagfOhC2LKZ5CGRWf3D-qFtfXlsFAuhGqSLmKEuy_FRgutCIKAil2MtzBshosAj0EEttXrctZZ8BsBXWjcekXIex9a96IneRYlTsTiS915ZMH-5wEB6kc_ZWXqOXuajdOnQ6a9es9yhA3neU_Gblv3YoNhkvY7AqHdD0ATeMNd4BWi8El7-atn7kwwNX1LmgSnOKLC2nYITL-lVa_mG6pP96qhvy9DOTEDXLpFxevLVl7IutcfF4G7U2GzTGqlj6CCMBt31T6GJBJAMGFTn8uNiuaca0dfzqZyVZNctkSYpCzVPbCh-h30XBqE_fyIIss3dqo-BsEPTEOWvUy7G8ZfCLNSivei1M6XM7VMa7jsEDqzt7nER_bQrsV9uebJ4Kz5qzp5EDYtXkPfEwQLmK_hAtL1t2Dh4FY2ICLpXZMI5v5MDfnvWGFpBIYvm62NWsDzjl4fivm5H9myc2Bc0N_z0Ge4vEWrsfY7OqnspG9fDFpZd-bx5x5PAWXQB85iR9XjWmta1EoJpoDr7pu9xLTWlT1Q_KsZsZFJUFaakTPbOdqSI65xU1PkMttBihLCCmyLcxw816fX-TfgG82V2o7r7Jd7F4agyX2JhM")
	gspC, _ := client.CreateGSPClient("29AAACW4202F1ZM", "test_dlr228", "test_dlr228")
	return gspC
}

func TestClient_Ping(t *testing.T) {
	client, _ := NewDefaultClient(true, "da3a9839-c072-46b2-ac4f-601f72d39674")
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(true, "da3a9839-c072-46b2-ac4f-601f72d39674")
	err := client.Authenticate("tech@kernel.finance", "Strawhats16!")
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
