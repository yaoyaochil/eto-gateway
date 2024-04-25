package response

type LoginUserInfo struct {
	Username  string `json:"username"`
	HeaderImg string `json:"headerImg"`
	NickName  string `json:"nickName"`
	ID        uint   `json:"id"`
	UUID      string `json:"uuid"`
}

type LoginResponse struct {
	User      LoginUserInfo `json:"user"`
	Token     string        `json:"token"`
	ExpiresAt int64         `json:"expiresAt"`
}
