package wordcount

import (
	"bytes"
	"reflect"
	"testing"
)

func TestStdin(t *testing.T) {
	t.Helper()
	desc := "複数行（改行あり）を読み込んでスライスに格納"
	s := `cat
dog
rabit`
	input := bytes.NewBufferString(s)
	expected := []string{"cat", "dog", "rabit"}
	actual, err := strStdin(input)
	if err != nil {
		t.Fatalf("failed to call strStdin():%s", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild test:%s, actual:%v, expected:%v", desc, actual, expected)
	}

}
func TestStrFields(t *testing.T) {
	t.Helper()
	desc := "文末のカンマとコロンを排除してmapに格納"
	input := []string{"cat..,.", "dog.", "cat dog...,,, el dog"}
	expected := map[string]int{"cat": 2, "dog": 3, "el": 1}
	actual := strFields(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild test:%s, actual:%v, expected:%v", desc, actual, expected)
	}

}

func TestRanking(t *testing.T) {
	t.Helper()
	desc := "単語が多い順にSORT"
	input := map[string]int{"dog": 2, "cat": 2, "el": 3}
	expected := List{}
	expected = []Entry{Entry{"el", 3}, Entry{"cat", 2}, Entry{"dog", 2}}
	actual := ranking(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild test:%s, actual:%v, expected:%v", desc, actual, expected)
	}

}
