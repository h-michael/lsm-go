package file

import (
	"os"
	"path"
)

func CheckFileExistence(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func CheckSymLinkExistence(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}

func BinDirName() (string, error) {
	mgrDir, err := mgrDirName()
	if err != nil {
		return "", err
	}

	return path.Join(mgrDir, "bin"), nil
}

func BuildTopDirName() (string, error) {
	mgrDir, err := mgrDirName()
	if err != nil {
		return "", err
	}

	return path.Join(mgrDir, "build"), nil
}

func BuildDirName(lsName string) (string, error) {
	buildTopDir, err := BuildTopDirName()
	if err != nil {
		return "", err
	}

	return path.Join(buildTopDir, lsName), nil
}

func CreateBuildDir(lsName string) error {
	buildTopDir, err := BuildTopDirName()
	if err != nil {
		return err
	}

	buildDir := path.Join(buildTopDir, lsName)
	if CheckFileExistence(buildDir) {
		return nil
	}

	if !CheckFileExistence(buildTopDir) {
		if err := createBuildTopDir(); err != nil {
			return err
		}
	}

	return os.Mkdir(buildDir, 0755)
}

func RemoveSymLink(symLinkPath string) error {
	if CheckSymLinkExistence(symLinkPath) {
		return os.Remove(symLinkPath)
	}

	return nil
}

func createBuildTopDir() error {
	buildTopDir, err := BuildTopDirName()
	if err != nil {
		return err
	}

	if CheckFileExistence(buildTopDir) {
		return nil
	}

	mgrDir, err := mgrDirName()
	if err != nil {
		return err
	}

	if !CheckFileExistence(mgrDir) {
		createMgrDir()
	}

	return os.Mkdir(buildTopDir, 0755)
}

func CreateBinDir() error {
	mgrDir, err := mgrDirName()
	if err != nil {
		return err
	}

	binDir, err := BinDirName()
	if err != nil {
		return err
	}

	if CheckFileExistence(binDir) {
		return nil
	}

	if !CheckFileExistence(mgrDir) {
		createMgrDir()
	}

	return os.Mkdir(binDir, 0755)
}

func createMgrDir() error {
	mgrDir, err := mgrDirName()
	if err != nil {
		return err
	}

	if CheckFileExistence(mgrDir) {
		return nil
	}

	return os.Mkdir(mgrDir, 0755)
}

func mgrDirName() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return path.Join(cacheDir, "lsm"), nil
}
