package faker_test

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"

	"github.com/gol4ng/faker"
)

func Test_MacProcessor(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"PPC", "U; PPC", "U; PPC", "U; PPC", "PPC", "U; Intel"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.MacProcessor())
	}
}

func Test_LinuxProcessor(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{"x86_64", "x86_64", "x86_64", "x86_64", "x86_64", "i686"}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.LinuxProcessor())
	}
}

func Test_UserAgent(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(10000000, 0) })
	defer monkey.UnpatchAll()

	faker.FakerRand.Seed(1)
	tests := []string{
		"Mozilla/5.0 (Windows CE; ku-IR; rv:1.9.0.20) Gecko/19700224 Firefox/36.0",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.0; Trident/3.0)",
		"Mozilla/5.0 (Windows; U; Windows NT 4.0) AppleWebKit/534.35.3 (KHTML, like Gecko) Version/4.0.3 Safari/534.35.3",
		"Mozilla/5.0 (X11; Linux x86_64; rv:6.0) Gecko/19700404 Firefox/35.0",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 5.1; Trident/4.0)",
		"Opera/8.59 (Windows CE; nds-DE) Presto/2.8.182 Version/11.00",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.UserAgent())
	}
}

func Test_Chrome(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Mozilla/5.0 (Macintosh; PPC Mac OS X 10_7_0) AppleWebKit/5311 (KHTML, like Gecko) Chrome/36.0.817.0 Mobile Safari/5311",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_8) AppleWebKit/5351 (KHTML, like Gecko) Chrome/38.0.860.0 Mobile Safari/5351",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Chrome())
	}
}

func Test_Firefox(t *testing.T) {
	monkey.Patch(time.Now, func() time.Time { return time.Unix(10000000, 0) })
	defer monkey.UnpatchAll()

	faker.FakerRand.Seed(1)
	tests := []string{
		"Mozilla/5.0 (X11; Linux i686; rv:5.0) Gecko/19700323 Firefox/36.0",
		"Mozilla/5.0 (X11; Linux x86_64; rv:5.0) Gecko/19700314 Firefox/35.0",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Firefox())
	}
}

func Test_Safari(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Mozilla/5.0 (iPhone; CPU iPhone OS 7_1_1 like Mac OS X; es-VE) AppleWebKit/533.47.4 (KHTML, like Gecko) Version/3.0.5 Mobile/8B117 Safari/6533.47.4",
		"Mozilla/5.0 (Windows; U; Windows NT 4.0) AppleWebKit/532.16.4 (KHTML, like Gecko) Version/4.0.2 Safari/532.16.4",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Safari())
	}
}

func Test_Opera(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Opera/8.76 (Windows NT 5.2; ku-TR) Presto/2.10.168 Version/10.00",
		"Opera/8.82 (Windows NT 6.0; es-NI) Presto/2.9.303 Version/11.00",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.Opera())
	}
}

func Test_InternetExplorer(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Mozilla/5.0 (compatible; MSIE 7.0; Windows NT 5.2; Trident/4.0)",
		"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 5.0; Trident/3.0)",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.InternetExplorer())
	}
}

func Test_WindowsPlatformToken(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Windows NT 5.01",
		"Windows NT 5.2",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.WindowsPlatformToken())
	}
}

func Test_MacPlatformToken(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"Macintosh; PPC Mac OS X 10_6_0",
		"Macintosh; U; PPC Mac OS X 10_7_8",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.MacPlatformToken())
	}
}

func Test_LinuxPlatformToken(t *testing.T) {
	faker.FakerRand.Seed(1)
	tests := []string{
		"X11; Linux x86_64",
		"X11; Linux x86_64",
	}
	for _, expected := range tests {
		assert.Equal(t, expected, faker.LinuxPlatformToken())
	}
}
