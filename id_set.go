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

type IdSet struct {
	ipt []idPropertyTitle
}

func (is *IdSet) Len() int {
	return len(is.ipt)
}

func (is *IdSet) Swap(i, j int) {
	is.ipt[i], is.ipt[j] = is.ipt[j], is.ipt[i]
}

func (is *IdSet) Less(i, j int) bool {
	if is.ipt[i].property == is.ipt[j].property {
		return is.ipt[i].title < is.ipt[j].title
	} else {
		return is.ipt[i].property < is.ipt[j].property
	}
}

func NewIdSet() *IdSet {
	return &IdSet{ipt: make([]idPropertyTitle, 0)}
}

func IdSetFromSlice(keys ...string) *IdSet {
	idSet := NewIdSet()
	for _, id := range keys {
		idSet.ipt = append(idSet.ipt, idPropertyTitle{id: id})
	}
	return idSet
}

func IdSetFromMap(keys map[string]bool) *IdSet {
	idSet := NewIdSet()
	for id := range keys {
		idSet.ipt = append(idSet.ipt, idPropertyTitle{id: id})
	}
	return idSet
}

func (is *IdSet) All() []string {
	keys := make([]string, 0, is.Len())

	for i := 0; i < is.Len(); i++ {
		keys = append(keys, is.ipt[i].id)
	}

	return keys
}

func (is *IdSet) Sort(rxa kvas.ReduxAssets, property string, desc bool) []string {
	for i := 0; i < is.Len(); i++ {
		is.ipt[i].property, _ = rxa.GetFirstVal(property, is.ipt[i].id)
		if property != TitleProperty {
			is.ipt[i].title, _ = rxa.GetFirstVal(TitleProperty, is.ipt[i].id)
		}
	}

	var sortInterface sort.Interface = is
	if desc {
		sortInterface = sort.Reverse(sortInterface)
	}

	sort.Sort(sortInterface)

	return is.All()
}

func (is *IdSet) Add(keys ...string) {
	for _, id := range keys {
		is.ipt = append(is.ipt, idPropertyTitle{id: id})
	}
}

func (is *IdSet) Remove(keys ...string) {
	for _, id := range keys {
		index := -1
		for i := 0; i < is.Len(); i++ {
			if is.ipt[i].id == id {
				index = i
				break
			}
		}
		if index >= 0 {
			is.ipt[index] = is.ipt[is.Len()-1]
			is.ipt = is.ipt[:is.Len()-1]
		}
	}
}

func (is *IdSet) AddSet(another *IdSet) {
	for _, ipt := range another.ipt {
		is.ipt = append(is.ipt, ipt)
	}
}

func (is *IdSet) AddMap(another map[string]bool) {
	for id := range another {
		is.ipt = append(is.ipt, idPropertyTitle{id: id})
	}
}

func (is *IdSet) Has(id string) bool {
	for _, ipt := range is.ipt {
		if ipt.id == id {
			return true
		}
	}
	return false
}

func (is *IdSet) Except(other *IdSet) []string {
	result := make([]string, 0, is.Len())
	for _, ipt := range is.ipt {
		if other.Has(ipt.id) {
			continue
		}
		result = append(result, ipt.id)
	}
	return result
}

func idSetFromSlugs(slugs []string, rxa kvas.ReduxAssets) (slugId *IdSet, err error) {

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
	idSet *IdSet,
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

	for _, ipt := range idSet.ipt {
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
