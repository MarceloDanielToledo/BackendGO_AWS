package models

type Relation struct {
	UserId         string `bson:"userid" json:"userid"`
	UserRelationId string `bson:"userrelationid" json:"userrelationid"`
}
