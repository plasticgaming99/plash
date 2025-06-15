// JUST PRINT TABLES PRETTY

package termgrid

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"golang.org/x/term"
)

type TgStyle int

const (
	BottomToUp = TgStyle(iota)
	LeftToRight
)

type Termgrid struct {
	Style   TgStyle
	Padding int
	Width   int
	Height  int
}

func (tg Termgrid) PrintSlice(stdout io.Writer, out []string) error {
	var err error
	var termWidth, termHeight int
	if tg.Width == 0 && tg.Height == 0 {
		termWidth, termHeight, err = term.GetSize(int(os.Stdout.Fd()))
		if err != nil {
			termWidth, termHeight = 60, 40
		}
	}

	maxLen := 0
	for i := 0; i < len(out); i++ {
		if len(out[i]) > maxLen {
			maxLen = len(out[i])
		}
	}

	colWidth := maxLen + tg.Padding
	cols := termWidth / colWidth

	rows := (len(out) + cols - 1) / cols

	if rows > termHeight {
		fmt.Fprintln(stdout, strings.Join(out, "\n"))
		return nil
	}

	sort.Strings(out)
	grid := make([][]string, rows)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]string, cols)
	}

	for i, entry := range out {
		row := i % rows
		col := i / rows
		grid[row][col] = entry
	}

	// 出力
	for _, row := range grid {
		for _, cell := range row {
			if cell == "" {
				continue
			}
			fmt.Printf("%-*s", colWidth, cell)
		}
		fmt.Println()
	}
	return nil
}
