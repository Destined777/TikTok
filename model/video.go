package model

type Video struct {
	Id				int64
	UserId  		int64
	PlayUrl			string
	CoverUrl		string
	FavouriteNum	int64
	CommentNum		int64
	Title			string
	CreatedAt		int64
}
