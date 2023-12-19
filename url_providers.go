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

func NewUrlProvider(pt ProductType, rdx kvas.ReadableRedux) (UrlProvider, error) {
	if IsGOGProduct(pt) {
		return NewGOGUrlProvider(pt)
	} else if IsSteamProduct(pt) {
		return NewSteamUrlProvider(pt, rdx)
	} else if IsPCGWProduct(pt) {
		return NewPCGWUrlProvider(pt, rdx)
	} else if IsHLTBProduct(pt) {
		return NewHLTBUrlProvider(pt, rdx)
	} else if IsProtonDBProduct(pt) {
		return NewSteamUrlProvider(pt, rdx)
	} else {
		return nil, errors.New(fmt.Sprintf("product type %s is not a url provider", pt))
	}
}
