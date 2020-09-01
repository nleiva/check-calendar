# Check Calendar

Simple app to verify if a timeslot is busy or not.

## Requirements

```bash
go get -u google.golang.org/api/calendar/v3
go get -u golang.org/x/oauth2/google
```

### 

You need to [Turn on the Google Calendar API](https://developers.google.com/calendar/quickstart/go#step_1_turn_on_the) and download your OAuth 2.0 client credentials (json file) from [API Credentials](https://console.developers.google.com/apis/credentials) as in the image below.

<p align="center">
  <img title="API Console" src="static/API.JPG"><br>
  <br>
</p>

## Test 

Execute `go run *.go args.json`, where `args.json` is something like:


```json
{
    "Name": "Test",
    "Time":  "2020-09-01T19:50:00Z"
}
```