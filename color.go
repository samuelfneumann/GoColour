package gocolour

var (
	reset  = "\033[0m"
	bold   = "\033[1m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	gray   = "\033[37m"
	white  = "\033[97m"
)

// colorize wraps a given message in a given color.
func colorize(color, message string) string {
	return color + message + reset
}

// Red colours the message red
func Red(message string) string {
	return colorize(red, message)
}

// Bold makes the message bold
func Bold(message string) string {
	return colorize(bold, message)
}

// Green colours the message green
func Green(message string) string {
	return colorize(green, message)
}

// Yellow colours the message yellow
func Yellow(message string) string {
	return colorize(yellow, message)
}

// Blue colours the message blue
func Blue(message string) string {
	return colorize(blue, message)
}

// Purple colours the message purple
func Purple(message string) string {
	return colorize(purple, message)
}

// Cyab colours the message cyan
func Cyan(message string) string {
	return colorize(cyan, message)
}

// Gray colours the message gray
func Gray(message string) string {
	return colorize(gray, message)
}

// White colours the message white
func White(message string) string {
	return colorize(white, message)
}
