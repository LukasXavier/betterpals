package team

import (
	"encoding/json"
	"errors"
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
        return team, err
    }
    return self[id], nil
}

func getIndividualTeam(id string) (t *Team, e error) {
    url := "https://www.leaguepals.com/loadIndividualTeam?noPre=false&id=" + id
    resp, e := http.Get(url)
    if e != nil {
        return nil, e
    }
    defer resp.Body.Close()
    var ingest struct {
        Data []struct {
            Id string `json:"_id"`;
            Name string `json:"name"`;
            // Email string `json:"email"`;
            // FirstName string `json:"firstName"`;
            // LastName string `json:"lastName"`;
            // Dexterity int `json:"dexterity"`;
            Average int `json:"average"`;
            // Averages []struct {
                // League string `json:"league"`;
                // Team string `json:"team"`;
                // Id string `json:"_id"`;
                // IndividualPoints int `json:"individualPoints"`;
                // TotalPointsCarryOn int `json:"totalPointsCarryOn"`;
                // GamesCarryOn int `json:"gamesCarryOn"`;
                // HighSeriesHdcp int `json:"highSeriesHdcp"`;
                // HighGameHdcp int `json:"highGameHdcp"`;
                // HighSeries int `json:"highSeries"`;
                // HighGame int `json:"highGame"`;
                // FixedAvg *sturct {} `json:"fixedAvg"`;
                // Average int `json:"average"`;
            // } `json:"averages"`;
            // Subs []struct {} `json:"subs"`;
            // NickNames []string `json:"nickNames"`;
            // IsJunior bool `json:"isJunior"`;
            // IsFemale bool `json:"isFemale"`;
            // Avatar string `json:"avatar"`;
            Games []int `json:"games"`;
            // IndPointsWon int `json:"indPointsWon"`;
            HighGame int `json:"highGame"`;
            HighGameHdcp int `json:"highGameHdcp"`;
            HighSeries int `json:"highSeries"`;
            HighSeriesHdcp int `json:"highSeriesHdcp"`;
            RealAvg int `json:"realAvg"`;
            EnteringAvg int `json:"enteringAvg"`;
            // GamesPlayed int `json:"gamesPlayed"`;
            TeamName string `json:"teamName"`;
            Team string `json:"team"`;
            League string `json:"league"`;
            LeagueName string `json:"leagueName"`;
            Center string `json:"center"`;
            CenterName string `json:"centerName"`;
            // TotalPointsScored int `json:"totalPointsScored"`;
        } `json:"data"`;
    }
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
