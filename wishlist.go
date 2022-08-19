package vangogh_local_data

import (
	"github.com/boggydigital/nod"
)

func AddToLocalWishlist(
	ids []string,
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rxa, err := ConnectReduxAssets(WishlistedProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
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
	tpw nod.TotalProgressWriter) ([]string, error) {

	processedIds := make([]string, 0, len(ids))

	rxa, err := ConnectReduxAssets(WishlistedProperty)
	if err != nil {
		return processedIds, err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
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

	err = Cut(processedIds, UserWishlistProducts)

	// don't check err because we're immediately returning it
	return processedIds, err
}
