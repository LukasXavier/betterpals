package schedule

import (
	"encoding/json"
	"log"
	"net/http"
)

type Schedule struct {

}

var self = make(map[string]*Schedule)
func New(id string) (*Schedule, error) {
    if self[id] == nil {
        schedule, err := getSchedule(id)
        if err != nil {
            return nil, err
        }
        self[id] = schedule
    }
    return self[id], nil
}

func getSchedule(id string) (*Schedule, error) {
    url := "https://www.leaguepals.com/laneSchedule?simple=true&league_id=" + id
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
    return nil, nil
}
