package file

import (
	"fmt"
	"os"
	"path"
)

func CheckExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
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
	if CheckExist(buildDir) {
		return nil
	}
	if !CheckExist(buildTopDir) {
		if err := createBuildTopDir(); err != nil {
			return err
		}
	}
	if err := os.Mkdir(buildDir, 0755); err != nil {
		return err
	}
	return nil
}

func createBuildTopDir() error {
	buildTopDir, err := BuildTopDirName()
	if err != nil {
		return err
	}
	if CheckExist(buildTopDir) {
		return nil
	}
	mgrDir, err := mgrDirName()
	if err != nil {
		return err
	}
	if !CheckExist(mgrDir) {
		createMgrDir()
	}
	if err := os.Mkdir(buildTopDir, 0755); err != nil {
		return err
	}
	return nil
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
	if CheckExist(binDir) {
		return nil
	}
	if !CheckExist(mgrDir) {
		createMgrDir()
	}
	if err := os.Mkdir(binDir, 0755); err != nil {
		return err
	}
	return nil
}

func createMgrDir() error {
	mgrDir, err := mgrDirName()
	if err != nil {
		return err
	}
	if CheckExist(mgrDir) {
		return nil
	}
	fmt.Printf("create %s\n", mgrDir)
	if err := os.Mkdir(mgrDir, 0755); err != nil {
		return err
	}
	return nil
}

func mgrDirName() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	return path.Join(cacheDir, "lsm"), nil
}
