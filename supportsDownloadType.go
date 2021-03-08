package vangogh_types

func SupportsDownloadType(pt ProductType, dt DownloadType) bool {
	if !ValidProductType(pt) ||
		!ValidDownloadType(dt) {
		return false
	}

	switch dt {
	case Image:
		return supportsImage(pt)
	case BoxArt:
		return supportsBoxArt(pt)
	case BackgroundImage:
		return supportsBackgroundImage(pt)
	case GalaxyBackgroundImage:
		return supportsGalaxyBackgroundImage(pt)
	case Logo:
		return supportsLogo(pt)
	case Icon:
		return supportsIcon(pt)
	default:
		return false
	}
}

func supportsImage(pt ProductType) bool {
	switch pt {
	case StoreProducts:
		fallthrough
	case AccountProducts:
		fallthrough
	case WishlistProducts:
		fallthrough
	case ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsBoxArt(pt ProductType) bool {
	switch pt {
	case ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsLogo(pt ProductType) bool {
	switch pt {
	case ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsIcon(pt ProductType) bool {
	switch pt {
	case ApiProductsV1:
		fallthrough
	case ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsBackgroundImage(pt ProductType) bool {
	switch pt {
	case Details:
		fallthrough
	case ApiProductsV1:
		fallthrough
	case ApiProductsV2:
		return true
	default:
		return false
	}
}

func supportsGalaxyBackgroundImage(pt ProductType) bool {
	switch pt {
	case ApiProductsV2:
		return true
	default:
		return false
	}
}
