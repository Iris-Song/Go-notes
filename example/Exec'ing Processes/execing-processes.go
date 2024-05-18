package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// For our example we’ll exec ls. Go requires an absolute path to the binary we want to execute,
	// so we’ll use exec.LookPath to find it (probably /bin/ls).
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
