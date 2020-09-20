package utilities

import "fmt"

var (
	Info    = Teal
	Warn    = Yellow
	Fatal   = Red
	Success = Green
)

var (
	White   = Color("\033[0;30m%s\033[0m")
	Red     = Color("\033[0;31m%s\033[0m")
	Green   = Color("\033[0;32m%s\033[0m")
	Yellow  = Color("\033[0;33m%s\033[0m")
	Purple  = Color("\033[0;34m%s\033[0m")
	Magenta = Color("\033[0;35m%s\033[0m")
	Teal    = Color("\033[0;36m%s\033[0m")

	LightWhite   = Color("\033[1;30m%s\033[0m")
	LightRed     = Color("\033[1;31m%s\033[0m")
	LightGreen   = Color("\033[1;32m%s\033[0m")
	LightYellow  = Color("\033[1;33m%s\033[0m")
	LightBlue    = Color("\033[1;34m%s\033[0m")
	LightMagenta = Color("\033[1;35m%s\033[0m")
	LightTeal    = Color("\033[1;36m%s\033[0m")
	Normal       = Color("\033[1;37m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	return func(args ...interface{}) string {
		return fmt.Sprintf(colorString, fmt.Sprint(args...))
	}
}
