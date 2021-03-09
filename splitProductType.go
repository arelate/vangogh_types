package vangogh_types

var splitProductTypes = map[ProductType]ProductType{
	StorePage:    StoreProducts,
	AccountPage:  AccountProducts,
	WishlistPage: WishlistProducts,
}

func SplitProductType(pt ProductType) ProductType {
	splitProductType, ok := splitProductTypes[pt]
	if ok {
		return splitProductType
	}

	return UnknownProductType
}
