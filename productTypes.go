package vangogh_types

type ProductType int

const (
	Unknown ProductType = iota
	Store
	StoreProducts
	Account
	AccountProducts
	Wishlist
	WishlistProducts
	Details
	ApiProducts
)

const (
	storeStr            = "store"
	storeProductsStr    = "store-products"
	accountStr          = "account"
	accountProductsStr  = "account-products"
	wishlistStr         = "wishlist"
	wishlistProductsStr = "wishlist-products"
	detailsStr          = "details"
	apiProductsStr      = "api-products"
	unknownStr          = "unknown"
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
	case ApiProducts:
		return apiProductsStr
	default:
		return unknownStr
	}
}

func Parse(productType string) ProductType {
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
	case apiProductsStr:
		return ApiProducts
	default:
		return Unknown
	}
}
