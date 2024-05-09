package package_uts

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MuseumCollection struct {
    ID               primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Name             string             `bson:"name,omitempty" json:"name,omitempty"`
    Description      string             `bson:"description,omitempty" json:"description,omitempty"`
    Items            []MuseumItem       `bson:"items,omitempty" json:"items,omitempty"`
    Location         string             `bson:"location,omitempty" json:"location,omitempty"`
    OpeningHours     string             `bson:"opening_hours,omitempty" json:"opening_hours,omitempty"`
    EstablishedYear  int                `bson:"established_year,omitempty" json:"established_year,omitempty"`
    Director         string             `bson:"director,omitempty" json:"director,omitempty"`
    Website          string             `bson:"website,omitempty" json:"website,omitempty"`
    VisitorCount     int                `bson:"visitor_count,omitempty" json:"visitor_count,omitempty"`
}

type MuseumItem struct {
    ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
    Name        string             `bson:"name,omitempty" json:"name,omitempty"`
    Description string             `bson:"description,omitempty" json:"description,omitempty"`
    Year        int                `bson:"year,omitempty" json:"year,omitempty"`
    Artist      string             `bson:"artist,omitempty" json:"artist,omitempty"`
    Medium      string             `bson:"medium,omitempty" json:"medium,omitempty"`
    Dimensions  string             `bson:"dimensions,omitempty" json:"dimensions,omitempty"`
    Origin      string             `bson:"origin,omitempty" json:"origin,omitempty"`
    Acquisition string             `bson:"acquisition,omitempty" json:"acquisition,omitempty"`
    Condition   string             `bson:"condition,omitempty" json:"condition,omitempty"`
}
