package main

import (
	"flag"
	"log"

	"github.com/niciki/go-l2-tasks/develop/dev05/pkg/arguments"
	filedata "github.com/niciki/go-l2-tasks/develop/dev05/pkg/fileData"
)

func main() {
	arg := arguments.MakeArguments()
	fd := filedata.MakeFileData()
	err := fd.FillFileData(flag.Arg(0), arg)
	if err != nil {
		log.Fatal(err)
	}
	fd.Process()
}
