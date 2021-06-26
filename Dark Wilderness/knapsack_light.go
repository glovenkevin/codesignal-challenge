package main

import "fmt"

func main() {
	fmt.Println(10, 2 , 12, 3, 2); // 10	
}

func knapsackLight(value1 int, weight1 int, value2 int, weight2 int, maxW int) int {
    if weight1>maxW && weight2>maxW {
        return 0
    } else if weight1+weight2 <= maxW {
        return value1+value2
    } else {
        
        if weight1>maxW {
            value1=0
        } else if weight2>maxW {
            value2=0
        }
        
        if value1>value2 {
            return value1
        } else {
            return value2    
        }
        
    }
}
