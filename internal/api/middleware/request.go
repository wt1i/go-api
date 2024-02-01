package middleware

import (
	"context"
	"net"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"go-api/internal/infras/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-god/gutils"
	"github.com/go-god/logger"
	"go.uber.org/zap"
)

// NotFoundHandler not found api router
func NotFoundHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "this page not found!",
		})
	}
}

// RecoverHandler recover handler
func RecoverHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx := c.Request.Context()
				logger.Info(ctx, "exec panic error",
					zap.String("module", "web"), zap.String("trace_error", string(debug.Stack())),
				)

				// broker pipe
				if isBrokenPipe(ctx, err) {
					c.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				// services error
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
					"code":    http.StatusInternalServerError,
					"message": "server inner error",
				})
				return
			}
		}()

		c.Next()
	}
}

// AccessLog access log
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// x-request-id
		reqID := c.GetHeader("x-request-id")
		if reqID == "" {
			reqID = gutils.Uuid()
		}

		userAgentKey := logger.CtxKey{Name: "user-agent"}
		userAgent := c.GetHeader("User-Agent")
		c.Request = utils.ContextSet(c.Request, logger.XRequestID, reqID)
		c.Request = utils.ContextSet(c.Request, logger.ReqClientIP, c.Request.RemoteAddr)
		c.Request = utils.ContextSet(c.Request, logger.RequestMethod, c.Request.Method)
		c.Request = utils.ContextSet(c.Request, logger.RequestURI, c.Request.RequestURI)
		c.Request = utils.ContextSet(c.Request, userAgentKey, userAgent)

		logger.Info(c.Request.Context(), "exec begin",
			zap.String("module", "web"), zap.String(userAgentKey.String(), userAgent),
		)

		c.Next()
		logger.Info(c.Request.Context(), "exec end", map[string]interface{}{
			"exec_time": time.Since(start).Seconds(),
		})
	}
}

func isBrokenPipe(ctx context.Context, err interface{}) bool {
	// Check for a broken connection, as it is not really a
	// condition that warrants a panic stack trace.
	var brokenPipe bool
	if ne, ok := err.(*net.OpError); ok {
		if se, exist := ne.Err.(*os.SyscallError); exist {
			errMsg := strings.ToLower(se.Error())
			// logger error
			logger.Error(ctx, "os syscall error", map[string]interface{}{
				"trace_error": errMsg,
			})

			if strings.Contains(errMsg, "broken pipe") ||
				strings.Contains(errMsg, "reset by peer") ||
				strings.Contains(errMsg, "request headers: small read buffer") ||
				strings.Contains(errMsg, "unexpected EOF") ||
				strings.Contains(errMsg, "i/o timeout") {
				brokenPipe = true
			}
		}
	}

	return brokenPipe
}
