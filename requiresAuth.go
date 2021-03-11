package vangogh_types

var productTypeRequiresAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
}

func ProductTypeRequiresAuth(pt ProductType) bool {
	for _, ra := range productTypeRequiresAuth {
		if ra == pt {
			return true
		}
	}
	return false
}

var downloadTypeRequiresAuth = []DownloadType{}

func DownloadTypeRequiresAuth(dt DownloadType) bool {
	for _, ra := range downloadTypeRequiresAuth {
		if ra == dt {
			return true
		}
	}
	return false
}
