package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Haydz6/rich-go/client"
	"github.com/gin-gonic/gin"
)

var LastGameId = 0
var LastGameIcon = ""
var LastPresenceUpdate = time.Now().Unix()
var GameElapsed = time.Now()
var LoggedIn = false

type SetPresenceBodyStruct struct {
	UserId   int
	PlaceId  int
	GameId   int
	State    string
	Details  string
	GameName string
}

func Login() {
	if LoggedIn {
		return
	}

	client.Login("1050943341212217344")
	LoggedIn = true
}

func SetPresence(c *gin.Context) {
	var Body SetPresenceBodyStruct

	if err := c.BindJSON(&Body); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	LastPresenceUpdate = time.Now().Unix()

	if LastGameId != Body.GameId {
		GameElapsed = time.Now()
		Success, NewGameIcon := GetGameIcon(Body.GameId)

		if Success {
			LastGameIcon = NewGameIcon
			LastGameId = Body.GameId
		}
	}

	Login()

	err := client.SetActivity(client.Activity{
		State:      Body.State,
		Details:    Body.Details,
		LargeImage: LastGameIcon,
		LargeText:  Body.GameName,
		SmallImage: "https://cdn.discordapp.com/app-icons/1050943341212217344/96e49f383082fb14b482911d50ff6261.png?size=512",
		SmallText:  "Roblox Studio",
		Timestamps: &client.Timestamps{
			Start: &GameElapsed,
		},
		Buttons: []*client.Button{
			&client.Button{Label: "View Game", Url: fmt.Sprintf("https://roblox.com/games/%d", Body.PlaceId)},
			&client.Button{Label: "Join Game", Url: fmt.Sprintf("roblox://placeId=%d&launchData=", Body.PlaceId)},
		},
	})

	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	c.String(http.StatusOK, "")
}

func KillActivity() {
	err := client.SetActivity(client.Activity{
		State: "end",
	})

	if err != nil {
		println(err)
	}
}

func TimeoutCheck() {
	for range time.Tick(time.Second) {
		println(time.Now().Unix() - LastPresenceUpdate)
		if time.Now().Unix()-LastPresenceUpdate >= 3 {
			if LoggedIn {
				LoggedIn = false
				KillActivity()
				client.Logout()
			}
		}
	}
}
