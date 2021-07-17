package main

import (
    "regexp"
    "strings"
    "fmt"
)

func main() {
    fmt.Println(isMAC48Address("00-1B-63-84-45-E6"))
	fmt.Println(isMAC48AddressOther("00-1B-63-84-45-E6"))
}

// My Solution
func isMAC48Address(inputString string) bool {
    var pattern = "[0-9|A-F]{2}"
    var regex = regexp.MustCompile(pattern)
    var listHexaDec = strings.Split(inputString, "-")
    var totalItem = len(listHexaDec)
    if (totalItem == 6) {
        
        var rtn = true
        for _ , val := range listHexaDec {
            var res = regex.FindStringSubmatch(val)
            if (res == nil || len(val)>2) {
                rtn = false
                break
            }
        }
        return rtn
        
    } else {
        return false
    }
}

// Other regex solution 
func isMAC48AddressOther(inputString string) bool {
    var pattern = "^([0-9A-F]{2}-){5}[0-9A-F]{2}$"
    var regex = regexp.MustCompile(pattern)
    var res = regex.FindStringSubmatch(inputString)
    return nil != res
}