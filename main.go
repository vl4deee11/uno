package main

import (
	"fmt"
	"time"
	"uno/communicator"
	"uno/engine"
)

func main() {
	com := communicator.Communicator{Url: "https://unoserver20210228202123.azurewebsites.net", Name: "vl4deee11"}
	engine.GenerateAllCards()
	discard := make([]engine.Card, 0)
	if err := com.Token(); err != nil {
		panic(err)
	}

	if err := com.StartGame(""); err != nil {
		panic(err)
	}

	for {
		board, err := com.Board()
		if err != nil {
			panic(err)
		}
		if !board.MyMove {
			continue
		}
		if board.Status != engine.InProcess {
			fmt.Println("Game is end")
			return
		}
		fmt.Println("Next step")
		t := time.Now()
		discard = append(discard, board.CurrCard)
		if board.CurrCard.Type == engine.Skip {
			if err := com.Move(nil); err != nil {
				panic(err)
			}
		}

		opponent := engine.GetOpponentHand(board.Hand, discard)
		nextCard := engine.GetNextCard(board.Hand, opponent, &board.CurrCard)
		fmt.Printf("Card on board = (%s)\n", board.CurrCard.String())
		fmt.Printf("Card my turn card = (%s)\n", nextCard.String())
		fmt.Println(time.Since(t))
		if err := com.Move(nextCard); err != nil {
			panic(err)
		}
	}
}
