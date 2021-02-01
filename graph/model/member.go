package github.com/daiv2/graphql-go-postgres

type Member struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Skill []*Skill `json:"skill"`
}

func (u *User) Create() error{

}