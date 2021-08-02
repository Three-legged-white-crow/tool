package clean

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"tools/pkg/path"
)

type Result struct {
	AimPath    string
	FailedList map[string]string
	TimeStart  int64
	TimeEnd    int64
}

func CustomClean(aimPath string, showDetail bool) (Result, error) {
	res := Result{AimPath: aimPath}

	startTime := time.Now().UnixNano()
	failedList, err := Clean(aimPath, showDetail)
	endTime := time.Now().UnixNano()

	res.TimeStart = startTime
	res.TimeEnd = endTime

	if err != nil {
		return res, err
	}

	res.FailedList = failedList
	return res, nil
}

func RPMClean(showDetail bool) (Result, error) {
	res := Result{}

	curRPMWorkDir, err := path.RPMWorkDir()
	if err != nil {
		if showDetail {
			fmt.Println("[RPM Dir]Failed to get rpm work dir, err:", err)
		}
		return res, err
	}

	res.AimPath = curRPMWorkDir
	startTime := time.Now().UnixNano()
	failedList, err := Clean(curRPMWorkDir, showDetail)
	endTime := time.Now().UnixNano()

	res.TimeStart = startTime
	res.TimeEnd = endTime

	if err != nil {
		return res, err
	}

	res.FailedList = failedList
	return res, nil
}

func Clean(cleanPath string, showDetail bool) (map[string]string, error) {
	cleanPath = strings.TrimSuffix(cleanPath, "/")
	fileList, err := ioutil.ReadDir(cleanPath)
	if err != nil {
		if showDetail {
			fmt.Println("[Read Dir]Failed to read dir:", cleanPath, "and err:", err)
		}
		return nil, err
	}

	var curFailedList = map[string]string{}
	var failedList = map[string]string{}
	var curPath string

	for _, fInfo := range fileList {
		curPath = cleanPath + "/" + fInfo.Name()

		if fInfo.IsDir() {
			failedList, err = Clean(curPath, showDetail)
			if err != nil {
				curFailedList[curPath] = err.Error()
				continue
			}

			for failedPath, errStr := range failedList {
				curFailedList[failedPath] = errStr
			}
			continue
		}

		err = os.Remove(curPath)
		if err != nil {
			if showDetail {
				fmt.Println("[Remove Faile]Failed to remove", curPath, "and err:", err)
			}
			curFailedList[curPath] = err.Error()
			continue

		}

		if showDetail {
			fmt.Println("[Remove Faile]Succeed to remove", curPath)
		}

	}

	return curFailedList, nil
}
