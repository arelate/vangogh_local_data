package vangogh_local_data

import (
	"errors"
	"github.com/boggydigital/kvas"
	"golang.org/x/exp/slices"
	"time"
)

type IdReduxAssets = map[string]map[string][]string

var IRAProxyReadOnlyError = errors.New("id redux assets proxy is read-only")

type IRAProxy struct {
	rdx     IdReduxAssets
	modTime int64
}

func NewIRAProxy(data IdReduxAssets) *IRAProxy {
	return &IRAProxy{
		rdx:     data,
		modTime: time.Now().UTC().Unix(),
	}
}

func NewEmptyIRAProxy(properties []string) *IRAProxy {
	dra := NewIRAProxy(make(IdReduxAssets))
	dra.rdx[""] = make(map[string][]string)
	for _, p := range properties {
		dra.rdx[""][p] = nil
	}
	return dra
}

func (irap *IRAProxy) Keys(asset string) []string {
	keys := make([]string, 0)
	for id, pvs := range irap.rdx {
		if _, ok := pvs[asset]; ok {
			keys = append(keys, id)
		}
	}
	return keys
}

func (irap *IRAProxy) Has(asset string) bool {
	for _, pvs := range irap.rdx {
		if _, ok := pvs[asset]; ok {
			return true
		}
	}
	return false
}

func (irap *IRAProxy) HasKey(asset, key string) bool {
	if pvs, ok := irap.rdx[key]; ok {
		if _, ok := pvs[asset]; ok {
			return true
		}
	}
	return false
}

func (irap *IRAProxy) HasVal(asset, key, val string) bool {
	if pvs, ok := irap.rdx[key]; ok {
		if vals, ok := pvs[asset]; ok {
			return slices.Contains(vals, val)
		}
	}
	return false
}

func (irap *IRAProxy) AddVal(asset, key, val string) error {
	return IRAProxyReadOnlyError
}

func (irap *IRAProxy) ReplaceValues(asset, key string, values ...string) error {
	return IRAProxyReadOnlyError
}

func (irap *IRAProxy) BatchReplaceValues(asset string, keyValues map[string][]string) error {
	return IRAProxyReadOnlyError
}

func (irap *IRAProxy) CutVal(asset, key, val string) error {
	return IRAProxyReadOnlyError
}

func (irap *IRAProxy) GetAllValues(asset, key string) ([]string, bool) {
	if pvs, ok := irap.rdx[key]; ok {
		if vals, ok := pvs[asset]; ok {
			return vals, true
		}
	}
	return nil, false
}

func (irap *IRAProxy) GetAllUnchangedValues(asset, key string) ([]string, bool) {
	return irap.GetAllValues(asset, key)
}

func (irap *IRAProxy) GetFirstVal(asset, key string) (string, bool) {
	if vals, ok := irap.GetAllValues(asset, key); ok {
		if len(vals) > 0 {
			return vals[0], true
		}
	}
	return "", false
}

func (irap *IRAProxy) IsSupported(assets ...string) error {
	for _, a := range assets {
		if !irap.Has(a) {
			return errors.New("unsupported asset " + a)
		}
	}
	return nil
}

func (irap *IRAProxy) Match(query map[string][]string, anyCase bool) map[string]bool {
	//FIXME
	return nil
}

func (irap *IRAProxy) RefreshReduxAssets() (kvas.ReduxAssets, error) {
	return irap, nil
}

func (irap *IRAProxy) ReduxAssetsModTime() (int64, error) {
	return irap.modTime, nil
}

func (irap *IRAProxy) Sort(ids []string, desc bool, sortBy ...string) ([]string, error) {
	//FIXME
	return ids, IRAProxyReadOnlyError
}
