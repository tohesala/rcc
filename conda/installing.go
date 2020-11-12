package conda

import (
	"github.com/robocorp/rcc/common"
	"github.com/robocorp/rcc/shell"
)

func MustConda() bool {
	return HasConda() || (ValidateLocations() && (DoDownload() || DoDownload() || DoDownload()) && DoInstall())
}

func DoDownload() bool {
	if common.Debug {
		defer common.Stopwatch("Download done in").Report()
	}

	common.Log("Downloading Miniconda, this may take awhile ...")

	err := DownloadConda()
	if err != nil {
		common.Log("FAILURE: %s", err)
		return false
	} else {
		common.Log("Verify checksum from https://docs.conda.io/en/latest/miniconda.html")
		return true
	}
}

func DoInstall() bool {
	if common.Debug {
		defer common.Stopwatch("Installation done in").Report()
	}

	if !ValidateLocations() {
		return false
	}

	common.Log("Installing Miniconda, this may take awhile ...")

	install := InstallCommand()
	if common.Debug {
		common.Log("Running: %v", install)
	}
	_, err := shell.New(nil, ".", install...).Transparent()
	if err != nil {
		common.Log("Error: %v", err)
		return false
	}
	return true
}
