package main

func isIsomorphic(s string, t string) bool {
	stMap := make(map[byte]byte)
	tsMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		s_Val, oks := stMap[s[i]]
		t_Val, okt := tsMap[t[i]]

		if (oks && s_Val != t[i]) || (okt && t_Val != s[i]) {
			return false
		}
		stMap[s[i]] = t[i]
		tsMap[t[i]] = s[i]
	}
	return true
}

func main() {
	isIsomorphic("abb", "egg")
}
