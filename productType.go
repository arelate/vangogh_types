package vangogh_types

type ProductType int

const (
	UnknownProductType ProductType = iota
	Store
	StoreProducts
	Account
	AccountProducts
	Wishlist
	WishlistProducts
	Details
	ApiProductsV1
	ApiProductsV2
)

const (
	unknownProductTypeStr = "unknown-product-type"
	storeStr              = "store"
	storeProductsStr      = "store-products"
	accountStr            = "account"
	accountProductsStr    = "account-products"
	wishlistStr           = "wishlist"
	wishlistProductsStr   = "wishlist-products"
	detailsStr            = "details"
	apiProductsV1Str      = "api-products-v1"
	apiProductsV2Str      = "api-products-v2"
)

func (pt ProductType) String() string {
	switch pt {
	case Store:
		return storeStr
	case StoreProducts:
		return storeProductsStr
	case Account:
		return accountStr
	case AccountProducts:
		return accountProductsStr
	case Wishlist:
		return wishlistStr
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
	case storeStr:
		return Store
	case storeProductsStr:
		return StoreProducts
	case accountStr:
		return Account
	case accountProductsStr:
		return AccountProducts
	case wishlistStr:
		return Wishlist
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
