package vangogh_local_data

import (
	"fmt"
	"github.com/boggydigital/pasu"
	"path/filepath"
	"strings"
)

const DefaultVangoghRootDir = "/var/lib/vangogh"

const (
	Backups    pasu.AbsDir = "backups"
	Metadata   pasu.AbsDir = "metadata"
	Input      pasu.AbsDir = "input"
	Output     pasu.AbsDir = "output"
	Images     pasu.AbsDir = "images"
	Videos     pasu.AbsDir = "videos"
	Items      pasu.AbsDir = "items"
	Downloads  pasu.AbsDir = "downloads"
	RecycleBin pasu.AbsDir = "recycle_bin"
	Logs       pasu.AbsDir = "logs"
)

var AllAbsDirs = []pasu.AbsDir{
	Backups,
	Metadata,
	Input,
	Output,
	Images,
	Videos,
	Items,
	Downloads,
	RecycleBin,
	Logs,
}

const (
	Redux           pasu.RelDir = "_redux"
	Checksums       pasu.RelDir = "_checksums"
	DLCs            pasu.RelDir = "dlc"
	Extras          pasu.RelDir = "extras"
	VideoThumbnails pasu.RelDir = "_thumbnails"
)

var RelToAbsDirs = map[pasu.RelDir]pasu.AbsDir{
	Redux:           Metadata,
	Checksums:       Downloads,
	DLCs:            Downloads,
	Extras:          Downloads,
	VideoThumbnails: Videos,
}

func AbsVideoDirByVideoId(videoId string) (string, error) {
	if videoId == "" || len(videoId) < 1 {
		return "", fmt.Errorf("videoId cannot be empty")
	}
	vdp, err := pasu.GetAbsDir(Videos)
	return filepath.Join(vdp, strings.ToLower(videoId[0:1])), err
}

func AbsVideoThumbnailsDirByVideoId(videoId string) (string, error) {
	if videoId == "" || len(videoId) < 1 {
		return "", fmt.Errorf("videoId cannot be empty")
	}
	vdp, err := pasu.GetAbsRelDir(VideoThumbnails)
	return filepath.Join(vdp, strings.ToLower(videoId[0:1])), err
}

func AbsImagesDirByImageId(imageId string) (string, error) {
	if imageId == "" {
		return "", fmt.Errorf("imageId cannot be empty")
	}

	imageId = strings.TrimPrefix(imageId, "/")

	if len(imageId) < 2 {
		return "", fmt.Errorf("imageId is too short")
	}

	idp, err := pasu.GetAbsDir(Images)
	return filepath.Join(idp, imageId[0:2]), err
}

func AbsItemPath(path string) (string, error) {
	if path == "" {
		return "", fmt.Errorf("item path cannot be empty")
	}

	//GOG.com quirk - some item URLs path has multiple slashes
	//e.g. https://items.gog.com//atom_rpg_trudograd/mp4/TGWMap_Night_%281%29.gif.mp4
	//so we need to keep trimming while there is something to trim
	for strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}
	if len(path) < 1 {
		return "", fmt.Errorf("sanitized item path cannot be empty")
	}

	idp, err := pasu.GetAbsDir(Items)
	if err != nil {
		return "", err
	}

	return filepath.Join(idp, path[0:1], path), nil
}

func AbsLocalProductTypeDir(pt ProductType) (string, error) {
	if !IsValidProductType(pt) {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd, err := pasu.GetAbsDir(Metadata)
	if err != nil {
		return "", err
	}
	return filepath.Join(amd, pt.String()), nil
}

func RelProductDownloadsDir(slug string) (string, error) {
	if slug == "" {
		return "", fmt.Errorf("vangogh_urls: empty slug")
	}
	if len(slug) < 1 {
		return "", fmt.Errorf("vangogh_urls: slug is too short")
	}
	return filepath.Join(strings.ToLower(slug[0:1]), slug), nil
}

func AbsProductDownloadsDir(slug string) (string, error) {
	rDir, err := RelProductDownloadsDir(slug)
	if err != nil {
		return rDir, err
	}
	return AbsDownloadDirFromRel(rDir)
}

func AbsDownloadDirFromRel(p string) (string, error) {
	adp, err := pasu.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}
	return filepath.Join(adp, p), nil
}
