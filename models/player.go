package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Player struct {
    ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    Name          string             `bson:"name" json:"name"`
    Age           int                `bson:"age" json:"age"`
    Batting       string             `bson:"batting" json:"batting"`
    Bowling       string             `bson:"bowling" json:"bowling"`
    Centuries     int                `bson:"centuries" json:"centuries"`
    DateOfBirth   string             `bson:"dateOfBirth" json:"dateOfBirth"`
    HatTricks     int                `bson:"hatTricks" json:"hatTricks"`
    JerseyNumber  int                `bson:"jerseyNumber" json:"jerseyNumber"`
    Role          string             `bson:"role" json:"role"`
    Runs          int                `bson:"runs" json:"runs"`
    Wickets       int                `bson:"wickets" json:"wickets"`
    TeamsPlayedFor []Team            `bson:"teamsPlayedFor" json:"teamsPlayedFor"`
}

// Team represents the structure of a team played for
type Team struct {
    Team    string `bson:"team" json:"team"`
    Matches int    `bson:"matches" json:"matches"`
}
