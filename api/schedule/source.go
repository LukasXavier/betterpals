package schedule

type response struct {
    Id string `json:"_id"`;
    LeagueId string `json:"league_id"`;
    Schedule []respSchedule `json:"schedule"`;
    DisabledLanes []any `json:"disabledLanes"`;
    Shift []any `json:"shifts"`;
    V int `json:"__v"`;
}

type respSchedule struct {
    Date string `json:"date"`;
    Id string `json:"_id"`;
    CalculatedStartLane int `json:"calculatedStartLane"`;
    NormalWeek int `json:"normalWeek"`;
    Matches []respMatch `json:"matches"`;
    SplitMatches []any `json:"splitMatches"`;
    Skipped bool `json:"isSkipped"`;
    USBC bool `json:"isUSBC"`;
    Locked bool `json:"isLocked"`;
    DiffMultiShift bool `json:"isDiffMultiShift"`;
    SwitchPairLane bool `json:"isSwitchPairLane"`;
    OneTeamPairLane bool `json:"isOneTeamPairLane"`;
    LockPositionRound bool `json:"isLockPositionRound"`;
    DivPositionRound bool `json:"isDivPositionRound"`;
    PositionRound bool `json:"isPositionRound"`;
    NoPoints bool `json:"noPoints"`;
    Name string `json:"customName"`;
    Custom bool `json:"isCustom"`;
    Movement respMovement `json:"movement"`;
    LockedShiftRefresh []any `json:"lockedShiftRefresh"`;
    LockedShift []any `json:"lockedShift"`;
}

type respMatch struct {
    Id string `json:"_id"`;
    Team1Id string `json:"team1_id"`;
    Team1Lane int `json:"team1_lane"`;
    Team2Id string `json:"team2_id"`;
    Team2Lane int `json:"team2_lane"`;
    Game int `json:"game"`;
    Shift int `json:"shift"`;
}

type respMovement struct {
    Moves []any `json:"moves"`;
}
