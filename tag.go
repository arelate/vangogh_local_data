package vangogh_local_data

import (
	"encoding/json"
	"fmt"
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/kvas"
	"github.com/boggydigital/nod"
	"net/http"
	"net/url"
)

func postTagResp(httpClient *http.Client, url *url.URL, respVal interface{}) error {
	resp, err := httpClient.Post(url.String(), "", nil)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(&respVal)
}

func TagIdByName(tagName string, rxa kvas.ReduxAssets) (string, error) {
	if err := rxa.IsSupported(TagNameProperty); err != nil {
		return "", err
	}

	tagIds := rxa.Match(map[string][]string{TagNameProperty: {tagName}}, true)
	if len(tagIds) == 0 {
		return "", fmt.Errorf("unknown tag-name %s", tagName)
	}
	if len(tagIds) > 1 {
		return "", fmt.Errorf("ambiguous tag-name %s, matching tag-ids: %v",
			tagName,
			tagIds)
	}
	tagId := ""
	for ti := range tagIds {
		tagId = ti
	}
	return tagId, nil
}

func CreateTag(httpClient *http.Client, tagName string, rxa kvas.ReduxAssets) error {

	if err := rxa.IsSupported(TagNameProperty); err != nil {
		return err
	}

	createTagUrl := gog_integration.CreateTagUrl(tagName)
	var ctResp gog_integration.CreateTagResp
	if err := postTagResp(httpClient, createTagUrl, &ctResp); err != nil {
		return err
	}
	if ctResp.Id == "" {
		return fmt.Errorf("invalid create tag response")
	}

	if !rxa.HasVal(TagNameProperty, ctResp.Id, tagName) {
		if err := rxa.AddVal(TagNameProperty, ctResp.Id, tagName); err != nil {
			return err
		}
	}

	return nil
}

func DeleteTag(httpClient *http.Client, tagName, tagId string, rxa kvas.ReduxAssets) error {

	if err := rxa.IsSupported(TagNameProperty); err != nil {
		return err
	}

	deleteTagUrl := gog_integration.DeleteTagUrl(tagId)
	var dtResp gog_integration.DeleteTagResp
	if err := postTagResp(httpClient, deleteTagUrl, &dtResp); err != nil {
		return err
	}
	if dtResp.Status != "deleted" {
		return fmt.Errorf("invalid delete tag response")
	}

	if rxa.HasVal(TagNameProperty, tagId, tagName) {
		if err := rxa.CutVal(TagNameProperty, tagId, tagName); err != nil {
			return err
		}
	}

	return nil
}

func AddTag(
	httpClient *http.Client,
	idSet map[string]bool,
	tagId string,
	rxa kvas.ReduxAssets,
	tpw nod.TotalProgressWriter) error {

	if err := rxa.IsSupported(TagIdProperty); err != nil {
		return err
	}

	if tpw != nil {
		tpw.TotalInt(len(idSet))
	}

	for id := range idSet {

		if rxa.HasVal(TagIdProperty, id, tagId) {
			if tpw != nil {
				tpw.Increment()
			}
			continue
		}

		addTagUrl := gog_integration.AddTagUrl(id, tagId)
		var artResp gog_integration.AddRemoveTagResp
		if err := postTagResp(httpClient, addTagUrl, &artResp); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return err
		}
		if !artResp.Success {
			if tpw != nil {
				tpw.Increment()
			}
			return fmt.Errorf("failed to add tag %s", tagId)
		}

		if err := rxa.AddVal(TagIdProperty, id, tagId); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return err
		}

		if tpw != nil {
			tpw.Increment()
		}
	}

	return nil
}

func RemoveTag(
	httpClient *http.Client,
	idSet map[string]bool,
	tagId string,
	rxa kvas.ReduxAssets,
	tpw nod.TotalProgressWriter) error {

	if err := rxa.IsSupported(TagIdProperty); err != nil {
		return err
	}

	if tpw != nil {
		tpw.TotalInt(len(idSet))
	}

	for id := range idSet {

		if !rxa.HasVal(TagIdProperty, id, tagId) {
			if tpw != nil {
				tpw.Increment()
			}
			continue
		}

		removeTagUrl := gog_integration.RemoveTagUrl(id, tagId)
		var artResp gog_integration.AddRemoveTagResp
		if err := postTagResp(httpClient, removeTagUrl, &artResp); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return err
		}
		if !artResp.Success {
			if tpw != nil {
				tpw.Increment()
			}
			return fmt.Errorf("failed to remove tag %s", tagId)
		}

		if err := rxa.CutVal(TagIdProperty, id, tagId); err != nil {
			if tpw != nil {
				tpw.Increment()
			}
			return err
		}

		if tpw != nil {
			tpw.Increment()
		}
	}

	return nil
}
