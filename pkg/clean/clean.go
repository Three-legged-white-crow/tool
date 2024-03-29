package clean

import (
	"fmt"
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

	curRPMWorkDir, err := path.RPMWorkDir(showDetail)
	if err != nil {
		fmt.Println("[RPM Dir]Failed to get rpm work dir, err:", err)
		return res, err
	}

	fmt.Println("[RPM Dir]Succeed to get rpm work dir:", curRPMWorkDir)

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

// Fixme: Not suitable for large directories, because it will use a lot of memory(OOM or hang os), or cause stack overflow

func Clean(cleanPath string, showDetail bool) (map[string]string, error) {
	cleanPath = strings.TrimSuffix(cleanPath, "/")
	fileList, err := os.ReadDir(cleanPath)
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
