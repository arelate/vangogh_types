package vangogh_types

func DetailProductType(pt ProductType) ProductType {
	switch pt {
	case Store:
		return StoreProducts
	case Account:
		return AccountProducts
	case Wishlist:
		return WishlistProducts
	default:
		return UnknownProductType
	}
}
