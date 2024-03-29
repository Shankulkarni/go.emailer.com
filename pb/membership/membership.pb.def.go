// Code generated by protoc-gen-defaults. DO NOT EDIT.

package membership

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/golang/protobuf/ptypes/empty"
)

// BaseMembershipServiceServer is the dummy implementation of the MembershipServiceServer. Embed this into your own implementation
// to add new methods without breaking builds.
type BaseMembershipServiceServer struct{}

// ListMembership is an unimplemented form of the method ListMembership
func (BaseMembershipServiceServer) ListMembership(context.Context, *empty.Empty) (*MembershipList, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}
