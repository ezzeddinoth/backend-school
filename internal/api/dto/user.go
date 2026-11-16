package dto

type TokenUser struct {
	Name              string   `json:"name"`
	PreferredUsername string   `json:"preferredUsername"`
	GivenName         string   `json:"givenName"`
	FamilyName        string   `json:"familyName"`
	Email             string   `json:"email"`
	Permissions       []string `json:"permissions"`
}
