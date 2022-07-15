package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
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
		values := strings.Split(val, ",")
		//account for empty strings
		if len(values) == 1 && values[0] == "" {
			values = []string{}
		}
		return values
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

func MediaFromUrl(u *url.URL) gog_integration.Media {
	return gog_integration.ParseMedia(ValueFromUrl(u, "media"))
}

func OperatingSystemsFromUrl(u *url.URL) []OperatingSystem {
	osStrings := ValuesFromUrl(u, "operating-system")
	return ParseManyOperatingSystems(osStrings)
}

func DownloadTypesFromUrl(u *url.URL) []DownloadType {
	dtStrings := ValuesFromUrl(u, "download-type")
	return ParseManyDownloadTypes(dtStrings)
}

func IdSetFromUrl(u *url.URL) (map[string]bool, error) {

	idSet := make(map[string]bool)
	for _, id := range ValuesFromUrl(u, "id") {
		idSet[id] = true
	}

	slugs := ValuesFromUrl(u, "slug")

	slugIds, err := idSetFromSlugs(slugs, nil)
	if err != nil {
		return idSet, err
	}
	for id := range slugIds {
		idSet[id] = true
	}

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
