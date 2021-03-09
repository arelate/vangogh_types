package vangogh_types

var requiresAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
}

func RequiresAuth(pt ProductType) bool {
	for _, ra := range requiresAuth {
		if ra == pt {
			return true
		}
	}
	return false
}
