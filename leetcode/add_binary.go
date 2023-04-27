package main

import (
	"bytes"
	"fmt"
	"strconv"
)

// source: https://leetcode.com/problems/add-binary/
func main() {
	tc := []struct {
		a, b, res string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
		{"1", "1", "10"},
		{"0", "1", "1"},
		{"100", "110010", "110110"},
	}
	for _, c := range tc {
		res := addBinary(c.a, c.b)
		if res == c.res {
			fmt.Println("test successful")
			continue
		}

		fmt.Printf("res: %s | expected: %s\n", res, c.res)
	}
}

func addBinary(a string, b string) string {
	var (
		buff  bytes.Buffer
		carry int
	)

	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		if i < 0 || j < 0 {
			var bit byte
			if i >= 0 {
				bit = a[i]
			}
			if j >= 0 {
				bit = b[j]
			}

			if (carry == 1 && bit == '1') || (carry == 0 && bit == '0') {
				buff.WriteByte('0')
				continue
			}

			buff.WriteByte('1')
			carry = 0
			continue
		}

		if a[i] == '1' && b[j] == '1' {
			if carry == 1 {
				buff.WriteByte('1')
				continue
			}

			carry = 1
			buff.WriteByte('0')
			continue
		}

		if a[i] == '1' || b[j] == '1' {
			if carry == 1 {
				buff.WriteByte('0')
				continue
			}
			buff.WriteByte('1')
		} else { // 0 && 0
			buff.WriteString(strconv.Itoa(carry))
			carry = 0
		}
	}

	if carry == 1 {
		buff.WriteByte('1')
	}

	var bb = buff.Bytes()
	for i, j := 0, buff.Len()-1; i < j; i, j = i+1, j-1 {
		bb[i], bb[j] = bb[j], bb[i]
	}

	return buff.String()
}
