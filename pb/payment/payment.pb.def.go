// Code generated by protoc-gen-defaults. DO NOT EDIT.

package payment

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
)

// BasePaymentServiceServer is the dummy implementation of the PaymentServiceServer. Embed this into your own implementation
// to add new methods without breaking builds.
type BasePaymentServiceServer struct{}

// SubscribeMembership is an unimplemented form of the method SubscribeMembership
func (BasePaymentServiceServer) SubscribeMembership(context.Context, *SubscribeMembershipRequest) (*Subscription, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

// CancelMembership is an unimplemented form of the method CancelMembership
func (BasePaymentServiceServer) CancelMembership(context.Context, *CancelMembershipRequest) (*empty.Empty, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
