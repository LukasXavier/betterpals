package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Team struct {
    Data []struct {
        Id string `json:"_id"`;
        Name string `json:"name"`;
        Average int `json:"average"`;
        Games []int `json:"games"`;
        TeamName string `json:"teamName"`;
        TeamId string `json:"team"`;
    } `json:"data"`
    Id string;
    Team string;
}

func getIndividualTeam(id string) (*Team, error) {
    // https://www.leaguepals.com/loadIndividualTeam?id=64775970bba9d14862bcf9ce&noPre=false
    url := "https://www.leaguepals.com/loadIndividualTeam?id=" + id + "&noPre=false"
    log.Print(url)
    if res, err := http.Get(url); err != nil {
        return nil, err
    } else {
        var t Team
        defer res.Body.Close()
        if err := json.NewDecoder(res.Body).Decode(&t); err != nil {
            return nil, err
        } else {
            return &t, nil
        }
    }
}
