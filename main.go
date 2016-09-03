package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

//RequestBody represent body request coming from api.ai as trigered from intent
type RequestBody struct {
	ID          string        `json:"id"`
	TimeStamp   string        `json:"timestamp"`
	Result      Result        `json:"result"`
	Contexts    []interface{} `json:"contexts"`
	Metadata    interface{}   `json:"metadata"`
	Fulfillment interface{}   `json:"fulfillment"`
}

//Result represent result of api.ai processing on input
type Result struct {
	Source           string     `json:"source"`
	ResolvedQuery    string     `json:"resolvedQuery"`
	Action           string     `json:"action"`
	ActionIncomplete bool       `json:"actionIncomplete"`
	Parameters       Parameters `json:"parameters"`
}

//Parameters represent animal parameter
type Parameters struct {
	Animal string `json:"animal"`
}

//Response represent response body in JSON formated as replied for webhook call from api.ai
type Response struct {
	Speech      string        `json:"speech"`
	DisplayText string        `json:"displayText"`
	Data        string        `json:"data"`
	ContextOut  []interface{} `json:"contextOut"`
	Source      string        `json:"source"`
}

func main() {
	//gin.SetMode(gin.ReleaseMode) //Use This mode in release or production server
	r := gin.Default() //Use this mode in debug mode or in development server

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/webhook", func(c *gin.Context) {
		decoder := json.NewDecoder(c.Request.Body)
		var t RequestBody
		err := decoder.Decode(&t)
		if err != nil {
			fmt.Println("error decode to json")
		}
		desc := getAnimalDetails(t.Result.Parameters.Animal)
		res := Response{
			Speech:      t.Result.Parameters.Animal + ", " + desc,
			DisplayText: t.Result.Parameters.Animal + ", " + desc,
			Source:      "static-value-copied-from-internet",
		}
		c.JSON(200, res)
	})

	r.GET("/animal-details/:name", func(c *gin.Context) {
	})

	r.Run(":9292") // listen and server on 0.0.0.0:8080
}

func getAnimalDetails(name string) string {
	if strings.EqualFold(name, "lion") {
		return "Lions have strong, compact bodies and powerful forelegs, teeth, and jaws for pulling down and killing prey. Their coats are yellow-gold. Adult males have shaggy manes that range in color from blond to reddish-brown to black, and length."
	} else if strings.EqualFold(name, "tiger") {
		return "A tiger's fur color varies from orange-red to tawny yellow, with a lot of black stripes that have different lengths and widths."
	} else if strings.EqualFold(name, "cat") {
		return "a male cat is called a tom, a female cat is called a molly or queen while young cats are called kittens."
	} else if strings.EqualFold(name, "bird") {
		return "Birds (Aves) are a group of endothermic vertebrates, characterised by feathers, toothless ..... yellow-headed caracara and galah, have spread naturally far beyond their original ranges as agricultural practices created suitable new habitat."
	} else if strings.EqualFold(name, "elephant") {
		return "Elephants are the largest land animals on Earth. They have characteristic long noses, or trunks; large, floppy ears; and wide, thick legs."
	}

	return "We have no idea about the animal name."
}
