package schedule

import (
	"encoding/json"
	"log"
	"net/http"
)

type Schedule struct {

}

var self = make(map[string]*response)
func New(id string) (*response, error) {
    if self[id] == nil {
        schedule, err := getSchedule(id)
        if err != nil {
            return nil, err
        }
        self[id] = schedule
    }
    return self[id], nil
}

func getSchedule(id string) (*response, error) {
    url := "https://www.leaguepals.com/laneSchedule?simple=false&league_id=" + id
    resp, err := http.Get(url)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    defer resp.Body.Close()
    var ingest response
    err = json.NewDecoder(resp.Body).Decode(&ingest)
    if err != nil {
        log.Print(err)
        return nil, err
    }
    return &ingest, nil
}
