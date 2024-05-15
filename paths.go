package vangogh_local_data

import (
	"github.com/arelate/southern_light/gog_integration"
	"github.com/boggydigital/pathways"
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

func AbsLocalChecksumPath(p string) (string, error) {
	ext := path.Ext(p)
	if !validatedExtensions[ext] {
		return "", nil
	}
	dir, filename := path.Split(p)
	adp, err := pathways.GetAbsDir(Downloads)
	if err != nil {
		return "", err
	}
	cdp, err := pathways.GetAbsRelDir(Checksums)
	if err != nil {
		return "", err
	}

	if strings.HasPrefix(dir, adp) {
		dir = strings.Replace(dir, adp, cdp, 1)
	} else {
		dir = filepath.Join(cdp, dir)
	}
	return filepath.Join(dir, filename+xmlExt), nil
}

func absLocalVideoPath(videoId string, videoDir string, ext string) string {
	videoPath := filepath.Join(videoDir, videoId+ext)

	if _, err := os.Stat(videoPath); err == nil {
		return videoPath
	}

	return ""

}

func AbsLocalVideoPath(videoId string) (string, error) {
	vdp, err := AbsVideoDirByVideoId(videoId)
	if err != nil {
		return "", err
	}
	return absLocalVideoPath(videoId, vdp, yt_urls.DefaultVideoExt), nil
}

func AbsLocalVideoThumbnailPath(videoId string) (string, error) {
	vtdp, err := AbsVideoThumbnailsDirByVideoId(videoId)
	if err != nil {
		return "", err
	}
	return absLocalVideoPath(videoId, vtdp, yt_urls.DefaultThumbnailExt), nil
}

func relRecycleBinPath(p string) (string, error) {
	rbdp, err := pathways.GetAbsDir(RecycleBin)
	if err != nil {
		return "", err
	}
	return filepath.Rel(rbdp, p)
}

func AbsSkipListPath() (string, error) {
	ifdp, err := pathways.GetAbsDir(Input)
	return filepath.Join(ifdp, skipListFilename), err
}

func AbsLocalImagePath(imageId string) (string, error) {
	exts := []string{gog_integration.JpgExt, gog_integration.PngExt}
	idp, err := AbsImagesDirByImageId(imageId)
	if err != nil {
		return "", err
	}
	for _, ext := range exts {
		aip := filepath.Join(idp, imageId+ext)
		if _, err := os.Stat(aip); err == nil {
			return aip, nil
		} else {
			continue
		}
	}
	return "", err
}

func AbsCookiePath() (string, error) {
	ifdp, err := pathways.GetAbsDir(Input)
	return filepath.Join(ifdp, cookiesFilename), err
}

func AbsAtomFeedPath() (string, error) {
	ofdp, err := pathways.GetAbsDir(Output)
	return filepath.Join(ofdp, atomFeedFilename), err
}
