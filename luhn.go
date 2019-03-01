package faker

import "strconv"

func checksum(value int) int {
	stringValue := strconv.Itoa(value)
	stringValueLen := len(stringValue)
	sum := 0
	b64 := uint64(10)

	for i := stringValueLen - 1; i >= 0; i -= 2 {
		c, _ := strconv.Atoi(stringValue[i : i+1])
		sum += c
	}
	for i := stringValueLen - 2; i >= 0; i -= 2 {
		c, _ := strconv.Atoi(stringValue[i : i+1])
		d := uint64(c * 2)
		for ; d > 0; d /= b64 {
			sum += int(d % b64)
		}
	}
	return sum % 10
}

func checkdigit(value int) int {
	checkdigit := checksum(value * 10)
	if checkdigit == 0 {
		return 0
	}
	return 10 - checkdigit
}

func Luhn(value int) int {
	checkdigit := checkdigit(value)
	return value*len(strconv.Itoa(checkdigit))*10 + checkdigit
}
