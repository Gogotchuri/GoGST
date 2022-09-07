package vayana

import (
	"github.com/gogotchuri/GoGST"
	"github.com/gogotchuri/GoGST/types"
	vayanaTypes "github.com/gogotchuri/GoGST/vayana/types"
	"testing"
)

func getGSPClient() GoGST.GSPClient {
	client, _ := NewDefaultClient(false, "3b5f5dfa-d244-40ed-bfb1-04d7820328b5")
	client.SetActiveToken("eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiJBdXRoIiwidWlkIjoiOTYyM2Y5MDYtNmMyMy00NGFiLTg5MzUtYjU0OWZiMDM4ZjEwIiwibW9iIjoiKzkxLTU1NTU1NTUiLCJlbWwiOiJ0ZWNoK3ZheWFuYTFAa2VybmVsLmZpbmFuY2UiLCJpc3MiOiJ2LXRoZW8iLCJuYW1lIjoiSWxpYSBHb2dvdGNodXJpIiwib3JncyI6WyJ7XCJvaWRcIjpcIjNiNWY1ZGZhLWQyNDQtNDBlZC1iZmIxLTA0ZDc4MjAzMjhiNVwiLFwicHJpbVwiOnRydWUsXCJhZG1cIjp0cnVlLFwic2VydlwiOltcImdzcFwiLFwiZWFwaVwiLFwidnNcIl19Il0sImV4cCI6MTY2MjU0OTUwNCwiaWF0IjoxNjYyNTI3OTA0fQ.wdMwVr-y50hg8ShFzgACwmOWENAKGoMlnAtkeyAm-ViBKgB1X_XC_FZMXvyreGMEq2VqJeGSqU9nglguwqEjFXq2r0t38Bh_fmYOjSK7zW7HYa2sNAGCwc1-cpV62pYp6SrxcYf_2aqvxqOvIEBU4k9WYJ20VoO6w9j3kdJDaDxamluSvswgSCbjIcFSQQYZ9eXyURT9QIeUGVSmBNEdD7VjeHzJ-Dh3U_W-u9md81QDdbEoaPQ_35rIJeQvTT_O3V-_SDP56eNlX_IsZ9HpC3uh5UyM56TSkO8YD-zOfzewhR-QeOVJfL48_7k6J6HvfY8oe7S1Ps4vCY0_iTYltg")
	gspC, _ := client.CreateGSPClient("29AAACW4202F1ZM", "test_dlr228", "test_dlr228")
	return gspC
}

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
	gspC := getGSPClient()
	resp, err := gspC.GetGSTINDetails("29AKLPM8755F1Z2")
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
			DocNo:                 "7001-45",
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
