package faker

import (
	"strconv"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func StringBool() string {
	if Intn32(2) == 0 {
		return "false"
	}
	return "true"
}

func Number(n int) string {
	if n > 1 {
		return NumericNotZero() + Numerics(n)
	}
	return Numerics(n)
}

func Numeric() string {
	return string(RuneBetween('0', '9'))
}

func Numerics(n int) string {
	return string(RunesBetween(n, '0', '9'))
}

func NumericNotZero() string {
	return string(RuneBetween('1', '9'))
}

func NumericsNotZero(n int) string {
	return string(RunesBetween(n, '1', '9'))
}

func NumericBetween(a int, b int) string {
	return string(strconv.Itoa(IntBetween(a, b)))
}

func Letter() string {
	return string(RuneBetween('a', 'z'))
}

func LetterIn(s string) string {
	return string(RuneIn([]rune(s)))
}

func Letters(n int) string {
	return string(RunesIn(n, []rune(letterBytes)))
}

func LettersIn(n int, s string) string {
	return string(RunesIn(n, []rune(s)))
}

func Alphanumeric() string {
	return string(RuneIn([]rune(alphanumeric)))
}

func Alphanumerics(n int) string {
	return string(RunesIn(n, []rune(alphanumeric)))
}

func Ascii() string {
	return string(RuneBetween('!', '~'))
}

func StringIn(s []string) string {
	return s[Intn(len(s))]
}
