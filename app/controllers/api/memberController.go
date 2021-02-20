package Controllers

import (
	models "dqh-test/app/models/common"
	services "dqh-test/app/services/api"
	Tools "dqh-test/libs/tools"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
	"time"
)

func Login(c *gin.Context)  {
	tel := strings.Trim(c.PostForm("tel"), " ")
	app_type := strings.Trim(c.PostForm("app_type"), " ")
	app_version := strings.Trim(c.PostForm("app_version"), " ")
	pwd := strings.Trim(c.PostForm("pwd"), " ")

	_ , err := models.GetVersionInfo(app_type, app_version)
	if err != nil{
		Tools.Error(c, 401, "请更新最新版本后，登录使用", "")
		return
	}

	if len(tel) == 0{
		Tools.Error(c, 402, "手机号码不能为空", "")
		return
	}
	//获取用户信息
	memberInfo ,err := models.GetMenberInfoByTel(tel)
	if err != nil{
		Tools.Error(c, 403, "用户不存在", "")
		return
	}
	if pwd != memberInfo.Pwd{
		Tools.Error(c, 404, "密码错误", "")
		return
	}

	//更新token
	memberInfo.Token = Tools.Md5(tel)
	if err := models.UpdateMemberInfo(memberInfo); err != nil{
		Tools.Error(c, 403,"非法请求", "")
		return
	}

	Tools.Success(c, 0, "成功", memberInfo)
	return

}

func Regist(c *gin.Context)  {
	tel := strings.Trim(c.PostForm("tel"), " ")
	bianma := strings.Trim(c.DefaultPostForm("bianma", "+86"), " ")
	yz_code := strings.Trim(c.PostForm("yz_code"), " ")
	pwd := strings.Trim(c.PostForm("pwd"), " ")
	re_tel := strings.Trim(c.PostForm("re_tel"), " ")
	if len(yz_code) == 0{
		Tools.Error(c, 401,"邀请码不能为空", "")
		return
	}
	if len(tel) == 0{
		Tools.Error(c, 401,"手机号不能为空", "")
		return
	}
	if len(pwd) == 0{
		Tools.Error(c, 401,"用户密码不能为空", "")
		return
	}
	var yznn = models.Yznn{YzAddress: tel,YzType: 1}
	err := models.GetYznnInfo(&yznn)
	if err != nil{
		Tools.Error(c, 402,err.Error(), "")
		return
	}
	if yznn.ETime < time.Now().Unix() {
		Tools.Error(c, 403,"验证码已过期", "")
		return
	}
	if yz_code != yznn.YzCode {
		Tools.Error(c, 405,"验证码错误", "")
		return
	}
	_ , err  = models.GetMenberInfoByTel(tel)
	if err == nil{
		Tools.Error(c, 405,"用户已存在", "")
		return
	}
	//新用户
	var newMember = models.Member{
		Bianma: bianma,
		Tel: tel,
		Pwd: pwd,
		Status: 1,
		IsStatus: 0,
		Addtime: time.Now().Unix(),
	}
	//推荐人信息
	tj_member,err := models.GetMenberInfoByIntrocode(re_tel)
	if err == nil{
		newMember.RePath = tj_member.RePath + strconv.Itoa(tj_member.Id) + ","
		newMember.ReId = tj_member.Id
		newMember.Level = tj_member.Level + 1
	}
	//生成邀请码
	newMember.IntroCode = Tools.GetRandomString(8)
	//获取BTC/USDT钱包地址
	btcWalletAddress,err := services.GetTokenAddress()
	if err == nil{
		newMember.Address = btcWalletAddress.Data.Address
		newMember.PrivateKey = btcWalletAddress.Data.PrivateKey
	}
	/*ercWalletAddress,err := services.Geterc20Address()
	if err == nil{
		newMember.ErcAddress = ercWalletAddress.Data.Address
		newMember.ErcPrivateKey = ercWalletAddress.Data.PrivateKey
	}*/

	/*if newMember.Insert() != nil{
		Tools.Error(c, 406,"注册失败", "")
		return
	}*/

	Tools.Success(c, 0, "注册成功", newMember)
	return
}
