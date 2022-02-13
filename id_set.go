package vangogh_local_data

import (
	"fmt"
	"github.com/boggydigital/gost"
	"github.com/boggydigital/kvas"
	"sort"
	"strings"
)

const (
	DefaultSort = TitleProperty
	DefaultDesc = false
)

type IdSet struct {
	gost.SortStrSet
}

func NewIdSet() IdSet {
	return IdSet{SortStrSet: gost.NewSortStrSet()}
}

func IdSetFromSlice(ids ...string) IdSet {
	idSet := NewIdSet()
	idSet.Add(ids...)
	return idSet
}

func IdSetFromMap(ids map[string]bool) IdSet {
	idSet := NewIdSet()
	for id := range ids {
		idSet.Add(id)
	}
	return idSet
}

func (is IdSet) Sort(rxa kvas.ReduxAssets, property string, desc bool) []string {
	ips := make(map[string]string, 0)

	for _, id := range is.All() {
		prop, _ := rxa.GetFirstVal(property, id)
		ips[id] = prop
	}

	return is.SortByStrVal(ips, desc)
}

func (is IdSet) AddSet(another IdSet) {
	for _, id := range another.All() {
		is.Add(id)
	}
}

func (is IdSet) AddMap(another map[string]bool) {
	for id := range another {
		is.Add(id)
	}
}

func idSetFromSlugs(slugs []string, rxa kvas.ReduxAssets) (slugId IdSet, err error) {

	if rxa == nil && len(slugs) > 0 {
		rxa, err = ConnectReduxAssets(SlugProperty)
		if err != nil {
			return NewIdSet(), err
		}
	}

	if rxa != nil {
		if err := rxa.IsSupported(SlugProperty); err != nil {
			return NewIdSet(), err
		}
	}

	idSet := NewIdSet()
	for _, slug := range slugs {
		if slug != "" && rxa != nil {
			idSet.AddMap(rxa.Match(map[string][]string{SlugProperty: {slug}}, true))
		}
	}

	return idSet, nil
}

func PropertyListsFromIdSet(
	ids IdSet,
	propertyFilter map[string][]string,
	properties []string,
	rxa kvas.ReduxAssets) (map[string][]string, error) {

	propSet := gost.NewStrSetWith(properties...)
	propSet.Add(TitleProperty)

	if rxa == nil {
		var err error
		rxa, err = ConnectReduxAssets(propSet.All()...)
		if err != nil {
			return nil, err
		}
	}

	itps := make(map[string][]string)

	for _, id := range ids.All() {
		itp, err := propertyListFromId(id, propertyFilter, propSet.All(), rxa)
		if err != nil {
			return itps, err
		}
		for idTitle, props := range itp {
			itps[idTitle] = props
		}
	}

	return itps, nil
}

func propertyListFromId(
	id string,
	propertyFilter map[string][]string,
	properties []string,
	rxa kvas.ReduxAssets) (map[string][]string, error) {

	if err := rxa.IsSupported(properties...); err != nil {
		return nil, err
	}

	title, ok := rxa.GetFirstVal(TitleProperty, id)
	if !ok {
		return nil, nil
	}

	itp := make(map[string][]string)
	idTitle := fmt.Sprintf("%s %s", id, title)
	itp[idTitle] = make([]string, 0)

	sort.Strings(properties)

	for _, prop := range properties {
		if prop == IdProperty ||
			prop == TitleProperty {
			continue
		}
		values, ok := rxa.GetAllValues(prop, id)
		if !ok || len(values) == 0 {
			continue
		}
		filterValues := propertyFilter[prop]

		if len(values) > 1 && IsPropertiesJoinPreferred(prop) {
			joinedValue := strings.Join(values, ",")
			if isPropertyValueFiltered(joinedValue, filterValues) {
				continue
			}
			itp[idTitle] = append(itp[idTitle], fmt.Sprintf("%s:%s", prop, joinedValue))
			continue
		}

		for _, val := range values {
			if isPropertyValueFiltered(val, filterValues) {
				continue
			}
			itp[idTitle] = append(itp[idTitle], fmt.Sprintf("%s:%s", prop, val))
		}
	}

	return itp, nil
}

func isPropertyValueFiltered(value string, filterValues []string) bool {
	value = strings.ToLower(value)
	for _, fv := range filterValues {
		if strings.Contains(value, fv) {
			return false
		}
	}
	// this makes sure we don't filter values if there is no filter
	return len(filterValues) > 0
}
