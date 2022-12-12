package vangogh_local_data

import (
	"errors"
	"fmt"
	"github.com/boggydigital/kvas"
	"net/url"
)

type UrlProvider interface {
	Url(gogId string) *url.URL
}

func NewUrlProvider(pt ProductType, rxa kvas.ReduxAssets) (UrlProvider, error) {
	if IsGOGProduct(pt) {
		return NewGOGUrlProvider(pt)
	} else if IsSteamProduct(pt) {
		return NewSteamUrlProvider(pt, rxa)
	} else if IsPCGWProduct(pt) {
		return NewPCGWUrlProvider(pt, rxa)
	} else {
		return nil, errors.New(fmt.Sprintf("product type %s is not a url provider", pt))
	}
}
