package user

type Repository interface {
    FindAll() ([]User, error)
    FindByID(id string) (*User, error)
    Save(v *User) error
}

type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &repository{db: db}
}