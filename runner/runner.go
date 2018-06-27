package runner

import (
	"io"
	"os/exec"
	"strconv"
)

func run() bool {
	only_build, err := strconv.Atoi(settings["only_build"])
	if err != nil {
		fatal(err)
	}
	if only_build != 0 {
		return true
	}

	runnerLog("Running...")

	cmd := exec.Command(buildPath())

	stderr, err := cmd.StderrPipe()
	if err != nil {
		fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(err)
	}

	err = cmd.Start()
	if err != nil {
		fatal(err)
	}

	go io.Copy(appLogWriter{}, stderr)
	go io.Copy(appLogWriter{}, stdout)

	go func() {
		<-stopChannel
		pid := cmd.Process.Pid
		runnerLog("Killing PID %d", pid)
		cmd.Process.Kill()
	}()

	return true
}
