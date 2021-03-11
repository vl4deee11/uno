package main

import (
	"fmt"
	"time"
	"uno/communicator"
	"uno/engine"
)

func main() {
	dlr := communicator.Communicator{Url: "https://unoserver20210228202123.azurewebsites.net", Name: "vl4deee11"}
	engine.GenerateAllCards()
	discard := make([]engine.Card, 0)
	if err := dlr.Token(); err != nil {
		panic(err)
	}

	if err := dlr.StartGame(""); err != nil {
		panic(err)
	}

	for {
		board, err := dlr.Board()
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
		_ = engine.GetOpponentHand(board.Hand, discard)
		fmt.Println(time.Since(t))

	}
}
