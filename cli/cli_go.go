package cli

import (
	"flag"
	"log"
	"os"
	"path"
)

func handleGoProject() {
	requireExecutable("go")

	goCmd := flag.NewFlagSet("go", flag.ExitOnError)
	dir := addDirectoryFlag(goCmd)
	createRepo := addGitFlag(goCmd)
	moduleName := goCmd.String("module", "", "Module to be passed to go mod init <MODULE>")
	createMakefile := goCmd.Bool("makefile", false, "Create a makefile with build and run commands")

	goCmd.Parse(os.Args[2:])

	if *moduleName != "" {
		log.Fatal("-module is a required param. Use flag -help for more information")
		os.Exit(1)
	}

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	workingDir := path.Join(cwd, *dir)

	err = os.Mkdir(*dir, 0775)
	assertErrorOr(err, func() {
		log.Println("Created directory", workingDir)
	})

	assertAndRunCommand(workingDir, "go", "mod", "init", *moduleName)
	log.Println("Initialized Go project")

	if *createRepo {
		assertAndRunCommand(workingDir, "git", "init")
		log.Println("Initialized git repository")
	}

	if *createMakefile {
		requireExecutable("make")
		file, err := os.Create(path.Join(workingDir, "Makefile"))
		assertError(err)
		_, err = file.WriteString("build:\n    @go build -o ./bin/out\n\nrun: build\n    @./bin/out")
		assertError(err)
		log.Println("Created Makefile")
	}

	createMainFile(workingDir)
}

func createMainFile(workingDir string) {
	file, err := os.Create(path.Join(workingDir, "main.go"))
	_, err = file.WriteString(`package main

import "fmt"

func main() {
    fmt.Println("Hello, World")
}`)
	assertError(err)
}
