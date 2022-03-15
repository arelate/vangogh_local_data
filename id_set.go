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

type idPropertyTitle struct {
	id       string
	property string
	title    string
}

type IdSet []idPropertyTitle

func (is IdSet) Len() int {
	return len(is)
}

func (is IdSet) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func (is IdSet) Less(i, j int) bool {
	if is[i].property == is[j].property {
		return is[i].title < is[j].title
	} else {
		return is[i].property < is[j].property
	}
}

func NewIdSet() IdSet {
	return make(IdSet, 0)
}

func IdSetFromSlice(keys ...string) IdSet {
	idSet := make(IdSet, 0, len(keys))
	for _, id := range keys {
		idSet = append(idSet, idPropertyTitle{id: id})
	}
	return idSet
}

func IdSetFromMap(keys map[string]bool) IdSet {
	idSet := make(IdSet, 0, len(keys))
	for id := range keys {
		idSet = append(idSet, idPropertyTitle{id: id})
	}
	return idSet
}

func (is IdSet) All() []string {
	keys := make([]string, 0, len(is))

	for i := 0; i < len(is); i++ {
		keys = append(keys, is[i].id)
	}

	return keys
}

func (is IdSet) Sort(rxa kvas.ReduxAssets, property string, desc bool) []string {
	for _, ips := range is {
		ips.property, _ = rxa.GetFirstVal(property, ips.id)
		if property != TitleProperty {
			ips.title, _ = rxa.GetFirstVal(TitleProperty, ips.id)
		}
	}

	var sortInterface sort.Interface = is
	if desc {
		sortInterface = sort.Reverse(sortInterface)
	}

	sort.Sort(sortInterface)

	return is.All()
}

func (is IdSet) Add(keys ...string) {
	for _, id := range keys {
		is = append(is, idPropertyTitle{id: id})
	}
}

func (is IdSet) Remove(keys ...string) {
	for _, id := range keys {
		index := -1
		for i, ipt := range is {
			if ipt.id == id {
				index = i
				break
			}
		}
		if index >= 0 {
			is[index] = is[len(is)-1]
			is = is[:len(is)-1]
		}
	}
}

func (is IdSet) AddSet(another IdSet) {
	for _, ipt := range another {
		is = append(is, ipt)
	}
}

func (is IdSet) AddMap(another map[string]bool) {
	for id := range another {
		is = append(is, idPropertyTitle{id: id})
	}
}

func (is IdSet) Has(id string) bool {
	for _, ipt := range is {
		if ipt.id == id {
			return true
		}
	}
	return false
}

func (is IdSet) Except(other IdSet) []string {
	result := make([]string, 0, len(is))
	for _, ipt := range is {
		if other.Has(ipt.id) {
			continue
		}
		result = append(result, ipt.id)
	}
	return result
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
	idSet IdSet,
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

	for _, ipt := range idSet {
		itp, err := propertyListFromId(ipt.id, propertyFilter, propSet.All(), rxa)
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
