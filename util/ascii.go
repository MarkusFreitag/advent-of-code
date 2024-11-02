package util

const (
	ASCIICodeUpperA = int('A') // 65
	ASCIICodeUpperZ = int('Z') // 90
	ASCIICodeLowerA = int('a') // 97
	ASCIICodeLowerZ = int('z') // 122
)

func CharToASCII[E byte | rune](value E) int { return int(value) }
func ASCIIToChar[E byte | rune](code int) E  { return E(code) }
