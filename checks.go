package vangogh_local_data

import (
	"github.com/boggydigital/kvas"
	"os"
	"path"
)

func IsGOGPagedProduct(pt ProductType) bool {
	return containsProductType(GOGPagedProducts(), pt)
}

func IsArrayProduct(pt ProductType) bool {
	return IsGOGArrayProduct(pt) ||
		IsSteamArrayProduct(pt) ||
		IsHLTBArrayProduct(pt)
}

func IsGOGArrayProduct(pt ProductType) bool {
	return containsProductType(GOGArrayProducts(), pt)
}

func IsSteamArrayProduct(pt ProductType) bool {
	return containsProductType(SteamArrayProducts(), pt)
}

func IsHLTBArrayProduct(pt ProductType) bool {
	return containsProductType(HLTBArrayProducts(), pt)
}

func IsGOGDetailProduct(pt ProductType) bool {
	return containsProductType(GOGDetailProducts(), pt)
}

func IsGOGProduct(pt ProductType) bool {
	return IsGOGPagedProduct(pt) ||
		IsGOGArrayProduct(pt) ||
		IsGOGDetailProduct(pt)
}

func IsSteamDetailProduct(pt ProductType) bool {
	return containsProductType(SteamDetailProducts(), pt)
}

func IsSteamProduct(pt ProductType) bool {
	return IsSteamArrayProduct(pt) ||
		IsSteamDetailProduct(pt)
}

func IsPCGWDetailProduct(pt ProductType) bool {
	return containsProductType(PCGWDetailProducts(), pt)
}

func IsPCGWProduct(pt ProductType) bool {
	return IsPCGWDetailProduct(pt)
}

func IsFastPageFetchProduct(pt ProductType) bool {
	return containsProductType(FastPageFetchProducts(), pt)
}

func IsProductRequiresAuth(pt ProductType) bool {
	return containsProductType(requireAuth, pt)
}

func IsImageRequiresAuth(it ImageType) bool {
	for _, itra := range imageTypeRequiresAuth {
		if itra == it {
			return true
		}
	}
	return false
}

func IsGetItemsSupported(pt ProductType) bool {
	return containsProductType(supportsGetItems, pt)
}

func IsImageTypeSupported(pt ProductType, it ImageType) bool {
	if !IsValidProductType(pt) ||
		!IsValidImageType(it) {
		return false
	}

	supportedIts, ok := supportedImageTypes[pt]
	if !ok {
		return false
	}

	for _, sit := range supportedIts {
		if sit == it {
			return true
		}
	}

	return false
}

func containsProductType(all []ProductType, pt ProductType) bool {
	for _, apt := range all {
		if apt == pt {
			return true
		}
	}
	return false
}

func IsPathSupportingValidation(filePath string) bool {
	ext := path.Ext(filePath)
	return validatedExtensions[ext]
}

func IsSupportedProperty(pt ProductType, property string) bool {
	for _, supportedProperty := range supportedProperties[pt] {
		if property == supportedProperty {
			return true
		}
	}
	return false
}

func IsProductDownloaded(id string, rxa kvas.ReduxAssets) (bool, error) {
	if err := rxa.IsSupported(SlugProperty); err != nil {
		return false, err
	}

	slug, ok := rxa.GetFirstVal(SlugProperty, id)
	if !ok {
		return false, nil
	}

	pDir, err := AbsProductDownloadsDir(slug)
	if err != nil {
		return false, err
	}

	if _, err := os.Stat(pDir); os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func IsPropertiesJoinPreferred(property string) bool {
	for _, nd := range joinNotDesirable() {
		if property == nd {
			return false
		}
	}
	return true
}
