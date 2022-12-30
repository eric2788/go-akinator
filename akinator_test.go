package akinator

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestAkinator(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		log.Fatal(err)
	}

	for r := range c.Next() {
		if r.Status != StatusOk {
			log.Fatal("Bad Status:", r.Status)
		}

		if r.Guessed {
			// Akinator made a guess.
			// ...
			fmt.Println("I guess", r.CharacterName, "is your character.")
			continue
		}

		// Akinator asked a question.
		// ...
		// For the next response to be called,
		// you must answer the akinator's question with r.AnswerYes(), r.AnswerNo(), etc.

		fmt.Println(r.Question)
		ans := rand.Intn(5)
		ansName := []string{"Yes", "No", "Don't Know", "Probably", "Probably Not"}
		fmt.Println("our answers: ", ansName[ans])
		r.answer(ans)
	}
}
