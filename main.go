package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/domainr/whois"
)

type Output struct {
	option string
	csv    *CsvWriter
}

func main() {
	outputCSVFlag := flag.String("c", "whois.csv", "output results to csv file")
	outputDumpFlag := flag.Bool("d", false, "print raw whois response to stdout")
	maxWorkersFlag := flag.Int("w", 8, "max workers")

	flag.Parse()

	out := &Output{}
	switch {
	case *outputDumpFlag:
		out.option = "stdout"
	default:
		out.option = "csv"
		outputFile, err := NewCsvWriter(*outputCSVFlag)
		if err != nil {
			log.Fatalf("error creating csv file: %s\n", err)
		}
		out.csv = outputFile
		// headers
		rec := &Record{}
		outputFile.Write(rec.GetHeaders())
	}

	ch := make(chan string)
	go func() {
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()

	var wg sync.WaitGroup
	for i := 0; i < *maxWorkersFlag; i++ {
		wg.Add(1)
		go whoisLookup(ch, &wg, out)
	}
	wg.Wait()

	if !*outputDumpFlag {
		out.csv.Flush()
		fmt.Printf("file saved to: %s\n", *outputCSVFlag)
	}
}

func whoisLookup(ch chan string, wg *sync.WaitGroup, out *Output) {
	defer wg.Done()
	for d := range ch {
		req, err := whois.NewRequest(d)
		if err != nil {
			continue
		}
		res, err := whois.DefaultClient.Fetch(req)
		if err == nil {
			switch out.option {
			case "stdout":
				fmt.Println(res.String())
				fmt.Println("------------------")
			case "csv":
				record := ParseResponse(d, res.String())
				out.csv.Write(record)

			}
		}
	}
}
