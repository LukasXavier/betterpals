package store

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "sync"

    Json "github.com/LukasXavier/betterpals/api/json"
)

type Schedule struct {
    LegaugeId string
    Weeks     []Week
}

type Week struct {
    Date    string
    Matches []Match
}

type Match struct {
    Team1Id string
    Team2Id string
    LaneNum int
}

type Team struct {
    Members []TeamMember
    Id      string
    Name    string
}

type TeamMember struct {
    Id      string
    Name    string
    Average int
    Games   []int
}

type Store struct {
    Mutex     sync.Mutex
    Schedules map[string]*Schedule
    Teams     map[string]*Team
}

func New() *Store {
    return &Store{
    	Schedules: make(map[string]*Schedule),
    	Teams:     make(map[string]*Team),
    }
}

func (s *Store) Sync(id string) {
    schedule, ok := s.Schedules[id]
    if !ok {
    	var err error
    	schedule, err = fetchSchedule(id)
    	if err != nil {
    		log.Println("store::Sync() - schedule is nil", err)
    		return
    	}
    	s.Schedules[id] = schedule
    }
    if schedule == nil {
    	log.Println("store::Sync() - schedule was nil")
    	return
    }

    if len(schedule.Weeks) == 0 {
    	log.Println("store::Sync() - schedule has no weeks")
    	return
    }

    ch := make(chan *Team)
    var wg sync.WaitGroup

    for _, match := range schedule.Weeks[0].Matches {
    	wg.Add(2)
    	if _, ok := s.Teams[match.Team1Id]; !ok {
    		go fetchTeam(match.Team1Id, ch, &wg)
    	} else {
    		wg.Done()
    	}
    	if _, ok := s.Teams[match.Team2Id]; !ok {
    		go fetchTeam(match.Team2Id, ch, &wg)
    	} else {
    		wg.Done()
    	}
    }

    go func() {
    	wg.Wait()
    	close(ch)
    }()

    for result := range ch {
    	s.Teams[result.Id] = result
    }
    log.Println("finished sync")

}

func (s *Store) GetTeam(id string) (*Team, error) {
    if team, ok := s.Teams[id]; !ok {
        ch := make(chan *Team)
        var wg sync.WaitGroup
    	wg.Add(1)
    	go fetchTeam(id, ch, &wg)
    	go func() {
    		wg.Wait()
    		close(ch)
    	}()

    	for result := range ch {
    		s.Teams[result.Id] = result
    	}
    	return s.Teams[id], nil
    } else {
    	return team, nil
    }
}

func fetchSchedule(id string) (*Schedule, error) {
    url := fmt.Sprintf("https://www.leaguepals.com/laneSchedule?simple=false&league_id=%s", id)
    resp, err := http.Get(url)
    if err != nil {
    	return nil, err
    }

    defer resp.Body.Close()

    var ingest Json.Schedule
    if err := json.NewDecoder(resp.Body).Decode(&ingest); err != nil {
    	return nil, err
    }

    weeks := make([]Week, len(ingest.Schedule))
    for i, week := range ingest.Schedule {
    	matches := make([]Match, len(week.Matches))
    	for j, match := range week.Matches {
    		matches[j] = Match{
    			Team1Id: match.Team1Id,
    			Team2Id: match.Team2Id,
    			LaneNum: match.Team1Lane,
    		}
    	}
    	weeks[i] = Week{
    		Date:    week.Date,
    		Matches: matches,
    	}
    }
    schedule := &Schedule{
    	LegaugeId: ingest.LeagueId,
    	Weeks:     weeks,
    }
    return schedule, nil
}

func fetchTeam(id string, ch chan<- *Team, wg *sync.WaitGroup) {
    defer wg.Done()

    url := fmt.Sprintf("https://www.leaguepals.com/loadIndividualTeam?noPre=false&id=%s", id)
    resp, err := http.Get(url)
    if err != nil {
    	log.Println("store::fetchTeam() - http error: ", err)
    	return
    }

    defer resp.Body.Close()

    var ingest Json.Team
    if err := json.NewDecoder(resp.Body).Decode(&ingest); err != nil {
    	log.Println("store::fetchTeam() - json decode error: ", err)
    	return
    }

    if len(ingest.Data) > 0 {
    	members := make([]TeamMember, 0, 4)
    	for _, data := range ingest.Data {
    		members = append(members, TeamMember{
    			Id:      data.Id,
    			Name:    data.Name,
    			Average: data.Average,
    			Games:   data.Games,
    		})
    	}
    	ch <- &Team{
    		Name:    ingest.Data[0].TeamName,
    		Id:      ingest.Data[0].Team,
    		Members: members,
    	}
    } else {
    	log.Println("store::fetchTeam() - server returned empty data")
    	return
    }
}
