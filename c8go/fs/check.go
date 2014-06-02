package fs

func IsValid(name string) bool {
	for _, r := range name {
		if '0' <= r && r <= '9' {
			continue
		}
		if 'A' <= r && r <= 'Z' {
			continue
		}
		if 'a' <= r && r <= 'z' {
			continue
		}
		if r == '-' || r == '_' || r == '.' {
			continue
		}
		return false
	}
	return true
}
