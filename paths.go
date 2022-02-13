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
	if strings.HasPrefix(dir, absDownloadsDir()) {
		dir = strings.Replace(dir, absDownloadsDir(), absChecksumsDir(), 1)
	} else {
		dir = filepath.Join(absChecksumsDir(), dir)
	}
	return filepath.Join(dir, filename+xmlExt)
}

func AbsLocalVideoPath(videoId string) string {
	dir := AbsDirByVideoId(videoId)

	videoPath := filepath.Join(dir, videoId+yt_urls.DefaultExt)

	if _, err := os.Stat(videoPath); err == nil {
		return videoPath
	}

	return ""
}

func relRecycleBinPath(p string) (string, error) {
	return filepath.Rel(AbsRecycleBinDir(), p)
}

func AbsSkipListPath() string {
	return filepath.Join(rootDir, skipListFilename)
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
