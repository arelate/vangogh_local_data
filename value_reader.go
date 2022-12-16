package vangogh_local_data

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/southern_light/gog_integration"
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/boggydigital/kvas"
	"golang.org/x/net/html"
	"io"
)

type ValueReader struct {
	productType ProductType
	valueSet    kvas.KeyValues
}

func NewReader(pt ProductType) (*ValueReader, error) {
	dst, err := AbsLocalProductTypeDir(pt)
	if err != nil {
		return nil, err
	}

	vs, err := kvas.ConnectLocal(dst, kvas.JsonExt)
	if err != nil {
		return nil, err
	}

	vr := &ValueReader{
		productType: pt,
		valueSet:    vs,
	}

	return vr, nil
}

func (vr *ValueReader) readValue(id string, val interface{}) error {
	spReadCloser, err := vr.valueSet.Get(id)
	if err != nil {
		return err
	}

	if spReadCloser == nil {
		return nil
	}

	defer spReadCloser.Close()

	if err := json.NewDecoder(spReadCloser).Decode(val); err != nil {
		return err
	}

	return nil
}

func (vr *ValueReader) Keys() []string {
	return vr.valueSet.Keys()
}

func (vr *ValueReader) Has(id string) bool {
	return vr.valueSet.Has(id)
}

func (vr *ValueReader) Set(id string, data io.Reader) error {
	return vr.valueSet.Set(id, data)
}

func (vr *ValueReader) Cut(id string) (bool, error) {
	return vr.valueSet.Cut(id)
}

func (vr *ValueReader) CreatedAfter(timestamp int64) []string {
	return vr.valueSet.CreatedAfter(timestamp)
}

func (vr *ValueReader) ModifiedAfter(timestamp int64, excludeCreated bool) []string {
	return vr.valueSet.ModifiedAfter(timestamp, excludeCreated)
}

func (vr *ValueReader) IsModifiedAfter(id string, timestamp int64) bool {
	return vr.valueSet.IsModifiedAfter(id, timestamp)
}

func (vr *ValueReader) CatalogProduct(id string) (catalogProduct *gog_integration.CatalogProduct, err error) {
	err = vr.readValue(id, &catalogProduct)
	return catalogProduct, err
}

func (vr *ValueReader) AccountProduct(id string) (accountProduct *gog_integration.AccountProduct, err error) {
	err = vr.readValue(id, &accountProduct)
	return accountProduct, err
}

func (vr *ValueReader) Details(id string) (details *gog_integration.Details, err error) {
	err = vr.readValue(id, &details)
	return details, err
}

func (vr *ValueReader) ApiProductV1(id string) (apiProductV1 *gog_integration.ApiProductV1, err error) {
	err = vr.readValue(id, &apiProductV1)
	return apiProductV1, err
}

func (vr *ValueReader) ApiProductV2(id string) (apiProductV2 *gog_integration.ApiProductV2, err error) {
	err = vr.readValue(id, &apiProductV2)
	return apiProductV2, err
}

func (vr *ValueReader) CatalogPage(page string) (catalogPage *gog_integration.CatalogPage, err error) {
	err = vr.readValue(page, &catalogPage)
	return catalogPage, err
}

func (vr *ValueReader) AccountPage(page string) (accountPage *gog_integration.AccountPage, err error) {
	err = vr.readValue(page, &accountPage)
	return accountPage, err
}

func (vr *ValueReader) UserWishlist() (userWishlist *gog_integration.UserWishlist, err error) {
	err = vr.readValue(UserWishlist.String(), &userWishlist)
	return userWishlist, err
}

func (vr *ValueReader) Licences() (licences *gog_integration.Licences, err error) {
	err = vr.readValue(Licences.String(), &licences)
	return licences, err
}

func (vr *ValueReader) OrderPage(page string) (orderPage *gog_integration.OrderPage, err error) {
	err = vr.readValue(page, &orderPage)
	return orderPage, err
}

func (vr *ValueReader) Order(id string) (order *gog_integration.Order, err error) {
	err = vr.readValue(id, &order)
	return order, err
}

func (vr *ValueReader) SteamAppList() (steamAppListResponse *steam_integration.GetAppListResponse, err error) {
	err = vr.readValue(SteamAppList.String(), &steamAppListResponse)
	return steamAppListResponse, err
}

func (vr *ValueReader) SteamGetAppNewsResponse(id string) (steamAppNewsResponse *steam_integration.GetNewsForAppResponse, err error) {
	err = vr.readValue(id, &steamAppNewsResponse)
	return steamAppNewsResponse, err
}

func (vr *ValueReader) SteamAppReviews(id string) (steamAppReviews *steam_integration.AppReviews, err error) {
	err = vr.readValue(id, &steamAppReviews)
	return steamAppReviews, err
}

func (vr *ValueReader) UserWishlistProduct(id string) (userWishlistProduct string, err error) {
	userWishlistProduct, err = id, nil
	return userWishlistProduct, err
}

// TODO: redo this similarly to HLTBRootPage (also an HTML doc)
// SteamStorePage reads HTML content of the locally downloaded Steam store page and
// return an HTML document for traversal. This approach is different from other data
// types that have defined schemas.
func (vr *ValueReader) SteamStorePage(id string) (steamStorePage *html.Node, err error) {
	spReadCloser, err := vr.valueSet.Get(id)
	if err != nil {
		return nil, err
	}

	if spReadCloser == nil {
		return nil, nil
	}

	defer spReadCloser.Close()

	return html.Parse(spReadCloser)
}

func (vr *ValueReader) HLTBRootPage() (*hltb_integration.RootPage, error) {
	spReadCloser, err := vr.valueSet.Get(HLTBRootPage.String())
	if err != nil {
		return nil, err
	}

	if spReadCloser == nil {
		return nil, nil
	}
	defer spReadCloser.Close()

	doc, err := html.Parse(spReadCloser)

	return &hltb_integration.RootPage{Doc: doc}, err
}

func (vr *ValueReader) HLTBData(id string) (data *hltb_integration.Data, err error) {
	err = vr.readValue(id, &data)
	return data, err
}

func (vr *ValueReader) PCGWCargo(id string) (cargo *pcgw_integration.Cargo, err error) {
	err = vr.readValue(id, &cargo)
	return cargo, err
}

func (vr *ValueReader) PCGWExternalLinks(id string) (pel *pcgw_integration.ParseExternalLinks, err error) {
	err = vr.readValue(id, &pel)
	return pel, err
}

func (vr *ValueReader) ReadValue(key string) (interface{}, error) {
	switch vr.productType {
	case CatalogProducts:
		return vr.CatalogProduct(key)
	case AccountProducts:
		return vr.AccountProduct(key)
	case Details:
		return vr.Details(key)
	case ApiProductsV1:
		return vr.ApiProductV1(key)
	case ApiProductsV2:
		return vr.ApiProductV2(key)
	case Orders:
		return vr.Order(key)
	case CatalogPage:
		return vr.CatalogPage(key)
	case AccountPage:
		return vr.AccountPage(key)
	case UserWishlist:
		return vr.UserWishlist()
	case UserWishlistProducts:
		return vr.UserWishlistProduct(key)
	case OrderPage:
		return vr.OrderPage(key)
	case Licences:
		return vr.Licences()
	case SteamAppNews:
		return vr.SteamGetAppNewsResponse(key)
	case SteamReviews:
		return vr.SteamAppReviews(key)
	case SteamStorePage:
		return vr.SteamStorePage(key)
	case SteamAppList:
		return vr.SteamAppList()
	case PCGWCargo:
		return vr.PCGWCargo(key)
	case PCGWExternalLinks:
		return vr.PCGWExternalLinks(key)
	case HLTBRootPage:
		return vr.HLTBRootPage()
	case HLTBData:
		return vr.HLTBData(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", vr.productType)
	}
}

func (vr *ValueReader) ProductType() ProductType {
	return vr.productType
}

func (vr *ValueReader) ProductsGetter(page string) (productsGetter gog_integration.ProductsGetter, err error) {
	switch vr.productType {
	case CatalogPage:
		productsGetter, err = vr.CatalogPage(page)
	case AccountPage:
		productsGetter, err = vr.AccountPage(page)
	case UserWishlist:
		productsGetter, err = vr.UserWishlist()
	case Licences:
		productsGetter, err = vr.Licences()
	case OrderPage:
		productsGetter, err = vr.OrderPage(page)
	default:
		err = fmt.Errorf("%s doesn't implement ProductGetter interface", vr.productType)
	}
	return productsGetter, err
}

func (vr *ValueReader) IndexCurrentModTime() (int64, error) {
	return vr.valueSet.IndexCurrentModTime()
}

func (vr *ValueReader) CurrentModTime(id string) (int64, error) {
	return vr.valueSet.CurrentModTime(id)
}
