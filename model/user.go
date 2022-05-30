package model

type LogUser struct {
	ID 			int64`gorm:"primary_key"`
	Username 	string
	Password	string
	Token  		string
	FollowNum	int64
	FollowerNum	int64
	IsFollow	bool
}