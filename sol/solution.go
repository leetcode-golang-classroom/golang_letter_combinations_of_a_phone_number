package sol

var digitMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	result := []string{}
	var dfs func(i int, cur []byte)
	dfs = func(i int, cur []byte) {
		if i == len(digits) {
			temp := make([]byte, len(cur))
			copy(temp, cur)
			result = append(result, string(temp))
			return
		}
		search := digitMap[digits[i]]
		sLen := len(search)
		for idx := 0; idx < sLen; idx++ {
			cur = append(cur, search[idx])
			dfs(i+1, cur)
			cur = cur[:len(cur)-1]
		}
	}
	dfs(0, []byte{})
	return result
}
