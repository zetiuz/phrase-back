package favorites

type FavoritesRequest struct {
	User_id string `json:"user_id"`
	Post_id int64  `json:"post_id"`
}
