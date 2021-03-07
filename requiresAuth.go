package vangogh_types

func RequiresAuth(pt ProductType) bool {
	switch pt {
	case AccountPage:
		fallthrough
	case WishlistPage:
		fallthrough
	case Details:
		return true
	default:
		return false
	}
}
