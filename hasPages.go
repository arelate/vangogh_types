package vangogh_types

func HasPages(pt ProductType) bool {
	switch pt {
	case Store:
		fallthrough
	case Account:
		fallthrough
	case Wishlist:
		return true
	default:
		return false
	}
}
