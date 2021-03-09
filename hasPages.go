package vangogh_types

var paginatedProductTypes = []ProductType{
	StorePage,
	AccountPage,
	WishlistPage,
}

func HasPages(pt ProductType) bool {
	for _, ppt := range paginatedProductTypes {
		if ppt == pt {
			return true
		}
	}
	return false
}
