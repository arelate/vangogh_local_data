package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/pasu"
)

func NewReduxWriter(properties ...string) (kvas.WriteableRedux, error) {
	rdp, err := pasu.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.NewReduxWriter(rdp, properties...)
}

func NewReduxReader(properties ...string) (kvas.ReadableRedux, error) {
	rdp, err := pasu.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.NewReduxReader(rdp, properties...)
}

func NewReduxVetter(properties ...string) (kvas.IndexVetter, error) {
	rdp, err := pasu.GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.NewReduxVetter(rdp, properties...)
}
