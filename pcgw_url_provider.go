package vangogh_local_data

import (
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/boggydigital/kvas"
	"net/url"
)

type PCGWUrlProvider struct {
	pt  ProductType
	rxa kvas.ReduxAssets
}

func NewPCGWUrlProvider(pt ProductType, rxa kvas.ReduxAssets) (*PCGWUrlProvider, error) {
	if err := rxa.IsSupported(PCGWPageIdProperty); err != nil {
		return nil, err
	}

	return &PCGWUrlProvider{
		pt:  pt,
		rxa: rxa,
	}, nil
}

func (pcgwup *PCGWUrlProvider) GOGIdToPCGWPageId(gogId string) string {
	if pageId, ok := pcgwup.rxa.GetFirstVal(PCGWPageIdProperty, gogId); ok {
		return pageId
	}
	return ""
}

func (pcgwup *PCGWUrlProvider) Url(gogId string) *url.URL {
	switch pcgwup.pt {
	case PCGWPageIdSteamAppId:
		return pcgw_integration.SteamAppIdCargoQueryUrl(gogId)
	case PCGWEngine:
		if pageId := pcgwup.GOGIdToPCGWPageId(gogId); pageId != "" {
			return pcgw_integration.EngineCargoQueryUrl(pageId)
		}
	case PCGWExternalLinks:
		if pageId := pcgwup.GOGIdToPCGWPageId(gogId); pageId != "" {
			return pcgw_integration.ParseExternalLinksUrl(pageId)
		}
	}
	return nil
}
