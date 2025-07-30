package system

// LoginRequest Login登录请求数据
type LoginRequest struct {
	Username  string `json:"username" binding:"required" msg:"请正确输入用户名"` //用户名
	Password  string `json:"password" binding:"required" msg:"请正确输入密码"`  //密码
	Point     string `json:"point"`                                      //滑块验证码用户的坐标
	Timestamp string `json:"timestamp"`
}

// RegisterRequest register注册请求数据
type RegisterRequest struct {
	Nickname  string `json:"nickname" binding:"required" msg:"请正确输入昵称"`  //昵称
	Username  string `json:"username" binding:"required" msg:"请正确输入用户名"` //用户名
	Password  string `json:"password" binding:"required" msg:"请正确输入密码"`  //密码
	Role      int    `json:"role" binding:"required" msg:"请正确输入角色"`      //角色
	Point     string `json:"point"`                                      //滑块验证码用户的坐标
	Timestamp string `json:"timestamp"`
}
