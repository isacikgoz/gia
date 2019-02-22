package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/isacikgoz/gia/editor"
	"github.com/waigani/diffparser"
)

func main() {
	args := parseArgs()
	file, err := generateDiffFile(args[2], args[1])
	if err == nil {
		editor, err := editor.NewEditor(file)
		if err != nil {
			log.Fatal(err)
		}
		patches, err := editor.Run()
		if err != nil {
			log.Fatal(err)
		}
		for _, patch := range patches {
			if err := applyPatch(args[2], patch); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		log.Fatal("can't create diff for " + args[1] + " in " + args[2] + " mode")
	}
}

func parseArgs() []string {
	if len(os.Args) > 1 {
		if os.Args[1] == "--help" || os.Args[1] == "-h" {
			fmt.Println(help())
			os.Exit(0)
		} else if os.Args[1] == "--version" || os.Args[1] == "-v" {
			fmt.Println(version())
			os.Exit(0)
		}
	}
	args := make([]string, 3)
	for i, arg := range os.Args {
		args[i] = arg
		if i > 2 {
			break
		}
	}
	return args
}

func applyPatch(fileMode, patch string) error {
	mode := []string{"apply", "--cached"}
	if fileMode == "--cached" {
		mode = []string{"apply", "--cached", "--reverse"}
	}
	cmd := exec.Command("git", mode...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, patch+"\n")
	}()
	if _, err := cmd.CombinedOutput(); err != nil {
		return err
	}
	return nil
}

func generateDiffFile(fileMode, filePath string) (*diffparser.DiffFile, error) {
	args := diffArgs(fileMode, filePath)
	cmd := exec.Command("git", args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}
	diff, err := diffparser.Parse(string(out))
	if err != nil {
		return nil, err
	}
	if len(diff.Files) == 0 {
		return nil, errors.New(string(out))
	}
	return diff.Files[0], nil
}

// diffArgs returns git command args for getting diff
func diffArgs(fileMode, filePath string) []string {
	var args []string
	switch fileMode {
	case "--cached":
		args = []string{"diff", "--cached", filePath}
	case "--no-index":
		args = []string{"diff", "--no-index", "/dev/null", filePath}
	case "--":
	default:
		args = []string{"diff", "--", filePath}
	}
	return args
}

func help() string {
	return `usage: gia [<file>] [<args>]

Flags:
	-h, --help     Show help.
	-v, --version  Show application version.
	
Args:
  --
    Default mode, used on unstaged files.

  --no-index
    Used on untracked files.

  --cached
    Used on staged files.`
}

func version() string {
	return "gia version 0.1.0"
}
