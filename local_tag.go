package vangogh_local_data

import "github.com/boggydigital/nod"

func AddLocalTag(ids []string, tagName string, tpw nod.TotalProgressWriter) error {
	rxa, err := ConnectReduxAssets(LocalTagsProperty)
	if err != nil {
		return err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if !rxa.HasVal(LocalTagsProperty, id, tagName) {
			if err := rxa.AddVal(LocalTagsProperty, id, tagName); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return err
			}
		}

		if tpw != nil {
			tpw.Increment()
		}
	}

	return nil
}

func RemoveLocalTag(ids []string, tagName string, tpw nod.TotalProgressWriter) error {
	rxa, err := ConnectReduxAssets(LocalTagsProperty)
	if err != nil {
		return err
	}

	if tpw != nil {
		tpw.TotalInt(len(ids))
	}

	for _, id := range ids {
		if rxa.HasVal(LocalTagsProperty, id, tagName) {
			if err := rxa.CutVal(LocalTagsProperty, id, tagName); err != nil {
				if tpw != nil {
					tpw.Increment()
				}
				return err
			}
		}

		if tpw != nil {
			tpw.Increment()
		}
	}

	return nil
}
