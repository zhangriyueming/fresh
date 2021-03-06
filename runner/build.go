package runner

import (
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func build() (string, bool) {
	buildLog("Building...")

	var cmd *exec.Cmd

	if len(settings["build_command"]) < 1 {
		if len(settings["build_args"]) < 1 {
			cmd = exec.Command("go", "build", "-o", buildPath(), root())
		} else {
			cmd = exec.Command("go", "build", "-o", buildPath(), root(), settings["build_args"])
		}
	} else {
		if len(settings["build_args"]) < 1 {
			cmd = exec.Command(settings["build_command"])
		} else {
			cmd = exec.Command(settings["build_command"], settings["build_args"])
		}
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

	io.Copy(os.Stdout, stdout)
	errBuf, _ := ioutil.ReadAll(stderr)

	err = cmd.Wait()
	if err != nil {
		return string(errBuf), false
	}

	return "", true
}
