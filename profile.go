package devlib

import (
	"os"
	"runtime/pprof"
)

func StartCPUProfile(filename string) *os.File {
	fp, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	pprof.StartCPUProfile(fp)
	return fp
}

func StopCPUProfile(fp *os.File) {
	defer fp.Close()
	pprof.StopCPUProfile()
}
