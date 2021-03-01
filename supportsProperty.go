package vangogh_types

func SupportsProperty(pt ProductType, property string) bool {
	if !ValidProductType(pt) {
		return false
	}

	switch property {
	case TitleProperty:
		return true
	case DeveloperProperty:
		fallthrough
	case PublisherProperty:
		return pt == StoreProducts ||
			pt == WishlistProducts ||
			pt == ApiProductsV2
	default:
		return false
	}
}
