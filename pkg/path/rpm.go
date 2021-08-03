package path

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	rpmBuildPath      = "rpmbuild"
	userRPMMacrosFile = ".rpmmacros"
	rpmMacroWorkDir   = "%_topdir"

	errNoRPMWorkDirUserCustom = "not find rpm work dir that user set"
)

func RPMWorkDir(showDetail bool) (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(userHome, "/") {
		userHome += "/"
	}

	rpmWorkDirUserCustom, err := RPMWorkDirUserCustom(userHome, showDetail)
	if err == nil {
		rpmWorkDirUserCustom += rpmBuildPath
		return rpmWorkDirUserCustom, nil
	}

	rpmWorkDirUserDefault := userHome + rpmBuildPath
	if showDetail {
		fmt.Println("[RPM Dir]Use default work dir:", rpmWorkDirUserDefault)
	}
	return rpmWorkDirUserDefault, nil

}

func RPMWorkDirUserCustom(curUserHome string, showDetail bool) (string, error) {

	userRPMMacrosFilePath := curUserHome + userRPMMacrosFile

	rpmMacrosFile, err := os.Open(userRPMMacrosFilePath)
	if err != nil {
		return "", err
	}
	defer rpmMacrosFile.Close()

	rpmMacrosFileReader := bufio.NewReader(rpmMacrosFile)

	for {
		lineContent, err := rpmMacrosFileReader.ReadString('\n')
		if err == io.EOF {
			return "", errors.New(errNoRPMWorkDirUserCustom)
		}

		if err != nil {
			return "", err
		}

		if !strings.Contains(lineContent, rpmMacroWorkDir) {
			continue
		}

		curRPMWorkDir := lineContent[len(rpmMacroWorkDir) : len(lineContent)-1]
		curRPMWorkDir = strings.TrimSpace(curRPMWorkDir)
		if !strings.HasSuffix(curRPMWorkDir, "/") {
			curRPMWorkDir += "/"
		}

		if showDetail {
			fmt.Println("[RPM Dir]Succeed to find work dir:", curRPMWorkDir, "from user rpmmacros file:", userRPMMacrosFilePath)
		}

		return curRPMWorkDir, nil
	}

}
