package main

func reverseString(s []byte) {
	b, e := 0, len(s)-1
	for b <= e {
		s[e], s[b] = s[b], s[e]
		b++
		e--
	}
}
