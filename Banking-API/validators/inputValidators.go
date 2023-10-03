package validators

import (
	"regexp"
)

func ValidateName(name string) bool {

	if name != "" && regexp.MustCompile("^[a-zA-Z ]{2,30}$").MatchString(name) {
		return true
	}
	return false
}

func ValidatePassword(password string) bool {
    if password != "" && regexp.MustCompile(`^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,30}$`).MatchString(password) {
        return true
    }
    return false
}




