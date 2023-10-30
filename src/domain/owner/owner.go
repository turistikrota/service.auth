package owner

import "time"

type Entity struct {
	UUID       string `json:"uuid"`
	NickName   string `json:"nickName"`
	Users      []User `json:"users"`
	OwnerType  Type   `json:"ownerType"`
	IsVerified bool   `json:"isVerified"`
	IsDeleted  bool   `json:"isDeleted"`
	IsEnabled  bool   `json:"isEnabled"`
}

type User struct {
	UUID   string    `json:"uuid"`
	Name   string    `json:"name"`
	Roles  []string  `json:"roles"`
	JoinAt time.Time `json:"joinAt"`
}

type Type string

type ownerTypes struct {
	Individual  Type
	Corporation Type
}

var Types = ownerTypes{
	Individual:  "individual",
	Corporation: "corporation",
}
