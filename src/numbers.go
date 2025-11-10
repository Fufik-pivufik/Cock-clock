package main

import (
	"fmt"
)

func Conc(char rune, num int) string {
	result := ""
	for i := 0; i < num; i++ {
		result += string(char)
	}

	return result
}

func PrintD(time string, date string, locale string) {
	width, height := GetTermSize()

	lines := make([]string, height)
	i := 0

	midD := (height)/2 - height%5
	for ; i < midD; i++ {
		for j := 0; j < width; j++ {
			lines[i] += " "
		}
	}

	l := width/2 - width%5 - 25
	for j := i; j < i+6; j++ {

		lines[j] = Conc(' ', l)
	}

	for _, char := range time {
		degitlines := Belems[char]
		for j := 0; j < 6 && i+j < height; j++ {
			lines[i+j] += degitlines[j] + " "
		}
	}

	for j := 0; j < 6; j++ {
		lines[i+j] += "    " + Earth[j]
	}
	i += 6
	dateline := i

	for ; i < height; i++ {
		for j := 0; j < width; j++ {
			lines[i] += " "
		}
	}

	for numl, line := range lines {
		if numl == dateline {
			dlinesl := width/2 - len(date) - len(locale) + 6
			dline := Conc(' ', dlinesl)
			fmt.Printf("\033[38;2;100;100;150m%s %s\033[0m", dline, date)
			fmt.Printf("\033[38;2;80;80;100m      %s\033[0m\n", locale)
		} else {
			fmt.Printf("\033[38;2;255;200;255m%s\033[0m\n", line)
		}
	}
}

// big numbers (hours, minutes)
// █

var Belems = map[rune][]string{
	'0': {
		"███████",
		"██   ██",
		"██   ██",
		"██   ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'1': {
		"  ███  ",
		"   ██  ",
		"   ██  ",
		"   ██  ",
		"  ████ ",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'2': {
		"███████",
		"     ██",
		"███████",
		"██     ",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'3': {
		"███████",
		"     ██",
		"███████",
		"     ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'4': {
		"██   ██",
		"██   ██",
		"███████",
		"     ██",
		"     ██",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'5': {
		"███████",
		"██     ",
		"███████",
		"     ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'6': {
		"███████",
		"██     ",
		"███████",
		"██   ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'7': {
		"███████",
		"     ██",
		"     ██",
		"     ██",
		"     ██",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'8': {
		"███████",
		"██   ██",
		"███████",
		"██   ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	'9': {
		"███████",
		"██   ██",
		"███████",
		"     ██",
		"███████",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	':': {
		"        ",
		"   ██   ",
		"        ",
		"   ██   ",
		"        ",
		"󱋰󱋰󱋰󱋰󱋰󱋰󱋰󱋰",
	},

	' ': {
		" ",
		" ",
		" ",
		" ",
		" ",
	},
}

var Earth []string = []string{
	"  ▄█████▀█▄  ",
	" █▀██████▀ █ ",
	"▐  ▀███▀▌   ▌",
	"▐    ▀▄     ▌",
	" █    ▄██▄ █ ",
	"  ▀▄▄█████▀  ",
}

// small numbers (date, seconds)
// 
