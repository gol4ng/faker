package en_GB

import "github.com/gol4ng/faker"

var formats = []string{
	"+44(0)##########",
	"+44(0)#### ######",
	"+44(0)#########",
	"+44(0)#### #####",
	"0##########",
	"0#########",
	"0#### ######",
	"0#### #####",
	"0### ### ####",
	"0### #######",
	"(0####) ######",
	"(0####) #####",
	"(0###) ### ####",
	"(0###) #######",
}

/**
 * An array of en_GB mobile (cell) phone number formats
 * @var array
 */
var mobileFormats = []string{
	// Local
	"07#########",
	"07### ######",
	"07### ### ###",
}

func PhoneNumber() string {
	return faker.Compute([]byte(faker.StringIn(formats)))
}

/**
 * Return a en_GB mobile phone number
 * @return string
 */
func MobileNumber() string {
	return faker.Compute([]byte(faker.StringIn(mobileFormats)))
}
