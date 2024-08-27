package cli

import "flag"

func addDirectoryFlag(flagSet *flag.FlagSet) *string {
	return flagSet.String("dir", ".", "Directory where the project will be initialized. Empty means project is created in working directory")
}

func addGitFlag(flagSet *flag.FlagSet) *bool {
	return flagSet.Bool("git", false, "Init a git repo")
}
