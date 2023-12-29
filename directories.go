package vangogh_local_data

import (
	"fmt"
	"github.com/boggydigital/pathology"
	"path/filepath"
	"strings"
)

const (
	DefaultVangoghRootDir = "/var/lib/vangogh"

	Backups    pathology.AbsDir = "backups"
	Metadata   pathology.AbsDir = "metadata"
	Input      pathology.AbsDir = "input"
	Output     pathology.AbsDir = "output"
	Images     pathology.AbsDir = "images"
	Videos     pathology.AbsDir = "videos"
	Items      pathology.AbsDir = "items"
	Downloads  pathology.AbsDir = "downloads"
	RecycleBin pathology.AbsDir = "recycle_bin"
	Logs       pathology.AbsDir = "logs"
)

var AllAbsDirs = []pathology.AbsDir{
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
	Redux           pathology.RelDir = "_redux"
	Checksums       pathology.RelDir = "_checksums"
	DLCs            pathology.RelDir = "dlc"
	Extras          pathology.RelDir = "extras"
	VideoThumbnails pathology.RelDir = "_thumbnails"
)

var RelToAbsDirs = map[pathology.RelDir]pathology.AbsDir{
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
	vdp, err := pathology.GetAbsDir(Videos)
	return filepath.Join(vdp, strings.ToLower(videoId[0:1])), err
}

func AbsVideoThumbnailsDirByVideoId(videoId string) (string, error) {
	if videoId == "" || len(videoId) < 1 {
		return "", fmt.Errorf("videoId cannot be empty")
	}
	vdp, err := pathology.GetAbsRelDir(VideoThumbnails)
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

	idp, err := pathology.GetAbsDir(Images)
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

	idp, err := pathology.GetAbsDir(Items)
	if err != nil {
		return "", err
	}

	return filepath.Join(idp, path[0:1], path), nil
}

func AbsLocalProductTypeDir(pt ProductType) (string, error) {
	if !IsValidProductType(pt) {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd, err := pathology.GetAbsDir(Metadata)
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
	adp, err := pathology.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}
	return filepath.Join(adp, p), nil
}
