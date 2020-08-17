package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"reflect"
	"strconv"
	"time"
)

const DATE_PATTERN = "%d-%02d-%02dT%02d:%02d:%02d-00:00"

type Transaction struct {
	Type        string
	SubType     string
	FromAccount string
	ToAccount   string
	Value       float64
	Time        time.Time
	DeviceType  string
}

func (t *Transaction) line() string {
	return fmt.Sprintf("%10s", t.Type) +
		fmt.Sprintf("%9s", t.SubType) +
		fmt.Sprintf("%29s", t.FromAccount) +
		fmt.Sprintf("%29s", t.ToAccount) +
		fmt.Sprintf("%029f", t.Value) +
		fmt.Sprintf(DATE_PATTERN, t.Time.Year(), t.Time.Month(), t.Time.Day(),
			t.Time.Hour(), t.Time.Minute(), t.Time.Second()) +
		fmt.Sprintf("%29s", t.DeviceType) + "\n"
}

type transactionBuilder struct {
	Type           string
	SubType        string
	OriginAccount  string
	DestinyAccount string
	TValue         float64
	TTime          time.Time
	Device         string
}

func (tb *transactionBuilder) PaymentType(ptype string) *transactionBuilder {
	tb.Type = ptype
	return tb
}

func (tb *transactionBuilder) PaymentSubType(stype string) *transactionBuilder {
	tb.SubType = stype
	return tb
}

func (tb *transactionBuilder) FromAccount(fromAccount string) *transactionBuilder {
	tb.OriginAccount = fromAccount
	return tb
}

func (tb *transactionBuilder) ToAccount(toAccount string) *transactionBuilder {
	tb.DestinyAccount = toAccount
	return tb
}

func (tb *transactionBuilder) Value(value float64) *transactionBuilder {
	tb.TValue = value
	return tb
}

func (tb *transactionBuilder) Time(time time.Time) *transactionBuilder {
	tb.TTime = time
	return tb
}

func (tb *transactionBuilder) DeviceType(device string) *transactionBuilder {
	tb.Device = device
	return tb
}

func (tb *transactionBuilder) Build() *Transaction {
	return &Transaction{
		Type:        tb.Type,
		SubType:     tb.SubType,
		FromAccount: tb.OriginAccount,
		ToAccount:   tb.DestinyAccount,
		Value:       tb.TValue,
		Time:        tb.TTime,
		DeviceType:  tb.Device,
	}
}

func New() *transactionBuilder {
	return &transactionBuilder{}
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()
	return keys[rand.Intn(len(keys))].Interface()
}

func main() {
	log.Printf("Starting file generator....")
	start := time.Now()
	f, _ := os.Create(os.Getenv("OUT_FOLDER") + time.Now().String() + ".txt")
	defer f.Close()
	w := bufio.NewWriter(f)
	defer w.Flush()
	ts := map[string][]string{"TED": {"TED"}, "DOC": {"DOC"}, "CARD": {"VISA", "MASTER"}}
	devices := []string{"POS", "SMART", "BROWSER"}
	fromAccount := os.Getenv("FROM_ACCOUNT")
	toAccount := os.Getenv("TO_ACCOUNT")
	from, _ := strconv.Atoi(fromAccount)
	to, _ := strconv.Atoi(toAccount)
	for i := from; i < to; i++ {
		value := (rand.Float64() * 100) + 100
		device := devices[rand.Intn(len(devices))]
		pType := MapRandomKeyGet(ts).(string)
		pSubType := ts[pType][rand.Intn(len(ts[pType]))]
		rt := rand.Intn(to-from) + from
		transaction := New().FromAccount(strconv.Itoa(i)).ToAccount(strconv.Itoa(rt)).PaymentType(pType).PaymentSubType(pSubType).Time(randomDate()).Value(value).DeviceType(device).Build()
		w.WriteString(transaction.line())
	}
	elapsed := time.Since(start)
	log.Printf("File generated took %s", elapsed)
}

func randomDate() time.Time {
	hour := rand.Intn(24)
	return time.Now().Add(time.Duration(-hour) * time.Hour)
}
