package mylog

import (
	"fmt"
	"time"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARN
	ERROR
)

const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorGray   = "\033[90m"
)

var (
	globalLevel  = INFO
	useTimestamp = true
	useBold      = false
)

// SetLevel mengatur level minimum log yang akan dicetak.
func SetLevel(l Level) { globalLevel = l }

// SetTimestamp mengaktifkan atau mematikan timestamp.
func SetTimestamp(enabled bool) { useTimestamp = enabled }

// SetBold mengaktifkan atau mematikan huruf tebal pada prefix level.
func SetBold(enabled bool) { useBold = enabled }

func print(lvl Level, color, prefix, format string, args ...any) {
	if lvl < globalLevel {
		return
	}

	msg := fmt.Sprintf(format, args...)

	prefixStyle := color
	if useBold {
		prefixStyle = colorBold + color
	}

	header := ""
	if useTimestamp {
		header = time.Now().Format("2006/01/02 15:04:05") + " "
	}

	fmt.Printf("%s%s[%s]%s %s\n", header, prefixStyle, prefix, colorReset, msg)
}

func Debug(format string, args ...any) { print(DEBUG, colorGray, "DBG", format, args...) }
func Info(format string, args ...any)  { print(INFO, colorCyan, "INF", format, args...) }
func Warn(format string, args ...any)  { print(WARN, colorYellow, "WRN", format, args...) }
func Error(format string, args ...any) { print(ERROR, colorRed, "ERR", format, args...) }
