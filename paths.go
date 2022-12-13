package vangogh_local_data

import (
	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/yt_urls"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	xmlExt           = ".xml"
	skipListFilename = "skiplist.txt"
	cookiesFilename  = "cookies.txt"
	atomFeedFilename = "atom.xml"
)

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

func AbsLocalVideoThumbnailPath(videoId string) string {
	return absLocalVideoPath(videoId, AbsVideoThumbnailsDirByVideoId, yt_urls.DefaultThumbnailExt)
}

func relRecycleBinPath(p string) (string, error) {
	return filepath.Rel(AbsRecycleBinDir(), p)
}

func AbsSkipListPath() string {
	return filepath.Join(absRootDir, skipListFilename)
}

func absLocalImagePath(imageId string, imageDirDelegate func(imageId string) string, ext string) string {
	imagePath := filepath.Join(imageDirDelegate(imageId), imageId+ext)

	if _, err := os.Stat(imagePath); err == nil {
		return imagePath
	} else {
		return ""
	}
}

func AbsLocalImagePath(imageId string) string {
	exts := []string{gog_integration.JpgExt, gog_integration.PngExt}
	for _, ext := range exts {
		aip := absLocalImagePath(imageId, AbsImagesDirByImageId, ext)
		if _, err := os.Stat(aip); err == nil {
			return aip
		}
	}
	return ""
}

func AbsTempPath(f string) string {
	return filepath.Join(absTempDir, f)
}

func AbsCookiePath() string {
	return AbsTempPath(cookiesFilename)
}

func AbsAtomFeedPath() string {
	return AbsTempPath(atomFeedFilename)
}
