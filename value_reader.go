package vangogh_local_data

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/gog_integration"
	"github.com/arelate/steam_integration"
	"github.com/boggydigital/kvas"
	"golang.org/x/net/html"
	"io"
)

type ValueReader struct {
	productType ProductType
	mediaType   gog_integration.Media
	valueSet    kvas.KeyValues
}

func NewReader(pt ProductType, mt gog_integration.Media) (*ValueReader, error) {
	dst, err := AbsLocalProductTypeDir(pt, mt)
	if err != nil {
		return nil, err
	}

	vs, err := kvas.ConnectLocal(dst, kvas.JsonExt)
	if err != nil {
		return nil, err
	}

	vr := &ValueReader{
		productType: pt,
		mediaType:   mt,
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

func (vr *ValueReader) StoreProduct(id string) (storeProduct *gog_integration.StoreProduct, err error) {
	err = vr.readValue(id, &storeProduct)
	return storeProduct, err
}

func (vr *ValueReader) CatalogProduct(id string) (catalogProduct *gog_integration.CatalogProduct, err error) {
	err = vr.readValue(id, &catalogProduct)
	return catalogProduct, err
}

func (vr *ValueReader) AccountProduct(id string) (accountProduct *gog_integration.AccountProduct, err error) {
	err = vr.readValue(id, &accountProduct)
	return accountProduct, err
}

func (vr *ValueReader) WishlistProduct(id string) (wishlistProduct *gog_integration.StoreProduct, err error) {
	err = vr.readValue(id, &wishlistProduct)
	return wishlistProduct, err
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

func (vr *ValueReader) StorePage(page string) (storePage *gog_integration.StorePage, err error) {
	err = vr.readValue(page, &storePage)
	return storePage, err
}

func (vr *ValueReader) CatalogPage(page string) (catalogPage *gog_integration.CatalogPage, err error) {
	err = vr.readValue(page, &catalogPage)
	return catalogPage, err
}

func (vr *ValueReader) AccountPage(page string) (accountPage *gog_integration.AccountPage, err error) {
	err = vr.readValue(page, &accountPage)
	return accountPage, err
}

func (vr *ValueReader) WishlistPage(page string) (wishlistPage *gog_integration.WishlistPage, err error) {
	err = vr.readValue(page, &wishlistPage)
	return wishlistPage, err
}

func (vr *ValueReader) Licences(id string) (licences *gog_integration.Licences, err error) {
	err = vr.readValue(id, &licences)
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

func (vr *ValueReader) SteamGetAppListResponse() (steamAppListResponse *steam_integration.GetAppListResponse, err error) {
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

//SteamStorePage reads HTML content of the locally downloaded Steam store page and
//return an HTML document for traversal. This approach is different from other data
//types that have defined schemas.
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

func (vr *ValueReader) ReadValue(key string) (interface{}, error) {
	switch vr.productType {
	case StoreProducts:
		return vr.StoreProduct(key)
	case CatalogProducts:
		return vr.CatalogProduct(key)
	case AccountProducts:
		return vr.AccountProduct(key)
	case WishlistProducts:
		return vr.WishlistProduct(key)
	case Details:
		return vr.Details(key)
	case ApiProductsV1:
		return vr.ApiProductV1(key)
	case ApiProductsV2:
		return vr.ApiProductV2(key)
	case Orders:
		return vr.Order(key)
	case StorePage:
		return vr.StorePage(key)
	case CatalogPage:
		return vr.CatalogPage(key)
	case AccountPage:
		return vr.AccountPage(key)
	case WishlistPage:
		return vr.WishlistPage(key)
	case OrderPage:
		return vr.OrderPage(key)
	case Licences:
		return vr.Licences(key)
	case SteamAppNews:
		return vr.SteamGetAppNewsResponse(key)
	case SteamReviews:
		return vr.SteamAppReviews(key)
	case SteamStorePage:
		return vr.SteamStorePage(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", vr.productType)
	}
}

func (vr *ValueReader) ProductType() ProductType {
	return vr.productType
}

func (vr *ValueReader) ProductsGetter(page string) (productsGetter gog_integration.ProductsGetter, err error) {
	switch vr.productType {
	case StorePage:
		productsGetter, err = vr.StorePage(page)
	case CatalogPage:
		productsGetter, err = vr.CatalogPage(page)
	case AccountPage:
		productsGetter, err = vr.AccountPage(page)
	case WishlistPage:
		productsGetter, err = vr.WishlistPage(page)
	case Licences:
		productsGetter, err = vr.Licences(page)
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
