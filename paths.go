package vangogh_local_data

import (
	"github.com/arelate/gog_integration"
	"github.com/boggydigital/yt_urls"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const xmlExt = ".xml"
const skipListFilename = "skiplist.txt"

var validatedExtensions = map[string]bool{
	".exe": true,
	".bin": true,
	".dmg": true,
	".pkg": true,
	".sh":  true,
}

func RemoteChecksumPath(p string) string {
	ext := path.Ext(p)
	if validatedExtensions[ext] {
		return p + xmlExt
	}
	return ""
}

func AbsLocalChecksumPath(p string) string {
	ext := path.Ext(p)
	if !validatedExtensions[ext] {
		return ""
	}
	dir, filename := path.Split(p)
	if strings.HasPrefix(dir, AbsDownloadsDir()) {
		dir = strings.Replace(dir, AbsDownloadsDir(), AbsChecksumsDir(), 1)
	} else {
		dir = filepath.Join(AbsChecksumsDir(), dir)
	}
	return filepath.Join(dir, filename+xmlExt)
}

func absLocalVideoPath(videoId string, videoDirDelegate func(videoId string) string, ext string) string {
	videoPath := filepath.Join(videoDirDelegate(videoId), videoId+ext)

	if _, err := os.Stat(videoPath); err == nil {
		return videoPath
	}

	return ""
}

func AbsLocalVideoPath(videoId string) string {
	return absLocalVideoPath(videoId, AbsVideoDirByVideoId, yt_urls.DefaultVideoExt)
}

func AbsLocalThumbnailPath(videoId string) string {
	return absLocalVideoPath(videoId, AbsThumbnailDirByVideoId, yt_urls.DefaultThumbnailExt)
}

func relRecycleBinPath(p string) (string, error) {
	return filepath.Rel(AbsRecycleBinDir(), p)
}

func AbsSkipListPath() string {
	return filepath.Join(absRootDir, skipListFilename)
}

func AbsLocalImagePath(imageId string) string {
	dir := AbsDirByImageId(imageId)

	jpgPath := filepath.Join(dir, imageId+gog_integration.JpgExt)

	if _, err := os.Stat(jpgPath); err == nil {
		return jpgPath
	} else if os.IsNotExist(err) {
		pngPath := filepath.Join(dir, imageId+gog_integration.PngExt)
		if _, err := os.Stat(pngPath); err == nil {
			return pngPath
		}
	}

	return ""
}
