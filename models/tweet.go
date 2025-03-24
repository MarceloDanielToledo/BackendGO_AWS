package models

type Tweet struct {
	ID      string `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID  string `bson:"userid" json:"userid,omitempty"`
	Message string `bson:"message" json:"message,omitempty"`
	Date    string `bson:"date" json:"date,omitempty"`
}
