package service

import (
	"fmt"

	// external
	"github.com/fatih/color"
	entrevista "github.com/sniperkit/entrevista/pkg"
)

func createInterview() *entrevista.Interview {
	interview := entrevista.NewInterview()
	interview.ShowOutput = func(message string) {
		fmt.Print(color.GreenString(message))
	}
	interview.ShowError = func(message string) {
		color.Red(message)
	}
	return interview
}
