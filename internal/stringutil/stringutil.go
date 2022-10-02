package stringutil

func Reverse(stringToReverse string) string {
	stringToReverseAsByte := []rune(stringToReverse)
	for i, j := 0, len(stringToReverseAsByte)-1; i < j; i, j = i+1, j-1 {
		stringToReverseAsByte[i], stringToReverseAsByte[j] = stringToReverseAsByte[j], stringToReverseAsByte[i]
	}
	return string(stringToReverseAsByte)
}
