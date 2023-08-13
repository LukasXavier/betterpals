package team

type response struct {
    Data []respData `json:"data"`;
}

type respData struct {
            Id string `json:"_id"`;
            Name string `json:"name"`;
            Email string `json:"email"`;
            FirstName string `json:"firstName"`;
            LastName string `json:"lastName"`;
            Dexterity int `json:"dexterity"`;
            Average int `json:"average"`;
            Averages []respAverage`json:"averages"`;
            Subs []any `json:"subs"`;
            NickNames []string `json:"nickNames"`;
            Junior bool `json:"isJunior"`;
            Female bool `json:"isFemale"`;
            Avatar string `json:"avatar"`;
            Games []int `json:"games"`;
            IndPointsWon int `json:"indPointsWon"`;
            HighGame int `json:"highGame"`;
            HighGameHdcp int `json:"highGameHdcp"`;
            HighSeries int `json:"highSeries"`;
            HighSeriesHdcp int `json:"highSeriesHdcp"`;
            RealAvg int `json:"realAvg"`;
            EnteringAvg int `json:"enteringAvg"`;
            GamesPlayed int `json:"gamesPlayed"`;
            TeamName string `json:"teamName"`;
            Team string `json:"team"`;
            League string `json:"league"`;
            LeagueName string `json:"leagueName"`;
            Center string `json:"center"`;
            CenterName string `json:"centerName"`;
            TotalPointsScored int `json:"totalPointsScored"`;
}

type respAverage struct {
    League string `json:"league"`;
    Team string `json:"team"`;
    Id string `json:"_id"`;
    IndividualPoints int `json:"individualPoints"`;
    TotalPointsCarryOn int `json:"totalPointsCarryOn"`;
    GamesCarryOn int `json:"gamesCarryOn"`;
    HighSeriesHdcp int `json:"highSeriesHdcp"`;
    HighGameHdcp int `json:"highGameHdcp"`;
    HighSeries int `json:"highSeries"`;
    HighGame int `json:"highGame"`;
    FixedAvg any `json:"fixedAvg"`;
    Average int `json:"average"`;
}

