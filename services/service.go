package services

import "time"

type service struct {
	Id                 int        `json:"id"`
	Name               string     `json:"name"`
	AuthorizationToken string     `json:"authorization_token"`
	CreatedAt          *time.Time `json:"created_at"`
	OrganizationId     int        `json:"organization_id"`
	CreatedBy          string     `json:"creator_id"`
}

type servicev struct {
	Name      string `json:"name"`
	CreatedBy string `json:"creator_id"`
}

type serviceToken struct {
	Token string `json:"authorization_token"`
}

func (s *service) IsEmpty() bool {
	return (s.Id == 0 && s.Name == "" && s.AuthorizationToken == "" && s.CreatedAt == nil)
}

func (s *serviceToken) IsEmpty() bool {
	return s.Token == ""
}
func isSliceEmpty(s []service) bool {
	for _, ss := range s {
		if !ss.IsEmpty() {
			return false
		}
	}
	return true
}
