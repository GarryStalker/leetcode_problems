package main

func main() {

}

// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

//III - 3
//XVI - 10 + 5 + 1 = 16
//XIV - 10 + 5 - 1 = 14

//<-> VIX - 5
//res:      V	I	X
//			5 	1	10
//if last value < current value, then current substring = last-current and insert into last value

//  M C M X C I V
//REVERSE FOR:
//	1	|  	2								|  	1
//	V 	|	I --> I < V THEN RES -= 1	|	4
//	5	|	1								|
//	3	|	4
//	C	|	X --> X < C THEN RES -= 10
// 	100	|	10

func romanToInt(s string) int {
	dict := make(map[string]int, 7)

	dict["I"] = 1
	dict["V"] = 5
	dict["X"] = 10
	dict["L"] = 50
	dict["C"] = 100
	dict["D"] = 500
	dict["M"] = 1000

	var res int
	res += dict[string(s[len(s)-1])]
	for i := len(s) - 2; i > 0; i-- {
		if dict[string(s[i])] < dict[string(s[i+1])] {
			res -= dict[string(s[i])]
			continue
		}
		res += dict[string(s[i])]
	}

	return res
}
