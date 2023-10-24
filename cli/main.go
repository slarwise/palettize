package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/alexflint/go-arg"
)

var args struct {
	Input   string `arg:"positional"`
	Output  string `arg:"required"`
	Palette string `arg:"required"`
}

func main() {
	arg.MustParse(&args)
	magickCmd, err := exec.LookPath("convert")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	inputFiles := []string{args.Input, args.Palette}
	for _, file := range inputFiles {
		if _, err := os.Stat(file); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	cmd := exec.Command(magickCmd, args.Input, "+dither", "-remap", args.Palette, args.Output)
	if err = cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
