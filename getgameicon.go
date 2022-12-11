package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GetGameIconResponseOneStruct struct {
	TargetId int    `json:"targetId"`
	State    string `json:"state"`
	ImageUrl string `json:"imageUrl"`
}

type GetGameIconResponseStruct struct {
	Data []GetGameIconResponseOneStruct `json:"data"`
}

func GetGameIcon(GameId int) (bool, string) {
	client := &http.Client{}
	request, err := http.NewRequest("GET", fmt.Sprintf("https://thumbnails.roblox.com/v1/games/icons?universeIds=%d&returnPolicy=0&size=512x512&format=Png&isCircular=false", GameId), bytes.NewBuffer([]byte("")))

	if err != nil {
		return false, err.Error()
	}

	response, err := client.Do(request)

	if err != nil {
		return false, err.Error()
	}

	if response.StatusCode >= 400 {
		return false, strconv.Itoa(response.StatusCode)
	}

	var Body GetGameIconResponseStruct
	if err := json.NewDecoder(response.Body).Decode(&Body); err != nil {
		return false, err.Error()
	}
	defer response.Body.Close()

	return true, Body.Data[0].ImageUrl
}
