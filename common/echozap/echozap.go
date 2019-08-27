package echozap

import (
	"bytes"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/color"
	"github.com/valyala/fasttemplate"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	// LoggerConfig defines the config for Logger middleware.
	LoggerConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper middleware.Skipper

		// Tags to constructed the logger format.
		//
		// - time_unix
		// - time_unix_nano
		// - time_rfc3339
		// - time_rfc3339_nano
		// - time_custom
		// - id (Request ID)
		// - remote_ip
		// - uri
		// - host
		// - method
		// - path
		// - protocol
		// - referer
		// - user_agent
		// - status
		// - error
		// - latency (In nanoseconds)
		// - latency_human (Human readable)
		// - bytes_in (Bytes received)
		// - bytes_out (Bytes sent)
		// - header:<NAME>
		// - query:<NAME>
		// - form:<NAME>
		//
		// Example "${remote_ip} ${status}"
		//
		// Optional. Default value DefaultLoggerConfig.Format.
		Format string `yaml:"format"`

		// Optional. Default value DefaultLoggerConfig.CustomTimeFormat.
		CustomTimeFormat string `yaml:"custom_time_format"`

		// Output is a writer where logs in JSON format are written.
		// Optional. Default value os.Stdout.
		Output io.Writer

		template *fasttemplate.Template
		colorer  *color.Color
		pool     *sync.Pool
		zap      *zap.Logger
		level    zapcore.Level
	}
)

var (
	// DefaultLoggerConfig is the default Logger middleware config.
	ZapLoggerConfig = LoggerConfig{
		Skipper: middleware.DefaultSkipper,
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
			`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
		Output:           os.Stdout,
		colorer:          color.New(),
	}
)

// Logger returns a middleware that logs HTTP requests.
func Logger(logger *zap.Logger) echo.MiddlewareFunc {
	ZapLoggerConfig.zap = logger
	return LoggerWithConfig(ZapLoggerConfig)
}

// LoggerWithConfig returns a Logger middleware with config.
// See: `Logger()`.
func LoggerWithConfig(config LoggerConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = ZapLoggerConfig.Skipper
	}
	if config.Format == "" {
		config.Format = ZapLoggerConfig.Format
	}
	if config.Output == nil {
		config.Output = ZapLoggerConfig.Output
	}
	config.template = fasttemplate.New(config.Format, "${", "}")
	config.colorer = color.New()
	config.colorer.SetOutput(config.Output)
	config.pool = &sync.Pool{
		New: func() interface{} {
			return bytes.NewBuffer(make([]byte, 256))
		},
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			if config.Skipper(c) {
				return next(c)
			}
			logger := config.zap
			req := c.Request()
			res := c.Response()
			start := time.Now()
			if err = next(c); err != nil {
				c.Error(err)
			}
			stop := time.Now()
			buf := config.pool.Get().(*bytes.Buffer)
			buf.Reset()
			defer config.pool.Put(buf)

			var (
				time_unix         string
				time_unix_nano    string
				time_rfc3339      string
				time_rfc3339_nano string
				time_custom       string

				id         string
				remote_ip  string
				host       string
				uri        string
				method     string
				path       string
				protocol   string
				referer    string
				user_agent string
				status     int
				level      zapcore.Level
				message    string

				latency       string
				latency_human string
				bytes_in      string
				bytes_out     string
			)

			if _, err = config.template.ExecuteFunc(buf, func(w io.Writer, tag string) (int, error) {
				//var status int
				//var latency string

				//zap.String("latency", time.Since(start).String()),
				//zap.String("id", id),
				//zap.String("method", req.Method),
				//zap.String("uri", req.RequestURI),
				//zap.String("host", req.Host),
				//zap.String("remote_ip", c.RealIP()),

				switch tag {
				case "time_unix":
					//return buf.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
					time_unix = strconv.FormatInt(time.Now().Unix(), 10)
				case "time_unix_nano":
					//return buf.WriteString(strconv.FormatInt(time.Now().UnixNano(), 10))
					time_unix_nano = strconv.FormatInt(time.Now().UnixNano(), 10)
				case "time_rfc3339":
					//return buf.WriteString(time.Now().Format(time.RFC3339))
					time_rfc3339 = time.Now().Format(time.RFC3339)
				case "time_rfc3339_nano":
					//return buf.WriteString(time.Now().Format(time.RFC3339Nano))
					time_rfc3339_nano = time.Now().Format(time.RFC3339Nano)
				case "time_custom":
					//return buf.WriteString(time.Now().Format(config.CustomTimeFormat))
					time_custom = time.Now().Format(config.CustomTimeFormat)
				case "id":
					id = req.Header.Get(echo.HeaderXRequestID)
					if id == "" {
						id = res.Header().Get(echo.HeaderXRequestID)
					}
					//return buf.WriteString(id)
				case "remote_ip":
					remote_ip = c.RealIP()
					//return buf.WriteString(c.RealIP())
				case "host":
					//return buf.WriteString(req.Host)
					host = req.Host
				case "uri":
					//return buf.WriteString(req.RequestURI)
					uri = req.RequestURI
				case "method":
					//return buf.WriteString(req.Method)
					method = req.Method
				case "path":
					p := req.URL.Path
					if p == "" {
						p = "/"
					}
					//return buf.WriteString(p)
					path = p
				case "protocol":
					//return buf.WriteString(req.Proto)
					protocol = req.Proto
				case "referer":
					//return buf.WriteString(req.Referer())
					referer = req.Referer()
				case "user_agent":
					//return buf.WriteString(req.UserAgent())
					user_agent = req.UserAgent()
				case "status":
					n := res.Status
					status = res.Status
					//s := config.colorer.Green(n)
					level = zapcore.DebugLevel
					switch {
					case n >= 500:
						//s = config.colorer.Red(n)
						level = zapcore.ErrorLevel
					case n >= 400:
						//s = config.colorer.Yellow(n)
						level = zapcore.ErrorLevel
					case n >= 300:
						//s = config.colorer.Cyan(n)
						level = zapcore.InfoLevel
					}
					//return buf.WriteString(s)
				case "error":
					if err != nil {
						level = zapcore.ErrorLevel
						//return buf.WriteString(err.Error())
						message = err.Error()
					}
				case "latency":
					l := stop.Sub(start)
					//return buf.WriteString(strconv.FormatInt(int64(l), 10))
					latency = strconv.FormatInt(int64(l), 10)
				case "latency_human":
					//return buf.WriteString(stop.Sub(start).String())
					latency_human = stop.Sub(start).String()
				case "bytes_in":
					cl := req.Header.Get(echo.HeaderContentLength)
					if cl == "" {
						cl = "0"
					}
					//return buf.WriteString(cl)
					bytes_in = cl
				case "bytes_out":
					//return buf.WriteString(strconv.FormatInt(res.Size, 10))
					bytes_out = strconv.FormatInt(res.Size, 10)
				default:
					switch {
					case strings.HasPrefix(tag, "header:"):
						return buf.Write([]byte(c.Request().Header.Get(tag[7:])))
					case strings.HasPrefix(tag, "query:"):
						return buf.Write([]byte(c.QueryParam(tag[6:])))
					case strings.HasPrefix(tag, "form:"):
						return buf.Write([]byte(c.FormValue(tag[5:])))
					case strings.HasPrefix(tag, "cookie:"):
						cookie, err := c.Cookie(tag[7:])
						if err == nil {
							return buf.Write([]byte(cookie.Value))
						}
					}
				}

				return 0, nil
			}); err != nil {
				return
			}

			//_, err = config.Output.Write(buf.Bytes())

			fields := []zapcore.Field{
				zap.Int("status", status),
				zap.String("method", method),
				zap.String("id", id),
				zap.String("remote_ip", remote_ip),
				zap.String("host", host),
				zap.String("uri", uri),
				zap.String("user_agent", user_agent),
				zap.String("latency", latency),
				zap.String("latency_human", latency_human),
				zap.String("bytes_in", bytes_in),
				zap.String("bytes_out", bytes_out),
			}
			logger.Check(level, message).Write(fields...)
			return
		}
	}
}
