package basename1

// Basename はディレクトリ要素と、接尾辞を取り除きます。
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func Basename(s string) string {
	// 最後の'/'とその前のすべてを破棄する。
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}

	// 最後の'.'より前のすべてを保持する。
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}

	return s
}
