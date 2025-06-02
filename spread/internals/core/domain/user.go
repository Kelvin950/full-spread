package domain

import "time"

type User struct {
	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dob"`
	Avatar      string `json:"avatar"`
	Firebaseuid *string `json:"firebase_uid"`
	Status      bool   `json:"status"`
	Createdat   time.Time  `json:"created_at"` 
	Updatedat  time.Time  `json:"updated_at"`
}

type Payload struct{

	ID          uint   `json:"id"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	Status      bool   `json:"status"`
}

type GoogleuserRes struct{

	Email string  `json:"string"`
}