package vangogh_data

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/gog_atu"
	"github.com/boggydigital/kvas"
)

type ValueReader struct {
	productType ProductType
	mediaType   gog_atu.Media
	valueSet    kvas.KeyValues
}

func NewReader(pt ProductType, mt gog_atu.Media) (*ValueReader, error) {
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

func (vr *ValueReader) StoreProduct(id string) (storeProduct *gog_atu.StoreProduct, err error) {
	err = vr.readValue(id, &storeProduct)
	return storeProduct, err
}

func (vr *ValueReader) AccountProduct(id string) (accountProduct *gog_atu.AccountProduct, err error) {
	err = vr.readValue(id, &accountProduct)
	return accountProduct, err
}

func (vr *ValueReader) WishlistProduct(id string) (wishlistProduct *gog_atu.StoreProduct, err error) {
	err = vr.readValue(id, &wishlistProduct)
	return wishlistProduct, err
}

func (vr *ValueReader) Details(id string) (details *gog_atu.Details, err error) {
	err = vr.readValue(id, &details)
	return details, err
}

func (vr *ValueReader) ApiProductV1(id string) (apiProductV1 *gog_atu.ApiProductV1, err error) {
	err = vr.readValue(id, &apiProductV1)
	return apiProductV1, err
}

func (vr *ValueReader) ApiProductV2(id string) (apiProductV2 *gog_atu.ApiProductV2, err error) {
	err = vr.readValue(id, &apiProductV2)
	return apiProductV2, err
}

func (vr *ValueReader) StorePage(page string) (storePage *gog_atu.StorePage, err error) {
	err = vr.readValue(page, &storePage)
	return storePage, err
}

func (vr *ValueReader) AccountPage(page string) (accountPage *gog_atu.AccountPage, err error) {
	err = vr.readValue(page, &accountPage)
	return accountPage, err
}

func (vr *ValueReader) WishlistPage(page string) (wishlistPage *gog_atu.WishlistPage, err error) {
	err = vr.readValue(page, &wishlistPage)
	return wishlistPage, err
}

func (vr *ValueReader) Licences(id string) (licences *gog_atu.Licences, err error) {
	err = vr.readValue(id, &licences)
	return licences, err
}

func (vr *ValueReader) OrderPage(page string) (orderPage *gog_atu.OrderPage, err error) {
	err = vr.readValue(page, &orderPage)
	return orderPage, err
}

func (vr *ValueReader) Order(id string) (order *gog_atu.Order, err error) {
	err = vr.readValue(id, &order)
	return order, err
}

func (vr *ValueReader) ReadValue(key string) (interface{}, error) {
	switch vr.productType {
	case StoreProducts:
		return vr.StoreProduct(key)
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
	case AccountPage:
		return vr.AccountPage(key)
	case WishlistPage:
		return vr.WishlistPage(key)
	case OrderPage:
		return vr.OrderPage(key)
	case Licences:
		return vr.Licences(key)
	default:
		return nil, fmt.Errorf("vangogh_values: cannot create %s value", vr.productType)
	}
}

func (vr *ValueReader) ProductType() ProductType {
	return vr.productType
}

func (vr *ValueReader) ProductsGetter(page string) (productsGetter gog_atu.ProductsGetter, err error) {
	switch vr.productType {
	case StorePage:
		productsGetter, err = vr.StorePage(page)
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

func (vr *ValueReader) CopyToType(id string, toPt ProductType, toMt gog_atu.Media) error {

	if !IsCopySupported(vr.productType, toPt) {
		return fmt.Errorf("vangogh_values: copy type from %s to %s is unsupported", vr.productType, toPt)
	}
	if vr.mediaType != toMt {
		return fmt.Errorf("vangogh_values: copy media from %s to %s is unsupported", vr.mediaType, toMt)
	}

	toDir, err := AbsLocalProductTypeDir(toPt, toMt)
	if err != nil {
		return err
	}

	vsToType, err := kvas.ConnectLocal(toDir, kvas.JsonExt)
	if err != nil {
		return nil
	}

	rc, err := vr.valueSet.Get(id)
	if err != nil {
		return err
	}

	defer rc.Close()

	if err := vsToType.Set(id, rc); err != nil {
		return err
	}

	return nil
}