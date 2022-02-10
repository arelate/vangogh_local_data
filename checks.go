package vangogh_data

import (
	"github.com/arelate/gog_atu"
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
	if !ValidProductType(pt) ||
		!ValidImageType(it) {
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
	if !ValidProductType(pt) {
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
