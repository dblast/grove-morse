package main

import (
	"fmt"
	"os"
	"time"

	"grovepi"
)

const (
	PULSE_PAUSE   = 100
	LETTER_PAUSE  = 300
	MESSAGE_PAUSE = 2000
	WORD_PAUSE    = 700
	DOT_LENGTH    = 100
	DASH_LENGTH   = 300
)

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	err := g.PinMode(grovepi.D4, "output")
	if err != nil {
		fmt.Println(err)
	}

	pause := func(ms int) {
		time.Sleep(time.Duration(ms) * time.Millisecond)
	}

	pulse := func(ms int) {
		g.DigitalWrite(grovepi.D4, 1)
		pause(ms)
		g.DigitalWrite(grovepi.D4, 0)
		pause(PULSE_PAUSE)
	}

	sendMorse := func(msg string) {
		for _, s := range msg {
			switch s {
			case '.':
				pulse(DOT_LENGTH)
			case '-':
				pulse(DASH_LENGTH)
			case ' ':
				pause(LETTER_PAUSE)
			case '\t':
				pause(WORD_PAUSE)
			}
		}
	}

	mapping := map[rune]string{
		'a': ".-",
		'b': "-...",
		'c': "-.-.",
		'd': "-..",
		'e': ".",
		'f': "..-.",
		'g': "--.",
		'h': "....",
		'i': "..",
		'j': ".---",
		'k': "-.-",
		'l': ".-..",
		'm': "--",
		'n': "-.",
		'o': "---",
		'p': ".--.",
		'q': "--.-",
		'r': ".-.",
		's': "...",
		't': "-",
		'u': "..-",
		'v': "...-",
		'w': ".--",
		'x': "-..-",
		'y': "-.--",
		'z': "--..",
		' ': "\t",
	}

	encodeMorse := func(m string) string {
		result := ""
		for _, r := range m {
			result += mapping[r]
			result += " "
		}
		return result
	}

	for {
		for _, message := range os.Args[1:] {
			fmt.Println("Message:", message)
			morseCode := encodeMorse(message)
			fmt.Println("MorseCode:", morseCode)
			sendMorse(morseCode)
			fmt.Println("Sent!")
		}
	}
}
