package cli

import (
	"log"
	"os"
	"os/exec"
)

func assertOrExit(condition bool, errorMessage string) {
	if !condition {
		log.Fatalf("Error: %s", errorMessage)
		os.Exit(1)
	}
}

func requireExecutable(executable string) string {
	path, err := exec.LookPath(executable)
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}
	return path
}

func assertError(err error) {
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	}
}

func assertErrorOr(err error, fn func()) {
	if err != nil {
		log.Fatalf("Error: %s", err)
		os.Exit(1)
	} else {
		fn()
	}
}

func assertAndRunCommand(workingDir string, command string, arg ...string) {
	cmd := exec.Command(command, arg...)
	cmd.Dir = workingDir
	err := cmd.Run()
	assertError(err)
}
