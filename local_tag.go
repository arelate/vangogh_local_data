package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
)

func addLocalTag(id, tag string, rxa kvas.ReduxAssets, tpw nod.TotalProgressWriter) error {
	if !rxa.HasVal(LocalTagsProperty, id, tag) {
		if err := rxa.AddVal(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}
	nod.Increment(tpw)
	return nil
}

func removeLocalTag(id, tag string, rxa kvas.ReduxAssets, tpw nod.TotalProgressWriter) error {
	if rxa.HasVal(LocalTagsProperty, id, tag) {
		if err := rxa.CutVal(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}

	nod.Increment(tpw)
	return nil
}

func AddLocalTags(ids, tags []string, tpw nod.TotalProgressWriter) error {
	rxa, err := ConnectReduxAssets(LocalTagsProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {
			if err := addLocalTag(id, tag, rxa, tpw); err != nil {
				return err
			}
		}
	}

	return nil
}

func RemoveLocalTags(ids, tags []string, tpw nod.TotalProgressWriter) error {
	rxa, err := ConnectReduxAssets(LocalTagsProperty)
	if err != nil {
		return err
	}

	nod.TotalInt(tpw, len(ids)*len(tags))

	for _, id := range ids {
		for _, tag := range tags {
			if err := removeLocalTag(id, tag, rxa, tpw); err != nil {
				return err
			}
		}
	}

	return nil
}

func DiffLocalTags(id string, newTags []string) (add []string, rem []string, err error) {
	return diffTagProperty(LocalTagsProperty, id, newTags)
}
