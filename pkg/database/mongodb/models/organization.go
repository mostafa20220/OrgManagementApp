package models

type Organization struct {
	ID                      string `bson:"_id"`
	Name                    string `bson:"name"`
	Description             string `bson:"description"`
	OrganizationMembers     []OrganizationMember `bson:"organization_members"`
	// Other fields as needed
}

type OrganizationMember struct {
	Name        string `bson:"name"`
	Email       string `bson:"email"`
	AccessLevel string `bson:"access_level"`
}
