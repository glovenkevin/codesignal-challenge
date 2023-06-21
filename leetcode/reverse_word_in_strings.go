package main

import (
	"bytes"
	"fmt"
)

// source: https://leetcode.com/problems/reverse-words-in-a-string/
func main() {
	tc := []struct {
		s   string
		res string
	}{
		{"the sky is blue", "blue is sky the"},
		{"a good   example", "example good a"},
	}

	for _, t := range tc {
		res := reverseWords(t.s)

		if res != t.res {
			fmt.Println("test failed")
			continue
		}

		fmt.Println("test success")
	}
}

func reverseWords(s string) string {
	var buff bytes.Buffer
	var bb [][]byte
	for _, r := range s {
		if r == ' ' {
			if buff.Len() > 0 {
				bb = append(bb, buff.Bytes())
			}
			buff = bytes.Buffer{}
			continue
		}

		buff.WriteRune(r)
	}
	if buff.Len() > 0 {
		bb = append(bb, buff.Bytes())
	}

	for i, j := 0, len(bb)-1; i < j; i, j = i+1, j-1 {
		bb[i], bb[j] = bb[j], bb[i]
	}

	return string(bytes.Join(bb, []byte{' '}))
}
