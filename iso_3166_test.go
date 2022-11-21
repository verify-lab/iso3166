package iso3166

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAlpha2CountryCode(t *testing.T) {
	assert.False(t, IsAlpha2CountryCode(""))
	assert.False(t, IsAlpha2CountryCode("USA"))
	assert.False(t, IsAlpha2CountryCode("XX"))
	assert.True(t, IsAlpha2CountryCode("US"))
	assert.True(t, IsAlpha2CountryCode("mx"))
}

func TestIsAlpha3CountryCode(t *testing.T) {
	assert.False(t, IsAlpha3CountryCode(""))
	assert.False(t, IsAlpha3CountryCode("US"))
	assert.False(t, IsAlpha3CountryCode("XXX"))
	assert.True(t, IsAlpha3CountryCode("USA"))
	assert.True(t, IsAlpha3CountryCode("MEX"))
}

func TestCountryByCode(t *testing.T) {
	cases := []struct {
		needle string
		num    int
		name   string
		alpha2 string
		alpha3 string
	}{
		{"UKR", 804, "Ukraine", "ua", "ukr"},
		{"MEX", 484, "Mexico", "mx", "mex"},
		{"ua", 804, "Ukraine", "ua", "ukr"},
		{"pl", 616, "Poland", "pl", "pol"},
		{"xxx", 0, "", "", ""},
		{"xx", 0, "", "", ""},
		{"123", 0, "", "", ""},
	}

	for _, ts := range cases {
		num, name, alpha2, alpha3 := CountryByCode(ts.needle)

		assert.Equal(t, ts.num, num)
		assert.Equal(t, ts.name, name)
		assert.Equal(t, ts.alpha2, alpha2)
		assert.Equal(t, ts.alpha3, alpha3)
	}
}
