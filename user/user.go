package user

type User struct {
	Desktops map[string]struct{}
	Laptops  map[string]struct{}
}

func NewUser() *User {
	return &User{
		Desktops: make(map[string]struct{}, 1),
		Laptops:  make(map[string]struct{}, 1),
	}
}

func (u *User) AddIfUniqueDesktop(id string) {
	if _, ok := u.Desktops[id]; !ok {
		u.Desktops[id] = struct{}{}
	}
}

func (u *User) AddIfUniqueLaptop(id string) {
	if _, ok := u.Laptops[id]; !ok {
		u.Laptops[id] = struct{}{}
	}
}

// Applications returns the number of
// applications a user needs for all
// desktops and laptops attributed to
// them
//
// Because desktops and laptops are
// allowed to be installed simultaneously
// it can be assumed that are effectively
// one entity if both exist and whichever
// clone has more is the total required
func (u *User) Applications() int {
	if len(u.Desktops) < len(u.Laptops) {
		return len(u.Laptops)
	}

	return len(u.Desktops)
}

type Users map[string]*User

func NewUsers() Users {
	return make(map[string]*User, 1)
}

// TotalApplications returns the number
// of applications needed for all users
// in a given users map
func (u *Users) TotalApplications() int {
	var total int
	for _, user := range *u {
		total += user.Applications()
	}

	return total
}
