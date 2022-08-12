package vangogh_local_data

import (
	"fmt"
	"github.com/arelate/gog_integration"
	"github.com/arelate/steam_integration"
	"net/url"
)

func ImagePropertyUrls(imageIds []string, it ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(imageIds))

	//	var getUrl func(string) (*url.URL, error)
	var ext string
	switch it {
	case Image:
		fallthrough
	case Screenshots:
		ext = gog_integration.JpgExt
	}

	for _, imageId := range imageIds {
		if imageId == "" {
			continue
		}
		imageUrl, err := gog_integration.ImageUrl(imageId, ext)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}

type DefaultProductUrl func(string, gog_integration.Media) *url.URL

var defaultProductUrls = map[ProductType]DefaultProductUrl{
	StorePage:     gog_integration.DefaultStorePageUrl,
	CatalogPage:   gog_integration.DefaultCatalogPageUrl,
	AccountPage:   gog_integration.DefaultAccountPageUrl,
	WishlistPage:  gog_integration.DefaultWishlistPageUrl,
	UserWishlist:  gog_integration.DefaultUserWishlistUrl,
	Details:       gog_integration.DetailsUrl,
	ApiProductsV1: gog_integration.ApiProductV1Url,
	ApiProductsV2: gog_integration.ApiProductV2Url,
	Licences:      gog_integration.DefaultLicencesUrl,
	OrderPage:     gog_integration.DefaultOrdersPageUrl,
	SteamAppList:  DefaultSteamAppListUrl,
	// steam data types typically require app level transformation of GOG.com id -> Steam AppID
	SteamAppNews:   nil,
	SteamReviews:   nil,
	SteamStorePage: nil,
}

func RemoteProductsUrl(pt ProductType) (ptUrl DefaultProductUrl, err error) {
	if !IsValidProductType(pt) {
		return nil, fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}

	ptUrl, ok := defaultProductUrls[pt]
	if !ok {
		err = fmt.Errorf("vangogh_urls: no remote source for %s\n", pt)
	}

	return ptUrl, err
}

// DefaultSteamAppListUrl is a vangogh_local_data specific wrapper of steam_integration URL func
func DefaultSteamAppListUrl(_ string, _ gog_integration.Media) *url.URL {
	return steam_integration.AppListUrl()
}
