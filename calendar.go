/**
 * Most of this file is borrowed from: https://developers.google.com/calendar/quickstart/go
 *
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// Determines the max time.
func maxTime(min string) (string, error) {
	t1, err := time.Parse(time.RFC3339, min)
	if err != nil {
		return "", fmt.Errorf("Unable to parse the time: %v", err)
	}
	// Change to minutes
	t2 := t1.Add(1 * time.Hour)
	t2b, err := t2.MarshalText()
	if err != nil {
		return "", fmt.Errorf("Unable to add the delta: %v", err)
	}
	return string(t2b), nil
}

func isItBusy(t string) (bool, error) {
	// INPUTS
	// Calendar name. Fixed for now.
	name := "primary"
	// Current execution time. Ex: "2020-08-31T18:44:00Z"
	min := t
	// Time + delta to determine the timeslot to verify.
	max, err := maxTime(min)
	if err != nil {
		return false, fmt.Errorf("Unable to determine end of time range: %v", err)
	}

	b, err := ioutil.ReadFile("desktop.json")
	if err != nil {
		return false, fmt.Errorf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		return false, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.New(client)
	if err != nil {
		return false, fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}

	cal := calendar.FreeBusyRequestItem{Id: name}

	freebusyRequest := calendar.FreeBusyRequest{
		TimeMin: min,
		TimeMax: max,
		Items:   []*calendar.FreeBusyRequestItem{&cal},
	}

	freebusyRequestCall := srv.Freebusy.Query(&freebusyRequest)

	freebusyRequestResponse, err := freebusyRequestCall.Do()

	if err != nil {
		return false, fmt.Errorf("Unable to get a freebusyRequestResponse: %v", err)
	}

	b, err = freebusyRequestResponse.MarshalJSON()

	if len(freebusyRequestResponse.Calendars[name].Busy) == 0 {
		// fmt.Println("No busy timeslots found.")
		return false, nil
	}
	return true, nil
	// for _, item := range freebusyRequestResponse.Calendars[name].Busy {
	// 	start := item.Start
	// 	end := item.End
	// 	fmt.Printf("From: %v To: %v\n", start, end)
	// }
}