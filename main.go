package main

import (
	"fmt"
	"mvdan.cc/sh/v3/syntax"
	"os"
	"strconv"
	"strings"
)

func extractParams(bashString string) ([]segment, error) {
	in := strings.NewReader(bashString)
	f, err := syntax.NewParser().Parse(in, "")
	if err != nil {
		return nil, err
	}

	stmt := f.Stmts[0]
	cmd := stmt.Cmd
	args := cmd.(*syntax.CallExpr).Args
	segments := []segment{}
	for _, word := range args {
		parts := word.Parts
		for _, part := range parts {
			// partsStr = append(partsStr, fmt.Sprintf("%-20T", part))
			start := part.Pos().Offset()
			end := part.End().Offset()
			segments = append(segments, segment{
				start: int(start),
				end:   int(end),
				value: bashString[start:end]})
		}

	}
	return segments, nil
}

type segment struct {
	value string
	start int
	end   int
}

func findSegmentByStrPos(segments []segment, strPos int) int {
	for i, segment := range segments {
		if strPos >= segment.start && strPos <= segment.end {
			return i
		}
	}
	return 0
}

func main() {
	cmd, cursorPos, key := func() (string, int, string) {
		cursorPos, _ := strconv.ParseInt(os.Args[2], 10, 0)
		return os.Args[1], int(cursorPos), os.Args[3]
	}()
	params, _ := extractParams(cmd)
	activeSeg := params[findSegmentByStrPos(params, cursorPos)]
	var newCursorVal int
	switch key {
	case "b":
		newCursorVal = cursorPos - (cursorPos - activeSeg.start)
	case "w":
		newCursorVal = cursorPos + (activeSeg.end - cursorPos)
	}
	fmt.Println(newCursorVal)
}
