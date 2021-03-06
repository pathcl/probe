package http

import (
	"fmt"
)

// probeError satisfies the errors.HTTPProbeError interface.
type probeError struct {
	address    string
	statusCode int
	body       string
	timeout    bool
	message    string
	cause      error
}

func (pe *probeError) Address() string {
	return pe.address
}

func (pe *probeError) StatusCode() int {
	return pe.statusCode
}

func (pe *probeError) Body() string {
	return pe.body
}

func (pe *probeError) Cause() error {
	return pe.cause
}

func (pe *probeError) Timeout() bool {
	return pe.timeout
}

func (pe *probeError) Error() string {
	msg := fmt.Sprintf(
		"%s (address=%s, timeout=%t, status-code=%d)",
		pe.message,
		pe.address,
		pe.timeout,
		pe.statusCode,
	)
	if pe.cause != nil {
		msg = fmt.Sprintf("%s: %v", msg, pe.cause)
	}
	return msg
}

func (pe *probeError) String() string {
	return pe.Error()
}
