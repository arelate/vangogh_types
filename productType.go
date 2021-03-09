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

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",
	StorePage:          "store-page",
	StoreProducts:      "store-products",
	AccountPage:        "account-page",
	AccountProducts:    "account-products",
	WishlistPage:       "wishlist-page",
	WishlistProducts:   "wishlist-products",
	Details:            "details",
	ApiProductsV1:      "api-products-v1",
	ApiProductsV2:      "api-products-v2",
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[UnknownProductType]
}

func ParseProductType(productType string) ProductType {
	for pt, str := range productTypeStrings {
		if str == productType {
			return pt
		}
	}
	return UnknownProductType
}

func ValidProductType(pt ProductType) bool {
	_, ok := productTypeStrings[pt]
	return ok && pt != UnknownProductType
}
