package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
	"time"
)

const (
	IdProperty                  = "id"
	TitleProperty               = "title"
	DevelopersProperty          = "developers"
	PublisherProperty           = "publisher"
	ImageProperty               = "image"
	ScreenshotsProperty         = "screenshots"
	RatingProperty              = "rating"
	IncludesGamesProperty       = "includes-games"
	IsIncludedByGamesProperty   = "is-included-by-games"
	RequiresGamesProperty       = "requires-games"
	IsRequiredByGamesProperty   = "is-required-by-games"
	GenresProperty              = "genres"
	FeaturesProperty            = "features"
	SeriesProperty              = "series"
	TagIdProperty               = "tag"
	TagNameProperty             = "tag-name"
	VideoIdProperty             = "video-id"
	MissingVideoUrlProperty     = "missing-video-url"
	OperatingSystemsProperty    = "os"
	LanguageCodeProperty        = "lang-code"
	LanguageNameProperty        = "lang-name"
	NativeLanguageNameProperty  = "native-lang-name"
	SlugProperty                = "slug"
	GOGReleaseDateProperty      = "gog-release-date"
	GOGOrderDateProperty        = "gog-order-date"
	GlobalReleaseDateProperty   = "global-release-date"
	TextProperty                = "text"
	AllTextProperty             = "all-text"
	ImageIdProperty             = "image-id"
	TypesProperty               = "types"
	LocalManualUrlProperty      = "local-manual-url"
	DownloadStatusErrorProperty = "download-status-error"
	StoreUrlProperty            = "store-url"
	ForumUrlProperty            = "forum-url"
	SupportUrlProperty          = "support-url"
	ChanglogProperty            = "changelog"
	DescriptionProperty         = "description"
	CopyrightsProperty          = "copyrights"
	WishlistedProperty          = "wishlisted"
	ProductTypeProperty         = "product-type"
)

func AllProperties() []string {
	all := []string{IdProperty}
	return append(all, ReduxProperties()...)
}

func IsValidProperty(property string) bool {
	for _, p := range AllProperties() {
		if p == property {
			return true
		}
	}
	return false
}

func TextProperties() []string {
	return []string{
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
	}
}

func UrlProperties() []string {
	return []string{
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
	}
}

func LongTextProperties() []string {
	return []string{
		DescriptionProperty,
		ChanglogProperty,
		CopyrightsProperty,
	}
}

func AllTextProperties() []string {
	return append(TextProperties(),
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
		GlobalReleaseDateProperty,
		GOGOrderDateProperty,
		GOGReleaseDateProperty,
	)
}

func VideoIdProperties() []string {
	return []string{
		VideoIdProperty,
	}
}

func ComputedProperties() []string {
	return []string{
		TypesProperty,
	}
}

func ImageIdProperties() []string {
	return []string{
		ImageProperty,
		ScreenshotsProperty,
	}
}

func ReduxProperties() []string {
	all := AllTextProperties()
	all = append(all, VideoIdProperties()...)
	all = append(all, ComputedProperties()...)
	all = append(all, ImageIdProperties()...)
	all = append(all, UrlProperties()...)
	all = append(all, LongTextProperties()...)
	return append(all, WishlistedProperty, ProductTypeProperty)
}

func DigestibleProperties() []string {
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

func SearchableProperties() []string {
	searchable := make([]string, 0, len(ReduxProperties())+3)
	searchable = append(searchable, TextProperty, AllTextProperty, ImageIdProperty)
	searchable = append(searchable, ReduxProperties()...)
	return searchable
}

var atomicProperties = kvas.ReduxAtomics{
	LanguageCodeProperty:       true,
	NativeLanguageNameProperty: true,
	SlugProperty:               true,
}

func IsPropertyAtomic(property string) bool {
	return atomicProperties.IsAtomic(property)
}

var transitiveProperties = kvas.ReduxTransitives{
	IncludesGamesProperty:     TitleProperty,
	IsIncludedByGamesProperty: TitleProperty,
	RequiresGamesProperty:     TitleProperty,
	IsRequiredByGamesProperty: TitleProperty,
	TagIdProperty:             TagNameProperty,
	LanguageCodeProperty:      NativeLanguageNameProperty,
	VideoIdProperty:           MissingVideoUrlProperty,
}

func IsPropertyTransitive(property string) bool {
	return transitiveProperties.IsTransitive(property)
}

var aggregateProperties = kvas.ReduxAggregates{
	TextProperty:    TextProperties(),
	AllTextProperty: AllTextProperties(),
	ImageIdProperty: ImageIdProperties(),
}

func IsPropertyAggregate(property string) bool {
	return aggregateProperties.IsAggregate(property)
}

func DetailAggregateProperty(property string) []string {
	return aggregateProperties.Detail(property)
}

func DetailAllAggregateProperties(properties ...string) map[string]bool {
	return aggregateProperties.DetailAll(properties...)
}

func joinNotDesirable() []string {
	return append(
		ImageIdProperties(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
	)
}

var imageTypeProperties = map[ImageType]string{
	Image:       ImageProperty,
	Screenshots: ScreenshotsProperty,
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
		StoreUrlProperty,
	},
	ApiProductsV1: {
		IdProperty,
		TitleProperty,
		ScreenshotsProperty,
		VideoIdProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GOGReleaseDateProperty,
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
		ChanglogProperty,
		DescriptionProperty,
	},
	ApiProductsV2: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublisherProperty,
		ImageProperty,
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
		GlobalReleaseDateProperty,
		GOGReleaseDateProperty,
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
		DescriptionProperty,
		CopyrightsProperty,
		ProductTypeProperty,
	},
	Details: {
		TitleProperty,
		FeaturesProperty,
		TagIdProperty,
		GOGReleaseDateProperty,
		ForumUrlProperty,
		ChanglogProperty,
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
		GlobalReleaseDateProperty,
		GOGReleaseDateProperty,
		StoreUrlProperty,
		ForumUrlProperty,
		SupportUrlProperty,
	},
}

func ConnectReduxAssets(properties ...string) (kvas.ReduxAssets, error) {
	return kvas.ConnectReduxAssets(AbsReduxDir(),
		&kvas.ReduxFabric{
			Aggregates:  aggregateProperties,
			Transitives: transitiveProperties,
			Atomics:     atomicProperties,
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
	case ChanglogProperty:
		return getSlice(value.(gog_integration.ChangelogGetter).GetChangelog)
	case CopyrightsProperty:
		return getSlice(value.(gog_integration.CopyrightsGetter).GetCopyrights)
	case DescriptionProperty:
		return getSlice(value.(gog_integration.DescriptionGetter).GetDescription)
	case DevelopersProperty:
		return value.(gog_integration.DevelopersGetter).GetDevelopers()
	case FeaturesProperty:
		return value.(gog_integration.FeaturesGetter).GetFeatures()
	case ForumUrlProperty:
		return getSlice(value.(gog_integration.ForumUrlGetter).GetForumUrl)
	case ImageProperty:
		return getImageIdSlice(value.(gog_integration.ImageGetter).GetImage)
	case IncludesGamesProperty:
		return value.(gog_integration.IncludesGamesGetter).GetIncludesGames()
	case IsIncludedByGamesProperty:
		return value.(gog_integration.IsIncludedInGamesGetter).GetIsIncludedInGames()
	case IsRequiredByGamesProperty:
		return value.(gog_integration.IsRequiredByGamesGetter).GetIsRequiredByGames()
	case GenresProperty:
		return value.(gog_integration.GenresGetter).GetGenres()
	case GlobalReleaseDateProperty:
		return dateSlice(value.(gog_integration.GlobalReleaseGetter).GetGlobalRelease)
	case GOGReleaseDateProperty:
		return dateSlice(value.(gog_integration.GOGReleaseGetter).GetGOGRelease)
	case LanguageCodeProperty:
		return value.(gog_integration.LanguageCodesGetter).GetLanguageCodes()
	case OperatingSystemsProperty:
		return value.(gog_integration.OperatingSystemsGetter).GetOperatingSystems()
	case ProductTypeProperty:
		return getSlice(value.(gog_integration.ProductTypeGetter).GetProductType)
	case PublisherProperty:
		return getSlice(value.(gog_integration.PublisherGetter).GetPublisher)
	case RatingProperty:
		return getSlice(value.(gog_integration.RatingGetter).GetRating)
	case RequiresGamesProperty:
		return value.(gog_integration.RequiresGamesGetter).GetRequiresGames()
	case SeriesProperty:
		return getSlice(value.(gog_integration.SeriesGetter).GetSeries)
	case ScreenshotsProperty:
		return getScreenshots(value)
	case SlugProperty:
		return getSlice(value.(gog_integration.SlugGetter).GetSlug)
	case StoreUrlProperty:
		return getSlice(value.(gog_integration.StoreUrlGetter).GetStoreUrl)
	case SupportUrlProperty:
		return getSlice(value.(gog_integration.SupportUrlGetter).GetSupportUrl)
	case TagIdProperty:
		return value.(gog_integration.TagIdsGetter).GetTagIds()
	case TitleProperty:
		return getSlice(value.(gog_integration.TitleGetter).GetTitle)
	case VideoIdProperty:
		return value.(gog_integration.VideoIdsGetter).GetVideoIds()
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
		imageIds = append(imageIds, gog_integration.ImageId(str))
	}
	return imageIds
}

func getScreenshots(value interface{}) []string {
	screenshotsGetter := value.(gog_integration.ScreenshotsGetter)
	if screenshotsGetter != nil {
		screenshots := screenshotsGetter.GetScreenshots()
		imageIds := make([]string, 0, len(screenshots))
		for _, scr := range screenshots {
			imageIds = append(imageIds, gog_integration.ImageId(scr))
		}
		return imageIds
	}
	return []string{}
}
