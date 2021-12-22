package devlib

import (
	"os"
	"runtime/pprof"
)

func StartCPUProfile(filename string) {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	pprof.StartCPUProfile(fp)
}

func StopCPUProfile() {
	pprof.StopCPUProfile()
}
