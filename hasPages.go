package vangogh_types

func HasPages(pt ProductType) bool {
	switch pt {
	case StorePage:
		fallthrough
	case AccountPage:
		fallthrough
	case WishlistPage:
		return true
	default:
		return false
	}
}
