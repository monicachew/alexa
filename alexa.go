package alexa

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	//"math"
	"os"
	"strconv"
)

// An EntriesFile represents a file containing compressed log entries.
type AlexaRank struct {
	init  bool
	m     map[string]int64
	count int
}

// Returns a rank for the site in question, or -1 if the host doesn't exist. Lower is better.
func (a *AlexaRank) GetRank(host string) (rank int64, err error) {
	if a.m == nil {
		return -1, errors.New("Alexa: Must initialize before getting rank")
	}
	val := a.m[host]
	if val == 0 {
		return -1, errors.New("Alexa: No rank for host")
	}
	return val, nil
}

// Returns a reputation between [0, 1] or -1 if the host does not exist in the ranking.
func (a *AlexaRank) GetReputation(host string) (reputation float32, err error) {
	rank, err := a.GetRank(host)
	if err != nil {
		return -1, err
	}
	// Rank starts at 1, normalize to 0 so the top-ranked site can have reputation = 1.0
	return 1.0 - float32(rank-1)/float32(a.count), nil
}

func (a *AlexaRank) Init(fileName string) {
	csv_in, err := os.Open(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open Alexa top 1M file: %s\n", err)
		os.Exit(1)
	}
	defer csv_in.Close()

	reader := csv.NewReader(csv_in)
	a.m = make(map[string]int64)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read records: %s\n", err)
			return
		}
		rank, _ := strconv.ParseInt(record[0], 0, 0)
		a.m[record[1]] = rank
	}
	a.count = len(a.m)
}
