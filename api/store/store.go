package store

import (
	"log"
	"sync"

	"github.com/LukasXavier/betterpals/api/schedule"
	"github.com/LukasXavier/betterpals/api/team"
)

type Store struct {
    Mutex sync.Mutex
    Schedules map[string]*schedule.Schedule
    Teams map[string]*team.Team
}

func New() *Store {
    return &Store {
        Schedules: make(map[string]*schedule.Schedule),
        Teams: make(map[string]*team.Team),
    }
}

func (s *Store) FetchSchedule(id string) *schedule.Schedule {
    s.Mutex.Lock()
    leagueSchedule := s.Schedules[id]
    s.Mutex.Unlock()

    if leagueSchedule == nil {
        leagueSchedule, err := schedule.New(id)
        if err != nil {
            log.Print("store::FetchSchedule() - it broke", err)
            return nil
        }
        s.Mutex.Lock()
        s.Schedules[id] = leagueSchedule
        s.Mutex.Unlock()
        return leagueSchedule
    }
    return leagueSchedule
}

func (s *Store) FetchTeam(id string) *team.Team {
    s.Mutex.Lock()
    leagueTeam := s.Teams[id]
    s.Mutex.Unlock()

    if leagueTeam == nil {
        leagueTeam, err := team.New(id)
        if err != nil {
            log.Print("store::FetchTeam() - it broke", err)
            return nil
        }
        s.Mutex.Lock()
        s.Teams[id] = leagueTeam
        s.Mutex.Unlock()
        return leagueTeam
    }
    return leagueTeam
}

func (s *Store) Sync(id string) {
    schedule := s.FetchSchedule(id)
    for _, match := range schedule.Weeks[0].Matches {
        go s.FetchTeam(match.Team1Id)
        go s.FetchTeam(match.Team2Id)
    }
}
