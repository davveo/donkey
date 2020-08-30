package middleware

import (
	"github.com/gin-gonic/gin"
)

func OperationMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		//var body []byte
		//if c.Request.Method != http.MethodGet {
		//	var err error
		//	body, err = ioutil.ReadAll(c.Request.Body)
		//	if err != nil {
		//		global.GVA_LOG.Error("read body from request error:", err)
		//	} else {
		//		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		//	}
		//}
		//userId, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		//if err != nil {
		//	userId = 0
		//}
		//record := model.SysOperationRecord{
		//	Ip:     context.ClientIP(),
		//	Method: context.Request.Method,
		//	Path:   context.Request.URL.Path,
		//	Agent:  context.Request.UserAgent(),
		//	Body:   string(body),
		//	UserId: userId,
		//}
		//writer := responseBodyWriter{
		//	ResponseWriter: context.Writer,
		//	body:           &bytes.Buffer{},
		//}
		//c.Writer = writer
		//now := time.Now()
		//
		//c.Next()
		//
		//latency := time.Now().Sub(now)
		//record.ErrorMessage = context.Errors.ByType(gin.ErrorTypePrivate).String()
		//record.Status = c.Writer.Status()
		//record.Latency = latency
		//record.Resp = writer.body.String()
		//
		//if err := service.CreateSysOperationRecord(record); err != nil {
		//	global.GVA_LOG.Error("create operation record error:", err)
		//}
	}
}
