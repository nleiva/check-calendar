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