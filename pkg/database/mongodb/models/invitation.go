package models

type Invitation struct {
	ID          string `bson:"_id"`
	OrganizationID string `bson:"organization_id"`
	UserEmail   string `bson:"user_email"`
}