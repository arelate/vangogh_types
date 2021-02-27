package vangogh_types

func MainProductTypes(pt ProductType) []ProductType {
	switch pt {
	case Details:
		return []ProductType{AccountProducts}
	case ApiProductsV1:
		return []ProductType{StoreProducts, AccountProducts}
	case ApiProductsV2:
		return []ProductType{StoreProducts, AccountProducts}
	default:
		return []ProductType{}
	}
}
