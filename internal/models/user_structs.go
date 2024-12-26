package models

var RepositoryURL = "http://localhost:8088"

type Tag struct {
	Caption string
}

type ProfileMap struct {
	ProfileMap map[int]Profile
}

type Profile struct {
	Active     bool
	Registered int64
	Statuses   []string
	LastVisit  int64
	IsAdmin    bool
	Achives    []Tag
}

type UserState struct {
	State string
	Step  string
}

type StateRepo struct {
	States map[int]*UserState
	Vars   map[string]map[int]string
}

type UserProfile struct {
	ID        int
	Name      string
	Username  string
	City      string
	Login     string
	PswdHache string
	Profile   Profile
}
