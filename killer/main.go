package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"strings"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

const (
	emptyValue = "empty"
	signalMin  = 1
	signalMax  = 31
	timeLayout = "2006-01-02 15:04:05.000"
)

func main() {

	name := flag.String(
		"name",
		emptyValue,
		"name of process",
	)

	cmdContent := flag.String(
		"cmd",
		emptyValue,
		`command content, please use " wrapper command content`)

	signalNum := flag.Int(
		"signal",
		0,
		"signal that send to process")

	loopMod := flag.Bool(
		"loop",
		false,
		"loop kill process")

	randStart := flag.Int(
		"rand-start",
		0,
		"start of range to rand, effective at loop mod, unit is second")

	randEnd := flag.Int(
		"rand-end",
		0,
		"end of range to rand, effective at loop mod and must specify, unit is second")

	flag.Parse()

	log.SetOutput(os.Stderr)
	log.Println("[Killer-Info]New kill request:",
		"name:", *name,
		"command:", *cmdContent,
		"signal:", *signalNum,
		"loopMod:", *loopMod,
		"randStart:", *randStart,
		"randEnd:", *randEnd,
	)

	var isNameEmpty, isCommandEmpty, isSignalInvalid bool

	if *name == emptyValue {
		isNameEmpty = true
	}

	if *cmdContent == emptyValue {
		isCommandEmpty = true
	}

	if isNameEmpty && isCommandEmpty {
		log.Println("[Killer-Error]Please specify name or command")
		return
	}

	if !isNameEmpty && isCommandEmpty {
		log.Println("[Killer-Wraning]Only specify name but not specify command, only list pid that matched")
	}

	if *signalNum < signalMin || *signalNum > signalMax {
		isSignalInvalid = true
	}

	req := killReq{
		name:            *name,
		isNameEmpty:     isNameEmpty,
		command:         *cmdContent,
		isCommandEmpty:  isCommandEmpty,
		signal:          *signalNum,
		isSignalInvalid: isSignalInvalid,
	}

	if !(*loopMod) {
		kill(req)
		return
	}

	if *randEnd == 0 {
		log.Println("[Killer-Error]Must specify end of range ues to rand")
		return
	}

	if *randStart < 0 {
		log.Println("[Killer-Error]Start of range use to rand must >= 0")
		return
	}

	if *randEnd < 0 {
		log.Println("[Killer-Error]End of range use to rand must > 0")
		return
	}

	var sleepRand int

	for {
		sleepRand = rand.New(rand.NewSource(time.Now().UnixNano())).Intn(*randEnd-*randStart+1) + *randStart
		log.Println("[Killer-Info]Sleep", sleepRand, "second and then kill process")
		time.Sleep(time.Duration(sleepRand) * time.Second)
		kill(req)
	}

}

type killReq struct {
	name            string
	isNameEmpty     bool
	command         string
	isCommandEmpty  bool
	signal          int
	isSignalInvalid bool
}

func kill(req killReq) {
	processList, err := process.Processes()
	if err != nil {
		return
	}

	var (
		pID          int32
		ppID         int32
		pName        string
		pCmd         string
		pStatus      string
		pCreateTime  int64
		tm           time.Time
		ts           string
		threadNum    int32
		numMatchName int
		numMatchCmd  int
	)
	for _, curProcess := range processList {
		if !req.isNameEmpty {
			pName, err = curProcess.Name()
			if !strings.Contains(pName, req.name) {
				continue
			}
			numMatchName += 1

		}

		pID = curProcess.Pid
		pStatus, err = curProcess.Status()
		ppID, err = curProcess.Ppid()
		pCreateTime, err = curProcess.CreateTime()
		if pCreateTime > 0 {
			tm = time.Unix(pCreateTime/1e3, 0)
			ts = tm.Format(timeLayout)
		}
		threadNum, err = curProcess.NumThreads()

		pCmd, err = curProcess.Cmdline()
		if !req.isCommandEmpty {
			if err != nil {
				continue
			}

			if pCmd != req.command {
				continue
			}

			numMatchCmd += 1

			log.Println("[Killer-Info]Get process, pid:", pID,
				"name:", pName,
				"command:", pCmd,
				"status:", pStatus,
				"ppid:", ppID,
				"create-time:", ts,
				"threads:", threadNum)

			if req.isSignalInvalid {
				log.Println("[Killer-Wraning]Specify signal is invalid, do nothing")
				return

			} else {
				err = curProcess.SendSignal(syscall.Signal(req.signal))
				if err != nil {
					log.Println("[Killer-Error]Failed to send signal:", req.signal, "to process and err:", err.Error())
					return
				}
				log.Println("[Killer-Info]Succeed to send signal:", req.signal, "to process")
				return
			}

		}

		log.Println("[Killer-Info]Pid:", pID,
			"name:", pName,
			"cmd:", pCmd,
			"status:", pStatus,
			"ppid:", ppID,
			"create-time:", ts,
			"threads:", threadNum)
	}

	log.Println("[Killer-Info]Number match name:", numMatchName, "Number match command:", numMatchCmd)
}
