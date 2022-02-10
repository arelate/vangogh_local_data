package vangogh_data

import (
	"github.com/arelate/gog_atu"
	"github.com/boggydigital/kvas"
	"time"
)

const (
	IdProperty                 = "id"
	TitleProperty              = "title"
	DevelopersProperty         = "developers"
	PublisherProperty          = "publisher"
	ImageProperty              = "image"
	BoxArtProperty             = "box-art"
	BackgroundProperty         = "background"
	GalaxyBackgroundProperty   = "galaxy-background"
	IconProperty               = "icon"
	LogoProperty               = "logo"
	ScreenshotsProperty        = "screenshots"
	RatingProperty             = "rating"
	IncludesGamesProperty      = "includes-games"
	IsIncludedByGamesProperty  = "is-included-by-games"
	RequiresGamesProperty      = "requires-games"
	IsRequiredByGamesProperty  = "is-required-by-games"
	GenresProperty             = "genres"
	FeaturesProperty           = "features"
	SeriesProperty             = "series"
	TagIdProperty              = "tag"
	TagNameProperty            = "tag-name"
	VideoIdProperty            = "video-id"
	MissingVideoUrlProperty    = "missing-video-url"
	OperatingSystemsProperty   = "os"
	LanguageCodeProperty       = "lang-code"
	LanguageNameProperty       = "lang-name"
	NativeLanguageNameProperty = "native-lang-name"
	SlugProperty               = "slug"
	GOGReleaseDate             = "gog-release-date"
	GOGOrderDate               = "gog-order-date"
	GlobalReleaseDate          = "global-release-date"
	TextProperties             = "text"
	AllTextProperties          = "all-text"
	ImageIdProperties          = "image-id"
	TypesProperty              = "types"
	LocalManualUrl             = "local-manual-url"
	DownloadStatusError        = "download-status-error"
)

func AllProperties() []string {
	all := []string{IdProperty}
	return append(all, Extracted()...)
}

func IsValidProperty(property string) bool {
	for _, p := range AllProperties() {
		if p == property {
			return true
		}
	}
	return false
}

func Text() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
	}
}

func AllText() []string {
	return append(Text(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		RatingProperty,
		TagIdProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		SlugProperty,
		GlobalReleaseDate,
		GOGOrderDate,
		GOGReleaseDate,
	)
}

func VideoId() []string {
	return []string{
		VideoIdProperty,
	}
}

func Computed() []string {
	return []string{
		TypesProperty,
	}
}

func ImageId() []string {
	return []string{
		ImageProperty,
		BoxArtProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		IconProperty,
		LogoProperty,
		ScreenshotsProperty,
	}
}

func Extracted() []string {
	all := AllText()
	all = append(all, VideoId()...)
	all = append(all, Computed()...)
	return append(all, ImageId()...)
}

func Digestible() []string {
	return []string{
		DevelopersProperty,
		PublisherProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		TagIdProperty,
		LanguageCodeProperty,
		OperatingSystemsProperty,
		MissingVideoUrlProperty,
	}
}

func Searchable() []string {
	searchable := make([]string, 0, len(Extracted())+3)
	searchable = append(searchable, TextProperties, AllTextProperties, ImageIdProperties)
	searchable = append(searchable, Extracted()...)
	return searchable
}

var fullMatch = map[string]bool{
	LanguageCodeProperty:       true,
	NativeLanguageNameProperty: true,
	SlugProperty:               true,
}

var replacementProperties = map[string]string{
	IncludesGamesProperty:     TitleProperty,
	IsIncludedByGamesProperty: TitleProperty,
	RequiresGamesProperty:     TitleProperty,
	IsRequiredByGamesProperty: TitleProperty,
	TagIdProperty:             TagNameProperty,
	LanguageCodeProperty:      NativeLanguageNameProperty,
	VideoIdProperty:           MissingVideoUrlProperty,
}

var collapsedExpanded = map[string][]string{
	TextProperties:    Text(),
	AllTextProperties: AllText(),
	ImageIdProperties: ImageId(),
}

func joinNotDesirable() []string {
	return append(
		ImageId(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
	)
}

var imageTypeProperties = map[ImageType]string{
	Image:            ImageProperty,
	BoxArt:           BoxArtProperty,
	Background:       BackgroundProperty,
	GalaxyBackground: GalaxyBackgroundProperty,
	Logo:             LogoProperty,
	Icon:             IconProperty,
	Screenshots:      ScreenshotsProperty,
}

func PropertyFromImageType(it ImageType) string {
	return imageTypeProperties[it]
}

var supportedProperties = map[ProductType][]string{
	AccountProducts: {
		IdProperty,
		TitleProperty,
		ImageProperty,
		RatingProperty,
		OperatingSystemsProperty,
		SlugProperty,
	},
	ApiProductsV1: {
		IdProperty,
		TitleProperty,
		IconProperty,
		BackgroundProperty,
		ScreenshotsProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GOGReleaseDate,
	},
	ApiProductsV2: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		ImageProperty,
		BoxArtProperty,
		IconProperty,
		LogoProperty,
		BackgroundProperty,
		GalaxyBackgroundProperty,
		ScreenshotsProperty,
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		FeaturesProperty,
		SeriesProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		GlobalReleaseDate,
		GOGReleaseDate,
	},
	Details: {
		TitleProperty,
		BackgroundProperty,
		FeaturesProperty,
		TagIdProperty,
		GOGReleaseDate,
	},
	StoreProducts: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		ImageProperty,
		ScreenshotsProperty,
		RatingProperty,
		GenresProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GlobalReleaseDate,
		GOGReleaseDate,
	},
}

func ConnectReduxAssets(properties ...string) (kvas.ReduxAssets, error) {
	return kvas.ConnectReduxAssets(AbsExtractsDir(),
		&kvas.ReduxFabric{
			Aggregates:  collapsedExpanded,
			Transitives: replacementProperties,
			Atomics:     fullMatch,
		},
		properties...)
}

func GetProperties(
	id string,
	reader *ValueReader,
	properties []string) (propValues map[string][]string, err error) {
	supProps := SupportedPropertiesOnly(reader.ProductType(), properties)
	value, err := reader.ReadValue(id)
	return fillProperties(value, supProps), err
}

func fillProperties(value interface{}, properties []string) map[string][]string {
	propValues := make(map[string][]string, 0)
	for _, prop := range properties {
		propValues[prop] = getPropertyValues(value, prop)
	}
	return propValues
}

func getPropertyValues(value interface{}, property string) []string {
	switch property {
	case BackgroundProperty:
		return getImageIdSlice(value.(gog_atu.BackgroundGetter).GetBackground)
	case BoxArtProperty:
		return getImageIdSlice(value.(gog_atu.BoxArtGetter).GetBoxArt)
	case DevelopersProperty:
		return value.(gog_atu.DevelopersGetter).GetDevelopers()
	case FeaturesProperty:
		return value.(gog_atu.FeaturesGetter).GetFeatures()
	case IconProperty:
		return getImageIdSlice(value.(gog_atu.IconGetter).GetIcon)
	case ImageProperty:
		return getImageIdSlice(value.(gog_atu.ImageGetter).GetImage)
	case IncludesGamesProperty:
		return value.(gog_atu.IncludesGamesGetter).GetIncludesGames()
	case IsIncludedByGamesProperty:
		return value.(gog_atu.IsIncludedInGamesGetter).GetIsIncludedInGames()
	case IsRequiredByGamesProperty:
		return value.(gog_atu.IsRequiredByGamesGetter).GetIsRequiredByGames()
	case GalaxyBackgroundProperty:
		return getImageIdSlice(value.(gog_atu.GalaxyBackgroundGetter).GetGalaxyBackground)
	case GenresProperty:
		return value.(gog_atu.GenresGetter).GetGenres()
	case GlobalReleaseDate:
		return dateSlice(value.(gog_atu.GlobalReleaseGetter).GetGlobalRelease)
	case GOGReleaseDate:
		return dateSlice(value.(gog_atu.GOGReleaseGetter).GetGOGRelease)
	case LanguageCodeProperty:
		return value.(gog_atu.LanguageCodesGetter).GetLanguageCodes()
	case LogoProperty:
		return getImageIdSlice(value.(gog_atu.LogoGetter).GetLogo)
	case OperatingSystemsProperty:
		return value.(gog_atu.OperatingSystemsGetter).GetOperatingSystems()
	case PublisherProperty:
		return getSlice(value.(gog_atu.PublisherGetter).GetPublisher)
	case RatingProperty:
		return getSlice(value.(gog_atu.RatingGetter).GetRating)
	case RequiresGamesProperty:
		return value.(gog_atu.RequiresGamesGetter).GetRequiresGames()
	case SeriesProperty:
		return getSlice(value.(gog_atu.SeriesGetter).GetSeries)
	case ScreenshotsProperty:
		return getScreenshots(value)
	case SlugProperty:
		return getSlice(value.(gog_atu.SlugGetter).GetSlug)
	case TagIdProperty:
		return value.(gog_atu.TagIdsGetter).GetTagIds()
	case TitleProperty:
		return getSlice(value.(gog_atu.TitleGetter).GetTitle)
	case VideoIdProperty:
		return value.(gog_atu.VideoIdsGetter).GetVideoIds()
	default:
		return []string{}
	}
}

func dateSlice(timestamper func() int64) []string {
	dates := make([]string, 0)
	if timestamper != nil {
		val := timestamper()
		if val > 0 {
			date := time.Unix(val, 0)
			dates = append(dates, date.Format("2006-01-02"))
		}
	}
	return dates
}

func getSlice(stringer func() string) []string {
	values := make([]string, 0)
	if stringer != nil {
		val := stringer()
		if val != "" {
			values = append(values, val)
		}
	}
	return values
}

func getImageIdSlice(stringer func() string) []string {
	strings := getSlice(stringer)
	imageIds := make([]string, 0, len(strings))
	for _, str := range strings {
		imageIds = append(imageIds, gog_atu.ImageId(str))
	}
	return imageIds
}

func getScreenshots(value interface{}) []string {
	screenshotsGetter := value.(gog_atu.ScreenshotsGetter)
	if screenshotsGetter != nil {
		screenshots := screenshotsGetter.GetScreenshots()
		imageIds := make([]string, 0, len(screenshots))
		for _, scr := range screenshots {
			imageIds = append(imageIds, gog_atu.ImageId(scr))
		}
		return imageIds
	}
	return []string{}
}
