package EInvTypes

import "strings"

func getNameFromNamespace(namespace string) string {
	namespace = strings.ReplaceAll(namespace, "SellerDetails.Company", "Seller")
	namespace = strings.ReplaceAll(namespace, "Seller.Address", "Seller")
	namespace = strings.ReplaceAll(namespace, "BuyerDetails.Company", "Buyer")
	namespace = strings.ReplaceAll(namespace, "BuyerDetails", "Buyer")
	namespace = strings.ReplaceAll(namespace, "Buyer.Address", "Buyer")
	namespace = strings.ReplaceAll(namespace, "ShipToDetails.Address", "Ship To Details")
	namespace = strings.ReplaceAll(namespace, "DispatchDetails.Address", "Dispatch Details")
	namespace = strings.ReplaceAll(namespace, "DocumentDetails.Type", "Document Type")
	namespace = strings.ReplaceAll(namespace, "DocumentDetails.DocumentNo", "Document No")
	namespace = strings.ReplaceAll(namespace, "DocumentDetails.Date", "Document Date")
	namespace = strings.ReplaceAll(namespace, "DocumentValue.", "")
	namespace = strings.ReplaceAll(namespace, "StateCode", "State Code")
	namespace = strings.ReplaceAll(namespace, "Address1", "Address One")
	namespace = strings.ReplaceAll(namespace, "Address2", "Address Two")
	namespace = strings.ReplaceAll(namespace, "PlaceOfSupply", "Place Of Supply")
	namespace = strings.ReplaceAll(namespace, "LegalName", "Legal Name")
	namespace = strings.ReplaceAll(namespace, "TradeName", "Trade Name")
	namespace = strings.ReplaceAll(namespace, "ProductDescription", "Product Description")
	namespace = strings.ReplaceAll(namespace, "Rate", " Rate")
	namespace = strings.ReplaceAll(namespace, "Amount", " Amount")
	namespace = strings.ReplaceAll(namespace, "Price", " Price")

	namespace = strings.ReplaceAll(namespace, ".", " ")

	return namespace
}
