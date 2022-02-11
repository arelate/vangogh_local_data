package vangogh_data

import (
	"github.com/boggydigital/gost"
	"github.com/boggydigital/kvas"
)

type IdSet struct {
	gost.SortStrSet
}

func NewIdSet() IdSet {
	return IdSet{SortStrSet: gost.NewSortStrSet()}
}

func IdSetWith(ids ...string) IdSet {
	idSet := NewIdSet()
	idSet.Add(ids...)
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
