package main

import (
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)

	logrus.Info("Starting containapp")

	fmt.Println("test")
}
