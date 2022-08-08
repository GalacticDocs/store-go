package storing_errors

import (
	"strings"

	"github.com/iVitaliya/colors-go"
	"github.com/iVitaliya/logger-go"
	"github.com/iVitaliya/logger-go/utils"
)

// =========================
//
// # INVALID CHARACTERS TO USE
//
// =========================
func invalidCharacters() []string {
	return []string{
		"?",
		"!",
		"\\",
		"/",
		"|",
		"\"",
		"'",
		":",
		";",
		"[",
		"]",
		"{",
		"}",
		"=",
		"\u200b",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
		"*",
		"(",
		")",
	}
}

// ==================
//
// # PROHIBITED TO USE
//
// ==================
func bannedKeys() []string {
	return []string{
		"pussy",
		"p*ssy",
		"p**sy",
		"dick",
		"d*ck",
		"cock",
		"cuck",
		"cuckold",
		"c*ck",
		"c**k",
		"penis",
		"fuck",
		"whore",
		"hoe",
		"cunt",
		"f*ck",
		"f**k",
		"f***",
	}
}

func CheckIfBadKey(key string) bool {
	for _, v := range invalidCharacters() {
		str := strings.Contains(key, v)
		if str {
			logger.Error(utils.FormatString("You defined \"{0}\" in your key as seperator which isn't an allowed character, please use one of the following characters instead as seperator: \"-\", \".\", \"_\"", []string{
				colors.Red(v),
			}))
			return false
		}
	}

	for _, v := range bannedKeys() {
		str := strings.Contains(key, v)
		if str {
			logger.Error(utils.FormatString("You defined \"{0}\" in your key which isn't an allowed word, please refractor your key!", []string{
				colors.Red(v),
			}))
			return false
		}
	}

	return true
}

func CheckIfBadJoinSeperator(value string) bool {
	badSeperators := []string{
		"\u200b",
		"!",
		"?",
		"+",
		"=",
		"@",
		"#",
		"$",
		"%",
		"^",
		"&",
		"*",
		"(",
		")",
		"_",
	}

	for _, v := range badSeperators {
		if v == value {
			logger.LogFormatted(logger.ErrorState, "The character \"{0}\" isn't allowed to be used as a joinable seperator, please use one of the following characters instead as seperator: \"-\", \".\", \",\", \"|\"", []string{
				colors.Red(v),
			})

			return false
		}
	}

	return true
}
