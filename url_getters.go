package vangogh_data

import (
	"github.com/arelate/gog_atu"
	"net/url"
	"strconv"
	"strings"
	"time"
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

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, "operating-system")
	return ParseManyOperatingSystems(osStrings)
}

func DownloadTypesFromUrl(u *url.URL) []DownloadType {
	dtStrings := ValuesFromUrl(u, "download-type")
	return ParseManyDownloadTypes(dtStrings)
}

func IdSetFromUrl(u *url.URL) (idSet IdSet, err error) {

	idSet = IdSetFromSlice(ValuesFromUrl(u, "id")...)

	slugs := ValuesFromUrl(u, "slug")

	slugIds, err := idSetFromSlugs(slugs, nil)
	if err != nil {
		return idSet, err
	}
	idSet.AddSet(slugIds)

	return idSet, err
}

func SinceFromUrl(u *url.URL) (int64, error) {
	str := ValueFromUrl(u, "since-hours-ago")
	var sha int
	var err error
	if str != "" {
		sha, err = strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
	}
	return time.Now().Unix() - int64(sha*60*60), err
}