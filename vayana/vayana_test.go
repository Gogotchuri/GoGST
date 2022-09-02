package vayana

import (
	"github.com/gogotchuri/GoGST/types"
	"testing"
)

func TestClient_Ping(t *testing.T) {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	err := client.Authenticate("tech+vayana1@kernel.finance", "Strawhats16!")
	if err != nil {
		t.Error(err)
	}
	err = client.Logout()
	if err != nil {
		t.Error("error logging out", err.Error())
	}
}

func TestClient_GetGSTINDetails(t *testing.T) {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MjAzMzk2MSwiaWF0IjoxNjYyMDEyMzYxfQ.akPq4rK3IW-NSUzXmvYK6vG0Hr436VwHS05IJcB2G6yHpV2bTtNNO0HZemLWMo2Lmjn5rEyfUfy3M8o-u7wkgSDVrhwx31tYw3-8RfXCRIgm_Ozm3AgezW-9XaK3eNbcNVP5X0wVAXPFHyAls2r6ONbzOaj7NLC4YGEO-ryRzLnFDUl8DfuSJfZjB7m8bhY32jts7xtKXbX6QzfLJvwjxBJNDSYcM3dUuCK_Ivw_DdERVPG3qCommI37KOXpo4PE9Ox0YTUW0VtU9p64s-CLogBzsK0SB3PgQYtVcs6Opx0bd3ew2inRZMRqfeKrfVh7zavmNfVuSwkoWwp1RWDeRQ")
	resp, err := client.GetGSTINDetails("29AKLPM8755F1Z2")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestClient_GenerateEWB(t *testing.T) {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MjExMzg1OSwiaWF0IjoxNjYyMDkyMjU5fQ.gy3oxyOcVyUJBivWhh24rTka9euI1eZgjnpUfKw4T3jmNfxRw0dlNzaTKgaed8oP56RUnKqHBlFmrnrHGnDJ--lzaypK41Tf1Xqt6uJBXCAqnDqY71R5lKft82vbPHlOT5eensIwrww5YtkkejSVc32ashjNZdknybUvQkjRdeIvk8Y3Qw2aSxUTiCvPo31lfaEEAyzjl3GxbhvA98gKjFd5eShCvTMNCirU5pgAUkptxO_cYeXj0Kqe-DzqM-2VZBESoKVK-DM7jyju8aidY__wd3_eazbz0l5pBVrSefK5WmUNiG_2LxYBExVC9PB4fD2xHvONTZgCP1dzeLi5Dw")
	resp, err := client.CreateEWaybill(types.EWBCreateRequest{
		EWBBase: types.EWBBase{
			SupplyType:            "O",
			SubSupplyType:         "1",
			DocType:               "INV",
			DocNo:                 "7001-47",
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
		t.Error(err)
		return
	}
	t.Logf("%+v", resp)
	ewb, err := client.GetEWayBill(resp.EwayBillNo)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", ewb)
}
