package vangogh_types

var detailMainProductTypes = map[ProductType][]ProductType{
	Details:       {AccountProducts},
	ApiProductsV1: {StoreProducts, AccountProducts},
	ApiProductsV2: {StoreProducts, AccountProducts},
}

func MainProductTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}
