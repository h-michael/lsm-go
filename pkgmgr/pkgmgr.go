package pkgmgr

import (
	"fmt"
	"os"
	"path"

	"github.com/h-michael/lsm/file"
)

func createLsSymLink(lsName string, getBinPath func(lsName string) (string, error)) error {
	binPath, err := getBinPath(lsName)
	if err != nil {
		return err
	}

	if !file.CheckFileExistence(binPath) {
		return fmt.Errorf("\"%s\" does not exist", binPath)
	}

	if err := file.CreateBinDir(); err != nil {
		return err
	}

	symLinkPath, err := lsSymLinkPath(lsName)
	if err != nil {
		return err
	}

	if err := file.RemoveSymLink(symLinkPath); err != nil {
		return err
	}

	if err := os.Symlink(binPath, symLinkPath); err != nil {
		return err
	}
	return nil
}

func lsSymLinkPath(lsName string) (string, error) {
	binDir, err := file.BinDirName()
	if err != nil {
		return "", err
	}
	return path.Join(binDir, lsName), nil
}
