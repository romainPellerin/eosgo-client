/**
 *  @copyright defined in eosgo-client/LICENSE.txt
 *  @author Romain Pellerin - romain@slyseed.com
 *
 *  Donation appreciated :)
 *  EOS wallet: 0x8d39307d9a0687c894058115365f6ad0f03b9ff9
 *	ETH wallet: 0x317b60152f0a90c10cea52d751ccb4dfad2fe9e7
 *  BTC wallet: 3HdFx8C3WNA6RyyGYKygECGbLD1BxyeqTN
 *  BCH wallet: 14e9fnmcHSZxxd3fnrkaG6EYph9JuWcF9T
 */

package common

import (
	"github.com/rs/zerolog/log"
	"context"
	"github.com/rs/zerolog"
	"eosgo-client/errors"
	"log/syslog"
	"io"
	"os"
	"fmt"
	"encoding/json"
	"net/http"
)

const (
	MODE_SYSLOG  = "SYSLOG"
	MODE_CONSOLE = "STDOUT"
)

func LoggerInit(level string) {

	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break;
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break;
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		break;

	}
}

func Logger(module string, context context.Context) zerolog.Logger {
	var writer io.Writer
	var err error

	switch Config.LOGGING_MODE {
	case MODE_CONSOLE:
		writer = zerolog.ConsoleWriter{os.Stdout, false}
		break

	case MODE_SYSLOG:
		writer, err = syslog.New(syslog.LOG_LOCAL3, "eosgo")
		break
	}

	if err != nil {
		panic("unable to register zerolog on syslog")
	}

	ctx := log.With().Timestamp().Str("module", module).Logger().Output(writer).WithContext(context)

	return *log.Ctx(ctx);
}

func InfoLog(logger zerolog.Logger, method string, message string) {
	logger.Info().Str("method", method).Msg(message)
}

func InfoLogWithParams(logger zerolog.Logger, method string,
	message string, params map[string]string) {
	logger.Info().Str("method", method).Interface("params", params).Msg(message)
}

func ErrorLog(logger zerolog.Logger, method string, appErr *errors.AppError, message string) {
	if appErr.Error != nil {
		logger.Error().Str("method", method).
			Err(appErr.Error).
			Dict("apperror", zerolog.Dict().
			Str("message", appErr.Message).
			Int("code", appErr.Code)).
			Msg(message)
	} else {
		logger.Error().Str("method", method).
			Dict("apperror", zerolog.Dict().
			Str("message", appErr.Message).
			Int("code", appErr.Code)).
			Msg(message)
	}
}

func ErrorLogWithParams(logger zerolog.Logger, method string,
	appErr *errors.AppError, message string, params map[string]string) {
	if appErr != nil && appErr.Error != nil {
		logger.Error().Str("method", method).
			Err(appErr.Error).
			Dict("apperror", zerolog.Dict().
			Str("message", appErr.Message).
			Int("code", appErr.Code)).
			Interface("params", params).
			Msg(message)
	} else if appErr != nil {
		logger.Error().Str("method", method).
			Dict("apperror", zerolog.Dict().
			Str("message", appErr.Message).
			Int("code", appErr.Code)).
			Interface("params", params).
			Msg(message)
	} else {
		logger.Error().Str("method", method).
			Dict("apperror", zerolog.Dict().
			Int("code", -1)).
			Str("message", message).
			Interface("params", params).
			Msg(message)
	}
}

// errorf writes a swagger-compliant error response.
func Errorf(logger zerolog.Logger, r *http.Request, w http.ResponseWriter, code int, format string, error error, appErr *errors.AppError, a ...interface{}) {

	var out struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	out.Code = code
	out.Message = fmt.Sprintf(format, a...)

	b, err := json.Marshal(out)

	if err != nil {

		HTTPErrorLog(logger, r, 500,
			"Could not format JSON for original message.",
			"marshalling failed", err, appErr)

		http.Error(w, `{"code": 500, "message": "Could not format JSON for original message."}`, 500)

	} else {

		if appErr == nil {
			HTTPErrorLog(logger, r, code, string(b), "error triggered", error, nil)
		} else {
			HTTPErrorLog(logger, r, code, string(b), "error triggered", error, appErr)
		}

		http.Error(w, string(b), code)
	}

	defer r.Body.Close()
}

// errorf writes a swagger-compliant error response.
func ErrorfWithParams(logger zerolog.Logger, r *http.Request, w http.ResponseWriter, code int, format string, error error, appErr *errors.AppError, params map[string]string, a ...interface{}) {

	var out struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}

	out.Code = code
	out.Message = fmt.Sprintf(format, a...)

	b, err := json.Marshal(out)

	if err != nil {

		HTTPErrorLogWithParams(logger, r, 500,
			"Could not format JSON for original message.",
			"marshalling failed", err, appErr, params)

		http.Error(w, `{"code": 500, "message": "Could not format JSON for original message."}`, 500)

	} else {

		HTTPErrorLogWithParams(logger, r, code, string(b), "error triggered", error, appErr, params)

		http.Error(w, string(b), code)
	}

	defer r.Body.Close()
}

func HTTPErrorLog(logger zerolog.Logger, r *http.Request, code int, content string, message string, err error, appErr *errors.AppError) {

	if err != nil {
		if appErr != nil && &appErr.Message != nil && &appErr.Code != nil {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Err(err).
				Dict("apperror", zerolog.Dict().
				Str("message", appErr.Message).
				Int("code", appErr.Code)).
				Msg(message)
		} else {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Err(err).
				Msg(message)
		}
	} else if appErr != nil {
		if &appErr.Message != nil && &appErr.Code != nil {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Dict("apperror", zerolog.Dict().
				Str("message", appErr.Message).
				Int("code", appErr.Code)).
				Msg(message)
		} else {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Msg(message)
		}
	} else {
		logger.Info().Str("url", r.URL.String()).
			Dict("response", zerolog.Dict().
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Int("code", code).
			Str("content", content)).
			Msg(message)
	}
}

func HTTPErrorLogWithParams(logger zerolog.Logger, r *http.Request, code int, content string, message string, err error, appErr *errors.AppError, params map[string]string) {
	if err != nil {
		if appErr != nil && &appErr.Message != nil && &appErr.Code != nil {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Interface("params", params).
				Err(err).
				Dict("apperror", zerolog.Dict().
				Str("message", appErr.Message).
				Int("code", appErr.Code)).
				Msg(message)
		} else {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Interface("params", params).
				Msg(message)
		}
	} else if appErr != nil {
		if &appErr.Message != nil && &appErr.Code != nil {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Interface("params", params).
				Dict("apperror", zerolog.Dict().
				Str("message", appErr.Message).
				Int("code", appErr.Code)).
				Msg(message)
		} else {
			logger.Info().Str("url", r.URL.String()).
				Dict("response", zerolog.Dict().
				Str("host", r.Host).
				Str("ip", r.RemoteAddr).
				Int("code", code).
				Str("content", content)).
				Interface("params", params).
				Msg(message)
		}
	} else {
		logger.Info().Str("url", r.URL.String()).
			Dict("response", zerolog.Dict().
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Int("code", code).
			Str("content", content)).
			Interface("params", params).
			Msg(message)
	}
}

func HTTPRequestLog(logger zerolog.Logger, r *http.Request, message string) {
	fmt.Printf("User agent %s", r.UserAgent());
	if r.UserAgent() == "" {
		logger.Info().Str("url", r.URL.String()).
			Dict("request", zerolog.Dict().
			Str("method", r.Method).
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Str("method", r.Method)).
			Msg(message)

	} else {
		logger.Info().Str("url", r.URL.String()).
			Dict("request", zerolog.Dict().
			Str("method", r.Method).
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Str("method", r.Method).
			Str("user-agent", r.UserAgent())).
			Msg(message)
	}
}

func HTTPRequestWithParamsLog(logger zerolog.Logger, r *http.Request, message string, params map[string]string) {
	fmt.Printf("User agent %s", r.UserAgent());
	if r.UserAgent() == "" {
		logger.Info().Str("url", r.URL.String()).
			Dict("request", zerolog.Dict().
			Str("method", r.Method).
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Str("method", r.Method)).
			Interface("params", params).
			Msg(message)

	} else {
		logger.Info().Str("url", r.URL.String()).
			Dict("request", zerolog.Dict().
			Str("method", r.Method).
			Str("host", r.Host).
			Str("ip", r.RemoteAddr).
			Str("method", r.Method).
			Str("user-agent", r.UserAgent())).
			Interface("params", params).
			Msg(message)
	}
}

func HTTPResponseLog(logger zerolog.Logger, r *http.Request, code int, contentType string, content string, message string) {
	logger.Info().Str("url", r.URL.String()).
		Dict("response", zerolog.Dict().
		Str("host", r.Host).
		Str("ip", r.RemoteAddr).
		Int("code", code).
		Str("content-type", contentType).
		Str("content", content)).
		Msg(message)

}
