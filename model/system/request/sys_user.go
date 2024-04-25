package request

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePassword struct {
	Password    string `json:"password"`
	NewPassword string `json:"newPassword"`
}
