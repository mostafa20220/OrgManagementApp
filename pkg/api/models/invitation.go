package models

type Invitation struct {
	OrganizationID string `json:"organization_id"`
	UserEmail   string `json:"user_email"`
}