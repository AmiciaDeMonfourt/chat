package errors

import (
	"fmt"

	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCError struct {
	Reason   string `json:"reason"`
	Desc     string `json:"desc"`
	MetaData string `json:"metadata"`
}

func (e *GRPCError) Error() string {
	return e.Reason + e.Desc + e.MetaData
}

func NewGRPC(reason, domain, desc string, code codes.Code) error {
	md := &errdetails.ErrorInfo{
		Domain:   domain,
		Metadata: map[string]string{"error": desc},
	}

	st, err := status.New(code, reason).WithDetails(md)
	if err != nil {
		return fmt.Errorf("unexpected error: %v", err)
	}

	return st.Err()
}
