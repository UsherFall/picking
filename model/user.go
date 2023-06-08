package model

type User struct {
	UserId    string `json: userId`
	NickName  string `json: nickName`
	Role      uint32 `json: role`
	LoginTime int64  `json: loginTime`
}
