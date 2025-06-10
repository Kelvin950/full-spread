package domain

import "time"

type User struct {
	ID          uint      `json:"id,omitempty"`
	Email       string    `json:"email,omitempty"`
	Username    string    `json:"username,omitempty"`
	DateOfBirth string    `json:"dob,omitempty"`
	Avatar      string    `json:"avatar,omitempty"`
	Firebaseuid *string   `json:"firebase_uid,omitempty"`
	Status      bool      `json:"status,omitempty"`
	Createdat   time.Time `json:"created_at,omitempty"`
	Updatedat   time.Time `json:"updated_at,omitempty"`
}

type Payload struct {
	ID       uint   `json:"id,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`
	Status   bool   `json:"status,omitempty"`
}

type GoogleuserRes struct {
	Email string `json:"string,omitempty"`
}

type Creator struct {
	ID            uint      `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	TotalEarnings float32   `json:"total_earnings,omitempty"`
	HeaderUrl     string    `json:"header_url,omitempty"`
	AvatarUrl     string    `json:"avatar_url,omitempty"`
	PhoneNumber   string    `json:"phone_number,omitempty"`
	UserID        uint      `json:"user_id,omitempty"`
	Delete        bool      `json:"delete,omitempty"`
	Deactivate    bool      `json:"deactivate,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Members struct {
	ID        uint      `json:"id,omitempty"`
	MemberID  uint      `json:"user_id,omitempty"`
	CreatorId uint      `json:"creator_id,omitempty"`
	Creators  Creator   `json:"creators,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type MembershipSubscription struct {
	ID        uint                     `json:"id,omitempty"`
	MemberID  uint                     `json:"user_id,omitempty"`
	CreatorId uint                     `json:"creator_id,omitempty"`
	Creators  []CreatorandSubscription `json:"creators,omitempty"`
	CreatedAt time.Time                `json:"created_at,omitempty"`
	UpdatedAt time.Time                `json:"updated_at,omitempty"`
}
type CreatorandSubscription struct {
	Creator
	SubscriptionID        uint      `json:"subscription_id,omitempty"`
	SubscriptionStartDate time.Time `json:"start_date,omitempty"`
	EndDate               time.Time `json:"end_date,omitempty"`
	Status                bool      `json:"status,omitempty"`
}
type Subscription struct {
	ID        uint      `json:"id,omitempty"`
	MemberID  uint      `json:"member_id,omitempty"`
	Status    bool      `json:"status,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
