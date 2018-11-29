package main

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"strconv"
)

func login(campusId int, password string, apiKey string) *SupinfoStudent {
	url := fmt.Sprintf("https://campus-api.supinfo.com/AppScho/Login?language=fr&key=%s", apiKey)
	url += fmt.Sprintf("&boosterid=%s&password=%s", strconv.Itoa(campusId), password)
	log.Println("Reaching URL : " + url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `SUPINFO_INT_PLANNINGS_MARKS_Test/1.0 CFNetwork/976 Darwin/18.2.0"`)
	response, err := client.Do(req)

	if err != nil || response == nil {
		log.Fatalln(errors.Wrap(err, "Failed"))
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Fatalln(errors.Errorf("Invalid response status code %d", response.StatusCode))
	}

	decoder := json.NewDecoder(response.Body)
	var loginResponse LoginResponse
	err = decoder.Decode(&loginResponse)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed decoding json response"))
	}

	if loginResponse.Token == "" {
		response, _ := json.Marshal(loginResponse)
		log.Printf("MethodCallMessage : %s", loginResponse.MethodCallMessage)
		log.Fatalf("Unsuccessfull login, response : %s", response)
	}

	return &SupinfoStudent{
		Token:         loginResponse.Token,
		BoosterId:     loginResponse.BoosterId,
		CampusClassId: loginResponse.CampusClassId,
		FirstName:     loginResponse.FirstName,
		LastName:      loginResponse.LastName,
	}
}

func getPlaning(student SupinfoStudent, apiKey string) string {
	url := fmt.Sprintf("https://campus-api.supinfo.com/AppScho/planning?language=fr&key=%s", apiKey)
	url += fmt.Sprintf("&boosterid=%s&token=%s&campusclassid=%s", strconv.Itoa(student.BoosterId), student.Token, student.CampusClassId)
	log.Println("Reaching URL : " + url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", `SUPINFO_INT_PLANNINGS_MARKS_Test/1.0 CFNetwork/976 Darwin/18.2.0"`)
	response, err := client.Do(req)

	if err != nil || response == nil {
		log.Fatalln(errors.Wrap(err, "Failed"))
	}
	defer response.Body.Close()

	decoder := json.NewDecoder(response.Body)
	var planningResponse PlanningResponse
	err = decoder.Decode(&planningResponse)

	if err != nil {
		log.Fatalln(errors.Wrap(err, "Failed decode json response"))
	}

	if planningResponse.IcsPlanning == "" {
		response, _ := json.Marshal(planningResponse)
		log.Printf("MethodCallMessage : %s", planningResponse.MethodCallMessage)
		log.Fatalf("Fail, response : %s", response)
	}

	return planningResponse.IcsPlanning
}
