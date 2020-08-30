package captcha

import (
	"crypto/md5"
	"fmt"
	"github.com/davveo/donkey/utils/common"
	"image/color"
	"strings"

	"github.com/mojocn/base64Captcha"
)

var (
	driver base64Captcha.Driver
	store  = base64Captcha.DefaultMemStore // TODO 后期考虑使用redis存储
)

type Captcha struct {
	SeKey   string
	CodeSet string
	Expire  int64
}

func NewCaptcha() *Captcha {
	return &Captcha{
		SeKey:   "this is a test string",
		CodeSet: "2345678abcdefhijkmnpqrstuvwxyzABCDEFGHJKLMNPQRTUVWXY",
		Expire:  120,
	}
}

func GenerateCaptcha(captchaType string) (id, b64s string, err error) {
	c := color.RGBA{R: 0, G: 0, B: 0, A: 0}
	fonts := []string{"wqy-microhei.ttc"}
	sourceString := "1234567890qwertyuioplkjhgfdsazxcvbnm"
	sourceChinese := "设想,你在,处理,消费者,的音,频输,出音,频可,能无,论什,么都,没有,任何,输出,或者,它可,能是,单声道,立体声,或是,环绕立,体声的,,不想要,的值"

	switch captchaType {
	case "audio":
		driver = base64Captcha.NewDriverAudio(6, "zh")
	case "string":
		driver = base64Captcha.NewDriverString(
			60, 180, 0,
			0, 6, sourceString, &c, fonts)
	case "math":
		driver = base64Captcha.NewDriverMath(
			60, 190, 0, 0, &c, fonts)
	case "chinese":
		c = color.RGBA{R: 125, G: 125, B: 0, A: 118}

		driver = base64Captcha.NewDriverChinese(
			60, 320, 0, 0, 2, sourceChinese, &c, fonts)
	default:
		driver = base64Captcha.NewDriverDigit(
			70, 190, 6, 0.45, 85)
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = captcha.Generate()
	return
}

func VerfiyCaptcha(idKey, verifyVal string) bool {
	return store.Verify(idKey, verifyVal, true)
}

// 返回验证码标识
func (c *Captcha) GetKey(str string) string {
	return c.authCode(str)
}

// 验证验证码是否正确
// code 用户验证码
// id 已经生成的验证码标示
func (c *Captcha) Check(code, id string) bool {
	return false
}

func (c *Captcha) GetImage() {

}

func (c *Captcha) WriteCurve() {

}

func (c *Captcha) getMd5Str(str string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}

func (c *Captcha) getPart(str string, start, end int) string {
	return common.SubString(str, start, end)
}

func (c *Captcha) authCode(str string) string {
	var (
		build strings.Builder
	)
	str1, str2 := c.getMd5Str(c.SeKey), c.getMd5Str(str)
	key := c.getPart(str1, 5, 8)
	iStr := c.getPart(str2, 8, 10)

	build.WriteString(key)
	build.WriteString(iStr)
	return c.getMd5Str(build.String())
}
