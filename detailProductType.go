package vangogh_types

func DetailProductType(pt ProductType) ProductType {
	switch pt {
	case StorePage:
		return StoreProducts
	case AccountPage:
		return AccountProducts
	case WishlistPage:
		return WishlistProducts
	default:
		return UnknownProductType
	}
}
