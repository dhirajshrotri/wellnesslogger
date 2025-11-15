package models

import "time"

type Exercise struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    int       `json:"userId" bson:"userId"`
	Exercise  string    `json:"exercise" bson:"exercise"`
	Category  string    `json:"category" bson:"category"`
	Sets      int       `json:"sets" bson:"sets"`
	Reps      int       `json:"reps" bson:"reps"`
	Date      time.Time `json:"date" bson:"date"`
}
