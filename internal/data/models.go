package data

import "database/sql"

type Models struct {
	userModel UserModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		userModel: UserModel{conn: db},
	}
}
