package vangogh_data

import (
	"github.com/arelate/gog_atu"
	"net/url"
	"strings"
)

func ValueFromUrl(u *url.URL, arg string) string {
	if u == nil {
		return ""
	}

	q := u.Query()
	return q.Get(arg)
}

func ValuesFromUrl(u *url.URL, arg string) []string {
	if FlagFromUrl(u, arg) {
		val := ValueFromUrl(u, arg)
		return strings.Split(val, ",")
	}
	return nil
}

func FlagFromUrl(u *url.URL, arg string) bool {
	if u == nil {
		return false
	}

	q := u.Query()
	return q.Has(arg)
}

func PropertyFromUrl(u *url.URL) string {
	return ValueFromUrl(u, "property")
}

func PropertiesFromUrl(u *url.URL) []string {
	return ValuesFromUrl(u, "property")
}

func ImageTypesFromUrl(u *url.URL) []ImageType {
	imageTypes := ValuesFromUrl(u, "image-type")
	its := make([]ImageType, 0, len(imageTypes))
	for _, imageType := range imageTypes {
		its = append(its, ParseImageType(imageType))
	}
	return its
}

func ProductTypeFromUrl(u *url.URL) ProductType {
	return ParseProductType(ValueFromUrl(u, "product-type"))
}

func MediaFromUrl(u *url.URL) gog_atu.Media {
	return gog_atu.ParseMedia(ValueFromUrl(u, "media"))
}
