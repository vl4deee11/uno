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
	com := communicator.Communicator{Url: "https://unoserver20210228202123.azurewebsites.net", Name: "vl4deee11"}
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

	for {
		board, err := com.Board()
		if err != nil {
			log.Fatal(err)
		}
		if !board.MyMove {
			fmt.Println("Opponent turn")
			continue
		}
		fmt.Println("My turn")
		if board.Status != engine.InProcess {
			fmt.Println("Game is end")
			return
		}
		t := time.Now()
		discard = append(discard, board.CurrCard)

		opponent := engine.GetOpponentHand(board.Hand, discard)
		nextCards := engine.GetNextCard(board.Hand, opponent, &board.CurrCard)
		fmt.Printf("Card on board = (%s)\n", board.CurrCard.String())
		for i := 0; i < len(nextCards); i++ {
			fmt.Printf("Card my turn = (%s)\n", nextCards[i].String())
		}
		fmt.Println(time.Since(t))
		for i := 0; i < len(board.Hand); i++ {
			fmt.Printf("Card on hand = (%s)\n", board.Hand[i].String())
		}
		if err := com.Move(nextCards); err != nil {
			log.Fatal(err)
		}
	}
}
