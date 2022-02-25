package Validate

import (
	"strconv"
)

// Finnish social security code must follow certain legal conventions. Validate incoming string that the code is according to convention
func ValidatesocialSecurityCode(socialSecurityCode string) bool {
	socialSecurityCodeLength := len(socialSecurityCode)
	if socialSecurityCodeLength != 11 {
		return false
	}
	if !validateCenturyBorn(string(socialSecurityCode)) {
		return false
	}
	if !validateBirthDay(socialSecurityCode) {
		return false
	}
	if !validateBirthMonth(socialSecurityCode) {
		return false
	}
	return true
}

//Social security code must include code to indicate century born.
func validateCenturyBorn(socialSecurityCode string) bool {
	acceptedCenturyCodes := [3]string{"-", "A", "+"}
	for _, centuryCode := range acceptedCenturyCodes {
		if centuryCode == string(socialSecurityCode[6]) {
			return true
		}
	}
	return false
}

//Validate that day born is at least less than 31 and more than 0
func validateBirthDay(socialSecurityCode string) bool {
	if dayBorn, err := strconv.Atoi(socialSecurityCode[0:1]); err == nil {
		if dayBorn > 31 && dayBorn > 0 {
			return false
		}
		return true
	} else {
		panic(err)
	}
}

//Validate that day born is at least less than 12 and more than 0
func validateBirthMonth(socialSecurityCode string) bool {
	if monthBorn, err := strconv.Atoi(socialSecurityCode[2:3]); err == nil {
		if monthBorn <= 12 && monthBorn > 0 {
			return false
		}
		return true
	} else {
		panic(err)
	}
}

//Finnish social security identity codes contain an end validation.
func validateEndValidation(socialSecurityCode string) bool {
	if monthBorn, err := strconv.Atoi(socialSecurityCode[6:10]); err == nil {
		if monthBorn <= 12 && monthBorn > 0 {
			return false
		}
		return true
	} else {
		panic(err)
	}
}
