package vangogh_types

func AllPagedProductTypes() []ProductType {
	return []ProductType{
		StorePage,
		AccountPage,
		WishlistPage}
}

func AllDetailProductTypes() []ProductType {
	return []ProductType{
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}

func AllLocalProductTypes() []ProductType {
	return []ProductType{
		StoreProducts,
		AccountProducts,
		WishlistProducts,
		Details,
		ApiProductsV1,
		ApiProductsV2,
	}
}
