package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(alternate("beabeefeab"))                                                                                                                                    // Expected: 5
	fmt.Println(alternate("asdcbsdcagfsdbgdfanfghbsfdab"))                                                                                                                  // Expected: 8
	fmt.Println(alternate("cwomzxmuelmangtosqkgfdqvkzdnxerhravxndvomhbokqmvsfcaddgxgwtpgpqrmeoxvkkjunkbjeyteccpugbkvhljxsshpoymkryydtmfhaogepvbwmypeiqumcibjskmsrpllgbvc")) // Expected: 8

	fmt.Println(alternate("ab")) // Expected: 2
}

func alternate(s string) int32 {

	sSize := len(s)
	cc := map[string]int{}
	for i := 0; i < sSize; i++ {
		cc[string(s[i])] += 1
		if i+1 != sSize && s[i] == s[i+1] {
			s = strings.ReplaceAll(s, string(s[i]), "")
			sSize = len(s)
			i = 0
			cc = map[string]int{}
		}
	}

	i := 0
	keys := make([]string, len(cc))
	for k := range cc {
		keys[i] = k
		i++
	}

	max := 0
	lk := len(keys)
	for i, key := range keys {

		for j := i + 1; j < lk; j++ {
			pair := keys[j]

			if math.Abs(float64(cc[key]-cc[pair])) > 1 {
				continue
			}

			rgp := regexp.MustCompile("[^" + key + pair + "]")
			newS := rgp.ReplaceAllString(s, "")

			rgp = regexp.MustCompile(key + "{2,}|" + pair + "{2,}")
			doublemMatch := rgp.FindString(newS)

			if len(newS) > max && doublemMatch == "" {
				max = len(newS)
			}
		}

	}

	return int32(max)
}
