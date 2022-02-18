// Command picoinit is a combo minimalistic init and service manager process
// to launch and manage multiple services within a single docker container.
package main

import (
	"os"

	"github.com/tuxdude/pico"
	"github.com/tuxdude/zzzlog"
)

var (
	log = zzzlog.NewLogger()
)

func run() int {
	inv, err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Invocation: %v", inv)

	init, err := pico.NewInit(log, inv.services...)
	if err != nil {
		log.Errorf("Failed to initialize and launch the services, reason: %v", err)
		return 1
	}
	exitCode := init.Wait()

	log.Infof("picoinit exiting with status code: %d", exitCode)
	return exitCode
}

func main() {
	os.Exit(run())
}
