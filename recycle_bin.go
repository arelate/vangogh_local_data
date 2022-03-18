package vangogh_local_data

import (
	"os"
	"path/filepath"
)

func MoveToRecycleBin(absPath string) error {
	relPath, err := filepath.Rel(Pwd(), absPath)
	if err != nil {
		return err
	}

	rbFilepath := filepath.Join(AbsRecycleBinDir(), relPath)
	rbDir, _ := filepath.Split(rbFilepath)
	if _, err := os.Stat(rbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(rbDir, 0755); err != nil {
			return err
		}
	}
	return os.Rename(absPath, rbFilepath)
}
