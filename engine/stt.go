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

	Say("say card color")
	colorC := Listen()
	fmt.Print(string(colorC))
	Say("say card type")
	typeC := Listen()
	fmt.Print(string(typeC))

	return []Card{}
}
