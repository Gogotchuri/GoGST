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
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MjA2MTE1NywiaWF0IjoxNjYyMDM5NTU3fQ.S4k22KHAUHPAw_rno-Nqr2CW_OLM5uNoqqsAAFVv2zMUNwYYQQxoWXa8HRqAL5oMGumBQmaRFMiwDDnt-gqx-oJxlV0P3rNDgWvHPmXxpLJVxPtE6aHZs-Vhcbc3MOmSmfgDnCTNwgm43GLMkBEH-TpMexfSoNAAz6Zp0Ph4pPzuNYOmWcuP-VxBsueaWCGx3FPuKhc_BKNmMMl4tfnE-mqE3OuL6yxYUD6M7-Bf0oWmo6Uj2UuY-n4-HHRxFHNGMy1Dol5nDFJ3tHaMVkWBqte3O119jkf33LL3sjSMTlh49ssTUUZmUuUDsictagAVSO4sOIye5ig42wtj6YuuYw")
	resp, err := client.CreateEWaybill(types.EWBCreateRequest{
		SupplyType:            "O",
		SubSupplyType:         "1",
		DocType:               "INV",
		DocNo:                 "7001-8",
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
		ToPincode:             560090,
		ToStateCode:           27,
		ActToStateCode:        27,
		TransactionType:       4,
		OtherValue:            100,
		TotalValue:            56099,
		CgstValue:             0,
		SgstValue:             0,
		IgstValue:             300,
		CessValue:             400,
		CessNonAdvolValue:     400,
		TotalInvoiceValue:     68358,
		TransMode:             "1",
		VehicleNo:             "PVC1234",
		TransDistance:         "100",
		VehicleType:           "R",
		TransporterId:         "29AAACW6874H1ZS",
		TransDocDate:          "15/05/2022",
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
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
}
