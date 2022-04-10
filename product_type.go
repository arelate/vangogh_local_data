package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
)

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
	Licences
	LicenceProducts
	OrderPage
	Orders
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
	Licences:           "licences",
	LicenceProducts:    "licence-products",
	OrderPage:          "order-page",
	Orders:             "orders",
}

//the list is intentionally scoped to very few types we anticipate
//will be interesting to output in human-readable form
var productTypeHumanReadableStrings = map[ProductType]string{
	StoreProducts:    "store",
	WishlistProducts: "wishlist",
	AccountProducts:  "account",
	Details:          "account",
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[UnknownProductType]
}

func (pt ProductType) HumanReadableString() string {
	return productTypeHumanReadableStrings[pt]
}

func ParseProductType(productType string) ProductType {
	for pt, str := range productTypeStrings {
		if str == productType {
			return pt
		}
	}
	return UnknownProductType
}

func IsValidProductType(pt ProductType) bool {
	_, ok := productTypeStrings[pt]
	return ok && pt != UnknownProductType
}

func PagedProducts() []ProductType {
	return []ProductType{
		StorePage,
		AccountPage,
		WishlistPage,
		OrderPage,
	}
}

func ArrayProducts() []ProductType {
	return []ProductType{
		Licences,
	}
}

func FastSyncProducts() []ProductType {
	return []ProductType{
		AccountPage,
		WishlistPage,
	}
}

var detailMainProductTypes = map[ProductType][]ProductType{
	Details: {LicenceProducts, AccountProducts},
	ApiProductsV1: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2,
	},
	ApiProductsV2: {
		StoreProducts,
		AccountProducts,
		ApiProductsV2, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
}

func DetailProducts() []ProductType {
	pts := make([]ProductType, 0, len(detailMainProductTypes))
	for pt := range detailMainProductTypes {
		pts = append(pts, pt)
	}
	return pts
}

func MainProductTypes(pt ProductType) []ProductType {
	return detailMainProductTypes[pt]
}

func RemoteProducts() []ProductType {
	remote := make([]ProductType, 0)
	remote = append(remote, PagedProducts()...)
	remote = append(remote, ArrayProducts()...)
	return append(remote, DetailProducts()...)
}

func LocalProducts() []ProductType {
	return []ProductType{
		StoreProducts,
		AccountProducts,
		WishlistProducts,
		Details,
		ApiProductsV1,
		ApiProductsV2,
		LicenceProducts,
		Orders,
	}
}

var requireAuth = []ProductType{
	AccountPage,
	WishlistPage,
	Details,
	Licences,
	OrderPage,
}

var splitProductTypes = map[ProductType]ProductType{
	StorePage:    StoreProducts,
	AccountPage:  AccountProducts,
	WishlistPage: WishlistProducts,
	Licences:     LicenceProducts,
	OrderPage:    Orders,
}

func SplitProductType(pt ProductType) ProductType {
	splitProductType, ok := splitProductTypes[pt]
	if ok {
		return splitProductType
	}

	return UnknownProductType
}

var supportsGetItems = []ProductType{
	Details,
	ApiProductsV1,
	ApiProductsV2,
	Licences,
}

// unsupported is used instead of supported in similar cases to
// avoid all, but one repetitive data
var unsupportedMedia = map[ProductType][]gog_integration.Media{
	ApiProductsV2: {gog_integration.Movie},
}

var supportsCopyFromTo = map[ProductType]ProductType{
	StoreProducts: WishlistProducts,
}

var supportedImageTypes = map[ProductType][]ImageType{
	StoreProducts:    {Image, Screenshots},
	AccountProducts:  {Image},
	WishlistProducts: {Image},
	ApiProductsV1:    {Screenshots},
	ApiProductsV2:    {Image, Screenshots},
}

func ProductTypesSupportingImageType(imageType ImageType) []ProductType {
	pts := make([]ProductType, 0)
	for pt, its := range supportedImageTypes {
		for _, it := range its {
			if it == imageType {
				pts = append(pts, pt)
				break
			}
		}
	}
	return pts
}

func SupportedPropertiesOnly(pt ProductType, properties []string) []string {
	supported := make([]string, 0, len(properties))
	for _, prop := range properties {
		if IsSupportedProperty(pt, prop) {
			supported = append(supported, prop)
		}
	}
	return supported
}

func Cut(idSet map[string]bool, pt ProductType, mt gog_integration.Media) error {
	ptDir, err := AbsLocalProductTypeDir(pt, mt)
	if err != nil {
		return err
	}
	kvPt, err := kvas.ConnectLocal(ptDir, kvas.JsonExt)
	if err != nil {
		return err
	}

	for id := range idSet {
		if _, err := kvPt.Cut(id); err != nil {
			return err
		}
	}

	return nil
}
