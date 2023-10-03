package typ

type CommentDetail struct {
	CommentID    uint
	Content      string
	UserID       uint
	UserName     string
	UserImageURL string
	LikeCount    int
	CreateAt     string
	Children     []CommentDetail
}

type ListCommentResp struct {
	CommentDetailList []CommentDetail
}
