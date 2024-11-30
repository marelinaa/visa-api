package domain

type Application struct {
	ID             int64    `json:"id"`
	FirstName      string   `json:"first_name" validate:"required"`
	LastName       string   `json:"last_name" validate:"required"`
	PhoneNumber    string   `json:"phone_num" validate:"required,e164"`
	Email          string   `json:"email"  validate:"required,email"`
	SubmissionDate string   `json:"sub_date"`
	Status         string   `json:"status"`
	Documents      []string `json:"docs"`
	ApplicantID    int64    `json:"applicant_id"`
	PaymentID      int64    `json:"payment_id"`
	OperatorID     int64    `json:"operator_id"`
}

type Application2 struct {
	ID             int64    `json:"id"`
	SubmissionDate string   `json:"sub_date"`
	Status         string   `json:"status"`
	Documents      []string `json:"docs"`
	ApplicantID    int64    `json:"applicant_id"`
	PaymentID      int64    `json:"payment_id"`
	OperatorID     int64    `json:"operator_id"`
}

type User struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_num" validate:"required,e164"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}

type Applicant struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_num" validate:"required,e164"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type Payment struct {
	ID             int64   `json:"id"`
	Sum            float64 `json:"sum"`
	PaymentDate    string  `json:"payment_date"`
	Status         string  `json:"status"`
	ApplicantionID int64   `json:"applicantion_id"`
}

type Operator struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	PhoneNumber  string `json:"phone_num" validate:"required,e164"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}
