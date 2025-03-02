package infrastructure

import "os"

func CreateMakefile() {
	os.Create("Makefile")
}
