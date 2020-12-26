package main

var (
	letterMap = []string{
		" ",   // 0
		"",    // 1
		"abc", // 2
		"def", // 3
		"ghi", // 4
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz", // 9
	}

	res = []string{}
)

func letterCombinations(digits string) []string {
	if digits == "" {
		return res
	}
	findNumberCombinations(digits, 0, "")
	return res
}

func findNumberCombinations(digits string, index int, s string) {
	if index == len(digits) { // termination
		res = append(res, s)
		return
	}

	// process on
	num := digits[index]
	letter := letterMap[num-'0']

	for i := 0; i < len(letter); i++ {
		findNumberCombinations(digits, index+1, s+string(letter[i]))
	}

}

func main() {
	letterCombinations("")
}
