package main

import (
	"os"
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/opencontainers/runc/libcontainer"
	_ "github.com/opencontainers/runc/libcontainer/nsenter"
)

// init to handle go lang issues with forking
func init() {
	logrus.Info("init called")
	if len(os.Args) < 2 || os.Args[1] != "init" {
		return
	}
	runtime.GOMAXPROCS(1)
	runtime.LockOSThread()

	factory, _ := libcontainer.New("")
	if err := factory.StartInitialization(); err != nil {
		logrus.Fatal(err)
	}
	panic("--this line should have never been executed, congratulations--")
}
