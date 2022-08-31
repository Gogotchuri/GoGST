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
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MTk3NDA3NCwiaWF0IjoxNjYxOTUyNDc0fQ.eF6YhM8A-PVVSzv3X1pCpS0Xknw9lOFQozSWXhEo5a0lFQhpIoSITsSda-zYD10VCc_3KE6IrvwuP8_kXVUwBOMeOR_pZCa3ZnQGTsrslWyHK0rKXpbKLE0S1GRmDMgWPX2Z0Vw-atZ5ZlGvffyvGLfESse2BER6Hi_OjqLDnQhJE3s-06wxPOr6UwUJ4Vy7Q4dOFCHzPdTgvcYWh6QJ-VGglXnTHJzrh4yV8sdTUEVmSmR9Pe2JWvvVWWvPILLYL0AqIDSzhn4enBw1kq52wszDGWF5iXNEXtEUG0VrRtwC259ldtKsKwPhFcd4WQow49GKwA2ss9BCNURKD74ucg")
	resp, err := client.GetGSTINDetails("27AAAPI3182M002")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}

func TestClient_GenerateEWB(t *testing.T) {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MTk3NDA3NCwiaWF0IjoxNjYxOTUyNDc0fQ.eF6YhM8A-PVVSzv3X1pCpS0Xknw9lOFQozSWXhEo5a0lFQhpIoSITsSda-zYD10VCc_3KE6IrvwuP8_kXVUwBOMeOR_pZCa3ZnQGTsrslWyHK0rKXpbKLE0S1GRmDMgWPX2Z0Vw-atZ5ZlGvffyvGLfESse2BER6Hi_OjqLDnQhJE3s-06wxPOr6UwUJ4Vy7Q4dOFCHzPdTgvcYWh6QJ-VGglXnTHJzrh4yV8sdTUEVmSmR9Pe2JWvvVWWvPILLYL0AqIDSzhn4enBw1kq52wszDGWF5iXNEXtEUG0VrRtwC259ldtKsKwPhFcd4WQow49GKwA2ss9BCNURKD74ucg")
	resp, err := client.CreateEWaybill(types.EWBCreateRequest{
		SupplyType:            "O",
		SubSupplyType:         "1",
		DocType:               "INV",
		DocNo:                 "7001-8",
		DocDate:               "15/12/2022",
		FromGstin:             "29AAACW6288M1ZH",
		FromTrdName:           "welton",
		FromAddr1:             "2ND CROSS NO 59  19  A",
		FromAddr2:             "GROUND FLOOR OSBORNE ROAD",
		FromPlace:             "FRAZER TOWN",
		FromPincode:           560090,
		FromStateCode:         29,
		SupplierFromStateCode: 29,
		ToGstin:               "",
		ToTrdName:             "sthuthya",
		ToAddr1:               "Shree Nilaya",
		ToAddr2:               "Dasarahosahalli",
		ToPlace:               "Beml Nagar",
		ToPincode:             560090,
		ToStateCode:           27,
		ActToStateCode:        29,
		TransactionType:       4,
		OtherValue:            100,
		TotalValue:            56099,
		CgstValue:             0,
		SgstValue:             0,
		IgstValue:             300.67,
		CessValue:             400.56,
		CessNonAdvolValue:     400,
		TotalInvoiceValue:     68358,
		TransMode:             "1",
		VehicleNo:             "PVC1234",
		TransDistance:         100,
		VehicleType:           "R",
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
				TaxableAmount: 5609889,
				CessRate:      3,
				CessNonadvol:  0,
			},
		},
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
