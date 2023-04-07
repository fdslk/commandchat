package commandchat

func ReverseSlice(s []interface{}) []interface{} {
	reversed := make([]interface{}, len(s))
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		reversed[i], reversed[j] = s[j], s[i]
	}
	return reversed
}
