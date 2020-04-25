package pkgmgr

import (
	"fmt"
	"os"
	"path"

	"github.com/h-michael/lsm/file"
)

func createSymLink(lsName string, getBinPath func(lsName string) (string, error)) error {
	binDir, err := file.BinDirName()
	if err != nil {
		return err
	}
	binPath, err := getBinPath(lsName)
	if err != nil {
		return err
	}

	if !file.CheckExist(binPath) {
		return fmt.Errorf("\"%s\" does not exist", binPath)
	}

	symLinkPath := path.Join(binDir, lsName)
	if _, err := os.Lstat(symLinkPath); err == nil {
		if err := os.Remove(symLinkPath); err != nil {
			return err
		}
	} else if os.IsNotExist(err) {
		return err
	}
	if err := file.CreateBinDir(); err != nil {
		return err
	}

	if err := os.Symlink(binPath, symLinkPath); err != nil {
		return err
	}
	return nil
}
