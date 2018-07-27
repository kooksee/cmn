package cmn

import (
	"fmt"
	"strings"
)

type color struct {
	aNSIReset      string
	aNSIBright     string
	aNSIDim        string
	aNSIUnderscore string
	aNSIBlink      string
	aNSIReverse    string
	aNSIHidden     string

	aNSIFgBlack   string
	aNSIFgRed     string
	aNSIFgGreen   string
	aNSIFgYellow  string
	aNSIFgBlue    string
	aNSIFgMagenta string
	aNSIFgCyan    string
	aNSIFgWhite   string

	aNSIBgBlack   string
	aNSIBgRed     string
	aNSIBgGreen   string
	aNSIBgYellow  string
	aNSIBgBlue    string
	aNSIBgMagenta string
	aNSIBgCyan    string
	aNSIBgWhite   string
}

var Color = color{
	aNSIReset:      "\x1b[0m",
	aNSIBright:     "\x1b[1m",
	aNSIDim:        "\x1b[2m",
	aNSIUnderscore: "\x1b[4m",
	aNSIBlink:      "\x1b[5m",
	aNSIReverse:    "\x1b[7m",
	aNSIHidden:     "\x1b[8m",

	aNSIFgBlack:   "\x1b[30m",
	aNSIFgRed:     "\x1b[31m",
	aNSIFgGreen:   "\x1b[32m",
	aNSIFgYellow:  "\x1b[33m",
	aNSIFgBlue:    "\x1b[34m",
	aNSIFgMagenta: "\x1b[35m",
	aNSIFgCyan:    "\x1b[36m",
	aNSIFgWhite:   "\x1b[37m",

	aNSIBgBlack:   "\x1b[40m",
	aNSIBgRed:     "\x1b[41m",
	aNSIBgGreen:   "\x1b[42m",
	aNSIBgYellow:  "\x1b[43m",
	aNSIBgBlue:    "\x1b[44m",
	aNSIBgMagenta: "\x1b[45m",
	aNSIBgCyan:    "\x1b[46m",
	aNSIBgWhite:   "\x1b[47m",
}

// color the string s with color 'color'
// unless s is already colored
func (c color) treat(s string, color string) string {
	if len(s) > 2 && s[:2] == "\x1b[" {
		return s
	} else {
		return color + s + c.aNSIReset
	}
}

func (c color) treatAll(color string, args ...interface{}) string {
	var parts []string
	for _, arg := range args {
		parts = append(parts, c.treat(fmt.Sprintf("%v", arg), color))
	}
	return strings.Join(parts, "")
}

func (c color) Black(args ...interface{}) string {
	return c.treatAll(c.aNSIFgBlack, args...)
}

func (c color) Red(args ...interface{}) string {
	return c.treatAll(c.aNSIFgRed, args...)
}

func (c color) Green(args ...interface{}) string {
	return c.treatAll(c.aNSIFgGreen, args...)
}

func (c color) Yellow(args ...interface{}) string {
	return c.treatAll(c.aNSIFgYellow, args...)
}

func (c color) Blue(args ...interface{}) string {
	return c.treatAll(c.aNSIFgBlue, args...)
}

func (c color) Magenta(args ...interface{}) string {
	return c.treatAll(c.aNSIFgMagenta, args...)
}

func (c color) Cyan(args ...interface{}) string {
	return c.treatAll(c.aNSIFgCyan, args...)
}

func (c color) White(args ...interface{}) string {
	return c.treatAll(c.aNSIFgWhite, args...)
}
