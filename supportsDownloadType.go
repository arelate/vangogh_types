package vangogh_types

func SupportsDownloadType(pt ProductType, dt DownloadType) bool {
	if dt == UnknownDownloadType || pt == UnknownProductType {
		return false
	}

	switch dt {
	case ProductImage:
		return pt == StoreProducts ||
			pt == AccountProducts ||
			pt == WishlistProducts
	default:
		return false
	}
}
