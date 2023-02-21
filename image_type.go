package vangogh_local_data

type ImageType int

const (
	UnknownImageType ImageType = iota
	Image
	Screenshots
	VerticalImage
)

var imageTypeStrings = map[ImageType]string{
	UnknownImageType: "unknown-image-type",
	Image:            "image",
	Screenshots:      "screenshots",
	VerticalImage:    "vertical-image",
}

func (it ImageType) String() string {
	str, ok := imageTypeStrings[it]
	if ok {
		return str
	}

	return imageTypeStrings[UnknownImageType]
}

func ParseImageType(imageType string) ImageType {
	for it, str := range imageTypeStrings {
		if str == imageType {
			return it
		}
	}
	return UnknownImageType
}

func IsValidImageType(it ImageType) bool {
	_, ok := imageTypeStrings[it]
	return ok && it != UnknownImageType
}

func AllImageTypes() []ImageType {
	imageTypes := make([]ImageType, 0, len(imageTypeStrings))
	for it, _ := range imageTypeStrings {
		if it == UnknownImageType {
			continue
		}
		imageTypes = append(imageTypes, it)
	}
	return imageTypes
}

// starting with empty collection and no image types require auth at the moment
var imageTypeRequiresAuth []ImageType

func ImageTypeFromProperty(property string) ImageType {
	for it, prop := range imageTypeProperties {
		if prop == property {
			return it
		}
	}
	return UnknownImageType
}

var imageTypesSupportingDehydration = map[ImageType]bool{
	Image:         true,
	VerticalImage: true,
}

func IsImageTypeDehydrationSupported(it ImageType) bool {
	_, ok := imageTypesSupportingDehydration[it]
	return ok
}

// values are result of running issa dehydration on an image set and finding
// sampling rate that produces most compact representation (min sum of string lengths)
var imageTypeDehydrationSamples = map[ImageType]int{
	VerticalImage: 16,
	Image:         20,
}

func ImageTypeDehydrationSamples(it ImageType) int {
	if s, ok := imageTypeDehydrationSamples[it]; ok {
		return s
	}
	return -1
}
