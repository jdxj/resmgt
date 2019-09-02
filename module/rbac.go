package module

type Role struct {
	ID   int
	Name string
}

type Permission struct {
	ID   int
	Name string
}

type RolePerm struct {
	Role int
	Perm int
}

func (rp RolePerm) TableName() string {
	return "role_perm"
}
