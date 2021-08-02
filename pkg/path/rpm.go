package path

import (
	"bufio"
	"errors"
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

func RPMWorkDir() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	if !strings.HasSuffix(userHome, "/") {
		userHome = userHome + "/"
	}

	rpmWorkDirUserCustom, err := RPMWorkDirUserCustom(userHome)
	if err == nil {
		return rpmWorkDirUserCustom, nil
	}

	rpmWorkDirUserDefault := userHome + rpmBuildPath
	return rpmWorkDirUserDefault, nil

}

func RPMWorkDirUserCustom(curUserHome string) (string, error) {

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
		return curRPMWorkDir, nil
	}

}
