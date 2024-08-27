package cli

import (
	"log"
	"os"
)

func HandleCli() {
	if len(os.Args) < 2 {
		log.Fatalln("Expected one or more subcommands")
	}

	switch os.Args[1] {
	case "go":
		handleGoProject()
	}

    log.Println("Successfully created", os.Args[1], "project")
}
