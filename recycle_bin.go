package vangogh_local_data

import (
	"github.com/boggydigital/nod"
	"os"
	"path/filepath"
)

func MoveToRecycleBin(typeRootDir, absPath string) error {
	if absPath == "" {
		nod.Log("move to recycle bin: empty filename")
		return nil
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		nod.Log("move to recycle bin: file not found: %s", absPath)
		return nil
	}

	root, _ := filepath.Split(typeRootDir)

	relPath, err := filepath.Rel(root, absPath)
	if err != nil {
		nod.Log("move to recycle bin: relPath error: root=%s target=%s", root, absPath)
		return err
	}

	rbdp, err := GetAbsDir(RecycleBin)
	if err != nil {
		return err
	}

	rbFilepath := filepath.Join(rbdp, relPath)
	rbDir, _ := filepath.Split(rbFilepath)
	if _, err := os.Stat(rbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(rbDir, 0755); err != nil {
			return err
		}
	}
	return os.Rename(absPath, rbFilepath)
}
