package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
)

type ProductType int

const (
	UnknownProductType ProductType = iota
	// GOG.com product types
	CatalogPage
	CatalogProducts
	AccountPage
	AccountProducts
	UserWishlist
	UserWishlistProducts
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
	// PCGamingWiki product types
	PCGWCargo
	PCGWWikiText
)

var productTypeStrings = map[ProductType]string{
	UnknownProductType: "unknown-product-type",
	// GOG.com product types
	CatalogPage:          "catalog-page",
	CatalogProducts:      "catalog-products",
	AccountPage:          "account-page",
	AccountProducts:      "account-products",
	UserWishlist:         "user-wishlist",
	UserWishlistProducts: "user-wishlist-products",
	Details:              "details",
	ApiProductsV1:        "api-products-v1",
	ApiProductsV2:        "api-products-v2",
	Licences:             "licences",
	LicenceProducts:      "licence-products",
	OrderPage:            "order-page",
	Orders:               "orders",
	// Steam product types
	SteamAppList:   "steam-app-list",
	SteamAppNews:   "steam-app-news",
	SteamReviews:   "steam-reviews",
	SteamStorePage: "steam-store-page",
	// PCGamingWiki product types
	PCGWCargo:    "pcgw-cargo",
	PCGWWikiText: "pcgw-wikitext",
}

// the list is intentionally scoped to very few types we anticipate
// will be interesting to output in human-readable form
var productTypeHumanReadableStrings = map[ProductType]string{
	CatalogProducts:      "store",
	UserWishlistProducts: "wishlist",
	AccountProducts:      "account",
	Details:              "account",
	SteamAppNews:         "news",
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

func GOGPagedProducts() []ProductType {
	return []ProductType{
		CatalogPage,
		AccountPage,
		OrderPage,
	}
}

func GOGArrayProducts() []ProductType {
	return []ProductType{
		Licences,
		UserWishlist,
		SteamAppList,
	}
}

func FastPageFetchProducts() []ProductType {
	return []ProductType{
		OrderPage,
	}
}

var gogDetailMainProductTypes = map[ProductType][]ProductType{
	Details: {LicenceProducts, AccountProducts},
	ApiProductsV1: {
		CatalogProducts,
		AccountProducts,
		ApiProductsV2,
	},
	ApiProductsV2: {
		CatalogProducts,
		AccountProducts,
		ApiProductsV2, // includes-games, is-included-in-games, requires-games, is-required-by-games
	},
}

var steamDetailMainProductTypes = map[ProductType][]ProductType{
	//Steam product types are updated on GOG.com store or account product changes
	SteamAppNews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamReviews: {
		CatalogProducts,
		AccountProducts,
	},
	SteamStorePage: {
		CatalogProducts,
		AccountProducts,
	},
}

var pcgwDetailMainProductTypes = map[ProductType][]ProductType{
	//PCGamingWiki product types are updated on GOG.com store or account product changes
	PCGWCargo: {
		CatalogProducts,
		AccountProducts,
	},
	PCGWWikiText: {
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

func PCGWDetailProducts() []ProductType {
	return detailProducts(pcgwDetailMainProductTypes)
}

func MainProductTypes(pt ProductType) []ProductType {
	if IsGOGDetailProduct(pt) {
		return gogMainProductTypes(pt)
	} else if IsSteamDetailProduct(pt) {
		return steamMainProductTypes(pt)
	} else if IsPCGWDetailProduct(pt) {
		return pcgwMainProductTypes(pt)
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

func pcgwMainProductTypes(pt ProductType) []ProductType {
	return pcgwDetailMainProductTypes[pt]
}

func GOGRemoteProducts() []ProductType {
	remote := make([]ProductType, 0)
	remote = append(remote, GOGPagedProducts()...)
	remote = append(remote, GOGArrayProducts()...)
	return append(remote, GOGDetailProducts()...)
}

func SteamRemoteProducts() []ProductType {
	return SteamDetailProducts()
}

func PCGWRemoteProducts() []ProductType {
	return PCGWDetailProducts()
}

func LocalProducts() []ProductType {
	lps := make([]ProductType, 0, len(splitProductTypes))
	for _, spt := range splitProductTypes {
		lps = append(lps, spt)
	}
	lps = append(lps, GOGDetailProducts()...)
	lps = append(lps, SteamDetailProducts()...)
	lps = append(lps, PCGWRemoteProducts()...)

	return lps
}

func RemoteProducts() []ProductType {
	rps := GOGRemoteProducts()
	rps = append(rps, SteamRemoteProducts()...)
	rps = append(rps, PCGWRemoteProducts()...)
	return rps
}

var requireAuth = []ProductType{
	AccountPage,
	UserWishlist,
	Details,
	Licences,
	OrderPage,
}

var splitProductTypes = map[ProductType]ProductType{
	CatalogPage:  CatalogProducts,
	AccountPage:  AccountProducts,
	Licences:     LicenceProducts,
	UserWishlist: UserWishlistProducts,
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
	UserWishlist,
	SteamAppList,
	SteamAppNews,
	SteamReviews,
	SteamStorePage,
	PCGWCargo,
	PCGWWikiText,
}

var supportedImageTypes = map[ProductType][]ImageType{
	CatalogProducts: {Image, Screenshots},
	AccountProducts: {Image},
	ApiProductsV1:   {Screenshots},
	ApiProductsV2:   {Image, Screenshots},
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

// TODO: is this really required? Likely should move to LocalKeyValues
func Cut(ids []string, pt ProductType) error {
	ptDir, err := AbsLocalProductTypeDir(pt)
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
