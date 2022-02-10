package vangogh_data

import (
	"fmt"
	"github.com/arelate/gog_atu"
	"net/url"
)

func ImagePropertyUrls(imageIds []string, it ImageType) ([]*url.URL, error) {
	urls := make([]*url.URL, 0, len(imageIds))

	//	var getUrl func(string) (*url.URL, error)
	var ext string
	switch it {
	case Logo:
		// transparency doesn't look right in .jpg variants
		fallthrough
	case Icon:
		// transparency doesn't look right in .jpg variants
		ext = gog_atu.PngExt
	case BoxArt:
		fallthrough
	case Image:
		fallthrough
	case Background:
		fallthrough
	case GalaxyBackground:
		fallthrough
	case Screenshots:
		ext = gog_atu.JpgExt
	}

	for _, imageId := range imageIds {
		if imageId == "" {
			continue
		}
		imageUrl, err := gog_atu.ImageUrl(imageId, ext)
		if err != nil {
			return urls, err
		}
		urls = append(urls, imageUrl)
	}

	return urls, nil
}

type DefaultProductUrl func(string, gog_atu.Media) *url.URL

var defaultProductUrls = map[ProductType]DefaultProductUrl{
	StorePage:     gog_atu.DefaultStorePageUrl,
	AccountPage:   gog_atu.DefaultAccountPageUrl,
	WishlistPage:  gog_atu.DefaultWishlistPageUrl,
	Details:       gog_atu.DetailsUrl,
	ApiProductsV1: gog_atu.ApiProductV1Url,
	ApiProductsV2: gog_atu.ApiProductV2Url,
	Licences:      gog_atu.DefaultLicencesUrl,
	OrderPage:     gog_atu.DefaultOrdersPageUrl,
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
