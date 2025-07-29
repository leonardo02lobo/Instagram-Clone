package models

type User struct {
	ID         int    "json:id"
	Username   string "json:username"
	Email      string "json:email"
	Password   string "json:password"
	ProfilePic string "json:profile_pic"
	Bio        string "json:bio"
	CreatedAt  string "json:created_at"
}
