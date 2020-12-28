package consts

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	20001:          "TOKEN校验失败",
	20002:          "TOKEN校验超时",
	20003:          "生成TOKEN失败",
	20004:          "错误的TOKEN",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
