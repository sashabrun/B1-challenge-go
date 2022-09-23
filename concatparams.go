package piscine

func ConcatParams(args []string) string {
	chainedecharactere := ""
	for i := 0; i < len(args); i++ {
		chainedecharactere = chainedecharactere + args[i] + "/n"
	}
	return chainedecharactere
}
