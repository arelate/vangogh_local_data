package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
)

type ProductType int

const (
	UnknownProductType ProductType = iota
	// GOG.com product types
	StorePage
	StoreProducts
	CatalogPage
	CatalogProducts
	AccountPage
	AccountProducts
	WishlistPage
	WishlistProducts
	UserWishlist
	Details
	ApiProductsV1
	ApiProductsV2
	Licences
	LicenceProducts
	OrderPage
	Orders
	// Steam product types
	SteamAppList
	SteamAppNews
	SteamReviews
	SteamStorePage
)

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",
	// GOG.com product types
	StorePage:        "store-page",
	StoreProducts:    "store-products",
	CatalogPage:      "catalog-page",
	CatalogProducts:  "catalog-products",
	AccountPage:      "account-page",
	AccountProducts:  "account-products",
	WishlistPage:     "wishlist-page",
	WishlistProducts: "wishlist-products",
	UserWishlist:     "user-wishlist",
	Details:          "details",
	ApiProductsV1:    "api-products-v1",
	ApiProductsV2:    "api-products-v2",
	Licences:         "licences",
	LicenceProducts:  "licence-products",
	OrderPage:        "order-page",
	Orders:           "orders",
	// Steam product types
	SteamAppList:   "steam-app-list",
	SteamAppNews:   "steam-app-news",
	SteamReviews:   "steam-reviews",
	SteamStorePage: "steam-store-page",
}

// the list is intentionally scoped to very few types we anticipate
// will be interesting to output in human-readable form
var productTypeHumanReadableStrings = map[ProductType]string{
	StoreProducts:    "store",
	CatalogProducts:  "store",
	WishlistProducts: "wishlist",
	AccountProducts:  "account",
	Details:          "account",
	SteamAppNews:     "news",
}

func (pt ProductType) String() string {
	str, ok := productTypeStrings[pt]
	if ok {
		return str
	}

	return productTypeStrings[UnknownProductType]
}

func (pt ProductType) HumanReadableString() string {
	if hs, ok := productTypeHumanReadableStrings[pt]; ok {
		return hs
	} else {
		return pt.String()
	}
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
		CatalogPage,
		AccountPage,
		WishlistPage,
		OrderPage,
	}
}

func ArrayProducts() []ProductType {
	return []ProductType{
		Licences,
		UserWishlist,
		SteamAppList,
	}
}

func FastPageFetchProducts() []ProductType {
	return []ProductType{
		OrderPage,
		WishlistPage,
	}
}

var gogDetailMainProductTypes = map[ProductType][]ProductType{
	Details: {LicenceProducts, AccountProducts},
	ApiProductsV1: {
		StoreProducts,
		CatalogProducts,
		AccountProducts,
		ApiProductsV2,
	},
	ApiProductsV2: {
		StoreProducts,
		CatalogProducts,
		AccountProducts,
		ApiProductsV2, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
}

var steamDetailMainProductTypes = map[ProductType][]ProductType{
	//Steam product types are updated on GOG.com store or account product changes
	SteamAppNews: {
		StoreProducts,
		CatalogProducts,
		AccountProducts,
	},
	SteamReviews: {
		StoreProducts,
		CatalogProducts,
		AccountProducts,
	},
	SteamStorePage: {
		StoreProducts,
		CatalogProducts,
		AccountProducts,
	},
}

func detailProducts(dmp map[ProductType][]ProductType) []ProductType {
	pts := make([]ProductType, 0, len(dmp))
	for pt := range dmp {
		pts = append(pts, pt)
	}
	return pts
}

func GOGDetailProducts() []ProductType {
	return detailProducts(gogDetailMainProductTypes)
}

func SteamDetailProducts() []ProductType {
	return detailProducts(steamDetailMainProductTypes)
}

func MainProductTypes(pt ProductType) []ProductType {
	if IsGOGDetailProduct(pt) {
		return gogMainProductTypes(pt)
	} else if IsSteamDetailProduct(pt) {
		return steamMainProductTypes(pt)
	} else {
		return nil
	}
}

func gogMainProductTypes(pt ProductType) []ProductType {
	return gogDetailMainProductTypes[pt]
}

func steamMainProductTypes(pt ProductType) []ProductType {
	return steamDetailMainProductTypes[pt]
}

func GOGRemoteProducts() []ProductType {
	remote := make([]ProductType, 0)
	remote = append(remote, PagedProducts()...)
	remote = append(remote, ArrayProducts()...)
	return append(remote, GOGDetailProducts()...)
}

func SteamRemoteProducts() []ProductType {
	return SteamDetailProducts()
}

func LocalProducts() []ProductType {
	return []ProductType{
		StoreProducts,
		CatalogProducts,
		AccountProducts,
		WishlistProducts,
		UserWishlist,
		Details,
		ApiProductsV1,
		ApiProductsV2,
		LicenceProducts,
		Orders,
		SteamAppNews,
		SteamReviews,
		SteamStorePage,
	}
}

var requireAuth = []ProductType{
	AccountPage,
	WishlistPage,
	UserWishlist,
	Details,
	Licences,
	OrderPage,
}

var splitProductTypes = map[ProductType]ProductType{
	StorePage:    StoreProducts,
	CatalogPage:  CatalogProducts,
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
	SteamAppList,
	SteamAppNews,
	SteamReviews,
	SteamStorePage,
}

// unsupported is used instead of supported in similar cases to
// avoid all, but one repetitive data
var unsupportedMedia = map[ProductType][]gog_integration.Media{
	ApiProductsV2:  {gog_integration.Movie},
	SteamAppList:   {gog_integration.Movie},
	SteamAppNews:   {gog_integration.Movie},
	SteamReviews:   {gog_integration.Movie},
	SteamStorePage: {gog_integration.Movie},
}

var supportedImageTypes = map[ProductType][]ImageType{
	StoreProducts:    {Image, Screenshots},
	CatalogProducts:  {Image, Screenshots},
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

func Cut(ids []string, pt ProductType, mt gog_integration.Media) error {
	ptDir, err := AbsLocalProductTypeDir(pt, mt)
	if err != nil {
		return err
	}
	kvPt, err := kvas.ConnectLocal(ptDir, kvas.JsonExt)
	if err != nil {
		return err
	}

	for _, id := range ids {
		if _, err := kvPt.Cut(id); err != nil {
			return err
		}
	}

	return nil
}
