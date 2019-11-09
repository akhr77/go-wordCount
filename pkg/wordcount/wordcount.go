package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

const (
	//ExitCodeOk 正常終了
	ExitCodeOk = 0
	//ExitCodeErr 異常終了
	ExitCodeErr = 1
)

// Run the wordcount
func Run() int {
	lines, err := strStdin(os.Stdin)
	if err != nil {
		return ExitCodeErr
	}
	m := strFields(lines)
	s := ranking(m)
	print(s)
	return ExitCodeOk
}

func strStdin(r io.Reader) (lines []string, err error) {
	lines = []string{}
	scanner := bufio.NewScanner(r)
	// 複数行の読み込み
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if scanner.Err() != nil {
		return nil, err
	}
	return lines, nil
}

func strFields(lines []string) map[string]int {
	m := make(map[string]int)
	// 行単位に処理
	for _, line := range lines {
		words := strings.Fields(line)
		// 文末の「.」「,」を取り除く
		for _, word := range words {
			for strings.HasSuffix(word, ",") || strings.HasSuffix(word, ".") {
				word = strings.TrimRight(word, ".")
				word = strings.TrimRight(word, ",")
			}
			m[	]++
		}
	}
	return m
}

// ranking sort map
func ranking(m map[string]int) (a List) {
	a = List{}
	for k, v := range m {
		e := Entry{k, v}
		a = append(a, e)
	}

	sort.Sort(a)

	return
}

// print display the TOP3
func print(a List) {
	for k, v := range a {
		if k <= 2 {
			fmt.Println(v.name)
		}
	}
}

// Entry name and value
type Entry struct {
	name  string
	value int
}

// List slice Entry
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	}
	return (l[i].value > l[j].value)
}
