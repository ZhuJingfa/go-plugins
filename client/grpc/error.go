package grpc

import (
	"google.golang.org/grpc/status"
	"micro/go-micro/errors"
)

func microError(err error) error {
	// no error
	switch err {
	case nil:
		return nil
	}

	// micro error
	if v, ok := err.(*errors.Error); ok {
		return v
	}

	// grpc error
	if s, ok := status.FromError(err); ok {
		return errors.Parse(s.Message())
	}

	// do nothing
	return err
}
