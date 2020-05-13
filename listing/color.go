package listing

const reset = "\033[0m"
const red = "\033[31m"
const green = "\033[32m"
const yellow = "\033[33m"
const blue = "\033[34m"
const purple = "\033[35m"
const cyan = "\033[36m"
const gray = "\033[37m"
const white = "\033[97m"

func colorize(s, color string) string {
	return color + s + reset
}
