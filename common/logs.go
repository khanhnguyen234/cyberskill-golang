package common

import (
	"github.com/fatih/color"
)

func LogStatus(err error, msg string) {
	if err != nil {
		color.Red("FAIL: %s %s", msg, err)
	} else {
		color.Green("SUCCESS: %s", msg)
	}
}

func LogError(err error, msg string) {
	if err != nil {
		color.Red("FAIL: %s %s", msg, err)
	}
}

func LogSuccess(msg string) {
	color.Green("SUCCESS: %s", msg)
}
