package vangogh_local_data

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	relVideoThumbnailsDir = "_thumbnails"
	relExtrasDir          = "extras"
	relDLCDir             = "dlc"
	relChecksumsDir       = "_checksums"
	relReduxDir           = "_redux"
)

var (
	//absRootDir       = ""

	absBackupsDir    = ""
	absDownloadsDir  = ""
	absImagesDir     = ""
	absInputFilesDir = ""
	absRecycleBinDir = ""
	absMetadataDir   = ""
	absOutputDir     = ""
	absItemsDir      = ""
	absVideosDir     = ""
	absTempDir       = ""
)

//func ChRoot(rd string) {
//	absRootDir = rd
//}
//
//func Pwd() string {
//	return absRootDir
//}

func SetBackupsDir(d string) {
	absBackupsDir = d
}

func SetDownloadsDir(d string) {
	absDownloadsDir = d
}

func SetImagesDir(d string) {
	absImagesDir = d
}

func SetInputFilesDir(d string) {
	absInputFilesDir = d
}

func SetItemsDir(d string) {
	absItemsDir = d
}

//func SetTempDir(d string) {
//	absTempDir = d
//}

func SetMetadataDir(d string) {
	absMetadataDir = d
}

func SetOutputDir(d string) {
	absOutputDir = d
}

func SetRecycleBinDir(d string) {
	absRecycleBinDir = d
}

func SetVideosDir(d string) {
	absVideosDir = d
}

//func AbsTempDir() string {
//	return absTempDir
//}

func AbsVideosDir() string {
	return absVideosDir
}

func AbsVideoThumbnailsDir() string {
	return filepath.Join(absVideosDir, relVideoThumbnailsDir)
}

func AbsImagesDir() string {
	return absImagesDir
}

func AbsItemsDir() string {
	return absItemsDir
}

func AbsMetadataDir() string {
	return absMetadataDir
}

func AbsReduxDir() string {
	return filepath.Join(AbsMetadataDir(), relReduxDir)
}

func AbsRecycleBinDir() string {
	return absRecycleBinDir
}

func AbsDownloadsDir() string {
	return absDownloadsDir
}

func RelExtrasDir() string {
	return relExtrasDir
}

func RelDLCDir() string {
	return relDLCDir
}

func AbsChecksumsDir() string {
	return filepath.Join(absDownloadsDir, relChecksumsDir)
}

func absDirByVideoId(videoId string, absDirDelegate func() string) string {
	if videoId == "" || len(videoId) < 1 {
		return ""
	}

	return filepath.Join(absDirDelegate(), strings.ToLower(videoId[0:1]))
}

func AbsVideoDirByVideoId(videoId string) string {
	return absDirByVideoId(videoId, AbsVideosDir)
}

func AbsVideoThumbnailsDirByVideoId(videoId string) string {
	return absDirByVideoId(videoId, AbsVideoThumbnailsDir)
}

func absDirByImageId(imageId string, absDirDelegate func() string) string {
	if imageId == "" {
		return ""
	}

	imageId = strings.TrimPrefix(imageId, "/")

	if len(imageId) < 2 {
		return ""
	}

	return filepath.Join(absDirDelegate(), imageId[0:2])
}

func AbsImagesDirByImageId(imageId string) string {
	return absDirByImageId(imageId, AbsImagesDir)
}

func AbsItemPath(path string) string {
	if path == "" {
		return ""
	}

	//GOG.com quirk - some item URLs path has multiple slashes
	//e.g. https://items.gog.com//atom_rpg_trudograd/mp4/TGWMap_Night_%281%29.gif.mp4
	//so we need to keep trimming while there is something to trim
	for strings.HasPrefix(path, "/") {
		path = strings.TrimPrefix(path, "/")
	}

	if len(path) < 1 {
		return ""
	}

	return filepath.Join(AbsItemsDir(), path[0:1], path)
}

func AbsLocalProductTypeDir(pt ProductType) (string, error) {
	if !IsValidProductType(pt) {
		return "", fmt.Errorf("no local destination for product type %s", pt)
	}

	return filepath.Join(AbsMetadataDir(), pt.String()), nil
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
	return filepath.Join(AbsDownloadsDir(), p)
}

func AbsBackupsDir() string {
	return absBackupsDir
}

func AbsOutputDir() string {
	return absOutputDir
}
