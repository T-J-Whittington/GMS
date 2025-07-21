package user

type User struct {
	ID        uint     	`json:"id"`
	CreatedAt time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
    Email     string	`json:"email" db:"email"`
    Password  string	`json:"-" db:"password"`
    FirstName string	`json:"first_name" db:"first_name"`
    LastName  string	`json:"last_name" db:"last_name"`
    Role      string	`json:"role" db:"role"`
    Active    bool  	`json:"active" db:"default:true"`
}