# Get a valid Google Calendar ics file from [https://www.campus-booster.net](https://www.campus-booster.net)

This script written in Go give you a valid .ics file usable into [Google Calendar](https://calendar.google.com)

## Requirements
* Campus Booster ID
* Password
* https://campus-api.supinfo.com/ API key

## Run as Go binary
```bash
$ go get -u github.com/SkYNewZ/supinfo-calendar
# Ensure that $GOPATH is set, or if $HOME/go/bin is into your $PATH
$ supinfo-calendar -u CAMPUS_BOOSTER_ID -p CAMPUS_BOOSTER_PASSWORD -k SUPINFO_API_KEY # Or use environment variables

```

> Optional: add `-o /tmp` to generate ics file in `/tmp`

## Run as docker container
```bash
$ docker run -it --rm -v $PWD:/out skynewz/supinfo-calendar -u CAMPUS_BOOSTER_ID -p CAMPUS_BOOSTER_PASSWORD -k SUPINFO_API_KEY
$ #OR
$ docker run -it --rm \
    -v $PWD:/out \
    -e CAMPUS_ID=<campusId> \
    -e CAMPUS_PASSWORD=<password> \
    -e SUPINFO_API_KEY=<key> \
    skynewz/supinfo-calendar

```

> .ics file location will depend on the defined volume

## Usage
```
Usage:
  supinfo-calendar [OPTIONS]

Application Options:
  -u, --campus-id= Your CampusBooster id [$CAMPUS_ID]
  -p, --password=  Your CampusBooster password [$CAMPUS_PASSWORD]
  -k, --key=       http://campus-api.supinfo.com API KEY [$SUPINFO_API_KEY]
  -o, --output-path= .ics downloaded file location (default: .) [$OUTPUT_PATH]

Help Options:
  -h, --help       Show this help message

```