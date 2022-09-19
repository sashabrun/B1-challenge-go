package piscine

func BasicJoin(elems []string) string {
	niquetamerelexo := ""
	for i := 0; i < len(elems); i++ {
		niquetamerelexo += elems[i]
	}
	return niquetamerelexo
}
