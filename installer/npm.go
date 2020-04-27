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

	if err := initBuildDir(lsName); err != nil {
		return err
	}

	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return err
	}

	if err := execNpm(buildDir, "install", lsName); err != nil {
		return err
	}

	return createLsSymLink(lsName, getBinPathViaNpm)
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

	return os.RemoveAll(buildDir)
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

	return ioutil.WriteFile(packageJsonPath, []byte(`{"name": ""}`), 0755)
}

func getBinPathViaNpm(lsName string) (string, error) {
	buildDir, err := file.BuildDirName(lsName)
	if err != nil {
		return "", err
	}

	return path.Join(buildDir, "node_modules", ".bin", lsName), nil
}

func execNpm(dir string, args ...string) error {
	if _, err := exec.LookPath("npm"); err != nil {
		return err
	}

	cmd := exec.Command("npm", args...)
	cmd.Dir = dir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

func NpmInstallGlobal(pkgName string) error {
	return execNpm("install", "-g", pkgName)
}
