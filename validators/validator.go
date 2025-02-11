package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// Validator instance
var Validate = validator.New()

// Fungsi validasi custom untuk username (regexp)
func UsernameRegexp(fl validator.FieldLevel) bool {
	// Cek apakah username hanya berisi angka
	re := regexp.MustCompile("^[0-9]+$")
	return re.MatchString(fl.Field().String())
}
