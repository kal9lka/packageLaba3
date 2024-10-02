package stringutils

func Stringmirror(str string) string {
	var newstr string
	for i := len(str); i > 0; i-- {
		newstr += string(str[i])
	}
	return newstr
}
