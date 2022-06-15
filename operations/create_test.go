package operations_test

import (
	"github.com/gogotchuri/GoGST/operations"
	"github.com/gogotchuri/GoGST/types"
	"testing"
)

func TestCreateEWaybill(t *testing.T) {
	ewb := types.EWBCreateRequest{
		SupplyType:            "O",
		SubSupplyType:         "1",
		SubSupplyDesc:         "",
		DocType:               "INV",
		DocNo:                 "SRI/04",
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
		OtherValue:            0,
		TotalValue:            56099,
		CgstValue:             0,
		SgstValue:             0,
		IgstValue:             300.67,
		CessValue:             400.56,
		CessNonAdvolValue:     400,
		TotalInvoiceValue:     68358,
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
				Quantity:      4,
				QtyUnit:       "BOX",
				CgstRate:      0,
				SgstRate:      0,
				IgstRate:      3,
				CessRate:      3,
				CessNonadvol:  0,
				TaxableAmount: 56098,
			},
		},
	}
	resp, err := operations.CreateEWaybill(ewb)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(resp)
}
