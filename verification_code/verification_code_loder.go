package verification_code

import "my_project/consts"

var verifyCodeLink = map[string]string{
	consts.BOZHOU_UNIVERSITY:"http://211.141.201.154/CheckCode.aspx",
	consts.HRBUST_UNIVERSITY: "http://jwzx.hrbust.edu.cn/academic/getCaptcha.do",
}

func GetVerifyCodeLinkBySchoolName(schoolName string)string{
	return verifyCodeLink[schoolName]
}