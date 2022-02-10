package vangogh_data

import (
	"fmt"
	"github.com/arelate/gog_atu"
	"path"
	"path/filepath"
	"strings"
)

const (
	metadataDir   = "metadata"
	imagesDir     = "images"
	videosDir     = "videos"
	recycleBinDir = "recycle_bin"
	downloadsDir  = "downloads"
	extrasDir     = "extras"
	dlcDir        = "dlc"
	checksumsDir  = "checksums"
	extractsDir   = "_extracts"
)

var rootDir = ""

func ChRoot(rd string) {
	rootDir = rd
}

func Pwd() string {
	return rootDir
}

func absVideosDir() string {
	return filepath.Join(rootDir, videosDir)
}

func absImagesDir() string {
	return filepath.Join(rootDir, imagesDir)
}

func AbsMetadataDir() string {
	return filepath.Join(rootDir, metadataDir)
}

func AbsExtractsDir() string {
	return filepath.Join(AbsMetadataDir(), extractsDir)
}

func AbsRecycleBinDir() string {
	return filepath.Join(rootDir, recycleBinDir)
}

func absDownloadsDir() string {
	return filepath.Join(rootDir, downloadsDir)
}

func RelExtrasDir() string {
	return extrasDir
}

func RelDLCDir() string {
	return dlcDir
}

func absChecksumsDir() string {
	return filepath.Join(rootDir, checksumsDir)
}

func AbsDirByVideoId(videoId string) string {
	if videoId == "" {
		return ""
	}
	if len(videoId) < 1 {
		return ""
	}

	return path.Join(absVideosDir(), strings.ToLower(videoId[0:1]))
}

func AbsDirByImageId(imageId string) string {
	if imageId == "" {
		return ""
	}
	if len(imageId) < 2 {
		return ""
	}

	imageId = strings.TrimPrefix(imageId, "/")

	return filepath.Join(absImagesDir(), imageId[0:2])
}

func AbsLocalProductTypeDir(pt ProductType, mt gog_atu.Media) (string, error) {
	if !ValidProductType(pt) {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}
	if !gog_atu.ValidMedia(mt) {
		return "", fmt.Errorf("no local destination for media %s", pt)
	}

	//Licence* and Order* are media agnostic, so we'll treat them like `game` media
	if pt == Licences ||
		pt == LicenceProducts ||
		pt == Orders ||
		pt == OrderPage {
		mt = gog_atu.Game
	}

	return filepath.Join(AbsMetadataDir(), pt.String(), mt.String()), nil
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
	return AbsDownloadDirFromRel(rDir), nil
}

func AbsDownloadDirFromRel(p string) string {
	return filepath.Join(absDownloadsDir(), p)
}
