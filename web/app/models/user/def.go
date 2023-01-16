package user

type user struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	ProfilePic   string `json:"profilePic"`
	ExternalType string `json:"externalType"`
	ExternalID   string `json:"externalID"`
}
