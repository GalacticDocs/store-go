package storing_errors

import (
	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	"github.com/iVitaliya/logger-go/utils"
)

func NoKeysError(msgType string, logEmptySpaces int) {
	logger.LogEmptySpace(logEmptySpaces)

	var str string
	switch msgType {
	case "clear":
		str = "There are no keys set in the map and thus there are no keys to clear from the map"
		break
	case "all":
		str = "There are no keys set in the map and thus I can't display any keys/values"
		break
	case "keys":
		str = "There are no keys set in the map and thus I can't display any keys"
		break
	case "values":
		str = "There are no keys set in the map and thus I can't display any values"
		break
	}

	logger.Error(str)
}

func NoKeyError(msgType string, key string) {
	logger.LogEmptySpace(2)

	var str string
	switch msgType {
	case "delete":
		str = utils.FormatString("Couldn't delete the key \"{0}\" as it either doesn't exist or it failed deleting it", []string{
			colors.Red(key),
		})
		break
	case "set":
		str = "The key couldn't be set in the map as the provided key was either empty or it contained invalid characters"
		break
	case "update":
		str = utils.FormatString("The key \"{0}\" couldn't be updated in the map as it couldn't be found as a existing key", []string{
			colors.Red(key),
		})
	}

	logger.Error(str)
}

func KeyNotFoundError(key string, msgType string) {
	logger.LogEmptySpace(2)

	var str string
	switch msgType {
	case "get":
		str = utils.FormatString("The key \"{0}\" couldn't be found in the map, are you sure you spelled it correctly?", []string{
			colors.Dim(colors.BrightGreen(key)),
		})
		break
	}

	logger.Error(str)
}
