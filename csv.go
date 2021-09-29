package main

import (
	"encoding/csv"
	"os"
	"sync"
)

type CsvWriter struct {
	mutex     *sync.Mutex
	csvWriter *csv.Writer
}

func NewCsvWriter(file string) (*CsvWriter, error) {
	csvFile, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	w := csv.NewWriter(csvFile)
	return &CsvWriter{mutex: &sync.Mutex{}, csvWriter: w}, nil
}

func (w *CsvWriter) Write(record []string) {
	w.mutex.Lock()
	w.csvWriter.Write(record)
	w.mutex.Unlock()
}

func (w *CsvWriter) Flush() {
	w.mutex.Lock()
	w.csvWriter.Flush()
	w.mutex.Unlock()
}
