package vangogh_local_data

import (
	"fmt"
	"path/filepath"
	"strings"
)

const (
	relVideosDir          = "videos"
	relVideoThumbnailsDir = "video_thumbnails"
	relRecycleBinDir      = "recycle_bin"
	relDownloadsDir       = "downloads"
	relExtrasDir          = "extras"
	relDLCDir             = "dlc"
	relChecksumsDir       = "checksums"
	relReduxDir           = "_redux"
	relBackupDir          = "backup"
)

var (
	absRootDir     = ""
	absMetadataDir = ""
	absItemsDir    = ""
	absImagesDir   = ""
	absTempDir     = ""
)

func ChRoot(rd string) {
	absRootDir = rd
}

func Pwd() string {
	return absRootDir
}

func SetTempDir(d string) {
	absTempDir = d
}

func SetImagesDir(d string) {
	absImagesDir = d
}

func SetItemsDir(d string) {
	absItemsDir = d
}

func SetMetadataDir(d string) {
	absMetadataDir = d
}

func AbsTempDir() string {
	return absTempDir
}

func absVideosDir() string {
	return filepath.Join(absRootDir, relVideosDir)
}

func absVideoThumbnailsDir() string {
	return filepath.Join(absRootDir, relVideoThumbnailsDir)
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
	return filepath.Join(absRootDir, relRecycleBinDir)
}

func AbsDownloadsDir() string {
	return filepath.Join(absRootDir, relDownloadsDir)
}

func RelExtrasDir() string {
	return relExtrasDir
}

func RelDLCDir() string {
	return relDLCDir
}

func AbsChecksumsDir() string {
	return filepath.Join(absRootDir, relChecksumsDir)
}

func absDirByVideoId(videoId string, absDirDelegate func() string) string {
	if videoId == "" || len(videoId) < 1 {
		return ""
	}

	return filepath.Join(absDirDelegate(), strings.ToLower(videoId[0:1]))
}

func AbsVideoDirByVideoId(videoId string) string {
	return absDirByVideoId(videoId, absVideosDir)
}

func AbsVideoThumbnailsDirByVideoId(videoId string) string {
	return absDirByVideoId(videoId, absVideoThumbnailsDir)
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

func AbsBackupDir() string {
	return filepath.Join(absRootDir, relBackupDir)
}
