package domain

import (
	"time"
)

type User struct {
	ID uint64 `json:"id"`

	UserID int64 `json:"user_id" gorm:"primaryKey"`
	// Authentificate and verification information
	Credentials UserCredentials `json:"credentials" gorm:"foreignkey:UserID;references:UserID"`

	// Personal information
	Biography UserBiography `json:"biography" gorm:"foreignkey:UserID;references:UserID"`

	// Contacts information about user
	Contact UserContactInfo `json:"contact" gorm:"foreignkey:UserID;references:UserID"`

	// Information about user's profile
	Profile UserProfile `json:"profile" gorm:"foreignkey:UserID;references:UserID"`

	// Profile photos
	Avatars []UserAvatar `json:"avatars" gorm:"foreignkey:UserID;references:UserID"`

	// Information about page of user's friends
	Friends []UserFriend `json:"friends" gorm:"foreignkey:UserID;references:UserID"`

	// Information about blocked pages
	Blocked []UserBlocked `json:"blocked_users" gorm:"foreignkey:UserID;references:UserID"`
}

type UserCredentials struct {
	UserID int64 `json:"user_id" gorm:"primaryKey"`
	// add email tag in validate
	Email    string `json:"email" validate:"required" gorm:"not null;unique"`
	Password string `json:"password" validate:"required" gorm:"-"`
	HashPass string `json:"-"`
}

type UserBiography struct {
	UserID     int64     `json:"user_id" gorm:"primaryKey"`
	FirstName  string    `json:"first_name" validate:"required"`
	SecondName string    `json:"second_name" validate:"required"`
	Birthday   time.Time `json:"birthday"`
	Age        int       `json:"age"`
}

type UserProfile struct {
	UserID   int64  `json:"user_id" gorm:"primaryKey"`
	AvatarID int64  `json:"avatar_id"`
	Username string `json:"username"`
	Status   string `json:"status"`

	IsClosed  bool      `json:"is_closed"`
	Online    bool      `json:"online"`
	IsBlocked bool      `json:"is_blocked"`
	LastSeen  time.Time `json:"last_seen"`
	CreatedAt time.Time `json:"created_at"`

	NumFriends     int `json:"num_friends"`
	NumSubscribers int `json:"num_subscribers"`
	NumGroups      int `json:"num_groups"`
}

type UserContactInfo struct {
	UserID  int64  `json:"user_id" gorm:"primaryKey"`
	Phone   string `json:"phone"`
	City    string `json:"city"`
	Country string `json:"country"`
}

type UserAvatar struct {
	AvatarID int64     `json:"avatar_id" gorm:"primaryKey"`
	UserID   int64     `json:"user_id"`
	Bucket   string    `json:"bucket"`
	Key      string    `json:"key"`
	AddedAt  time.Time `json:"added_at"`
	Likes    int       `json:"likes"`
}

type UserSettings struct {
	UserID            int64  `json:"user_id" gorm:"primaryKey"`
	EmailNotification bool   `json:"email_notification"`
	OpenToFriendship  bool   `json:"open_to_friendship"`
	ShowLastSeen      bool   `json:"show_last_seen"`
	Theme             string `json:"theme"`
}

type UserFriend struct {
	UserID    int64     `json:"user_id" gorm:"primaryKey"`
	FriendID  int64     `json:"fiend_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserBlocked struct {
	UserID    int64     `json:"user_id" gorm:"primaryKey"`
	BlockedID int64     `json:"blocked_id"`
	CreatedAt time.Time `json:"created_at"`
}

type NewUser struct {
	Biography   UserBiography   `json:"biography"`
	Credentials UserCredentials `json:"credentials"`
}
