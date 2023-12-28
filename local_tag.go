package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
)

func addLocalTag(id, tag string, rdx kvas.WriteableRedux, tpw nod.TotalProgressWriter) error {
	if err := rdx.MustHave(LocalTagsProperty); err != nil {
		return err
	}

	if !rdx.HasValue(LocalTagsProperty, id, tag) {
		if err := rdx.AddValues(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}
	nod.Increment(tpw)
	return nil
}

func removeLocalTag(id, tag string, rdx kvas.WriteableRedux, tpw nod.TotalProgressWriter) error {
	if err := rdx.MustHave(LocalTagsProperty); err != nil {
		return err
	}

	if rdx.HasValue(LocalTagsProperty, id, tag) {
		if err := rdx.CutValues(LocalTagsProperty, id, tag); err != nil {
			nod.Increment(tpw)
			return err
		}
	}

	nod.Increment(tpw)
	return nil
}

func AddLocalTags(ids, tags []string, tpw nod.TotalProgressWriter) error {
	rxa, err := NewReduxWriter(LocalTagsProperty)
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
	rxa, err := NewReduxWriter(LocalTagsProperty)
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
