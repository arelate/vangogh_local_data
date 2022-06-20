package vangogh_local_data

import (
	"fmt"
	"github.com/arelate/gog_integration"
)

var filterNewProductTypes = map[ProductType]bool{
	Orders: true,
	//not all licence-products have associated api-products-v1/api-products-v2,
	//so in some cases we won't get a meaningful information like a title
	LicenceProducts: true,
	//both ApiProductsVx are not interesting since they correspond to store-products or account-products
	ApiProductsV1: true,
	ApiProductsV2: true,
	//news would typically be created for new products, so new in store should reflect that
	SteamAppNews: true,
	//reviews are not that interesting by themselves
	SteamReviews: true,
	//Steam product pages (currently) are only used for tags
	SteamStorePage: true,
}

var filterUpdatedProductTypes = map[ProductType]bool{
	Orders: true,
	//not all licence-products have associated api-products-v1/api-products-v2,
	//so in some cases we won't get a meaningful information like a title
	LicenceProducts: true,
	//most of the Updates are price changes for a sale, not that interesting for recurring sync
	StoreProducts: true,
	// wishlist-products are basically store-products, so see above
	WishlistProducts: true,
	//meaningful Updates for account products come from details, not account-products
	AccountProducts: true,
	//same as above for those product types
	ApiProductsV1: true,
	ApiProductsV2: true,
	//reviews are not that interesting by themselves
	SteamReviews: true,
	//Steam product pages (currently) are only used for tags
	SteamStorePage: true,
}

func Updates(mt gog_integration.Media, since int64) (map[string]map[string]bool, error) {
	updates := make(map[string]map[string]bool, 0)

	for _, pt := range LocalProducts() {

		vr, err := NewReader(pt, mt)
		if err != nil {
			return updates, err
		}

		if !filterNewProductTypes[pt] {
			categorize(vr.CreatedAfter(since),
				fmt.Sprintf("new in %s", pt.HumanReadableString()),
				updates)
		}

		if !filterUpdatedProductTypes[pt] {
			categorize(vr.ModifiedAfter(since, true),
				fmt.Sprintf("updates in %s", pt.HumanReadableString()),
				updates)
		}
	}

	return updates, nil
}

func categorize(ids []string, cat string, updates map[string]map[string]bool) {
	for _, id := range ids {
		if updates[cat] == nil {
			updates[cat] = make(map[string]bool, 0)
		}
		updates[cat][id] = true
	}
}
