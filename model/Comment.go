package model


type Comment struct {
	ID			int64
	UserId		int64
	VideoId		int64
	Content		string
	CreatedAt	string
	IsDeleted	bool
}
