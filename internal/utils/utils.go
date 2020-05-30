package utils

func CutStringTail(org string, width int) (cutted string) {
	cutted = org
	if width < len(org) {
		cutted = org[:width]
	}
	return cutted
}
