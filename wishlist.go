package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
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
	mt gog_integration.Media) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	vrStoreProducts, rxa, err := initLocalWishlistOperators(mt)
	if err != nil {
		return processedIds, err
	}

	for _, id := range ids {
		if !vrStoreProducts.Has(id) {
			continue
		}

		if err := vrStoreProducts.CopyToType(id, WishlistProducts, mt); err != nil {
			return processedIds, err
		}

		// remove "false" reduction
		if rxa.HasVal(WishlistedProperty, id, FalseValue) {
			if err := rxa.CutVal(WishlistedProperty, id, FalseValue); err != nil {
				return processedIds, err
			}
		}

		if !rxa.HasVal(WishlistedProperty, id, TrueValue) {
			if err := rxa.AddVal(WishlistedProperty, id, TrueValue); err != nil {
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
	}

	return processedIds, nil
}

func RemoveFromLocalWishlist(
	ids []string,
	mt gog_integration.Media) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	vrStoreProducts, rxa, err := initLocalWishlistOperators(mt)
	if err != nil {
		return processedIds, err
	}

	for _, id := range ids {
		if !vrStoreProducts.Has(id) {
			continue
		}

		if rxa.HasVal(WishlistedProperty, id, TrueValue) {
			if err := rxa.CutVal(WishlistedProperty, id, TrueValue); err != nil {
				return processedIds, err
			}
		}

		if !rxa.HasVal(WishlistedProperty, id, FalseValue) {
			if err := rxa.AddVal(WishlistedProperty, id, FalseValue); err != nil {
				return processedIds, err
			}
		}

		processedIds = append(processedIds, id)
	}

	err = Cut(processedIds, WishlistProducts, mt)

	// don't check err because we're immediately returning it
	return processedIds, err
}
