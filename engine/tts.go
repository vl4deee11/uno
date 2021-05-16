package engine

import (
	"fmt"
	"os/exec"
	"strings"
)

const dir = "/Users/macbookpro/Downloads/learn/uno/"

var prefTTS = fmt.Sprintf("cd %s && . ./tts/venv/bin/activate && python3 tts/tts.py --text '", dir)
const postTTS = "' && afplay gen.mp3"

func SayTurn(cards []Card) {
	if len(cards) > 0 {
		phrases := make([]string, len(cards))
		for i := range cards {
			phrases[i] = cards[i].StringTTS()
		}
		exec.Command(
			"bash",
			"-c",
			fmt.Sprintf(
				"%si am going to make a move with the cards %s%s",
				prefTTS,
				strings.Join(phrases, ", and "),
				postTTS,
			),
		).Output()
		return
	}

	exec.Command(
		"bash",
		"-c",
		fmt.Sprintf(
			"%s%s%s",
			prefTTS,
			"i have no cards to turn, and i will rely on luck",
			postTTS,
		),
	).Output()
}

func Say(phrase string) {
	exec.Command("bash", "-c", fmt.Sprintf("%s%s%s", prefTTS, phrase, postTTS)).Output()
}
