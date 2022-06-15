package operations_test

import (
	"fmt"
	"github.com/gogotchuri/GoGST/operations"
	"github.com/gogotchuri/GoGST/types"
	"math/rand"
	"testing"
	"time"
)

func TestCreateEWaybill(t *testing.T) {
	client := operations.NewClient("43af8d80-5da7-4479-9b55-7fe342c9eb62",
		"https://my.gstzen.in/~gstzen/a/ewbapi/generate/", "29AAACW6288M1ZH")

	ewb := types.EWBCreateRequest{
		SupplyType:            "O",
		SubSupplyType:         "1",
		SubSupplyDesc:         "",
		DocType:               "INV",
		DocNo:                 "SRI/09",
		DocDate:               "15/12/2017",
		FromGstin:             "29AAACW6288M1ZH",
		FromTrdName:           "welton",
		FromAddr1:             "2ND CROSS NO 59  19  A",
		FromAddr2:             "GROUND FLOOR OSBORNE ROAD",
		FromPlace:             "FRAZER TOWN",
		FromPincode:           560090,
		SupplierFromStateCode: 29,
		FromStateCode:         29,
		ToGstin:               "02EHFPS5910D2Z0",
		ToTrdName:             "sthuthya",
		ToAddr1:               "Shree Nilaya",
		ToAddr2:               "Dasarahosahalli",
		ToPlace:               "Beml Nagar",
		ToPincode:             560090,
		ActToStateCode:        29,
		ToStateCode:           27,
		TransactionType:       4,
		TransporterId:         "",
		TransporterName:       "",
		TransDocNo:            "",
		TransMode:             "1",
		TransDistance:         100,
		TransDocDate:          "",
		VehicleNo:             "PVC1234",
		VehicleType:           "R",
		ItemList: []types.EWBItem{
			{
				ProductName:   "Wheat",
				ProductDesc:   "Wheat",
				HsnCode:       1001,
				Quantity:      1,
				QtyUnit:       "BOX",
				IgstRate:      5,
				TaxableAmount: 210,
			},
		},
		TotalValue:        210,
		IgstValue:         10.5,
		TotalInvoiceValue: 220.5,
	}
	rand.Seed(time.Now().UnixNano())
	ewb.DocNo += fmt.Sprintf("%d", rand.Intn(10000))
	resp, err := client.CreateEWaybill(ewb)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
