package main

import (
	"fmt"
	"github.com/connorjacobsen/blockscore-go"
	"testing"
)

var TestPersonId string
var QuestionSetId string

func TestCreate(t *testing.T) {
	var peopleParams = blockscore.PersonParams{
		NameFirst:          "John",
		NameMiddle:         "P",
		NameLast:           "Denver",
		DocumentType:       "ssn",
		DocumentValue:      "0000",
		BirthDay:           7,
		BirthMonth:         6,
		BirthYear:          1980,
		AddressStreet1:     "1234 Main Street",
		AddressStreet2:     "APT 12",
		AddressCity:        "Palo Alto",
		AddressSubdivision: "California",
		AddressPostalCode:  "94025",
		AddressCountryCode: "US",
		PhoneNumber:        "123-456-78910",
		IPAddress:          "127.0.0.1",
		Note:               "Hello, world",
	}
	cus := NewCustomer()
	personId, err := cus.VerifyCreate(&peopleParams)
	if err != nil {
		fmt.Println(err)
	} else {
		TestPersonId = personId
		fmt.Println(personId)
	}
}

func TestRetreive(t *testing.T) {

	personId := "5694005831313700030006d1"
	cus := &Customer{}
	resp, err := cus.RetrieveId(personId)
	if err != nil {
		t.Error("Retrieving the person")
	} else {
		fmt.Println(resp)
	}
}

func TestCreateQuestions(t *testing.T) {
	fmt.Println("Testing creation of questions with person id", TestPersonId)
	cus := NewCustomer()
	// Populating the object: Already tested so should work
	cus.RetrieveId(TestPersonId)
	resp, err := cus.GetQuestions(TestPersonId)
	if err == nil {
		fmt.Println("Question set id", QuestionSetId)
		QuestionSetId = resp
	} else {
		t.Error("Error Retrieving questions", err)
	}
}

func TestQuestionSetScore(t *testing.T) {
	cus := NewCustomer()
	// Populating the object: Already tested so should work
	cus.RetrieveId(TestPersonId)
	cus.QuestionSetId = QuestionSetId
	score, err := cus.GetQuestionScore([]int{1, 1, 1, 1, 1})
	if err == nil {
		fmt.Println(score)
	} else {
		t.Error("Error retreiving score", err)
	}
}
