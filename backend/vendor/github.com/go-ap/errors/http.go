package errors

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-ap/jsonld"
)

const errorsPackageName = "github.com/go-ap/errors"
const runtimeDebugPackageName = "runtime/debug"

type Error interface {
	error
	json.Unmarshaler
}

type notFound struct {
	Err
	s int
}

type methodNotAllowed struct {
	Err
	s int
}

type notValid struct {
	Err
	s int
}

type forbidden struct {
	Err
	s int
}

type notImplemented struct {
	Err
	s int
}

type conflict struct {
	Err
	s int
}

type gone struct {
	Err
	s int
}

type badRequest struct {
	Err
	s int
}

type unauthorized struct {
	Err
	s         int
	challenge string
}

type notSupported struct {
	Err
	s int
}

type timeout struct {
	Err
	s int
}

type badGateway struct {
	Err
	s int
}

func wrapErr(err error, s string, args ...interface{}) Err {
	e := Annotatef(err, s, args...)
	asErr := Err{}
	As(e, &asErr)
	return asErr
}

func FromResponse(resp *http.Response) error {
	if resp.StatusCode < http.StatusBadRequest {
		return nil
	}
	body := make([]byte, 0)
	defer resp.Body.Close()

	body, _ = ioutil.ReadAll(resp.Body)

	var withStatus error
	errors, err := UnmarshalJSON(body)
	if err != nil {
		return AnnotateFromStatus(nil, resp.StatusCode, string(body))
	}
	for _, err := range errors {
		if err == nil {
			withStatus = err
		}
		withStatus = Annotatef(err, err.Error())
	}
	return AnnotateFromStatus(withStatus, resp.StatusCode, resp.Status)
}

func AnnotateFromStatus(err error, status int, s string, args ...interface{}) error {
	switch status {
	case http.StatusBadRequest:
		return NewBadRequest(err, s, args...)
	case http.StatusUnauthorized:
		return NewUnauthorized(err, s, args...)
	// http.StatusPaymentRequired
	case http.StatusForbidden:
		return NewForbidden(err, s, args...)
	case http.StatusNotFound:
		return NewNotFound(err, s, args...)
	case http.StatusMethodNotAllowed:
		return NewMethodNotAllowed(err, s, args...)
	case http.StatusNotAcceptable:
		return NewNotValid(err, s, args...)
	// case http.StatusProxyAuthRequired
	// case http.StatusRequestTimeout
	case http.StatusConflict:
		return NewConflict(err, s, args...)
	case http.StatusGone:
		return NewGone(err, s, args...)
	// case http.StatusLengthRequres
	// case http.StatusPreconditionFailed
	// case http.StatusRequestEntityTooLarge
	// case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	// case http.StatusRequestedRangeNotSatisfiable
	// case http.StatusExpectationFailed
	// case http.StatusTeapot
	// case http.StatusMisdirectedRequest
	// case http.StatusUnprocessableEntity
	// case http.StatusLocked
	// case http.StatusFailedDependency
	// case http.StatusTooEarly
	// case http.StatusTooManyRequests
	// case http.StatusRequestHeaderFieldsTooLarge
	// case http.StatusUnavailableForLegalReason
	// case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NewNotImplemented(err, s, args...)
	case http.StatusBadGateway:
		return NewBadGateway(err, s, args...)
	// case http.StatusServiceUnavailable
	// case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NewNotSupported(err, s, args...)
	case http.StatusGatewayTimeout:
		return NewTimeout(err, s, args...)
	}
	return Annotatef(err, s, args...)
}

func NewFromStatus(status int, s string, args ...interface{}) error {
	switch status {
	case http.StatusBadRequest:
		return BadRequestf(s, args...)
	case http.StatusUnauthorized:
		return Unauthorizedf(s, args...)
	// http.StatusPaymentRequired
	case http.StatusForbidden:
		return Forbiddenf(s, args...)
	case http.StatusNotFound:
		return NotFoundf(s, args...)
	case http.StatusMethodNotAllowed:
		return MethodNotAllowedf(s, args...)
	case http.StatusNotAcceptable:
		return NotValidf(s, args...)
	// case http.StatusProxyAuthRequired
	// case http.StatusRequestTimeout
	case http.StatusConflict:
		return Conflictf(s, args...)
	case http.StatusGone:
		return Gonef(s, args...)
	// case http.StatusLengthRequres
	// case http.StatusPreconditionFailed
	// case http.StatusRequestEntityTooLarge
	// case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	// case http.StatusRequestedRangeNotSatisfiable
	// case http.StatusExpectationFailed
	// case http.StatusTeapot
	// case http.StatusMisdirectedRequest
	// case http.StatusUnprocessableEntity
	// case http.StatusLocked
	// case http.StatusFailedDependency
	// case http.StatusTooEarly
	// case http.StatusTooManyRequests
	// case http.StatusRequestHeaderFieldsTooLarge
	// case http.StatusUnavailableForLegalReason
	// case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NotImplementedf(s, args...)
	case http.StatusBadGateway:
		return BadGatewayf(s, args...)
	// case http.StatusServiceUnavailable
	// case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NotSupportedf(s, args...)
	case http.StatusGatewayTimeout:
		return Timeoutf(s, args...)
	}
	return Newf(s, args...)
}

func WrapWithStatus(status int, err error, s string, args ...interface{}) error {
	switch status {
	case http.StatusBadRequest:
		return NewBadRequest(err, s, args...)
	case http.StatusUnauthorized:
		return NewUnauthorized(err, s, args...)
	// http.StatusPaymentRequired
	case http.StatusForbidden:
		return NewForbidden(err, s, args...)
	case http.StatusNotFound:
		return NewNotFound(err, s, args...)
	case http.StatusMethodNotAllowed:
		return NewMethodNotAllowed(err, s, args...)
	case http.StatusNotAcceptable:
		return NewNotValid(err, s, args...)
	// case http.StatusProxyAuthRequired
	// case http.StatusRequestTimeout
	case http.StatusConflict:
		return NewConflict(err, s, args...)
	case http.StatusGone:
		return NewGone(err, s, args...)
	// case http.StatusLengthRequres
	// case http.StatusPreconditionFailed
	// case http.StatusRequestEntityTooLarge
	// case http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	// case http.StatusRequestedRangeNotSatisfiable
	// case http.StatusExpectationFailed
	// case http.StatusTeapot
	// case http.StatusMisdirectedRequest
	// case http.StatusUnprocessableEntity
	// case http.StatusLocked
	// case http.StatusFailedDependency
	// case http.StatusTooEarly
	// case http.StatusTooManyRequests
	// case http.StatusRequestHeaderFieldsTooLarge
	// case http.StatusUnavailableForLegalReason
	// case http.StatusInternalServerError
	case http.StatusNotImplemented:
		return NewNotImplemented(err, s, args...)
	case http.StatusBadGateway:
		return NewBadGateway(err, s, args...)
	// case http.StatusServiceUnavailable
	// case http.StatusGatewayTimeout
	case http.StatusHTTPVersionNotSupported:
		return NewNotSupported(err, s, args...)
	case http.StatusGatewayTimeout:
		return NewTimeout(err, s, args...)
	}
	return wrapErr(err, s, args...)
}
func NotFoundf(s string, args ...interface{}) *notFound {
	return &notFound{Err: wrapErr(nil, s, args...), s: http.StatusNotFound}
}
func NewNotFound(e error, s string, args ...interface{}) *notFound {
	return &notFound{Err: wrapErr(e, s, args...), s: http.StatusNotFound}
}
func MethodNotAllowedf(s string, args ...interface{}) *methodNotAllowed {
	return &methodNotAllowed{Err: wrapErr(nil, s, args...), s: http.StatusMethodNotAllowed}
}
func NewMethodNotAllowed(e error, s string, args ...interface{}) *methodNotAllowed {
	return &methodNotAllowed{Err: wrapErr(e, s, args...), s: http.StatusMethodNotAllowed}
}
func NotValidf(s string, args ...interface{}) *notValid {
	return &notValid{Err: wrapErr(nil, s, args...)}
}
func NewNotValid(e error, s string, args ...interface{}) *notValid {
	return &notValid{Err: wrapErr(e, s, args...)}
}
func Conflictf(s string, args ...interface{}) *conflict {
	return &conflict{Err: wrapErr(nil, s, args...), s: http.StatusConflict}
}
func NewConflict(e error, s string, args ...interface{}) *conflict {
	return &conflict{Err: wrapErr(e, s, args...), s: http.StatusConflict}
}
func Gonef(s string, args ...interface{}) *gone {
	return &gone{Err: wrapErr(nil, s, args...), s: http.StatusGone}
}
func NewGone(e error, s string, args ...interface{}) *gone {
	return &gone{Err: wrapErr(e, s, args...), s: http.StatusGone}
}
func Forbiddenf(s string, args ...interface{}) *forbidden {
	return &forbidden{Err: wrapErr(nil, s, args...), s: http.StatusForbidden}
}
func NewForbidden(e error, s string, args ...interface{}) *forbidden {
	return &forbidden{Err: wrapErr(e, s, args...), s: http.StatusForbidden}
}
func NotImplementedf(s string, args ...interface{}) *notImplemented {
	return &notImplemented{Err: wrapErr(nil, s, args...), s: http.StatusNotImplemented}
}
func NewNotImplemented(e error, s string, args ...interface{}) *notImplemented {
	return &notImplemented{Err: wrapErr(e, s, args...), s: http.StatusNotImplemented}
}
func BadRequestf(s string, args ...interface{}) *badRequest {
	return &badRequest{Err: wrapErr(nil, s, args...), s: http.StatusBadRequest}
}
func NewBadRequest(e error, s string, args ...interface{}) *badRequest {
	return &badRequest{Err: wrapErr(e, s, args...), s: http.StatusBadRequest}
}
func Unauthorizedf(s string, args ...interface{}) *unauthorized {
	return &unauthorized{Err: wrapErr(nil, s, args...), s: http.StatusUnauthorized}
}
func NewUnauthorized(e error, s string, args ...interface{}) *unauthorized {
	return &unauthorized{Err: wrapErr(e, s, args...), s: http.StatusUnauthorized}
}
func NotSupportedf(s string, args ...interface{}) *notSupported {
	return &notSupported{Err: wrapErr(nil, s, args...), s: http.StatusHTTPVersionNotSupported}
}
func NewNotSupported(e error, s string, args ...interface{}) *notSupported {
	return &notSupported{Err: wrapErr(e, s, args...), s: http.StatusHTTPVersionNotSupported}
}
func Timeoutf(s string, args ...interface{}) *timeout {
	return &timeout{Err: wrapErr(nil, s, args...), s: http.StatusRequestTimeout}
}
func NewTimeout(e error, s string, args ...interface{}) *timeout {
	return &timeout{Err: wrapErr(e, s, args...), s: http.StatusRequestTimeout}
}
func BadGatewayf(s string, args ...interface{}) *badGateway {
	return &badGateway{Err: wrapErr(nil, s, args...), s: http.StatusBadGateway}
}
func NewBadGateway(e error, s string, args ...interface{}) *badGateway {
	return &badGateway{Err: wrapErr(e, s, args...), s: http.StatusBadGateway}
}
func IsBadRequest(e error) bool {
	_, okp := e.(*badRequest)
	_, oks := e.(badRequest)
	return okp || oks || As(e, &badRequest{})
}
func IsForbidden(e error) bool {
	_, okp := e.(*forbidden)
	_, oks := e.(forbidden)
	return okp || oks || As(e, &forbidden{})
}
func IsNotSupported(e error) bool {
	_, okp := e.(*notSupported)
	_, oks := e.(notSupported)
	return okp || oks
}
func IsConflict(e error) bool {
	_, okp := e.(*conflict)
	_, oks := e.(conflict)
	return okp || oks || As(e, &conflict{})
}
func IsGone(e error) bool {
	_, okp := e.(*gone)
	_, oks := e.(gone)
	return okp || oks || As(e, &gone{})
}
func IsMethodNotAllowed(e error) bool {
	_, okp := e.(*methodNotAllowed)
	_, oks := e.(methodNotAllowed)
	return okp || oks || As(e, &methodNotAllowed{})
}
func IsNotFound(e error) bool {
	_, okp := e.(*notFound)
	_, oks := e.(notFound)
	return okp || oks || As(e, &notFound{})
}
func IsNotImplemented(e error) bool {
	_, okp := e.(*notImplemented)
	_, oks := e.(notImplemented)
	return okp || oks || As(e, &notImplemented{})
}
func IsUnauthorized(e error) bool {
	_, okp := e.(*unauthorized)
	_, oks := e.(unauthorized)
	return okp || oks || As(e, &unauthorized{})
}
func IsTimeout(e error) bool {
	_, okp := e.(*timeout)
	_, oks := e.(timeout)
	return okp || oks || As(e, &timeout{})
}
func IsNotValid(e error) bool {
	_, okp := e.(*notValid)
	_, oks := e.(notValid)
	return okp || oks || As(e, &notValid{})
}

func IsBadGateway(e error) bool {
	_, okp := e.(*badGateway)
	_, oks := e.(badGateway)
	return okp || oks || As(e, &badGateway{})
}
func (n notFound) Is(e error) bool {
	return IsNotFound(e)
}
func (n notValid) Is(e error) bool {
	return IsNotValid(e)
}
func (n notImplemented) Is(e error) bool {
	return IsNotImplemented(e)
}
func (n notSupported) Is(e error) bool {
	return IsNotSupported(e)
}
func (b badRequest) Is(e error) bool {
	return IsBadRequest(e)
}
func (t timeout) Is(e error) bool {
	return IsTimeout(e)
}
func (u unauthorized) Is(e error) bool {
	return IsUnauthorized(e)
}
func (m methodNotAllowed) Is(e error) bool {
	return IsMethodNotAllowed(e)
}
func (f forbidden) Is(e error) bool {
	return IsForbidden(e)
}
func (b badGateway) Is(e error) bool {
	return IsBadGateway(e)
}
func (g gone) Is(e error) bool {
	return IsGone(e)
}
func (c conflict) Is(e error) bool {
	return IsConflict(e)
}
func (n notFound) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notValid) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notImplemented) Unwrap() error {
	return n.Err.Unwrap()
}
func (n notSupported) Unwrap() error {
	return n.Err.Unwrap()
}
func (b badRequest) Unwrap() error {
	return b.Err.Unwrap()
}
func (t timeout) Unwrap() error {
	return t.Err.Unwrap()
}
func (u unauthorized) Unwrap() error {
	return u.Err.Unwrap()
}
func (m methodNotAllowed) Unwrap() error {
	return m.Err.Unwrap()
}
func (f forbidden) Unwrap() error {
	return f.Err.Unwrap()
}
func (b badGateway) Unwrap() error {
	return b.Err.Unwrap()
}
func (g gone) Unwrap() error {
	return g.Err.Unwrap()
}
func (c conflict) Unwrap() error {
	return c.Err.Unwrap()
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a notFound to its underlying type Err.
func (n *notFound) As(err interface{}) bool {
	switch x := err.(type) {
	case **notFound:
		*x = n
	case *notFound:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a notValid to its underlying type Err.
func (n *notValid) As(err interface{}) bool {
	switch x := err.(type) {
	case **notValid:
		*x = n
	case *notValid:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a notImplemented to its underlying type Err.
func (n *notImplemented) As(err interface{}) bool {
	switch x := err.(type) {
	case **notImplemented:
		*x = n
	case *notImplemented:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a notSupported to its underlying type Err.
func (n *notSupported) As(err interface{}) bool {
	switch x := err.(type) {
	case **notSupported:
		*x = n
	case *notSupported:
		*x = *n
	case *Err:
		return n.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a badRequest to its underlying type Err.
func (b *badRequest) As(err interface{}) bool {
	switch x := err.(type) {
	case **badRequest:
		*x = b
	case *badRequest:
		*x = *b
	case *Err:
		return b.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a timeout to its underlying type Err.
func (t *timeout) As(err interface{}) bool {
	switch x := err.(type) {
	case **timeout:
		*x = t
	case *timeout:
		*x = *t
	case *Err:
		return t.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a unauthorized to its underlying type Err.
func (u *unauthorized) As(err interface{}) bool {
	switch x := err.(type) {
	case **unauthorized:
		*x = u
	case *unauthorized:
		*x = *u
	case *Err:
		return u.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a methodNotAllowed to its underlying type Err.
func (m *methodNotAllowed) As(err interface{}) bool {
	switch x := err.(type) {
	case **methodNotAllowed:
		*x = m
	case *methodNotAllowed:
		*x = *m
	case *Err:
		return m.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a forbidden to its underlying type Err.
func (f *forbidden) As(err interface{}) bool {
	switch x := err.(type) {
	case **forbidden:
		*x = f
	case *forbidden:
		*x = *f
	case *Err:
		return f.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a badGateway to its underlying type Err.
func (b *badGateway) As(err interface{}) bool {
	switch x := err.(type) {
	case **badGateway:
		*x = b
	case *badGateway:
		*x = *b
	case *Err:
		return b.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a gone error to its underlying type Err.
func (g *gone) As(err interface{}) bool {
	switch x := err.(type) {
	case **gone:
		*x = g
	case *gone:
		*x = *g
	case *Err:
		return g.Err.As(x)
	default:
		return false
	}
	return true
}

// As is used by the errors.As() function to coerce the method's parameter to the one of the receiver
//
//	if the underlying logic of the receiver's type can understand it.
//
// In this case we're converting a conflict error to its underlying type Err.
func (c *conflict) As(err interface{}) bool {
	switch x := err.(type) {
	case **conflict:
		*x = c
	case *conflict:
		*x = *c
	case *Err:
		return c.Err.As(x)
	default:
		return false
	}
	return true
}

// Challenge adds a challenge token to be added to the HTTP response
func (u *unauthorized) Challenge(c string) *unauthorized {
	u.challenge = c
	return u
}

// Challenge returns the challenge of the err parameter if it's an unauthorized type error
func Challenge(err error) string {
	un := unauthorized{}
	if ok := As(err, &un); ok {
		return un.challenge
	}
	return ""
}

// ErrorHandlerFn
type ErrorHandlerFn func(http.ResponseWriter, *http.Request) error

// ServeHTTP implements the http.Handler interface for the ItemHandlerFn type
func (h ErrorHandlerFn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var dat []byte
	var status int

	if err := h(w, r); err != nil {
		if status, dat = RenderErrors(r, err); status == 0 {
			status = http.StatusInternalServerError
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dat)
}

// HandleError is a generic method to return an HTTP handler that passes an error up the chain
func HandleError(e error) ErrorHandlerFn {
	return func(w http.ResponseWriter, r *http.Request) error {
		return e
	}
}

// NotFound is a generic method to return an 404 error HTTP handler that
var NotFound = ErrorHandlerFn(func(w http.ResponseWriter, r *http.Request) error {
	return NotFoundf("%s not found", r.URL.Path)
})

type Http struct {
	Code    int        `jsonld:"status,omitempty"`
	Message string     `jsonld:"message"`
	Trace   StackTrace `jsonld:"trace,omitempty"`
}

func HttpErrors(err error) []Http {
	https := make([]Http, 0)

	load := func(err error) Http {
		var trace StackTrace
		var msg string
		switch e := err.(type) {
		case *Err:
			msg = e.Error()
			if IncludeBacktrace {
				trace = e.StackTrace()
			}
		default:
			local := new(Err)
			if ok := As(err, local); ok {
				if IncludeBacktrace {
					trace = local.StackTrace()
				}
			}
			msg = err.Error()
		}

		return Http{
			Message: msg,
			Trace:   trace,
			Code:    HttpStatus(err),
		}
	}
	https = append(https, load(err))
	for {
		if err = Unwrap(err); err != nil {
			https = append(https, load(err))
		} else {
			break
		}
	}

	return https
}

func HttpStatus(e error) int {
	if IsBadRequest(e) {
		return http.StatusBadRequest
	}
	if IsUnauthorized(e) {
		return http.StatusUnauthorized
	}
	// http.StatusPaymentRequired
	if IsForbidden(e) {
		return http.StatusForbidden
	}
	if IsNotFound(e) {
		return http.StatusNotFound
	}
	if IsMethodNotAllowed(e) {
		return http.StatusMethodNotAllowed
	}
	if IsNotValid(e) {
		return http.StatusNotAcceptable
	}
	//  http.StatusProxyAuthRequired
	//  http.StatusRequestTimeout
	if IsConflict(e) {
		return http.StatusConflict
	}
	if IsGone(e) {
		return http.StatusGone
	}
	//  TODO(marius): http.StatusGone
	//  http.StatusLengthRequires
	//  http.StatusPreconditionFailed
	//  http.StatusRequestEntityTooLarge
	//  http.StatusRequestURITooLong
	//  TODO(marius): http.StatusUnsupportedMediaType
	//  http.StatusRequestedRangeNotSatisfiable
	//  http.StatusExpectationFailed
	//  http.StatusTeapot
	//  http.StatusMisdirectedRequest
	//  http.StatusUnprocessableEntity
	//  http.StatusLocked
	//  http.StatusFailedDependency
	//  http.StatusTooEarly
	//  http.StatusTooManyRequests
	//  http.StatusRequestHeaderFieldsTooLarge
	//  http.StatusUnavailableForLegalReason

	//  http.StatusInternalServerError
	if IsNotImplemented(e) {
		return http.StatusNotImplemented
	}
	if IsBadGateway(e) {
		return http.StatusBadGateway
	}
	//  http.StatusServiceUnavailable
	//  http.StatusGatewayTimeout
	if IsNotSupported(e) {
		return http.StatusHTTPVersionNotSupported
	}

	if IsTimeout(e) {
		return http.StatusGatewayTimeout
	}

	return 0
}

func errorFromStatus(status int) Error {
	switch status {
	case http.StatusBadRequest:
		return new(badRequest)
	case http.StatusUnauthorized:
		return new(unauthorized)
	// case http.StatusPaymentRequired:
	case http.StatusForbidden:
		return new(forbidden)
	case http.StatusNotFound:
		return new(notFound)
	case http.StatusMethodNotAllowed:
		return new(methodNotAllowed)
	case http.StatusNotAcceptable:
		return new(notValid)
	// case http.StatusProxyAuthRequired:
	//  case http.StatusRequestTimeout:
	case http.StatusConflict:
		return new(conflict)
	//  case http.StatusGone: // TODO(marius):
	//  case http.StatusLengthRequres:
	//  case http.StatusPreconditionFailed:
	//  case http.StatusRequestEntityTooLarge:
	//  case http.StatusRequestURITooLong:
	//  case http.StatusUnsupportedMediaType: // TODO(marius):
	//  case http.StatusRequestedRangeNotSatisfiable:
	//  case http.StatusExpectationFailed:
	//  case http.StatusTeapot:
	//  case http.StatusMisdirectedRequest:
	//  case http.StatusUnprocessableEntity:
	//  case http.StatusLocked:
	//  case http.StatusFailedDependency:
	//  case http.StatusTooEarly:
	//  case http.StatusTooManyRequests:
	//  case http.StatusRequestHeaderFieldsTooLarge:
	//  case http.StatusUnavailableForLegalReason:
	//  case http.StatusInternalServerError:
	case http.StatusNotImplemented:
		return new(notImplemented)
	case http.StatusBadGateway:
		return new(badGateway)
	// case http.StatusServiceUnavailable:
	case http.StatusHTTPVersionNotSupported:
		return new(notSupported)
	case http.StatusGatewayTimeout:
		return new(badGateway)
	case http.StatusInternalServerError:
		fallthrough
	default:
		return new(Err)
	}
}

// TODO(marius): get a proper ctxt
func ctxt(r *http.Request) jsonld.Context {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	return jsonld.Context{
		jsonld.ContextElement{
			Term: "errors",
			IRI:  jsonld.IRI(fmt.Sprintf("%s://%s/ns#errors", scheme, r.Host)),
		},
	}
}

// RenderErrors outputs the json encoded errors, with the JsonLD ctxt for current
func RenderErrors(r *http.Request, errs ...error) (int, []byte) {
	errMap := make([]Http, 0)
	var status int
	for _, err := range errs {
		more := HttpErrors(err)
		errMap = append(errMap, more...)
		status = HttpStatus(err)
	}
	var dat []byte
	var err error

	m := struct {
		Errors []Http `jsonld:"errors"`
	}{Errors: errMap}
	if dat, err = jsonld.WithContext(ctxt(r)).Marshal(m); err != nil {
		return http.StatusInternalServerError, dat
	}
	return status, dat
}
