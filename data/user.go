package data

import "time"

type User struct {
	ID        int64
	Uuid      string
	Name      *string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Session struct {
	ID        int64
	Uuid      string
	Email     string
	UserID    int64
	CreatedAt time.Time
}

// create a new session for an existing user
func (user *User) CreateSession() (session Session, err error) {
	statement := `INSERT INTO session (uuid, email, user_id) VALUES ($1, $2, $3) returning *`
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	err = stmt.QueryRow(createUUID(), user.ID, user.Email, user.Name).Scan(&session.ID, &session.Uuid, &session.Email, &session.UserID, &session.CreatedAt)
	return
}

func (user *User) Session() (session *Session, err error) {
	session = &Session{}
	statement := `SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = $1`
	err = Db.QueryRow(statement, user.ID).Scan(session.ID, session.Uuid, session.Email, session.UserID, session.CreatedAt)
	return
}

// check cookie
func (session *Session) Check() (valid bool, err error) {
	err = Db.QueryRow(`SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = $1`, session.Uuid).Scan(&session)

	if err != nil {
		valid = false
		return
	}

	if session.ID != 0 {
		valid = true
	}

	return
}

func UserByEmail(email string) (user *User, err error) {
	err = Db.QueryRow(`SELECT id, uuid, email, name, password FROM users WHERE email = $1`, email).
		Scan(&user)
	return
}
