package faker

import (
	"strconv"
)

/**
 * @example +27113456789
 * @return string
 */
func E164PhoneNumber() string {
	return Compute([]byte("+%############"))
}

/**
 * International Mobile Equipment Identity (IMEI)
 *
 * @link http://en.wikipedia.org/wiki/International_Mobile_Station_Equipment_Identity
 * @link http://imei-number.com/imei-validation-check/
 * @example '720084494799532'
 * @return int $imei
 */
func Imei() string {
	value, _ := strconv.Atoi(Compute([]byte("##############")))
	return strconv.Itoa(Luhn(value))
}
