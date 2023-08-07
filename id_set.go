package vangogh_local_data

import (
	"fmt"
	"github.com/boggydigital/kvas"
	"golang.org/x/exp/maps"
	"sort"
	"strings"
)

const (
	DefaultSort = TitleProperty
	DefaultDesc = false
)

func idSetFromSlugs(slugs []string, rxa kvas.ReduxAssets) (map[string]bool, error) {

	var err error
	if rxa == nil && len(slugs) > 0 {
		rxa, err = ConnectReduxAssets(SlugProperty)
		if err != nil {
			return map[string]bool{}, err
		}
	}

	if rxa != nil {
		if err := rxa.IsSupported(SlugProperty); err != nil {
			return map[string]bool{}, err
		}
	}

	idSet := make(map[string]bool)
	for _, slug := range slugs {
		if slug != "" && rxa != nil {
			for id := range rxa.Match(map[string][]string{SlugProperty: {slug}}, true, false) {
				idSet[id] = true
			}
		}
	}

	return idSet, nil
}

func PropertyListsFromIdSet(
	idSet map[string]bool,
	propertyFilter map[string][]string,
	properties []string,
	rxa kvas.ReduxAssets) (map[string][]string, error) {

	propSet := make(map[string]bool)
	for _, p := range properties {
		propSet[p] = true
	}
	propSet[TitleProperty] = true

	if rxa == nil {
		var err error
		rxa, err = ConnectReduxAssets(maps.Keys(propSet)...)
		if err != nil {
			return nil, err
		}
	}

	itps := make(map[string][]string)

	for id := range idSet {
		itp, err := propertyListFromId(id, propertyFilter, maps.Keys(propSet), rxa)
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

		//if len(values) > 1 && IsPropertiesJoinPreferred(prop) {
		//	joinedValue := strings.Join(values, ",")
		//	if isPropertyValueFiltered(joinedValue, filterValues) {
		//		continue
		//	}
		//	itp[idTitle] = append(itp[idTitle], fmt.Sprintf("%s:%s", prop, joinedValue))
		//	continue
		//}

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
