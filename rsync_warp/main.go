package rsync_warp

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	args := flag.String("args", "", "args of rsync command")
	flag.Parse()

	c := exec.Command("rsync")
	stdout, err := c.StdoutPipe()
	if err != nil {
		log.Fatal("Failed to get pipe of stdout, err:", err.Error())
	}
	if err = c.Start(); err != nil {
		log.Fatal("Failed to start process of cmd, err:", err.Error())
	}

	if err = c.Wait(); err != nil {
		processExitErr, ok := err.(*exec.ExitError)
		if ok {
			log.Println("Command fails to run or doesn't complete successfully, err:", err.Error())
			processExitCode := processExitErr.ExitCode()
			os.Exit(processExitCode)
		}

		log.Fatal("IO error that read stdout of command process, err:", err)
	}

	fmt.Println("Succeed to complete cmd:", "rsync", *args)
}

func filterLogRsync(ctx context.Context, reader io.Reader) {
	for {

		select {
		case <-ctx.Done():
			log.Println("[Filter]End filter because of receive end")
			return
		default:

		}

		content, err := io.ReadAll(reader)
		if err != nil {
			continue
		}

		bytes.Index()
		bytes.LastIndex()

	}

}
