package engine

import (
	"fmt"
	"os/exec"
	"regexp"
)

var cmdSTT = fmt.Sprintf("cd %s && . ./stt/venv/bin/activate && python3 stt/stt.py", dir)

func Listen() []byte {
	b, _ := exec.Command(
		"bash",
		"-c",
		cmdSTT,
	).Output()

	return b
}

func HelpToCard(cardsOnHand []Card) []Card {
	Say("do you want to help me ?")

	help := Listen()
	if m, err := regexp.Match(`\bno\b`, help); err != nil || m {
		return nil
	}

	c, f := tryToRecognizeColor()
	if !f {
		Say("i cannot recognize what did you say")
		return nil
	}

	t, f := tryToRecognizeType()
	if !f {
		Say("i cannot recognize what did you say")
		return nil
	}

	var n int8 = -1
	if t == Numeric {
		n, f = tryToRecognizeNumber()
		if !f {
			Say("i cannot recognize what did you say")
			return nil
		}
	}

	card := Card{}
	f = false

	for i := range cardsOnHand {
		if cardsOnHand[i].Type == t && cardsOnHand[i].Color == c && cardsOnHand[i].Num == n {
			card = cardsOnHand[i]
			f = true
		}
	}

	if f {
		return []Card{card}
	}
	return nil
}

func tryToRecognizeColor() (Color, bool) {
	Say("say card color")
	colorC := Listen()
	fmt.Print(string(colorC))
	return Red, true
}

func tryToRecognizeType() (CardType, bool) {
	Say("say card type")
	typeC := Listen()
	fmt.Print(string(typeC))

	return Numeric, true
}

func tryToRecognizeNumber() (int8, bool) {
	Say("say card number")
	n := Listen()
	fmt.Print(string(n))

	return 1, true
}
