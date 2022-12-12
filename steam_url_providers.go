package vangogh_local_data

import (
	"github.com/arelate/steam_integration"
	"github.com/boggydigital/kvas"
	"net/url"
	"strconv"
)

var steamProductTypeUrlGetters = map[ProductType]func(uint32) *url.URL{
	SteamAppNews:   steam_integration.NewsForAppUrl,
	SteamReviews:   steam_integration.AppReviewsUrl,
	SteamStorePage: steam_integration.StorePageUrl,
}

type SteamUrlProvider struct {
	pt  ProductType
	rxa kvas.ReduxAssets
}

func NewSteamUrlProvider(pt ProductType, rxa kvas.ReduxAssets) (*SteamUrlProvider, error) {
	if err := rxa.IsSupported(SteamAppIdProperty); err != nil {
		return nil, err
	}

	return &SteamUrlProvider{
		pt:  pt,
		rxa: rxa,
	}, nil
}

func (sup *SteamUrlProvider) GOGIdToSteamAppId(gogId string) uint32 {
	if appIdStr, ok := sup.rxa.GetFirstVal(SteamAppIdProperty, gogId); ok {
		if appId, err := strconv.ParseUint(appIdStr, 10, 32); err == nil {
			return uint32(appId)
		}
	}
	return 0
}

func (sup *SteamUrlProvider) Url(gogId string) *url.URL {
	switch sup.pt {
	case SteamAppList:
		return steam_integration.AppListUrl()
	default:
		if appId := sup.GOGIdToSteamAppId(gogId); appId > 0 {
			if sug, ok := steamProductTypeUrlGetters[sup.pt]; ok {
				return sug(appId)
			}
		}
	}
	return nil
}
