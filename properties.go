package vangogh_local_data

import (
	"github.com/arelate/southern_light/gog_integration"
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/arelate/southern_light/ign_integration"
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/arelate/southern_light/steam_integration"
	"github.com/boggydigital/kvas"
	"strconv"
	"time"
)

const (
	IdProperty                      = "id"
	TitleProperty                   = "title"
	DevelopersProperty              = "developers"
	PublishersProperty              = "publishers"
	ImageProperty                   = "image"
	VerticalImageProperty           = "vertical-image"
	ScreenshotsProperty             = "screenshots"
	RatingProperty                  = "rating"
	IncludesGamesProperty           = "includes-games"
	IsIncludedByGamesProperty       = "is-included-by-games"
	RequiresGamesProperty           = "requires-games"
	IsRequiredByGamesProperty       = "is-required-by-games"
	GenresProperty                  = "genres"
	StoreTagsProperty               = "store-tags"
	FeaturesProperty                = "features"
	SeriesProperty                  = "series"
	TagIdProperty                   = "tag"
	TagNameProperty                 = "tag-name"
	VideoIdProperty                 = "video-id"
	MissingVideoUrlProperty         = "missing-video-url"
	MissingVideoThumbnailProperty   = "missing-video-thumbnail"
	OperatingSystemsProperty        = "os"
	LanguageCodeProperty            = "lang-code"
	LanguageNameProperty            = "lang-name"
	NativeLanguageNameProperty      = "native-lang-name"
	SlugProperty                    = "slug"
	GOGReleaseDateProperty          = "gog-release-date"
	GOGOrderDateProperty            = "gog-order-date"
	GlobalReleaseDateProperty       = "global-release-date"
	TextProperty                    = "text"
	AllTextProperty                 = "all-text"
	ImageIdProperty                 = "image-id"
	TypesProperty                   = "types"
	LocalManualUrlProperty          = "local-manual-url"
	DownloadStatusErrorProperty     = "download-status-error"
	StoreUrlProperty                = "store-url"
	ForumUrlProperty                = "forum-url"
	SupportUrlProperty              = "support-url"
	ChangelogProperty               = "changelog"
	DescriptionOverviewProperty     = "description-overview"
	DescriptionFeaturesProperty     = "description-features"
	AdditionalRequirementsProperty  = "additional-requirements"
	CopyrightsProperty              = "copyrights"
	WishlistedProperty              = "wishlisted"
	OwnedProperty                   = "owned"
	ProductTypeProperty             = "product-type"
	InDevelopmentProperty           = "in-development"
	PreOrderProperty                = "pre-order"
	ComingSoonProperty              = "coming-soon"
	BasePriceProperty               = "base-price"
	PriceProperty                   = "price"
	IsFreeProperty                  = "is-free"
	IsDiscountedProperty            = "is-discounted"
	DiscountPercentageProperty      = "discount-percentage"
	SteamAppIdProperty              = "steam-app-id"
	LocalTagsProperty               = "local-tags"
	SortProperty                    = "sort"
	DescendingProperty              = "desc"
	SteamReviewScoreDescProperty    = "steam-review-score-desc"
	SteamTagsProperty               = "steam-tags"
	DehydratedImageProperty         = "dehydrated-image"
	DehydratedImageModifiedProperty = "dehydrated-image-modified"
	MissingDehydratedImageProperty  = "missing-dehydrated-image"
	SyncEventsProperty              = "sync-events"
	LastSyncUpdatesProperty         = "last-sync-updates"
	ValidationResultProperty        = "validation-result"
	ValidationCompletedProperty     = "validation-completed"
	PCGWPageIdProperty              = "pcgw-page-id"
	HLTBIdProperty                  = "hltb-id"
	HLTBBuildIdProperty             = "hltb-next-build"
	HLTBHoursToCompleteMainProperty = "hltb-comp-main"
	HLTBHoursToCompletePlusProperty = "hltb-comp-plus"
	HLTBHoursToComplete100Property  = "hltb-comp-100"
	HLTBReviewScoreProperty         = "hltb-review-score"
	HLTBGenresProperty              = "hltb-genres"
	HLTBPlatformsProperty           = "hltb-platforms"
	IGDBIdProperty                  = "igdb-id"
	StrategyWikiIdProperty          = "strategy-wiki-id"
	MobyGamesIdProperty             = "moby-games-id"
	WikipediaIdProperty             = "wikipedia-id"
	WineHQIdProperty                = "winehq-id"
	VNDBIdProperty                  = "vndb-id"
	IGNWikiSlugProperty             = "ign-wiki-slug"
	EnginesProperty                 = "engines"
	EnginesBuildsProperty           = "engines-builds"

	// property values
	TrueValue  = "true"
	FalseValue = "false"
	OKValue    = "OK"
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
		PublishersProperty,
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
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
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
		ChangelogProperty,
		CopyrightsProperty,
		AdditionalRequirementsProperty,
	}
}

func AllTextProperties() []string {
	return append(TextProperties(),
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		HLTBGenresProperty,
		StoreTagsProperty,
		FeaturesProperty,
		SeriesProperty,
		RatingProperty,
		TagIdProperty,
		LocalTagsProperty,
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
		VerticalImageProperty,
	}
}

func AvailabilityProperties() []string {
	return []string{
		InDevelopmentProperty,
		PreOrderProperty,
		ComingSoonProperty,
	}
}

func AccountStatusProperties() []string {
	return []string{
		WishlistedProperty,
		OwnedProperty,
	}
}

func AdvancedProductProperties() []string {
	return []string{
		ProductTypeProperty,
		HLTBHoursToCompleteMainProperty,
		HLTBHoursToCompletePlusProperty,
		HLTBHoursToComplete100Property,
		HLTBPlatformsProperty,
		HLTBReviewScoreProperty,
	}
}

func EnginesProperties() []string {
	return []string{
		EnginesProperty,
		EnginesBuildsProperty,
	}
}

func PriceProperties() []string {
	return []string{
		BasePriceProperty,
		PriceProperty,
		IsFreeProperty,
		IsDiscountedProperty,
		DiscountPercentageProperty,
	}
}

func ExternalDataSourcesProperties() []string {
	return []string{
		SteamAppIdProperty,
		SteamReviewScoreDescProperty,
		SteamTagsProperty,
		PCGWPageIdProperty,
		HLTBIdProperty,
		HLTBBuildIdProperty,
		IGDBIdProperty,
		StrategyWikiIdProperty,
		MobyGamesIdProperty,
		WikipediaIdProperty,
		WineHQIdProperty,
		VNDBIdProperty,
		IGNWikiSlugProperty,
	}
}

func MediaContentProperties() []string {
	return []string{
		DehydratedImageProperty,
	}
}

func SyncProperties() []string {
	return []string{
		LastSyncUpdatesProperty,
		SyncEventsProperty,
	}
}

func ValidationProperties() []string {
	return []string{
		ValidationResultProperty,
		ValidationCompletedProperty,
	}
}

func ReduxProperties() []string {
	all := AllTextProperties()
	all = append(all, VideoIdProperties()...)
	all = append(all, ComputedProperties()...)
	all = append(all, ImageIdProperties()...)
	all = append(all, UrlProperties()...)
	all = append(all, LongTextProperties()...)
	all = append(all, AvailabilityProperties()...)
	all = append(all, AccountStatusProperties()...)
	all = append(all, AdvancedProductProperties()...)
	all = append(all, PriceProperties()...)
	all = append(all, ExternalDataSourcesProperties()...)
	all = append(all, MediaContentProperties()...)
	all = append(all, SyncProperties()...)
	all = append(all, ValidationProperties()...)
	all = append(all, EnginesProperties()...)
	return all
}

func DigestibleProperties() []string {
	return []string{
		DevelopersProperty,
		PublishersProperty,
		GenresProperty,
		StoreTagsProperty,
		FeaturesProperty,
		SeriesProperty,
		TagIdProperty,
		LanguageCodeProperty,
		OperatingSystemsProperty,
		MissingVideoUrlProperty,
		SteamReviewScoreDescProperty,
		SteamTagsProperty,
		HLTBPlatformsProperty,
		HLTBGenresProperty,
		EnginesProperty,
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
	Image:         ImageProperty,
	Screenshots:   ScreenshotsProperty,
	VerticalImage: VerticalImageProperty,
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
		ChangelogProperty,
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
		InDevelopmentProperty,
		PreOrderProperty,
	},
	ApiProductsV2: {
		AdditionalRequirementsProperty,
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublishersProperty,
		ImageProperty,
		VerticalImageProperty,
		ScreenshotsProperty,
		IncludesGamesProperty,
		IsIncludedByGamesProperty,
		RequiresGamesProperty,
		IsRequiredByGamesProperty,
		GenresProperty,
		StoreTagsProperty,
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
		DescriptionOverviewProperty,
		DescriptionFeaturesProperty,
		CopyrightsProperty,
		ProductTypeProperty,
		InDevelopmentProperty,
		PreOrderProperty,
	},
	Details: {
		TitleProperty,
		FeaturesProperty,
		TagIdProperty,
		GOGReleaseDateProperty,
		ForumUrlProperty,
		ChangelogProperty,
	},
	CatalogProducts: {
		IdProperty,
		TitleProperty,
		DevelopersProperty,
		PublishersProperty,
		ImageProperty,
		VerticalImageProperty,
		ScreenshotsProperty,
		FeaturesProperty,
		RatingProperty,
		GenresProperty,
		OperatingSystemsProperty,
		SlugProperty,
		GlobalReleaseDateProperty,
		ProductTypeProperty,
		StoreTagsProperty,
		BasePriceProperty,
		PriceProperty,
		IsFreeProperty,
		IsDiscountedProperty,
		DiscountPercentageProperty,
		ComingSoonProperty,
		PreOrderProperty,
		InDevelopmentProperty,
	},
	SteamReviews: {
		SteamReviewScoreDescProperty,
	},
	SteamStorePage: {
		SteamTagsProperty,
	},
	PCGWPageId: {
		PCGWPageIdProperty,
	},
	PCGWEngine: {
		EnginesProperty,
		EnginesBuildsProperty,
	},
	PCGWExternalLinks: {
		SteamAppIdProperty,
		HLTBIdProperty,
		IGDBIdProperty,
		StrategyWikiIdProperty,
		MobyGamesIdProperty,
		WikipediaIdProperty,
		WineHQIdProperty,
		VNDBIdProperty,
	},
	HLTBRootPage: {
		HLTBBuildIdProperty,
	},
	HLTBData: {
		HLTBHoursToCompleteMainProperty,
		HLTBHoursToCompletePlusProperty,
		HLTBHoursToComplete100Property,
		SteamAppIdProperty,
		GlobalReleaseDateProperty,
		HLTBGenresProperty,
		HLTBPlatformsProperty,
		HLTBReviewScoreProperty,
		IGNWikiSlugProperty,
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
	case AdditionalRequirementsProperty:
		return getSlice(value.(gog_integration.AdditionalRequirementsGetter).GetAdditionalRequirements)
	case BasePriceProperty:
		return getSlice(value.(gog_integration.BasePriceGetter).GetBasePrice)
	case ChangelogProperty:
		return getSlice(value.(gog_integration.ChangelogGetter).GetChangelog)
	case ComingSoonProperty:
		return boolSlice(value.(gog_integration.ComingSoonGetter).GetComingSoon)
	case CopyrightsProperty:
		return getSlice(value.(gog_integration.CopyrightsGetter).GetCopyrights)
	case DescriptionFeaturesProperty:
		return getSlice(value.(gog_integration.DescriptionFeaturesGetter).GetDescriptionFeatures)
	case DescriptionOverviewProperty:
		return getSlice(value.(gog_integration.DescriptionOverviewGetter).GetDescriptionOverview)
	case DevelopersProperty:
		return value.(gog_integration.DevelopersGetter).GetDevelopers()
	case DiscountPercentageProperty:
		return intSlice(value.(gog_integration.DiscountPercentageGetter).GetDiscountPercentage)
	case EnginesProperty:
		return value.(pcgw_integration.EnginesGetter).GetEngines()
	case EnginesBuildsProperty:
		return value.(pcgw_integration.EnginesBuildsGetter).GetEnginesBuilds()
	case FeaturesProperty:
		return value.(gog_integration.FeaturesGetter).GetFeatures()
	case ForumUrlProperty:
		return getSlice(value.(gog_integration.ForumUrlGetter).GetForumUrl)
	case IGDBIdProperty:
		return getSlice(value.(pcgw_integration.IGDBIdGetter).GetIGDBId)
	case IGNWikiSlugProperty:
		return getSlice(value.(ign_integration.IGNWikiSlugGetter).GetIGNWikiSlug)
	case ImageProperty:
		return getImageIdSlice(value.(gog_integration.ImageGetter).GetImage)
	case IncludesGamesProperty:
		return value.(gog_integration.IncludesGamesGetter).GetIncludesGames()
	case InDevelopmentProperty:
		return boolSlice(value.(gog_integration.InDevelopmentGetter).GetInDevelopment)
	case IsDiscountedProperty:
		return boolSlice(value.(gog_integration.IsDiscountedGetter).IsDiscounted)
	case IsFreeProperty:
		return boolSlice(value.(gog_integration.IsFreeGetter).IsFree)
	case IsIncludedByGamesProperty:
		return value.(gog_integration.IsIncludedInGamesGetter).GetIsIncludedInGames()
	case IsRequiredByGamesProperty:
		return value.(gog_integration.IsRequiredByGamesGetter).GetIsRequiredByGames()
	case GenresProperty:
		return value.(gog_integration.GenresGetter).GetGenres()
	case GlobalReleaseDateProperty:
		return getSlice(value.(gog_integration.GlobalReleaseGetter).GetGlobalRelease)
	case GOGReleaseDateProperty:
		return getSlice(value.(gog_integration.GOGReleaseGetter).GetGOGRelease)
	case HLTBIdProperty:
		return getSlice(value.(pcgw_integration.HLTBIdGetter).GetHLTBId)
	case HLTBBuildIdProperty:
		return getSlice(value.(hltb_integration.BuildIdGetter).GetBuildId)
	case HLTBHoursToCompleteMainProperty:
		return getSlice(value.(hltb_integration.HoursToCompleteMainGetter).GetHoursToCompleteMain)
	case HLTBHoursToCompletePlusProperty:
		return getSlice(value.(hltb_integration.HoursToCompletePlusGetter).GetHoursToCompletePlus)
	case HLTBHoursToComplete100Property:
		return getSlice(value.(hltb_integration.HoursToComplete100Getter).GetHoursToComplete100)
	case HLTBReviewScoreProperty:
		return intSlice(value.(hltb_integration.ReviewScoreGetter).GetReviewScore)
	case HLTBGenresProperty:
		return value.(gog_integration.GenresGetter).GetGenres()
	case HLTBPlatformsProperty:
		return value.(hltb_integration.PlatformsGetter).GetPlatforms()
	case LanguageCodeProperty:
		return value.(gog_integration.LanguageCodesGetter).GetLanguageCodes()
	case MobyGamesIdProperty:
		return getSlice(value.(pcgw_integration.MobyGamesIdGetter).GetMobyGamesId)
	case OperatingSystemsProperty:
		return value.(gog_integration.OperatingSystemsGetter).GetOperatingSystems()
	case PCGWPageIdProperty:
		return getSlice(value.(pcgw_integration.PageIdGetter).GetPageId)
	case PreOrderProperty:
		return boolSlice(value.(gog_integration.PreOrderGetter).GetPreOrder)
	case PriceProperty:
		return getSlice(value.(gog_integration.PriceGetter).GetPrice)
	case ProductTypeProperty:
		return getSlice(value.(gog_integration.ProductTypeGetter).GetProductType)
	case PublishersProperty:
		return value.(gog_integration.PublishersGetter).GetPublishers()
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
	case SteamAppIdProperty:
		return uint32Slice(value.(steam_integration.SteamAppIdGetter).GetSteamAppId)
	case SteamReviewScoreDescProperty:
		return getSlice(value.(steam_integration.ReviewScoreDescGetter).GetReviewScoreDesc)
	case SteamTagsProperty:
		return value.(steam_integration.SteamTagsGetter).GetSteamTags()
	case StoreTagsProperty:
		return value.(gog_integration.StoreTagsGetter).GetStoreTags()
	case StoreUrlProperty:
		return getSlice(value.(gog_integration.StoreUrlGetter).GetStoreUrl)
	case StrategyWikiIdProperty:
		return getSlice(value.(pcgw_integration.StrategyWikiIdGetter).GetStrategyWikiId)
	case SupportUrlProperty:
		return getSlice(value.(gog_integration.SupportUrlGetter).GetSupportUrl)
	case TagIdProperty:
		return value.(gog_integration.TagIdsGetter).GetTagIds()
	case TitleProperty:
		return getSlice(value.(gog_integration.TitleGetter).GetTitle)
	case VerticalImageProperty:
		return getImageIdSlice(value.(gog_integration.VerticalImageGetter).GetVerticalImage)
	case VideoIdProperty:
		return value.(gog_integration.VideoIdsGetter).GetVideoIds()
	case VNDBIdProperty:
		return getSlice(value.(pcgw_integration.VNDBIdGetter).GetVNDBId)
	case WikipediaIdProperty:
		return getSlice(value.(pcgw_integration.WikipediaIdGetter).GetWikipediaId)
	case WineHQIdProperty:
		return getSlice(value.(pcgw_integration.WineHQIdGetter).GetWineHQId)
	default:
		return []string{}
	}
}

func boolSlice(confirmer func() bool) []string {
	facts := make([]string, 0)
	if confirmer != nil {
		val := FalseValue
		if confirmer() {
			val = TrueValue
		}
		facts = append(facts, val)
	}
	return facts
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

func intSlice(integer func() int) []string {
	values := make([]string, 0)
	if integer != nil {
		values = append(values, strconv.FormatInt(int64(integer()), 10))
	}
	return values
}

func uint32Slice(integer func() uint32) []string {
	values := make([]string, 0)
	if integer != nil {
		values = append(values, strconv.FormatUint(uint64(integer()), 10))
	}
	return values
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
