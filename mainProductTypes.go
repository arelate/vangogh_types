package vangogh_types

func MainProductTypes(pt ProductType) []ProductType {
	switch pt {
	case Details:
		return []ProductType{AccountProducts}
	case ApiProducts:
		return []ProductType{StoreProducts, AccountProducts}
	default:
		return []ProductType{}
	}
}
