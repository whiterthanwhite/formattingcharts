package main

import (
	"encoding/json"
	"flag"
	"io"
	"log"
	"os"
	"sort"
	"time"
)

var (
	intervalType = flag.Int("i", 1, "interval type for graph formatting")
)

func main() {
	flag.Parse()
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err.Error())
	}
	var transactions []*Transaction
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		log.Fatal(err.Error())
	}
	formattedTransactions, err := FormattingCharts(transactions, IntervalType(*intervalType))
	if err != nil {
		log.Fatal(err.Error())
	}
	formattedTransactionsJSON, err := json.Marshal(&formattedTransactions)
	if err != nil {
		log.Fatal(err.Error())
	}
	io.WriteString(os.Stdout, string(formattedTransactionsJSON))
}

type UnknownIntervalError struct{}

func (u UnknownIntervalError) Error() string {
	return `unknown interval`
}

type IntervalType int

const (
	Hour  IntervalType = 0
	Day   IntervalType = 1
	Week  IntervalType = 2
	Month IntervalType = 3
)

type Transaction struct {
	Value     int       `json:"value"`
	Timestamp time.Time `json:"timestamp"`
}

func (t *Transaction) UnmarshalJSON(body []byte) error {
	type temp struct {
		Value     int
		Timestamp int64
	}

	var te temp
	err := json.Unmarshal(body, &te)
	if err != nil {
		return err
	}
	_ = te

	t.Value = te.Value
	t.Timestamp = time.Unix(te.Timestamp, 0)

	return nil
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	type temp struct {
		Value     int
		Timestamp int64
	}
	te := temp{
		Value:     t.Value,
		Timestamp: t.Timestamp.Unix(),
	}
	teJSON, err := json.Marshal(&te)
	if err != nil {
		return nil, err
	}
	return teJSON, nil
}

func FormattingCharts(transactions []*Transaction, intervalType IntervalType) ([]*Transaction, error) {
	var truncateDuration time.Duration
	switch intervalType {
	case Hour:
		truncateDuration = time.Hour
	case Day:
		truncateDuration = 24 * time.Hour
	case Week:
		truncateDuration = 7 * 24 * time.Hour
	case Month:
		truncateDuration = 4 * 7 * 24 * time.Hour
	default:
		return nil, UnknownIntervalError{}
	}

	groupedTransactions := make(map[time.Time]int)

	for _, transaction := range transactions {
		groupedTransactions[transaction.Timestamp.Truncate(truncateDuration)] = transaction.Value
	}

	newTransactions := make([]*Transaction, 0)
	for timestamp, value := range groupedTransactions {
		newTransaction := &Transaction{
			Value:     value,
			Timestamp: timestamp,
		}
		newTransactions = append(newTransactions, newTransaction)
	}

	sort.Slice(newTransactions, func(i, j int) bool {
		return !newTransactions[i].Timestamp.Before(newTransactions[j].Timestamp)
	})

	return newTransactions, nil
}
