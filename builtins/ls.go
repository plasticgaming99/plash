package builtins

// ls implements ls

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strings"

	"github.com/plasticgaming99/plash/_lib/argparse"
	"github.com/plasticgaming99/plash/_lib/termgrid"
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

	slices.SortFunc(files, func(a, b os.DirEntry) int {
		return strings.Compare(a.Name(), b.Name())
	})

	names := make([]string, len(files))
	var n string
	for i, entry := range files {
		n = entry.Name()
		if strings.HasPrefix(n, ".") && !opt.All {
			continue
		}
		names[i] = entry.Name()
	}

	tg := termgrid.Termgrid{
		Style:   termgrid.BottomToUp,
		Padding: 2,
	}

	tg.PrintSlice(stdout, names)

	return nil
}
