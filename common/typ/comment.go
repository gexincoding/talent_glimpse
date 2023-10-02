package typ

type CommentDetail struct {
	CommentID uint
	Children  []CommentDetail
	LikeCount int
	CreateAt  string
}

type ListCommentResp struct {
	CommentDetailList []CommentDetail
}
