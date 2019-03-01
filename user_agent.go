package faker

var windowsPlatformTokens = []string{"Windows NT 6.2", "Windows NT 6.1", "Windows NT 6.0", "Windows NT 5.2", "Windows NT 5.1", "Windows NT 5.01", "Windows NT 5.0", "Windows NT 4.0", "Windows 98; Win 9x 4.90", "Windows 98", "Windows 95", "Windows CE",}

var linuxProcessor = []string{"i686", "x86_64"}

var macProcessor = []string{"Intel", "PPC", "U; Intel", "U; PPC"}

func MacProcessor() string {
	return StringIn(macProcessor)
}

func LinuxProcessor() string {
	return StringIn(linuxProcessor)
}

func UserAgent() string {
	userAgents := []func() string{
		Chrome,
		Firefox,
		Safari,
		Opera,
		InternetExplorer,
	}
	return userAgents[Intn(len(userAgents))]()
}

func Chrome() string {
	saf := NumericBetween(531, 536) + NumericBetween(0, 2)

	platformTokens := []string{
		LinuxPlatformToken(),
		WindowsPlatformToken(),
		MacPlatformToken(),
	}
	return "Mozilla/5.0 " + "(" + StringIn(platformTokens) + ") AppleWebKit/" + saf + " (KHTML, like Gecko) Chrome/" + NumericBetween(36, 40) + ".0." + NumericBetween(800, 899) + ".0 Mobile Safari/" + saf
}

func Firefox() string {
	ver := "Gecko/" + TimePast().Format("20060102") + " Firefox/" + NumericBetween(35, 37) + ".0"

	platforms := []string{
		"(" + WindowsPlatformToken() + "; " + StringIn(localeDash) + "; rv:1.9." + NumericBetween(0, 2) + ".20) " + ver,
		"(" + LinuxPlatformToken() + "; rv:" + NumericBetween(5, 7) + ".0) " + ver,
		"(" + MacPlatformToken() + " rv:" + NumericBetween(2, 6) + ".0) " + ver,
	}
	return "Mozilla/5.0 " + StringIn(platforms)
}

func Safari() string {
	saf := NumericBetween(531, 535) + "." + NumericBetween(1, 50) + "." + NumericBetween(1, 7)
	ver := ""
	if Bool() {
		ver = NumericBetween(4, 5) + "." + NumericBetween(0, 1)
	} else {
		ver = NumericBetween(4, 5) + ".0." + NumericBetween(1, 5)
	}

	mobileDevices := []string{
		"iPhone; CPU iPhone OS",
		"iPad; CPU OS",
	}

	platforms := []string{
		"(Windows; U; " + WindowsPlatformToken() + ") AppleWebKit/" + saf + " (KHTML, like Gecko) Version/" + ver + " Safari/" + saf,
		"(" + MacPlatformToken() + " rv:" + NumericBetween(2, 6) + ".0; " + StringIn(localeDash) + ") AppleWebKit/" + saf + " (KHTML, like Gecko) Version/" + ver + " Safari/" + saf,
		"(" + StringIn(mobileDevices) + " " + NumericBetween(7, 8) + "_" + NumericBetween(0, 2) + "_" + NumericBetween(1, 2) + " like Mac OS X; " + StringIn(localeDash) + ") AppleWebKit/" + saf + " (KHTML, like Gecko) Version/" + NumericBetween(3, 4) + ".0.5 Mobile/8B" + NumericBetween(111, 119) + " Safari/6" + saf,
	}
	return "Mozilla/5.0 " + StringIn(platforms)
}

func Opera() string {
	platformTokens := []string{
		LinuxPlatformToken(),
		WindowsPlatformToken(),
	}
	return "Opera/" + NumericBetween(8, 9) + "." + NumericBetween(10, 99) + " " + "(" + StringIn(platformTokens) + "; " + StringIn(localeDash) + ") Presto/2." + NumericBetween(8, 12) + "." + NumericBetween(160, 355) + " Version/" + NumericBetween(10, 12) + ".00"
}

func InternetExplorer() string {
	return "Mozilla/5.0 (compatible; MSIE " + NumericBetween(5, 11) + ".0; " + WindowsPlatformToken() + "; Trident/" + NumericBetween(3, 5) + "." + NumericBetween(0, 1) + ")"
}

func WindowsPlatformToken() string {
	return StringIn(windowsPlatformTokens)
}

func MacPlatformToken() string {
	return "Macintosh; " + StringIn(macProcessor) + " Mac OS X 10_" + NumericBetween(5, 8) + "_" + NumericBetween(0, 9)
}

func LinuxPlatformToken() string {
	return "X11; Linux " + StringIn(linuxProcessor)
}
