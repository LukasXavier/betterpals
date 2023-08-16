package schedule

import (
	"encoding/json"
	"log"
	"net/http"
)

type Schedule struct {
    LeagueId string
    Weeks []Week
}

type Week struct {
    Date string
    Matches []Match
}

type Match struct {
    Team1Id string
    Team2Id string
    LaneNum int
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
    weeks := make([]Week, len(ingest.Schedule))
    for i, weekInfo := range ingest.Schedule {
        matches := make([]Match, len(weekInfo.Matches))
        for j, matchInfo := range weekInfo.Matches {
            matches[j] = Match{
                Team1Id: matchInfo.Team1Id,
                Team2Id: matchInfo.Team2Id,
                LaneNum: matchInfo.Team1Lane,
            }
        }
        weeks[i] = Week {
            Date: weekInfo.Date,
            Matches: matches,
        }
    }
    result := &Schedule {
        LeagueId: ingest.LeagueId,
        Weeks: weeks,
    }
    return result, nil
}
