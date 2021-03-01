package vangogh_types

func SupportsDownloadType(pt ProductType, dt DownloadType) bool {
	if dt == UnknownDownloadType || pt == UnknownProductType {
		return false
	}

	switch pt {
	case StoreProducts:
		fallthrough
	case AccountProducts:
		fallthrough
	case WishlistProducts:
		fallthrough
	case ApiProductsV1:
		fallthrough
	case ApiProductsV2:
		return dt == ProductImage
	default:
		return false
	}
}
