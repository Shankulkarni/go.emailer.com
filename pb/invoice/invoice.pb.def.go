// Code generated by protoc-gen-defaults. DO NOT EDIT.

package invoice

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.emailer.com/pb/user"
)

// BaseInvoiceServiceServer is the dummy implementation of the InvoiceServiceServer. Embed this into your own implementation
// to add new methods without breaking builds.
type BaseInvoiceServiceServer struct{}

// GetInvoices is an unimplemented form of the method GetInvoices
func (BaseInvoiceServiceServer) GetInvoices(context.Context, *GetInvoicesRequest) (*InvoiceList, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}

// GetUpcomingInvoice is an unimplemented form of the method GetUpcomingInvoice
func (BaseInvoiceServiceServer) GetUpcomingInvoice(context.Context, *user.UserIdentifier) (*Invoice, error) {
	return nil, status.Error(codes.Unimplemented, codes.Unimplemented.String())
}