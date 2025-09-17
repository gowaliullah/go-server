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

var users []User

func (u User) Store() User {
	if u.ID != 0 {
		return u
	}

	u.ID = len(users) + 1

	users = append(users, u)
	return u
}

func Find(eamil, pass string) *User {
	for _, u := range users {
		if u.Email == eamil && u.Password == pass {
			return &u
		}
	}
	return nil
}
