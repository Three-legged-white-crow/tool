package main

import (
	"flag"
	"fmt"

	"tools/pkg/clean"
	"tools/pkg/value"
)

func main() {
	cleanPath := flag.String("p", "", "directory that need clean")
	showDetail := flag.Bool("v", false, "show detail of clean")
	rpmMode := flag.Bool("rpm", false, "clean rpmbuild work directory")
	flag.Parse()

	if len(*cleanPath) > 0 {
		res, err := clean.CustomClean(*cleanPath, *showDetail)
		if err != nil {
			fmt.Println("[Error]Failed to clean specified path:", res.AimPath, ", err:", err)
		} else {
			printCleanResult(res)
		}
		fmt.Println("-------------------------------------------------------------")
		fmt.Println()
	}

	if *rpmMode {
		res, err := clean.RPMClean(*showDetail)
		if err != nil {
			fmt.Println("[Error]Failed to clean rpm work directory:", res.AimPath, ", err:", err)
		} else {
			printCleanResult(res)
		}
		fmt.Println("-------------------------------------------------------------")
		fmt.Println()
	}

}

func printCleanResult(res clean.Result) {
	fmt.Println()
	fmt.Println("* Clean aim path:", res.AimPath)
	timeUse := value.TimeValueFormat(res.TimeEnd - res.TimeStart)
	fmt.Println("* Time use:", timeUse)

	if len(res.FailedList) == 0 {
		return
	}

	fmt.Println("* Failed to clean:")
	for pathName := range res.FailedList {
		fmt.Println("  - " + pathName)
	}
}
