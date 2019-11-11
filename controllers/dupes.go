package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jessemillar/butler-lewis/models"
	"github.com/labstack/echo"
)

func KillDupes(c echo.Context) error {
	secret := os.Getenv("BUTLER_LEWIS_SECRET")

	// Tiny "security" thing
	secretQuery := c.QueryParam("secret")
	if secret != secretQuery {
		return c.JSON(401, nil)
	}

	trelloKey := os.Getenv("TRELLO_KEY")
	trelloToken := os.Getenv("TRELLO_TOKEN")
	trelloList := "5b1bf0d6b20211b1d693e4a6"

	response, err := http.Get("https://api.trello.com/1/lists/" + trelloList + "/cards?key=" + trelloKey + "&token=" + trelloToken)
	if err != nil {
		log.Print(err)
	} else {
		defer response.Body.Close()

		// TODO Does reading close the body?
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
		}

		var allCards []models.Card
		json.Unmarshal(body, &allCards)

		var uniqueCards = removeDuplicateCards(allCards)
		var duplicateCards []models.Card

		for _, c := range uniqueCards {
			duplicateCards = append(duplicateCards, getDuplicateCards(c, allCards)...)
		}

		// Archive all duplicate cards
		for _, c := range duplicateCards {
			// Mark card as done
			client := &http.Client{}
			request, err := http.NewRequest("PUT", "https://api.trello.com/1/cards/"+c.ID+"?closed=true&key="+trelloKey+"&token="+trelloToken, nil)
			_, err = client.Do(request)
			if err != nil {
				log.Print(err)
			}
		}
	}

	// TODO Do I need/want to send a response back?
	return c.JSON(200, nil)
}

func removeDuplicateCards(allCards []models.Card) []models.Card {
	uniqueCardNames := make(map[string]bool)
	uniqueCards := []models.Card{}

	for _, c := range allCards {
		if _, value := uniqueCardNames[c.Name]; !value {
			uniqueCardNames[c.Name] = true
			uniqueCards = append(uniqueCards, c)
		}
	}

	return uniqueCards
}

func getDuplicateCards(card models.Card, allCards []models.Card) []models.Card {
	var duplicateCards []models.Card

	for _, c := range allCards {
		if strings.TrimSpace(c.Name) == strings.TrimSpace(card.Name) && c.ID != card.ID {
			// Add the dupe to the slice that's returned
			duplicateCards = append(duplicateCards, c)
		}
	}

	return duplicateCards
}
