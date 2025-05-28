package system

// LoginRequest Login登录请求数据
type LoginRequest struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Point     string `json:"point"` //滑块验证码用户的坐标
	Timestamp string `json:"timestamp"`
}

// RegisterRequest register注册请求数据
type RegisterRequest struct {
	Nickname  string `json:"nickname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      int    `json:"role"`
	Point     string `json:"point"` //滑块验证码用户的坐标
	Timestamp string `json:"timestamp"`
}
