# Single Responsibility Principle(SRP)

This states that a class or type should have one primary responsibility and as a result one reason to change. That reason being related to its primary responsibility.

## Analogy

We're writing a simple application where we're making a journal to record thoughts. This journal has entries represented by:

```
var entryCount = 0

type Journal struct {
	entries []string
}
```

Lets imagine ways of adding and removing entries:

```
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
```

We can easily break the Single Responsibility Principle by adding functions which deal with another concern.

There's another term in addition to the S.R.P, **Seperation of Concern**. Seperation of Concern basically means different concerns or problems that the system solves has to reside in different constructs.

Whether attached to different structures or packages, is upto us, but they have to be split up so we cannot just take everything and put it in the same package.

That is an anti-pattern which is referred to as **God Object** which is basically when we take everything and put it into a single package for example.

Lets assume we also want to add persistance to our journal application to save the journal entries to a file.

```
func (j *Journal) String() string {
    return strings.Join(j.entries, "\n")
}

// Save journal entries
func (j *Journal) Save (filename string) {
    _ = ioutil.WriteFile(filename, []byte(j.String()), 0644)
}
```

We can also implement methods for loading from a file:

```
func (j *Journal) Load (filename string) {
    //...
}
```

We can similarly expand the idea and decide to also load from the web:

```
func (j *Journal) LoadFromWeb (url *url.URL) {
    //...
}
```

The above breaks the **S.R.P**. This is because the responsibility of the Journal deals with the management of entries and not to handle persistence. Persistence can be handled by a seperate component whether its a seperate package or struct that has some methods related to persistence.

```
type Persistence struct {
    lineSeperator string
}

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
```

## Recap

The whole premise of the S.R.P. is the class or type has a single primary responsibility. In this case, the Journal has the responsibility of storing entries and allowing the manipulation of entries. 

When it comes to other concerns like persistence, we adopt the idea of Seperation of Concerns. We take those concerns and put them somewhere else because those concerns can be cross-cutting concerns. They can influence not only how Journals are saved but also the way some other types of structures are saved.
