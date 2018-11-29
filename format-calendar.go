package main

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const ICSFileName = "calendar.ics"

func write(calendar string, path string) {
	path = path + "/" + ICSFileName
	d1 := []byte(calendar)
	err := ioutil.WriteFile(path, d1, 0644)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed writing in file"))
	}

	log.Printf("File location : %s", path)
}

func createCalendar(old *string, path string) {
	array := formatCalendar(*old)
	for i, line := range array {

		// fix UID
		if strings.Contains(line, "UID") {
			array[i] = "UID:" + uuid.New().String()
		}

		// fix timezone on date
		if strings.Contains(line, "DTSTART") || strings.Contains(line, "DTEND") {
			if !strings.Contains(line, "Z") {
				array[i] = array[i] + "Z"
			}
		}
		
		// fix name
		if strings.Contains(line, "X-WR-CALNAME") {
			array[i] = "X-WR-CALNAME:Supinfo - M.Sc.1"
		}
	}

	// append element to copy
	array = append(array, "0", "0", "0")
	copy(array[9+3:], array[9:]) //delay for 2 index
	array[9] = "X-WR-TIMEZONE:Europe/Paris"
	array[10] = "X-PUBLISHED-TTL:PT12H"
	array[11] = "TZID:Europe/Paris"

	write(strings.Join(array, "\r\n"), path)
}

func formatCalendar(calendar string) []string {
	c := strings.Split(calendar, "\r\n")
	log.Printf("%s lines in file", strconv.Itoa(len(c)))
	return c
}
