package main

import "fmt"

var romertall = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func main() {
	//x := 2 * 3
	//	//fmt.Printf("2 * 3 er lik %d", x)
	rt := []string{"CLXVI", "MCDCCVI", "MCCCCCCVI"}
	for _, s := range rt {
		v := decode(s)
		fmt.Printf("Verdien av romertallet %s er %d \n", s, v)
	}
}

func decode(s string) int {
	sum := 0
	for i:= 0;i < len(s); i++ {
		v := romertall[string(s[i])]
		if i == len(s)-1 {
			sum += v
			continue
		}
		v2 := romertall[string(s[i+1])]
		if v2 <= v {
			sum += v
			continue
		}
		sum = sum + (v2-v)
		i++
	}
	return sum
}