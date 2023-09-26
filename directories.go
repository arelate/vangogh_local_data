package vangogh_local_data

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type AbsDir int

const (
	Backups AbsDir = iota
	Downloads
	Images
	InputFiles
	Items
	Metadata
	OutputFiles
	RecycleBin
	Videos
)

var absDirsStrings = map[AbsDir]string{
	Backups:     "backups",
	Downloads:   "downloads",
	Images:      "images",
	InputFiles:  "input_files",
	Items:       "items",
	Metadata:    "metadata",
	OutputFiles: "output_files",
	RecycleBin:  "recycle_bin",
	Videos:      "videos",
}

var absDirsPaths = map[AbsDir]string{}

type RelDir int

const (
	Checksums RelDir = iota
	DLCs
	Extras
	Redux
	VideoThumbnails
)

var relDirsStrings = map[RelDir]string{
	Checksums:       "_checksums",
	DLCs:            "dlc",
	Extras:          "extras",
	Redux:           "_redux",
	VideoThumbnails: "_thumbnails",
}

var relToAbsDir = map[RelDir]AbsDir{
	Checksums:       Downloads,
	Redux:           Metadata,
	VideoThumbnails: Videos,
}

func GetRelDir(rd RelDir) (string, error) {
	if rds, ok := relDirsStrings[rd]; ok && rds != "" {
		return rds, nil
	} else {
		return "", fmt.Errorf("unknown rel dir")
	}
}

func SetAbsDirs(kv map[string]string) error {
	for adk, ads := range absDirsStrings {
		if d, ok := kv[ads]; ok && d != "" {
			// make sure directory exists
			if _, err := os.Stat(d); err != nil {
				return err
			}
			absDirsPaths[adk] = d
		} else {
			return fmt.Errorf("missing required abs dir %s", ads)
		}
	}
	return nil
}

func GetAbsDir(ad AbsDir) (string, error) {
	if _, ok := absDirsStrings[ad]; !ok {
		return "", fmt.Errorf("unknown abs dir")
	}

	if adp, ok := absDirsPaths[ad]; ok && adp != "" {
		return adp, nil
	}
	return "", fmt.Errorf("abs dir %s not set", absDirsStrings[ad])
}

func GetAbsRelDir(rd RelDir) (string, error) {
	if _, ok := relDirsStrings[rd]; ok {
		return "", fmt.Errorf("unknown rel dir")
	}

	if ad, ok := relToAbsDir[rd]; ok {

		adp, err := GetAbsDir(ad)
		if err != nil {
			return "", err
		}

		return filepath.Join(adp, relDirsStrings[rd]), nil
	} else {
		return "", fmt.Errorf("%s dir relativity not set", relDirsStrings[rd])
	}
}

func AbsVideoDirByVideoId(videoId string) (string, error) {
	if videoId == "" || len(videoId) < 1 {
		return "", fmt.Errorf("videoId cannot be empty")
	}
	vdp, err := GetAbsDir(Videos)
	return filepath.Join(vdp, strings.ToLower(videoId[0:1])), err
}

func AbsVideoThumbnailsDirByVideoId(videoId string) (string, error) {
	if videoId == "" || len(videoId) < 1 {
		return "", fmt.Errorf("videoId cannot be empty")
	}
	vdp, err := GetAbsRelDir(VideoThumbnails)
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

	idp, err := GetAbsDir(Images)
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

	idp, err := GetAbsDir(Items)
	if err != nil {
		return "", err
	}

	return filepath.Join(idp, path[0:1], path), nil
}

func AbsLocalProductTypeDir(pt ProductType) (string, error) {
	if !IsValidProductType(pt) {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	amd, err := GetAbsDir(Metadata)
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
	adp, err := GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}
	return filepath.Join(adp, p), nil
}
