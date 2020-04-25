package installer

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/h-michael/lsm/file"
)

func InstallViaNpm(lsName string) error {
	if err := file.CreateBuildDir(lsName); err != nil {
		return err
	}

	err := initBuildDir(lsName)
	if err != nil {
		return err
	}

	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return err
	}

	if err := execNpm(buildDir, "install", lsName); err != nil {
		return err
	}

	if err := createLsSymLink(lsName, getBinPathViaNpm); err != nil {
		return err
	}
	return nil
}

func UninstallViaNpm(lsName string) error {
	symLink, err := lsSymLinkPath(lsName)
	if err != nil {
		return err
	}
	if err := file.RemoveSymLink(symLink); err != nil {
		return err
	}

	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return err
	}

	if err := os.RemoveAll(buildDir); err != nil {
		return err
	}

	return nil
}

func initBuildDir(lsName string) error {
	if err := file.CreateBuildDir(lsName); err != nil {
		return err
	}
	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return err
	}
	packageJsonPath := path.Join(buildDir, "package.json")
	if file.CheckFileExistence(packageJsonPath) {
		return nil
	}
	if err := execNpm(buildDir, "init", "-y"); err != nil {
		return err
	}
	body := []byte("{\"name\": \"\"}")
	if err := ioutil.WriteFile(packageJsonPath, body, 0755); err != nil {
		return err
	}
	return nil
}

func getBinPathViaNpm(lsName string) (string, error) {
	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return "", err
	}
	return path.Join(buildDir, "node_modules", ".bin", lsName), nil
}

func execNpm(dir string, args ...string) error {
	_, err := exec.LookPath("npm")
	if err != nil {
		return err
	}

	cmd := exec.Command("npm", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

func NpmInstallGlobal(pkgName string) error {
	if err := execNpm("install", "-g", pkgName); err != nil {
		return err
	}

	return nil
}
