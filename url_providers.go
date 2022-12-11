package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
	"net/url"
)

type UrlProvider interface {
	Url(gogId string) *url.URL
}

func NewUrlProvider(pt ProductType, rxa kvas.ReduxAssets) UrlProvider {
	if IsGOGDetailProduct(pt) {
		return &GOGUrlProvider{pt: pt}
	} else if IsSteamDetailProduct(pt) {
		return &SteamUrlProvider{
			pt:  pt,
			rxa: rxa,
		}
	}
	return nil
}
