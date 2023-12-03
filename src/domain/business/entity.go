package business

import "time"

type Entity struct {
	UUID         string `json:"uuid"`
	NickName     string `json:"nickName"`
	Users        []User `json:"users"`
	BusinessType Type   `json:"businessType"`
	IsVerified   bool   `json:"isVerified"`
	IsDeleted    bool   `json:"isDeleted"`
	IsEnabled    bool   `json:"isEnabled"`
}

type User struct {
	UUID   string    `json:"uuid"`
	Name   string    `json:"name"`
	Roles  []string  `json:"roles"`
	JoinAt time.Time `json:"joinAt"`
}

type Type string

type businessTypes struct {
	Individual  Type
	Corporation Type
}

var Types = businessTypes{
	Individual:  "individual",
	Corporation: "corporation",
}
