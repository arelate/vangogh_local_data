package vangogh_data

import (
	"github.com/arelate/gog_atu"
	"github.com/boggydigital/kvas"
	"os"
	"path"
)

func IsPagedProduct(pt ProductType) bool {
	return containsProductType(PagedProducts(), pt)
}

func IsArrayProduct(pt ProductType) bool {
	return containsProductType(ArrayProducts(), pt)
}

func IsDetailProduct(pt ProductType) bool {
	return containsProductType(DetailProducts(), pt)
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

func IsCopySupported(from, to ProductType) bool {
	return supportsCopyFromTo[from] == to
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

func IsMediaSupported(pt ProductType, mt gog_atu.Media) bool {
	if !gog_atu.ValidMedia(mt) {
		return false
	}
	if !IsValidProductType(pt) {
		return false
	}

	ums, ok := unsupportedMedia[pt]
	if !ok {
		return true
	}

	for _, um := range ums {
		if um == mt {
			return false
		}
	}

	return true
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
