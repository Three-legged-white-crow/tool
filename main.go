package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type result struct {
	failedList map[string]int8
}

func main() {
	cleanPath := flag.String("p", "", "directory that wait clean")
	showDetail := flag.Bool("v", false, "show detail of clean")
	flag.Parse()

	aimPath := strings.TrimSuffix(*cleanPath, "/")
	res, err := clean(aimPath, *showDetail)
	if err != nil {
		fmt.Println(err)
		return
	}

	failedList := res.failedList

	if len(failedList) > 0 {
		fmt.Println("Failed to clean:")
		for pathName := range failedList {
			fmt.Println("  - " + pathName)
		}
		fmt.Println()
	}

}

func clean(cleanPath string, showDetail bool) (result, error) {
	fileList, err := ioutil.ReadDir(cleanPath)
	if err != nil {
		if showDetail {
			fmt.Println("[Read Dir]Failed to read dir", cleanPath, "and err:", err)
		}
		return result{}, err
	}

	var failedList = map[string]int8{}

	for _, fInfo := range fileList {
		curPath := cleanPath + "/" + fInfo.Name()

		if fInfo.IsDir() {
			res, err := clean(curPath, showDetail)
			if err != nil {
				failedList[curPath] = 0
				continue
			}

			for f := range res.failedList {
				failedList[f] = 0
			}
			continue
		}

		err = os.Remove(curPath)
		if err != nil {
			if showDetail {
				fmt.Println("[Remove Faile]Failed to remove", curPath, "and err:", err)
			}
			failedList[curPath] = 0

		} else {
			if showDetail {
				fmt.Println("[Remove Faile]Succeed to remove", curPath)
			}
		}

	}

	return result{failedList: failedList}, nil
}
