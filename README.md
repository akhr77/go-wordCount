# wordcount

Gopher道場の入門課題 `wordcount` の実装です。

## Installation

```bash
$ go install github.com/akhr77/go-dojo7/cmd/wordcount
```

## Usage
```bash
$ echo -e "cat cat cat dog dog dog.. dog.,. cat　 cat el\ncat" | wordcount
cat
dog
el
```

- 与えられた文字列を空白または改行文字で区切る
- 文末のドット(.)とカンマ（,）は無視する
- 受け取った単語のTOP3を表示する
- TOP3の中に同数の単語が含まれる場合、アルファベット順に表示する
