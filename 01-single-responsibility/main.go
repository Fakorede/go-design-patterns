package main

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"strings"
)

var entryCount = 0

// Journal struct
type Journal struct {
	entries []string
}

// AddEntry adds entry to the journal
func (j *Journal) AddEntry(text string) int {
	entryCount++
	entry := fmt.Sprintf("%d: %s", entryCount, text)
	j.entries = append(j.entries, entry)
	return entryCount
}

// RemoveEntry removes entry from the journal
func (j *Journal) RemoveEntry(index int) {
	// ...
}

func (j *Journal) String() string {
	return strings.Join(j.entries, "\n")
}

// Save journal entries
func (j *Journal) Save(filename string) {
	_ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}

// Load from file
func (j *Journal) Load(filename string) {
	//...
}

// LoadFromWeb load from web
func (j *Journal) LoadFromWeb(url *url.URL) {
	//...
}

// Persistence struct
type Persistence struct {
	lineSeperator string
}

// SaveToFile method
func (p *Persistence) SaveToFile(j *Journal, filename string) {
	_ = ioutil.WriteFile(filename, []byte(strings.Join(j.entries, p.lineSeperator)), 0644)
}

func main() {
	j := Journal{}
	j.AddEntry("I learnt the S.R.P. today!")
	j.AddEntry("Looking forward to tomorrow...")
	fmt.Println(j.String())

	p := Persistence{"\n"}
	p.SaveToFile(&j, "journal.txt")
}
