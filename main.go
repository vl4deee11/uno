package main

import (
	"fmt"
	"log"
	"time"
	"uno/communicator"
	"uno/engine"
)

func main() {
	com := communicator.Communicator{Url: "https://unoserver20210228202123.azurewebsites.net", Name: "vl4deee11"}
	engine.GenerateAllCards()
	discard := make([]engine.Card, 0)
	if err := com.Token(); err != nil {
		log.Fatal(err)
	}

	if err := com.StartGame(""); err != nil {
		log.Fatal(err)
	}

	for {
		board, err := com.Board()
		if err != nil {
			log.Fatal(err)
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
		nextCards := engine.GetNextCard(board.Hand, opponent, &board.CurrCard)
		fmt.Printf("Card on board = (%s)\n", board.CurrCard.String())
		for i:=0; i<len(nextCards);i++ {
			fmt.Printf("Card my turn = (%s)\n", nextCards[i].String())
		}
		fmt.Println(time.Since(t))
		if err := com.Move(nextCards); err != nil {
			log.Fatal(err)
		}
	}
}
