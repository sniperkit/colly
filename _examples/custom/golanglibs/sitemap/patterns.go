package main

import (
	"plugin"
	// "github.com/mantyr/urls"
	// "github.com/opennota/linkify"
)

var (
	Numeric              = `^(\d+)$`
	AlphaNumeric         = `^([0-9A-Za-z]+)$`
	Alpha                = `^([A-Za-z]+)$`
	AlphaCapsOnly        = `^([A-Z]+)$`
	AlphaNumericCapsOnly = `^([0-9A-Z]+)$`
	Url                  = `^((http?|https?|ftps?):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	Email                = `^(.+@([\da-z\.-]+)\.([a-z\.]{2,6}))$`
	HashtagHex           = `^#([a-f0-9]{6}|[a-f0-9]{3})$`
	ZeroXHex             = `^0x([a-f0-9]+|[A-F0-9]+)$`
	IPv4                 = `^(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`
	IPv6                 = `^([0-9A-Fa-f]{0,4}:){2,7}([0-9A-Fa-f]{1,4}$|((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})$`
)
