package vangogh_types

func RequiresAuth(pt ProductType) bool {
	switch pt {
	case Account:
		fallthrough
	case Wishlist:
		fallthrough
	case Details:
		return true
	default:
		return false
	}
}
