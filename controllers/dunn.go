package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/jessemillar/dunn/models"
	"github.com/labstack/echo"
)

func ArchiveCard(c echo.Context) error {
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

		name := c.QueryParam("name")

		// Find the card that matches the passed name
		for _, c := range allCards {
			if c.Name == name {
				// Mark card as done
				client := &http.Client{}
				request, err := http.NewRequest("PUT", "https://api.trello.com/1/cards/"+c.ID+"?closed=true&key="+trelloKey+"&token="+trelloToken, nil)
				_, err = client.Do(request)
				if err != nil {
					log.Print(err)
				}
			}
		}
	}

	// TODO Do I need/want to send a response back?
	return c.JSON(200, nil)
}
