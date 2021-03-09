package vangogh_types

var supportedDownloadTypes = map[ProductType][]DownloadType{
	StoreProducts:    {Image},
	AccountProducts:  {Image},
	WishlistProducts: {Image},
	Details:          {BackgroundImage},
	ApiProductsV1:    {Icon, BackgroundImage},
	ApiProductsV2:    {Image, BoxArt, Logo, Icon, BackgroundImage, GalaxyBackgroundImage},
}

func SupportsDownloadType(pt ProductType, dt DownloadType) bool {
	if !ValidProductType(pt) ||
		!ValidDownloadType(dt) {
		return false
	}

	sdts, ok := supportedDownloadTypes[pt]
	if !ok {
		return false
	}

	for _, sdt := range sdts {
		if sdt == dt {
			return true
		}
	}

	return false
}
