package response

import (
	"encoding/json"
	"net/http"

	"project-v/pkg/errors"
)

func HttpStatusFromCode(code errors.Code) int {
	switch code {
	case errors.OK:
		return http.StatusOK // 200
	case errors.Canceled:
		return http.StatusRequestTimeout // 408
	case errors.Unknown:
		return http.StatusInternalServerError // 500
	case errors.InvalidArgument:
		return http.StatusBadRequest // 400

	case errors.DeadlineExceeded:
		return http.StatusGatewayTimeout // 504
	case errors.NotFound:
		return http.StatusNotFound // 404
	case errors.AlreadyExists:
		return http.StatusConflict // 409

	case errors.PermissionDenied:
		return http.StatusForbidden // 403
	case errors.Unauthenticated:
		return http.StatusUnauthorized // 401
	case errors.ResourceExhausted:
		return http.StatusTooManyRequests // 429
	case errors.FailedPrecondition:
		return http.StatusPreconditionFailed // 412
	case errors.Aborted:
		return http.StatusConflict // 409

	case errors.OutOfRange:
		return http.StatusBadRequest // 400
	case errors.Unimplemented:
		return http.StatusNotImplemented // 501
	case errors.Internal:
		return http.StatusInternalServerError // 500
	case errors.Unavailable:
		return http.StatusServiceUnavailable // 503
	case errors.DataLoss:
		return http.StatusInternalServerError // 500

	}

	// ll.Info("Unknown error code: %v", l.Int("code", int(code)))
	return http.StatusInternalServerError
}

func Write(w http.ResponseWriter, data string) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(data))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Json(w http.ResponseWriter, data interface{}) {
	responseJSON(w, data)
}

func responseJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func JsonError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

}
