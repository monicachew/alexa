package alexa

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

// An EntriesFile represents a file containing compressed log entries.
type AlexaRank struct {
	init bool
	m map[string]int64
}

func (a* AlexaRank) GetRank(host string) (rank int64, err error) {
	if a.m == nil {
		return 0, errors.New("Alexa: Must initialize before getting rank")
	}
	val := a.m[host]
	if val == 0 {
		return 0, errors.New("Alexa: No rank for host")
	}
	return val, nil
}

func (a* AlexaRank) Init(fileName string) {
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
}
