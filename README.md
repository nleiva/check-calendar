# Check Calendar

Simple app to versify if a timeslot is busy or not.

## Requirements

```bash
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/google
```

## Run it

First time you run it

```bash
$ go run main.go
Go to the following link in your browser then type the authorization code: 
<https://accounts.google.com/o/oauth2/auth?access_type=offline&client_id=...>
<PASTE CODE>
```

## Go Time

Temporary [Example](https://play.golang.org/p/cmb-E0Njf8p)

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	today := time.Now()
	plusOne := today.Add(1 * time.Minute)

	// Using time.Before() method
	g1 := today.Before(plusOne)
	fmt.Println("today before plusOne:", g1)

	// Using time.After() method
	g2 := plusOne.After(today)
	fmt.Println("plusOne after today:", g2)

	t1, err := time.Parse(time.RFC3339, "2020-08-31T18:44:52Z")
	if err != nil {
		fmt.Println("Oops 1")
	}
	t2:= t1.Add(1 * time.Minute)
	
	t2b, err := t2.MarshalText()
	
	if err != nil {
		fmt.Println("Oops 2")
	}
	
	fmt.Printf("%v", string(t2b))

}
```