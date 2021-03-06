package vangogh_local_data

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var exclude = map[string]bool{
	".DS_Store":   true, // https://en.wikipedia.org/wiki/.DS_Store
	"desktop.ini": true, // https://en.wikipedia.org/wiki/INI_file#History
}

func filenameAsId(p string) (string, error) {
	_, idFile := path.Split(p)
	if !strings.HasSuffix(idFile, ".download") {
		return strings.TrimSuffix(idFile, path.Ext(idFile)), nil
	}
	return "", nil
}

func LocalImageIds() (map[string]bool, error) {
	return walkFiles(absImagesDir(), filenameAsId)
}

func LocalVideoIds() (map[string]bool, error) {
	return walkFiles(absVideosDir(), filenameAsId)
}

func LocalVideoThumbnailIds() (map[string]bool, error) {
	return walkFiles(absVideoThumbnailsDir(), filenameAsId)
}

func RecycleBinDirs() (map[string]bool, error) {
	return walkDirectories(AbsRecycleBinDir())
}

func RecycleBinFiles() (map[string]bool, error) {
	return walkFiles(AbsRecycleBinDir(), relRecycleBinPath)
}

func LocalSlugDownloads(slug string) (map[string]bool, error) {
	pDir, err := AbsProductDownloadsDir(slug)
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(pDir); os.IsNotExist(err) {
		return map[string]bool{}, nil
	}
	return walkFiles(
		pDir,
		func(p string) (string, error) {
			return filepath.Rel(pDir, p)
		})
}

func walkFiles(dir string, transformDelegate func(string) (string, error)) (map[string]bool, error) {
	fileSet := make(map[string]bool)
	err := filepath.WalkDir(
		dir,
		func(p string, de fs.DirEntry, err error) error {
			if de != nil && de.IsDir() {
				return nil
			}
			_, fn := filepath.Split(p)
			if exclude[fn] {
				return nil
			}
			tPath, err := transformDelegate(p)
			if err != nil {
				return err
			}
			if tPath != "" {
				fileSet[tPath] = true
			}
			return err
		})

	return fileSet, err
}

func walkDirectories(rootDir string) (map[string]bool, error) {
	rbd := AbsRecycleBinDir()
	dirSet := make(map[string]bool)
	err := filepath.WalkDir(
		rootDir,
		func(p string, de fs.DirEntry, err error) error {
			if de != nil && !de.IsDir() {
				return nil
			}
			if p == "" || p == rbd {
				return nil
			}
			dirSet[p] = true
			return nil
		})

	return dirSet, err
}
