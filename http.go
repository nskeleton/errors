package errors

import (
	"net/http"
)

var (
	fromHTTP map[int]int // Populated in init

	toHTTP = map[int]int{
		StatusBadReq:              http.StatusBadRequest,
		StatusForbidden:           http.StatusForbidden,
		StatusNotFound:            http.StatusNotFound,
		StatusWrongAcceptType:     http.StatusNotAcceptable,
		StatusReqTimeout:          http.StatusRequestTimeout,
		StatusPreconditionFailed:  http.StatusPreconditionFailed,
		StatusTooManyReqs:         http.StatusTooManyRequests,
		StatusInternal:            http.StatusInternalServerError,
		StatusUnimplemented:       http.StatusNotImplemented,
		StatusUpstreamUnreachable: http.StatusBadGateway,
		StatusUnavailable:         http.StatusServiceUnavailable,
	}
)
