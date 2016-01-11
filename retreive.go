package main

import (
	"errors"
	"github.com/connorjacobsen/blockscore-go"
)

// In order to execute the Unit Test, you must set your BlockScore
// API key as environment variable: BLOCKSCORE_API_KEY=xxxx
// Copy pasted from the test framework.
func Init() {
	if err := blockscore.SetKeyEnv(); err != nil {
		panic(err)
	}
}

type Customer struct {
	PersonId      string
	QuestionSetId string
}

func NewCustomer() (cus *Customer) {
	Init()
	cus = &Customer{
		PersonId:      "",
		QuestionSetId: "",
	}
	return
}

//This is to verify the creation of the api.
// string - id of customer
// err - custom error while creating
func (cus *Customer) VerifyCreate(params *blockscore.PersonParams) (string, error) {
	resp, err := blockscore.People.Create(params)
	if err == nil {
		cus.PersonId = resp.Id
		questionId, err := cus.GetQuestions(resp.Id)
		if err == nil {
			cus.QuestionSetId = questionId
		} else {
			err = errors.New("Error creating questions")
		}
	} else {
		err = errors.New("Error creating customer")
	}
	return resp.Id, err
}

//Is a utility method. Added initially so kept it.
func (cus *Customer) RetrieveId(ID string) (string, error) {
	resp, err := blockscore.People.Retrieve(ID)
	return resp.Status, err
}

// Returns the set of questions the user is concerned about.
func (cus *Customer) GetQuestions(personID string) (string, error) {
	resp, err := blockscore.QuestionSets.Create(personID)
	// Setting for answering to get it.
	cus.QuestionSetId = resp.Id
	return resp.Id, err
}

// Hit the api to get the score using the predefined set of Questions.
func (cus *Customer) GetQuestionScore(answers []int) (float64, error) {
	scoreParams := blockscore.ScoreParams{
		Answers: []blockscore.ScoreAnswer{},
	}

	for index, answer := range answers {

		scoreAnswerObj := blockscore.ScoreAnswer{QuestionId: index + 1, AnswerId: answer}
		scoreParams.Answers = append(scoreParams.Answers, scoreAnswerObj)
	}
	resp, err := blockscore.QuestionSets.Score(cus.QuestionSetId, &scoreParams)
	return resp.Score, err
}
