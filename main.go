package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
	"uno/communicator"
	"uno/engine"
)

func main() {
	com := communicator.Communicator{Url: "https://unoserver20210412203209.azurewebsites.net", Name: "vl4deee11"}
	engine.GenerateAllCards()
	discard := make([]engine.Card, 0)
	if err := com.Token(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter token: ")
	oToken, _ := reader.ReadString('\n')
	oToken = oToken[:len(oToken)-1]
	if err := com.StartGame(oToken); err != nil {
		log.Fatal(err)
	}

	engine.Say("hello, my name is AI and now, i will play uno")
	for {
		board, err := com.Board()
		if err != nil {
			break
		}
		if !board.MyMove {
			fmt.Println("Opponent turn")
			continue
		}
		fmt.Println("My turn")
		if board.Status != engine.InProcess {
			fmt.Println("Game is end")
			break
		}
		t := time.Now()
		discard = append(discard, board.CurrCard)

		opponent := engine.GetOpponentHand(board.Hand, discard)

		engine.Say("i think about next turn")

		nextCards := engine.GetNextCard(board.Hand, opponent, &board.CurrCard)
		fmt.Printf("Card on board = (%s)\n", board.CurrCard.String())
		for i := 0; i < len(nextCards); i++ {
			fmt.Printf("Card my turn = (%s)\n", nextCards[i].String())
		}
		fmt.Println(time.Since(t))

		cardsN := len(board.Hand) - len(nextCards)
		if cardsN == 1 {
			engine.Say("uno")
		} else {
			engine.Say(fmt.Sprintf("i have %d cards in my hands", cardsN))
		}
		for i := 0; i < len(board.Hand); i++ {
			fmt.Printf("Card on hand = (%s)\n", board.Hand[i].String())
		}

		engine.SayTurn(nextCards)
		if err := com.Move(nextCards); err != nil {
			break
		}
	}

	engine.Say("game is end, bye")
}
