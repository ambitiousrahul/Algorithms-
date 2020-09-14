package main

import (
	polynomialhash "PolynomialHash"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fptr := flag.String("fpath", "/Assignment/mobydick.txt", "file path to read from")
	flag.Parse()
	data, err := ioutil.ReadFile(*fptr)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	words := strings.Fields(string(data))
	totalWordsCount := len(words)
	hashTableWords := make([]polynomialhash.WordsFrequency, totalWordsCount)
	for i := 0; i < totalWordsCount; i++ {
		if words[i] != "" {
			hashcode := polynomialhash.GetPolynomialHashCode(words[i], totalWordsCount)

			fmt.Println("Line number", i, words[i], hashcode)
			hashTableWords[hashcode].Word = words[i]
			hashTableWords[hashcode].Frequency++
		}
	}
	maxheap := polynomialhash.NewMaxheapWithElements(polynomialhash.TableSize, hashTableWords[:])
	maxheap.BuildMaxHeap()
	fmt.Println("The most frequent used words are")
	for loopingIndex := 0; loopingIndex < 20; loopingIndex++ {
		wordandFrequency := maxheap.ExtractMaximum()
		fmt.Println(wordandFrequency.Word, wordandFrequency.Frequency)
	}
}
