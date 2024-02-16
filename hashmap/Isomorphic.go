package hashmap

func IsIsomorphic(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1ToS2 := make(map[byte]byte)
	s2ToS1 := make(map[byte]byte)
	for i := 0; i < len(s1); i++ {
		ch, ok := s1ToS2[s1[i]]
		if ok && ch != s2[i] {
			return false
		}
		ch, ok = s2ToS1[s2[i]]
		if ok && ch != s1[i] {
			return false
		}
		s1ToS2[s1[i]] = s2[i]
		s2ToS1[s2[i]] = s1[i]
	}
	return true
}
