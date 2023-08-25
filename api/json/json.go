package json

type Team struct {
    Data []TeamData `json:"data"`
}

type TeamData struct {
    Id                string    `json:"_id"`
    Name              string    `json:"name"`
    Email             string    `json:"email"`
    FirstName         string    `json:"firstName"`
    LastName          string    `json:"lastName"`
    Dexterity         int       `json:"dexterity"`
    Average           int       `json:"average"`
    Averages          []Average `json:"averages"`
    Subs              []any     `json:"subs"`
    NickNames         []string  `json:"nickNames"`
    Junior            bool      `json:"isJunior"`
    Female            bool      `json:"isFemale"`
    Avatar            string    `json:"avatar"`
    Games             []int     `json:"games"`
    IndPointsWon      int       `json:"indPointsWon"`
    HighGame          int       `json:"highGame"`
    HighGameHdcp      int       `json:"highGameHdcp"`
    HighSeries        int       `json:"highSeries"`
    HighSeriesHdcp    int       `json:"highSeriesHdcp"`
    RealAvg           int       `json:"realAvg"`
    EnteringAvg       int       `json:"enteringAvg"`
    GamesPlayed       int       `json:"gamesPlayed"`
    TeamName          string    `json:"teamName"`
    Team              string    `json:"team"`
    League            string    `json:"league"`
    LeagueName        string    `json:"leagueName"`
    Center            string    `json:"center"`
    CenterName        string    `json:"centerName"`
    TotalPointsScored int       `json:"totalPointsScored"`
}

type Average struct {
    League             string `json:"league"`
    Team               string `json:"team"`
    Id                 string `json:"_id"`
    IndividualPoints   int    `json:"individualPoints"`
    TotalPointsCarryOn int    `json:"totalPointsCarryOn"`
    GamesCarryOn       int    `json:"gamesCarryOn"`
    HighSeriesHdcp     int    `json:"highSeriesHdcp"`
    HighGameHdcp       int    `json:"highGameHdcp"`
    HighSeries         int    `json:"highSeries"`
    HighGame           int    `json:"highGame"`
    FixedAvg           any    `json:"fixedAvg"`
    Average            int    `json:"average"`
}

type Schedule struct {
    Id            string         `json:"_id"`
    LeagueId      string         `json:"league_id"`
    Schedule      []ScheduleData `json:"schedule"`
    DisabledLanes []any          `json:"disabledLanes"`
    Shift         []any          `json:"shifts"`
    V             int            `json:"__v"`
}

type ScheduleData struct {
    Date                string   `json:"date"`
    Id                  string   `json:"_id"`
    CalculatedStartLane int      `json:"calculatedStartLane"`
    NormalWeek          int      `json:"normalWeek"`
    Matches             []Match  `json:"matches"`
    SplitMatches        []any    `json:"splitMatches"`
    Skipped             bool     `json:"isSkipped"`
    USBC                bool     `json:"isUSBC"`
    Locked              bool     `json:"isLocked"`
    DiffMultiShift      bool     `json:"isDiffMultiShift"`
    SwitchPairLane      bool     `json:"isSwitchPairLane"`
    OneTeamPairLane     bool     `json:"isOneTeamPairLane"`
    LockPositionRound   bool     `json:"isLockPositionRound"`
    DivPositionRound    bool     `json:"isDivPositionRound"`
    PositionRound       bool     `json:"isPositionRound"`
    NoPoints            bool     `json:"noPoints"`
    Name                string   `json:"customName"`
    Custom              bool     `json:"isCustom"`
    Movement            Movement `json:"movement"`
    LockedShiftRefresh  []any    `json:"lockedShiftRefresh"`
    LockedShift         []any    `json:"lockedShift"`
}

type Match struct {
    Id        string `json:"_id"`
    Team1Id   string `json:"team1_id"`
    Team1Lane int    `json:"team1_lane"`
    Team2Id   string `json:"team2_id"`
    Team2Lane int    `json:"team2_lane"`
    Game      int    `json:"game"`
    Shift     int    `json:"shift"`
}

type Movement struct {
    Moves []any `json:"moves"`
}
