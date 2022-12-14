package vangogh_local_data

import (
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/boggydigital/kvas"
	"net/url"
)

type HLTBUrlProvider struct {
	pt  ProductType
	rxa kvas.ReduxAssets
}

func NewHLTBUrlProvider(pt ProductType, rxa kvas.ReduxAssets) (*HLTBUrlProvider, error) {
	if err := rxa.IsSupported(HowLongToBeatIdProperty); err != nil {
		return nil, err
	}

	return &HLTBUrlProvider{
		pt:  pt,
		rxa: rxa,
	}, nil
}

func (hup *HLTBUrlProvider) GOGIdToHLTBId(gogId string) string {
	if hltbId, ok := hup.rxa.GetFirstVal(HowLongToBeatIdProperty, gogId); ok {
		return hltbId
	}
	return ""
}

func (hup *HLTBUrlProvider) Url(gogId string) *url.URL {
	switch hup.pt {
	case HLTBRootPage:
		return hltb_integration.RootUrl()
		//default:
		//	if appId := sup.GOGIdToSteamAppId(gogId); appId > 0 {
		//		if sug, ok := steamProductTypeUrlGetters[sup.pt]; ok {
		//			return sug(appId)
		//		}
		//	}
	}
	return nil
}
