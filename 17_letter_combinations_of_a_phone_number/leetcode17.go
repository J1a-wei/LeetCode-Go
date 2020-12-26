package leetcode17

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
		return []string{}
	}

	res = []string{}
	findCombination(&digits, 0, "")
	return res
}

func findCombination(digits *string, index int, s string) {
	if index == len(*digits) {
		res = append(res, s)
		return
	}
	num := (*digits)[index]
	letter := letterMap[num-'0'] // 小技巧要记住

	for i := 0; i < len(letter); i++ {
		findCombination(digits, index+1, s+string(letter[i]))
	}
}
