package team

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type TeamMember struct {
    Id string;
    Name string;
    Average int;
    Games []int;
}

type Team struct {
    Members []TeamMember;
    Id string;
    Name string;
}

var self = make(map[string]*Team)
func New(id string) (*Team, error) {
    if self[id] == nil {
        team, err := getIndividualTeam(id)
        if err != nil {
            return nil, err
        }
        self[id] = team
    }
    return self[id], nil
}

func getIndividualTeam(id string) (t *Team, e error) {
    url := "https://www.leaguepals.com/loadIndividualTeam?noPre=false&id=" + id
    log.Print(url)
    resp, e := http.Get(url)
    if e != nil {
        return nil, e
    }
    defer resp.Body.Close()
    var ingest response
    e = json.NewDecoder(resp.Body).Decode(&ingest)
    if e != nil {
        return nil, e
    }
    if len(ingest.Data) > 0 {
        members := make([]TeamMember, 0, 4)
        for _, data := range ingest.Data {
            members = append(members, TeamMember{
                Id: data.Id,
                Name: data.Name,
                Average: data.Average,
                Games: data.Games,
            })
        }
        t = &Team{
            Name: ingest.Data[0].TeamName,
            Id: ingest.Data[0].Team,
            Members: members,
        }
    } else {
        return nil, errors.New("Server returned empty data")
    }
    return
}
