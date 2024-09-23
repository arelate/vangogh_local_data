package vangogh_local_data

import (
	"fmt"
	"github.com/arelate/southern_light/gog_integration"
	"github.com/arelate/southern_light/hltb_integration"
	"github.com/arelate/southern_light/ign_integration"
	"github.com/arelate/southern_light/pcgw_integration"
	"github.com/arelate/southern_light/protondb_integration"
	"github.com/arelate/southern_light/steam_integration"
	"strconv"
	"time"
)

const (
	IdProperty                                    = "id"
	TitleProperty                                 = "title"
	DevelopersProperty                            = "developers"
	PublishersProperty                            = "publishers"
	ImageProperty                                 = "image"
	VerticalImageProperty                         = "vertical-image"
	ScreenshotsProperty                           = "screenshots"
	RatingProperty                                = "rating"
	IncludesGamesProperty                         = "includes-games"
	IsIncludedByGamesProperty                     = "is-included-by-games"
	RequiresGamesProperty                         = "requires-games"
	IsRequiredByGamesProperty                     = "is-required-by-games"
	GenresProperty                                = "genres"
	StoreTagsProperty                             = "store-tags"
	FeaturesProperty                              = "features"
	SeriesProperty                                = "series"
	TagIdProperty                                 = "tag"
	TagNameProperty                               = "tag-name"
	VideoIdProperty                               = "video-id"
	MissingVideoUrlProperty                       = "missing-video-url"
	MissingVideoThumbnailProperty                 = "missing-video-thumbnail"
	OperatingSystemsProperty                      = "os"
	LanguageCodeProperty                          = "lang-code"
	LanguageNameProperty                          = "lang-name"
	NativeLanguageNameProperty                    = "native-lang-name"
	SlugProperty                                  = "slug"
	GOGReleaseDateProperty                        = "gog-release-date"
	GOGOrderDateProperty                          = "gog-order-date"
	GlobalReleaseDateProperty                     = "global-release-date"
	TypesProperty                                 = "types"
	LocalManualUrlProperty                        = "local-manual-url"
	DownloadStatusErrorProperty                   = "download-status-error"
	StoreUrlProperty                              = "store-url"
	ForumUrlProperty                              = "forum-url"
	SupportUrlProperty                            = "support-url"
	ChangelogProperty                             = "changelog"
	DescriptionOverviewProperty                   = "description-overview"
	DescriptionFeaturesProperty                   = "description-features"
	AdditionalRequirementsProperty                = "additional-requirements"
	CopyrightsProperty                            = "copyrights"
	WishlistedProperty                            = "wishlisted"
	OwnedProperty                                 = "owned"
	ProductTypeProperty                           = "product-type"
	InDevelopmentProperty                         = "in-development"
	PreOrderProperty                              = "pre-order"
	ComingSoonProperty                            = "coming-soon"
	BasePriceProperty                             = "base-price"
	PriceProperty                                 = "price"
	IsFreeProperty                                = "is-free"
	IsDiscountedProperty                          = "is-discounted"
	DiscountPercentageProperty                    = "discount-percentage"
	SteamAppIdProperty                            = "steam-app-id"
	LocalTagsProperty                             = "local-tags"
	SortProperty                                  = "sort"
	DescendingProperty                            = "desc"
	SteamReviewScoreDescProperty                  = "steam-review-score-desc"
	SteamTagsProperty                             = "steam-tags"
	SteamDeckAppCompatibilityCategoryProperty     = "steam-deck-app-compatibility-category"
	SteamDeckAppCompatibilityResultsProperty      = "steam-deck-app-compatibility-results"
	SteamDeckAppCompatibilityDisplayTypesProperty = "steam-deck-app-compatibility-display-types"
	SteamDeckAppCompatibilityBlogUrlProperty      = "steam-deck-app-compatibility-blog-url"
	DehydratedImageProperty                       = "dehydrated-image"
	DehydratedImageModifiedProperty               = "dehydrated-image-modified"
	DehydratedVerticalImageProperty               = "dehydrated-vertical-image"
	DehydratedVerticalImageModifiedProperty       = "dehydrated-vertical-image-modified"
	SyncEventsProperty                            = "sync-events"
	LastSyncUpdatesProperty                       = "last-sync-updates"
	ValidationResultProperty                      = "validation-result"
	ValidationCompletedProperty                   = "validation-completed"
	PCGWPageIdProperty                            = "pcgw-page-id"
	HLTBIdProperty                                = "hltb-id"
	HLTBBuildIdProperty                           = "hltb-next-build"
	HLTBHoursToCompleteMainProperty               = "hltb-comp-main"
	HLTBHoursToCompletePlusProperty               = "hltb-comp-plus"
	HLTBHoursToComplete100Property                = "hltb-comp-100"
	HLTBReviewScoreProperty                       = "hltb-review-score"
	HLTBGenresProperty                            = "hltb-genres"
	HLTBPlatformsProperty                         = "hltb-platforms"
	IGDBIdProperty                                = "igdb-id"
	StrategyWikiIdProperty                        = "strategy-wiki-id"
	MobyGamesIdProperty                           = "moby-games-id"
	WikipediaIdProperty                           = "wikipedia-id"
	WineHQIdProperty                              = "winehq-id"
	VNDBIdProperty                                = "vndb-id"
	IGNWikiSlugProperty                           = "ign-wiki-slug"
	EnginesProperty                               = "engines"
	EnginesBuildsProperty                         = "engines-builds"
	ProtonDBTierProperty                          = "protondb-tier"
	ProtonDBConfidenceProperty                    = "protondb-confidence"

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
		TagNameProperty,
		LocalTagsProperty,
		OperatingSystemsProperty,
		LanguageCodeProperty,
		SlugProperty,
		GlobalReleaseDateProperty,
		GOGOrderDateProperty,
		GOGReleaseDateProperty,
		CopyrightsProperty,
	)
}

func VideoIdProperties() []string {
	return []string{
		VideoIdProperty,
		MissingVideoUrlProperty,
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
		SteamDeckAppCompatibilityCategoryProperty,
		SteamDeckAppCompatibilityResultsProperty,
		SteamDeckAppCompatibilityDisplayTypesProperty,
		SteamDeckAppCompatibilityBlogUrlProperty,
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
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
	}
}

func DehydratedImagesProperties() []string {
	return []string{
		DehydratedImageProperty,
		DehydratedVerticalImageProperty,
	}
}

func SyncProperties() []string {
	return []string{
		LastSyncUpdatesProperty,
		SyncEventsProperty,
	}
}

func DownloadProperties() []string {
	return []string{
		LocalManualUrlProperty,
	}
}

func ValidationProperties() []string {
	return []string{
		ValidationResultProperty,
		ValidationCompletedProperty,
	}
}

func ReduxProperties() []string {
	all := TextProperties()
	all = append(all, AllTextProperties()...)
	all = append(all, VideoIdProperties()...)
	all = append(all, ComputedProperties()...)
	all = append(all, ImageIdProperties()...)
	all = append(all, DehydratedImagesProperties()...)
	all = append(all, UrlProperties()...)
	all = append(all, LongTextProperties()...)
	all = append(all, AvailabilityProperties()...)
	all = append(all, AccountStatusProperties()...)
	all = append(all, AdvancedProductProperties()...)
	all = append(all, PriceProperties()...)
	all = append(all, ExternalDataSourcesProperties()...)
	all = append(all, SyncProperties()...)
	all = append(all, DownloadProperties()...)
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
		SteamDeckAppCompatibilityCategoryProperty,
		HLTBPlatformsProperty,
		HLTBGenresProperty,
		EnginesProperty,
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
	}
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
	SteamDeckCompatibilityReport: {
		SteamDeckAppCompatibilityCategoryProperty,
		SteamDeckAppCompatibilityResultsProperty,
		SteamDeckAppCompatibilityDisplayTypesProperty,
		SteamDeckAppCompatibilityBlogUrlProperty,
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
	ProtonDBSummary: {
		ProtonDBTierProperty,
		ProtonDBConfidenceProperty,
	},
}

func GetProperties(
	id string,
	reader *ProductReader,
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
		if gar, ok := value.(gog_integration.AdditionalRequirementsGetter); ok {
			return getSlice(gar.GetAdditionalRequirements)
		}
	case BasePriceProperty:
		if gbp, ok := value.(gog_integration.BasePriceGetter); ok {
			return getSlice(gbp.GetBasePrice)
		}
	case ChangelogProperty:
		if gcl, ok := value.(gog_integration.ChangelogGetter); ok {
			return getSlice(gcl.GetChangelog)
		}
	case ComingSoonProperty:
		if gcc, ok := value.(gog_integration.ComingSoonGetter); ok {
			return boolSlice(gcc.GetComingSoon)
		}
	case CopyrightsProperty:
		if gc, ok := value.(gog_integration.CopyrightsGetter); ok {
			return getSlice(gc.GetCopyrights)
		}
	case DescriptionFeaturesProperty:
		if gdf, ok := value.(gog_integration.DescriptionFeaturesGetter); ok {
			return getSlice(gdf.GetDescriptionFeatures)
		}
	case DescriptionOverviewProperty:
		if gdo, ok := value.(gog_integration.DescriptionOverviewGetter); ok {
			return getSlice(gdo.GetDescriptionOverview)
		}
	case DevelopersProperty:
		return value.(gog_integration.DevelopersGetter).GetDevelopers()
	case DiscountPercentageProperty:
		if gdp, ok := value.(gog_integration.DiscountPercentageGetter); ok {
			return intSlice(gdp.GetDiscountPercentage)
		}
	case EnginesProperty:
		return value.(pcgw_integration.EnginesGetter).GetEngines()
	case EnginesBuildsProperty:
		return value.(pcgw_integration.EnginesBuildsGetter).GetEnginesBuilds()
	case FeaturesProperty:
		return value.(gog_integration.FeaturesGetter).GetFeatures()
	case ForumUrlProperty:
		if gfu, ok := value.(gog_integration.ForumUrlGetter); ok {
			return getSlice(gfu.GetForumUrl)
		}
	case IGDBIdProperty:
		if gii, ok := value.(pcgw_integration.IGDBIdGetter); ok {
			return getSlice(gii.GetIGDBId)
		}
	case IGNWikiSlugProperty:
		if gis, ok := value.(ign_integration.IGNWikiSlugGetter); ok {
			return getSlice(gis.GetIGNWikiSlug)
		}
	case ImageProperty:
		if gi, ok := value.(gog_integration.ImageGetter); ok {
			return getImageIdSlice(gi.GetImage)
		}
	case IncludesGamesProperty:
		return value.(gog_integration.IncludesGamesGetter).GetIncludesGames()
	case InDevelopmentProperty:
		if gid, ok := value.(gog_integration.InDevelopmentGetter); ok {
			return boolSlice(gid.GetInDevelopment)
		}
	case IsDiscountedProperty:
		if id, ok := value.(gog_integration.IsDiscountedGetter); ok {
			return boolSlice(id.IsDiscounted)
		}
	case IsFreeProperty:
		if iff, ok := value.(gog_integration.IsFreeGetter); ok {
			return boolSlice(iff.IsFree)
		}
	case IsIncludedByGamesProperty:
		return value.(gog_integration.IsIncludedInGamesGetter).GetIsIncludedInGames()
	case IsRequiredByGamesProperty:
		return value.(gog_integration.IsRequiredByGamesGetter).GetIsRequiredByGames()
	case GenresProperty:
		return value.(gog_integration.GenresGetter).GetGenres()
	case GlobalReleaseDateProperty:
		if gdr, ok := value.(gog_integration.GlobalReleaseGetter); ok {
			return getSlice(gdr.GetGlobalRelease)
		}
	case GOGReleaseDateProperty:
		if ggr, ok := value.(gog_integration.GOGReleaseGetter); ok {
			return getSlice(ggr.GetGOGRelease)
		}
	case HLTBIdProperty:
		if ghi, ok := value.(pcgw_integration.HLTBIdGetter); ok {
			return getSlice(ghi.GetHLTBId)
		}
	case HLTBBuildIdProperty:
		if gbi, ok := value.(hltb_integration.BuildIdGetter); ok {
			return getSlice(gbi.GetBuildId)
		}
	case HLTBHoursToCompleteMainProperty:
		if ghcm, ok := value.(hltb_integration.HoursToCompleteMainGetter); ok {
			return getSlice(ghcm.GetHoursToCompleteMain)
		}
	case HLTBHoursToCompletePlusProperty:
		if ghcp, ok := value.(hltb_integration.HoursToCompletePlusGetter); ok {
			return getSlice(ghcp.GetHoursToCompletePlus)
		}
	case HLTBHoursToComplete100Property:
		if ghc100, ok := value.(hltb_integration.HoursToComplete100Getter); ok {
			return getSlice(ghc100.GetHoursToComplete100)
		}
	case HLTBReviewScoreProperty:
		if grs, ok := value.(hltb_integration.ReviewScoreGetter); ok {
			return intSlice(grs.GetReviewScore)
		}
	case HLTBGenresProperty:
		return value.(gog_integration.GenresGetter).GetGenres()
	case HLTBPlatformsProperty:
		return value.(hltb_integration.PlatformsGetter).GetPlatforms()
	case LanguageCodeProperty:
		return value.(gog_integration.LanguageCodesGetter).GetLanguageCodes()
	case MobyGamesIdProperty:
		if gmi, ok := value.(pcgw_integration.MobyGamesIdGetter); ok {
			return getSlice(gmi.GetMobyGamesId)
		}
	case OperatingSystemsProperty:
		return value.(gog_integration.OperatingSystemsGetter).GetOperatingSystems()
	case PCGWPageIdProperty:
		if gpi, ok := value.(pcgw_integration.PageIdGetter); ok {
			return getSlice(gpi.GetPageId)
		}
	case PreOrderProperty:
		if gpo, ok := value.(gog_integration.PreOrderGetter); ok {
			return boolSlice(gpo.GetPreOrder)
		}
	case PriceProperty:
		if gp, ok := value.(gog_integration.PriceGetter); ok {
			return getSlice(gp.GetPrice)
		}
	case ProductTypeProperty:
		if gpt, ok := value.(gog_integration.ProductTypeGetter); ok {
			return getSlice(gpt.GetProductType)
		}
	case ProtonDBConfidenceProperty:
		if sum, ok := value.(protondb_integration.ConfidenceGetter); ok {
			return getSlice(sum.GetConfidence)
		}
	case ProtonDBTierProperty:
		if sum, ok := value.(fmt.Stringer); ok {
			return getSlice(sum.String)
		}
	case PublishersProperty:
		return value.(gog_integration.PublishersGetter).GetPublishers()
	case RatingProperty:
		if gr, ok := value.(gog_integration.RatingGetter); ok {
			return getSlice(gr.GetRating)
		}
	case RequiresGamesProperty:
		return value.(gog_integration.RequiresGamesGetter).GetRequiresGames()
	case SeriesProperty:
		if gs, ok := value.(gog_integration.SeriesGetter); ok {
			return getSlice(gs.GetSeries)
		}
	case ScreenshotsProperty:
		return getScreenshots(value)
	case SlugProperty:
		if gs, ok := value.(gog_integration.SlugGetter); ok {
			return getSlice(gs.GetSlug)
		}
	case SteamAppIdProperty:
		if gsai, ok := value.(steam_integration.SteamAppIdGetter); ok {
			return uint32Slice(gsai.GetSteamAppId)
		}
	case SteamReviewScoreDescProperty:
		if grsd, ok := value.(steam_integration.ReviewScoreDescGetter); ok {
			return getSlice(grsd.GetReviewScoreDesc)
		}
	case SteamTagsProperty:
		return value.(steam_integration.SteamTagsGetter).GetSteamTags()
	case SteamDeckAppCompatibilityCategoryProperty:
		if dacr, ok := value.(fmt.Stringer); ok {
			return getSlice(dacr.String)
		}
	case SteamDeckAppCompatibilityResultsProperty:
		if dacr, ok := value.(steam_integration.ResultsGetter); ok {
			return dacr.GetResults()
		}
	case SteamDeckAppCompatibilityDisplayTypesProperty:
		if dacr, ok := value.(steam_integration.DisplayTypesGetter); ok {
			return dacr.GetDisplayTypes()
		}
	case SteamDeckAppCompatibilityBlogUrlProperty:
		if dacr, ok := value.(steam_integration.BlogUrlGetter); ok {
			return getSlice(dacr.GetBlogUrl)
		}
	case StoreTagsProperty:
		return value.(gog_integration.StoreTagsGetter).GetStoreTags()
	case StoreUrlProperty:
		if gsu, ok := value.(gog_integration.StoreUrlGetter); ok {
			return getSlice(gsu.GetStoreUrl)
		}
	case StrategyWikiIdProperty:
		if gswi, ok := value.(pcgw_integration.StrategyWikiIdGetter); ok {
			return getSlice(gswi.GetStrategyWikiId)
		}
	case SupportUrlProperty:
		if gsu, ok := value.(gog_integration.SupportUrlGetter); ok {
			return getSlice(gsu.GetSupportUrl)
		}
	case TagIdProperty:
		return value.(gog_integration.TagIdsGetter).GetTagIds()
	case TitleProperty:
		if gt, ok := value.(gog_integration.TitleGetter); ok {
			return getSlice(gt.GetTitle)
		}
	case VerticalImageProperty:
		if gvi, ok := value.(gog_integration.VerticalImageGetter); ok {
			return getImageIdSlice(gvi.GetVerticalImage)
		}
	case VideoIdProperty:
		return value.(gog_integration.VideoIdsGetter).GetVideoIds()
	case VNDBIdProperty:
		if gvi, ok := value.(pcgw_integration.VNDBIdGetter); ok {
			return getSlice(gvi.GetVNDBId)
		}
	case WikipediaIdProperty:
		if gwi, ok := value.(pcgw_integration.WikipediaIdGetter); ok {
			return getSlice(gwi.GetWikipediaId)
		}
	case WineHQIdProperty:
		if gwi, ok := value.(pcgw_integration.WineHQIdGetter); ok {
			return getSlice(gwi.GetWineHQId)
		}
	}
	return []string{}
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
