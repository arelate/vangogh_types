package vangogh_types

import "github.com/arelate/gog_types"

// unsupported is used instead of supported in similar cases to
// avoid all, but one repetitive data
var unsupporedMedia = map[ProductType][]gog_types.Media{
	ApiProductsV2: {gog_types.Movie},
}

func SupportsMedia(pt ProductType, mt gog_types.Media) bool {
	if !gog_types.ValidMedia(mt) {
		return false
	}
	if !ValidProductType(pt) {
		return false
	}

	ums, ok := unsupporedMedia[pt]
	if !ok {
		return true
	}

	for _, um := range ums {
		if um == mt {
			return false
		}
	}

	return true
}
