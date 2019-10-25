package version

import (
	"errors"
	"strconv"
	"strings"
)

//MaxNumber Max number
const MaxNumber = 9

// SeperateChar Seperate Char
const SeperateChar = "."

// DefaultVersion default version
const DefaultVersion = "1.0.0"

// UpdateVersion updates version of a service when update
func UpdateVersion(version string) (string, error) {
	if version == "" {
		return DefaultVersion, nil
	}
	versions := strings.Split(version, SeperateChar)

	for i := len(versions) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(versions[i])
		if err != nil {
			return "", errors.New("version:invalid")
		}
		if number < MaxNumber {
			number++
			versions[i] = strconv.Itoa(number)
			return strings.Join(versions, SeperateChar), nil
		}
		number = 0
		versions[i] = strconv.Itoa(number)
	}
	versions = append([]string{"1"}, versions...)
	// another solution : return strings.Join("1"+versions, SeperateChar), nil
	return strings.Join(versions, SeperateChar), nil
}
