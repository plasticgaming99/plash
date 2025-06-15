package builtins

// ls implements ls

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/plasticgaming99/plash/_lib/argparse"
)

type option struct {
	All bool `arg:"-a,--all" help:"Print hidden files"`
}

const help = `Usage ls [args...]
-a, --all
  Print hidden files`

func Ls(stdin io.Reader, stdout io.Writer, stderr io.Writer, args []string) error {
	opt := &option{}
	argparse.ParseArgs(opt, args)

	path := "."
	if len(args) >= 1 {
		path = args[0]
	}

	var err error

	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Fprintln(stderr, "ls:", err)
		return err
	}

	var nam string

	for _, file := range files {
		nam = file.Name()
		if strings.HasPrefix(nam, ".") && !opt.All {
			continue
		}
		io.WriteString(stdout, nam+"  ")
	}
	fmt.Println()

	/*cb := func(path string, d os.DirEntry, err error) error {
		io.WriteString(stdout, d.Name()+"\n")
		return nil
	}*/

	//filepath.WalkDir(path, cb)
	return nil
}
