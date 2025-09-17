package repo

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Get(userId int) (*User, error)
	List() ([]*User, error)
	Delete(userId int) error
	Update(user User) (*User, error)
}

type userRepo struct {
	users []*User
}

func NewUser() UserRepo {
	return &userRepo{}
}

func (r *userRepo) Create(user User) (*User, error) {
	if user.ID != 0 {
		return &user, nil
	}

	user.ID = len(r.users) + 1

	r.users = append(r.users, &user)
	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	for _, u := range r.users {
		if u.Email == email && u.Password == pass {
			return u, nil
		}
	}
	return nil, nil
}

// func (r *userRepo) Get(userId int) (*User, error)              {}
// func (r *userRepo) List() ([]*User, error)          {}
// func (r *userRepo) Delete(userId int) error         {}
// func (r *userRepo) Update(user User) (*User, error) {}

func Find(eamil, pass string) *User {
	for _, u := range users {
		if u.Email == eamil && u.Password == pass {
			return &u
		}
	}
	return nil
}
