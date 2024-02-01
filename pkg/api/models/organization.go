package models

type Organization struct {
	Name                    string `json:"name"`
	Description             string `json:"description"`
	OrganizationMembers     []OrganizationMember `json:"organization_members"`
}

type OrganizationMember struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	AccessLevel string `json:"access_level"`
}