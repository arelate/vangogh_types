package vangogh_types

func SupportsProperty(pt ProductType, property string) bool {
	if !ValidProductType(pt) {
		return false
	}

	switch property {
	case Title:
		return true
	case Developer:
		fallthrough
	case Publisher:
		return pt == StoreProducts ||
			pt == WishlistProducts ||
			pt == ApiProductsV2
	default:
		return false
	}
}
