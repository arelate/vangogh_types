package vangogh_types

import "github.com/arelate/gog_types"

func SupportsMedia(pt ProductType, mt gog_types.Media) bool {
	if !gog_types.ValidMedia(mt) {
		return false
	}
	if !ValidProductType(pt) {
		return false
	}

	if pt == ApiProductsV2 {
		return mt == gog_types.Game
	}

	return true
}
