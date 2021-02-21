package vangogh_types

import "github.com/arelate/gog_types"

func SupportsMedia(pt ProductType, mt gog_types.Media) bool {
	if mt == gog_types.Unknown {
		return false
	}

	switch pt {
	case Store:
		fallthrough
	case Account:
		fallthrough
	case Wishlist:
		fallthrough
	case StoreProducts:
		fallthrough
	case AccountProducts:
		fallthrough
	case WishlistProducts:
		return true
	case ApiProducts:
		return mt == gog_types.Game
	default:
		return false
	}
}
