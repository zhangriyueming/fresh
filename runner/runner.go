package runner

import (
	"io"
	"os/exec"
	"strconv"
	"strings"
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

	var cmd *exec.Cmd
	if len(settings["run_args"]) < 1 {
		cmd = exec.Command(buildPath())
	} else {
		args := strings.Fields(settings["run_args"])
		cmd = exec.Command(buildPath(), args...)
	}

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
