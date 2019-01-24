package api

type LoginPostData struct {
	Username      string      `json:"username"`
	StdNum        string      `json:"std_num"`
}

type LoginReturnData struct {
	Token         string      `json:"token"`
	UID           int         `json:"uid"`
	StdNum        string      `json:"std_num"`
}
