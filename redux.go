package vangogh_local_data

import "github.com/boggydigital/kvas"

func ReduxWriter(properties ...string) (kvas.WriteableRedux, error) {
	rdp, err := GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.ReduxWriter(rdp, properties...)
}

func ReduxReader(properties ...string) (kvas.ReadableRedux, error) {
	rdp, err := GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.ReduxReader(rdp, properties...)
}

func ReduxVetter(properties ...string) (kvas.IndexVetter, error) {
	rdp, err := GetAbsRelDir(Redux)
	if err != nil {
		return nil, err
	}
	return kvas.ReduxVetter(rdp, properties...)
}
