package phoneUtils

import (
	logger "github.com/sirupsen/logrus"
	"regexp"
)

var keymap = map[string]map[string]bool{
	"0": {"7": true, "8": true, "9": true, "0": true},
	"1": {"2": true, "4": true, "5": true, "1": true},
	"2": {"1": true, "3": true, "4": true, "5": true, "6": true, "2": true},
	"3": {"2": true, "5": true, "6": true, "3": true},
	"4": {"1": true, "2": true, "5": true, "7": true, "8": true, "4": true},
	"5": {"1": true, "2": true, "3": true, "4": true, "6": true, "7": true, "8": true, "9": true, "5": true},
	"6": {"2": true, "3": true, "5": true, "8": true, "9": true, "6": true},
	"7": {"4": true, "5": true, "8": true, "0": true, "7": true},
	"8": {"4": true, "5": true, "6": true, "7": true, "9": true, "0": true, "8": true},
	"9": {"5": true, "6": true, "8": true, "0": true, "9": true},
}

//Regex patter to standardize the number string
var NumericalRegex = "[^0-9]+"

//Given a number as a string this will evaluate whether it is easy to dial based on the standard phone keyboard layout
//Returns true if easy to dial else false
func IsEasyDial(number string) bool {
	number, isValid := validateAndStandardizeNumber(number)
	logger.Debugf("Phone number after standardization: %s\n", number)

	//Fail if validation checks fails
	if !isValid {
		return false
	}

	//If the number is a single digit, it is easy dial. Return true
	if len(number) == 1 {
		return true
	}

	//Set it to default false
	isNumberEasyDial := false
	//Loop through from the second digit till the last digit in the given number to check for proximity
	//If proximity check fails breaks the loop
	for i := 1; i < len(number); i++ {
		previousDigit := string(number[i-1])
		currentDigit := string(number[i])
		logger.Tracef("CurrentDigit: %s, keymap of previous digit: %v ", currentDigit, keymap[previousDigit][currentDigit])

		//Check the digit is in the proximity using the indexed keyboard layout in the map
		isCurrentDigitAdjacent, _ := keymap[previousDigit][currentDigit]

		//If any digit fails proximity check we can fast fail here
		if isCurrentDigitAdjacent {
			isNumberEasyDial = isCurrentDigitAdjacent
		} else {
			isNumberEasyDial = false
			logger.Infof("Current digit: %s is not adjacent to previous digit: %s", currentDigit, previousDigit)
			break
		}
	}
	return isNumberEasyDial
}

//Helper function to validate the allowed length and standardize the phone number
//Returns the formatted string
func validateAndStandardizeNumber(number string) (string, bool) {
	number, isValid := standardizeNumber(number)

	//Fail validation if standardize fails and propagate result
	if !isValid {
		return number, false
	}

	//Fail validation if standardize number is empty
	if number == "" {
		return number, false
	}

	//Check for allowed length if not log a warning
	if len(number) < 7 || len(number) > 10 {
		logger.Warnf("Given number: %s is not in the expected length range of 7-10", number)
	}
	return number, true
}

//Helper function to make sure the number does not have any alphabetic and special characters
//Returns the formatted string
func standardizeNumber(number string) (string, bool) {
	//Remove any non-numerical character in the given number
	reg, err := regexp.Compile(NumericalRegex)
	if err != nil {
		logger.Error(err)
		return "", false
	}

	return reg.ReplaceAllString(number, ""), true
}
