package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

const path = "entries.json"

type RaffleNodes struct {
	RaffleNodes []raffleEntry `json:"raffle_entry"`
}

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
	// TODO: Fill in definition
	Raffle_id string `json:"id"`
	Names     string `json:"name"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	file, _ := ioutil.ReadFile(path)
	data := RaffleNodes{}

	_ = json.Unmarshal([]byte(file), &data)

	for i := 0; i < len(data.RaffleNodes); i++ {
		fmt.Println("Raffle Id: ", data.RaffleNodes[i].Raffle_id)
		fmt.Println("Name: ", data.RaffleNodes[i].Names)
	}
	return data.RaffleNodes
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
