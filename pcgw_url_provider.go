package vangogh_local_data

import (
	"github.com/arelate/pcgw_integration"
	"github.com/boggydigital/kvas"
	"net/url"
)

type PCGWUrlProvider struct {
	pt  ProductType
	rxa kvas.ReduxAssets
}

func NewPCGWUrlProvider(pt ProductType, rxa kvas.ReduxAssets) (*PCGWUrlProvider, error) {
	if err := rxa.IsSupported(PCGWPageId); err != nil {
		return nil, err
	}

	return &PCGWUrlProvider{
		pt:  pt,
		rxa: rxa,
	}, nil
}

func (pcgwup *PCGWUrlProvider) GOGIdToPCGWPageId(gogId string) string {
	if pageId, ok := pcgwup.rxa.GetFirstVal(PCGWPageId, gogId); ok {
		return pageId
	}
	return ""
}

func (pcgwup *PCGWUrlProvider) Url(gogId string) *url.URL {
	switch pcgwup.pt {
	case PCGWCargo:
		return pcgw_integration.CargoQueryUrl(gogId)
	case PCGWWikiText:
		if pageId := pcgwup.GOGIdToPCGWPageId(gogId); pageId != "" {
			return pcgw_integration.ParseUrl(pageId)
		}
	}
	return nil
}
