package vangogh_types

type ProductType int

const (
	UnknownProductType ProductType = iota
	StorePage
	StoreProducts
	AccountPage
	AccountProducts
	WishlistPage
	WishlistProducts
	Details
	ApiProductsV1
	ApiProductsV2
)

const (
	unknownProductTypeStr = "unknown-product-type"
	storePageStr          = "store-page"
	storeProductsStr      = "store-products"
	accountPageStr        = "account-page"
	accountProductsStr    = "account-products"
	wishlistPageStr       = "wishlist-page"
	wishlistProductsStr   = "wishlist-products"
	detailsStr            = "details"
	apiProductsV1Str      = "api-products-v1"
	apiProductsV2Str      = "api-products-v2"
)

func (pt ProductType) String() string {
	switch pt {
	case StorePage:
		return storePageStr
	case StoreProducts:
		return storeProductsStr
	case AccountPage:
		return accountPageStr
	case AccountProducts:
		return accountProductsStr
	case WishlistPage:
		return wishlistPageStr
	case WishlistProducts:
		return wishlistProductsStr
	case Details:
		return detailsStr
	case ApiProductsV1:
		return apiProductsV1Str
	case ApiProductsV2:
		return apiProductsV2Str
	default:
		return unknownProductTypeStr
	}
}

func ParseProductType(productType string) ProductType {
	switch productType {
	case storePageStr:
		return StorePage
	case storeProductsStr:
		return StoreProducts
	case accountPageStr:
		return AccountPage
	case accountProductsStr:
		return AccountProducts
	case wishlistPageStr:
		return WishlistPage
	case wishlistProductsStr:
		return WishlistProducts
	case detailsStr:
		return Details
	case apiProductsV1Str:
		return ApiProductsV1
	case apiProductsV2Str:
		return ApiProductsV2
	default:
		return UnknownProductType
	}
}

func ValidProductType(pt ProductType) bool {
	switch pt {
	case StorePage:
		fallthrough
	case StoreProducts:
		fallthrough
	case AccountPage:
		fallthrough
	case AccountProducts:
		fallthrough
	case WishlistPage:
		fallthrough
	case WishlistProducts:
		fallthrough
	case Details:
		fallthrough
	case ApiProductsV1:
		fallthrough
	case ApiProductsV2:
		return true
	default:
		return false
	}
}
