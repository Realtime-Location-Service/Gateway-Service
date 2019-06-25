package model

import "strings"

// User ...
type User struct {
	UserID       string  `json:"user_id"`
	Domain       string  `json:"domain,omitempty"`
	Role         string  `json:"role,omitempty"`
	AppKey       string  `json:"app_key,omitempty"`
	CompanyID    int     `json:"company_id,omitempty"`
	Subordinates []*User `json:"subordinates,omitempty"`
}

// SubordinateIDs return subordinate user ids
func (u *User) SubordinateIDs() string {
	ids := []string{}
	for _, uu := range u.Subordinates {
		if uu.UserID != "" {
			ids = append(ids, uu.UserID)
		}
	}
	return strings.Join(ids, ",")

}
