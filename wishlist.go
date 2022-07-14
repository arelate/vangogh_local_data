package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
)

func initLocalWishlistOperators(mt gog_integration.Media) (*ValueReader, kvas.ReduxAssets, error) {
	vrStoreProducts, err := NewReader(StoreProducts, mt)
	if err != nil {
		return vrStoreProducts, nil, err
	}

	rxa, err := ConnectReduxAssets(WishlistedProperty)

	// don't check err because we're immediately returning it
	return vrStoreProducts, rxa, err
}

func AddToLocalWishlist(
	ids []string,
	mt gog_integration.Media,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	vrStoreProducts, rxa, err := initLocalWishlistOperators(mt)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if !vrStoreProducts.Has(id) {
			if tpw != nil {
				tpw.Increment()
			}
			continue
		}

		if err := vrStoreProducts.CopyToType(id, WishlistProducts, mt); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return processedIds, err
		}

		// remove "false" reduction
		if rxa.HasVal(WishlistedProperty, id, FalseValue) {
			if err := rxa.CutVal(WishlistedProperty, id, FalseValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		if !rxa.HasVal(WishlistedProperty, id, TrueValue) {
			if err := rxa.AddVal(WishlistedProperty, id, TrueValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	return processedIds, nil
}

func RemoveFromLocalWishlist(
	ids []string,
	mt gog_integration.Media,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	vrStoreProducts, rxa, err := initLocalWishlistOperators(mt)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if !vrStoreProducts.Has(id) {
			if tpw != nil {
				tpw.Increment()
			}
			continue
		}

		if rxa.HasVal(WishlistedProperty, id, TrueValue) {
			if err := rxa.CutVal(WishlistedProperty, id, TrueValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		if !rxa.HasVal(WishlistedProperty, id, FalseValue) {
			if err := rxa.AddVal(WishlistedProperty, id, FalseValue); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
		if tpw != nil {
			tpw.Increment()
		}
	}

	err = Cut(processedIds, WishlistProducts, mt)

	// don't check err because we're immediately returning it
	return processedIds, err
}
