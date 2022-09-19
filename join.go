package piscine

func Join(strs []string, sep string) string {
	fdplexo := ""
	for i := 0; i < len(strs); i++ {
		fdplexo += strs[i]
		if i < len(strs)-1 {
			fdplexo += sep
		}
	}
	return fdplexo
}
