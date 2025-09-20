package repo

import "github.com/jmoiron/sqlx"

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
	Find(email, pass string) (*User, error)

	// Get(userId int) (*User, error)
	// List() ([]*User, error)
	// Delete(userId int) error
	// Update(user User) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUser(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}
}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES (
			:first_name,
			:last_name,
			:email,
			:password,
			:is_shop_owner
		)
		RETURNING id
	`

	var userId int
	rows, err := r.dbCon.NamedQuery(query, user)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		rows.Scan(&userId)
	}

	user.ID = userId

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
