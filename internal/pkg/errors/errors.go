package errors

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/confluentinc/ccloudapis/core/v1"
)

/*
 * Invariants:
 * - Confluent SDK (http package) always returns a corev1.Error.
 * - Pkg always return an HTTP Error constant (top of this file)
 *
 * Error Flow:
 * - API error responses (json) are parsed into corev1.Error objects.
 *   - Note: API returns 404s for unauthorized resources, so HTTP package has to remap 404 -> 401 where appropriate.
 * - Pkg call ConvertAPIError() to transforms corev1.Error into HTTP Error constants
 */

type NotAuthenticatedError error
type KafkaError error
type UnknownKafkaContextError error

var (
	ErrNotImplemented = fmt.Errorf("not implemented")
	ErrIncorrectAuth  = fmt.Errorf("incorrect auth")
	ErrUnauthorized   = fmt.Errorf("unauthorized")
	ErrExpiredToken   = fmt.Errorf("expired")
	ErrMalformedToken = fmt.Errorf("malformed")
	ErrNotFound       = fmt.Errorf("not found")
	ErrNoContext      = fmt.Errorf("context not set")
	ErrNoKafkaContext = fmt.Errorf("kafka not set")
)

// ConvertAPIError transforms a corev1.Error into one of the standard errors if it matches.
// TODO: the SDK should expose typed errors so clients don't need to do this nonsense
func ConvertAPIError(err error) error {
	if e, ok := errors.Cause(err).(*v1.Error); ok {
		switch e.Message {
		// these messages are returned by the API itself
		case "token is expired":
			return ErrExpiredToken
		case "malformed token":
			return ErrMalformedToken
		// except this one.. its the special case of errUnauthorized from http/auth.go
		case "unauthorized":
			return ErrUnauthorized
		// except this one.. its the special case of errNotFound from http/client.go
		case "cluster not found":
			return ErrNotFound
			// TODO: assert invariant for default case: we're missing an corev1.Error -> HTTP Error constant mapping
		}
	}
	return err
}

func Wrap(err error, msg string) error {
	return errors.Wrap(err, msg)
}

func Wrapf(err error, fmt string, args ...interface{}) error {
	return errors.Wrapf(err, fmt, args...)
}

func New(msg string) error {
	return errors.New(msg)
}

func Errorf(fmt string, args ...interface{}) error {
	return errors.Errorf(fmt, args...)
}
