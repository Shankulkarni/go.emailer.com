package invoice

import (
	context "context"
	x "database/sql"
	sqrl "github.com/elgris/sqrl"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	driver "go.appointy.com/chaku/driver"
	sql "go.appointy.com/chaku/driver/sql"
	errors "go.appointy.com/chaku/errors"
)

type fieldsToTable map[string]string
type objectTable map[string]fieldsToTable

var objectTableMap = objectTable{
	"invoice": {
		"id":            "invoice",
		"user_id":       "invoice",
		"membership":    "invoice",
		"address":       "invoice",
		"user_profile":  "invoice",
		"invoice_start": "invoice",
		"invoice_end":   "invoice",
		"generated_on":  "invoice",
		"paid":          "invoice",
		"failed":        "invoice",
		"failed_reason": "invoice",
	},
}

func (m *Invoice) PackageName() string {
	return "go_emailer_com"
}

func (m *Invoice) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Invoice) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Invoice) ObjectName() string {
	return "invoice"
}

func (m *Invoice) Fields() []string {
	return []string{
		"id", "user_id", "membership", "address", "user_profile", "invoice_start", "invoice_end", "generated_on", "paid", "failed", "failed_reason",
	}
}

func (m *Invoice) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Invoice) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Invoice) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "user_id":
		return m.UserId, nil
	case "membership":
		return m.Membership, nil
	case "address":
		return m.Address, nil
	case "user_profile":
		return m.UserProfile, nil
	case "invoice_start":
		return m.InvoiceStart, nil
	case "invoice_end":
		return m.InvoiceEnd, nil
	case "generated_on":
		return m.GeneratedOn, nil
	case "paid":
		return m.Paid, nil
	case "failed":
		return m.Failed, nil
	case "failed_reason":
		return m.FailedReason, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Invoice) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "user_id":
		return &m.UserId, nil
	case "membership":
		return &m.Membership, nil
	case "address":
		return &m.Address, nil
	case "user_profile":
		return &m.UserProfile, nil
	case "invoice_start":
		return &m.InvoiceStart, nil
	case "invoice_end":
		return &m.InvoiceEnd, nil
	case "generated_on":
		return &m.GeneratedOn, nil
	case "paid":
		return &m.Paid, nil
	case "failed":
		return &m.Failed, nil
	case "failed_reason":
		return &m.FailedReason, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Invoice) New(field string) error {
	switch field {
	case "id":
		return nil
	case "user_id":
		return nil
	case "membership":
		return nil
	case "address":
		return nil
	case "user_profile":
		return nil
	case "invoice_start":
		if m.InvoiceStart == nil {
			m.InvoiceStart = &timestamp.Timestamp{}
		}
		return nil
	case "invoice_end":
		if m.InvoiceEnd == nil {
			m.InvoiceEnd = &timestamp.Timestamp{}
		}
		return nil
	case "generated_on":
		if m.GeneratedOn == nil {
			m.GeneratedOn = &timestamp.Timestamp{}
		}
		return nil
	case "paid":
		return nil
	case "failed":
		return nil
	case "failed_reason":
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Invoice) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "user_id":
		return "string"
	case "membership":
		return "json"
	case "address":
		return "json"
	case "user_profile":
		return "json"
	case "invoice_start":
		return "timestamp"
	case "invoice_end":
		return "timestamp"
	case "generated_on":
		return "timestamp"
	case "paid":
		return "bool"
	case "failed":
		return "bool"
	case "failed_reason":
		return "string"
	default:
		return ""
	}
}

func (_ *Invoice) GetEmptyObject() (m *Invoice) {
	m = &Invoice{}
	return
}

func (m *Invoice) GetPrefix() string {
	return "inv"
}

func (m *Invoice) GetID() string {
	return m.Id
}

func (m *Invoice) SetID(id string) {
	m.Id = id
}

func (m *Invoice) IsRoot() bool {
	return true
}

func (m *Invoice) IsFlatObject(f string) bool {
	return false
}

type InvoiceStore struct {
	d driver.Driver
}

func NewInvoiceStore(d driver.Driver) InvoiceStore {
	return InvoiceStore{d: d}
}

func NewPostgresInvoiceStore(db *x.DB, usr driver.IUserInfo) InvoiceStore {
	return InvoiceStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s InvoiceStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s InvoiceStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s InvoiceStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s InvoiceStore) CreateInvoices(ctx context.Context, list ...*Invoice) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Invoice{}, &Invoice{}, "")
}

func (s InvoiceStore) DeleteInvoice(ctx context.Context, cond invoiceCondition) error {
	return s.d.Delete(ctx, cond.invoiceCondToDriverCond(s.d), &Invoice{}, &Invoice{})
}

func (s InvoiceStore) UpdateInvoice(ctx context.Context, req *Invoice, cond invoiceCondition, fields []string) error {
	return s.d.Update(ctx, cond.invoiceCondToDriverCond(s.d), req, &Invoice{}, fields...)
}

func (s InvoiceStore) ListInvoices(ctx context.Context, fields []string, cond invoiceCondition) ([]*Invoice, error) {
	if len(fields) == 0 {
		fields = (&Invoice{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.invoiceCondToDriverCond(s.d), &Invoice{}, &Invoice{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Invoice, 0, 1000)
	for res.Next(ctx) {
		obj := &Invoice{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperInvoice(list), nil
}

func (s InvoiceStore) GetInvoice(ctx context.Context, fields []string, cond invoiceCondition) (*Invoice, error) {
	if len(fields) == 0 {
		fields = (&Invoice{}).Fields()
	}
	objList, err := s.ListInvoices(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type invoiceCondition interface {
	invoiceCondToDriverCond(d driver.Driver) driver.Conditioner
}

type InvoiceAnd []invoiceCondition

func (p InvoiceAnd) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.invoiceCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type InvoiceOr []invoiceCondition

func (p InvoiceOr) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.invoiceCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type InvoiceParentEq struct {
	Parent string
}

func (c InvoiceParentEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceIdEq struct {
	Id string
}

func (c InvoiceIdEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdEq struct {
	UserId string
}

func (c InvoiceUserIdEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipEq struct {
	Membership *Membership
}

func (c InvoiceMembershipEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressEq struct {
	Address *Address
}

func (c InvoiceAddressEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileEq struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartEq struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndEq struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnEq struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidEq struct {
	Paid bool
}

func (c InvoicePaidEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedEq struct {
	Failed bool
}

func (c InvoiceFailedEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonEq struct {
	FailedReason string
}

func (c InvoiceFailedReasonEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdNotEq struct {
	Id string
}

func (c InvoiceIdNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdNotEq struct {
	UserId string
}

func (c InvoiceUserIdNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipNotEq struct {
	Membership *Membership
}

func (c InvoiceMembershipNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressNotEq struct {
	Address *Address
}

func (c InvoiceAddressNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileNotEq struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartNotEq struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndNotEq struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnNotEq struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidNotEq struct {
	Paid bool
}

func (c InvoicePaidNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedNotEq struct {
	Failed bool
}

func (c InvoiceFailedNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonNotEq struct {
	FailedReason string
}

func (c InvoiceFailedReasonNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdGt struct {
	Id string
}

func (c InvoiceIdGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdGt struct {
	UserId string
}

func (c InvoiceUserIdGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipGt struct {
	Membership *Membership
}

func (c InvoiceMembershipGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressGt struct {
	Address *Address
}

func (c InvoiceAddressGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileGt struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartGt struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndGt struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnGt struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidGt struct {
	Paid bool
}

func (c InvoicePaidGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedGt struct {
	Failed bool
}

func (c InvoiceFailedGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonGt struct {
	FailedReason string
}

func (c InvoiceFailedReasonGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdLt struct {
	Id string
}

func (c InvoiceIdLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdLt struct {
	UserId string
}

func (c InvoiceUserIdLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipLt struct {
	Membership *Membership
}

func (c InvoiceMembershipLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressLt struct {
	Address *Address
}

func (c InvoiceAddressLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileLt struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartLt struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndLt struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnLt struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidLt struct {
	Paid bool
}

func (c InvoicePaidLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedLt struct {
	Failed bool
}

func (c InvoiceFailedLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonLt struct {
	FailedReason string
}

func (c InvoiceFailedReasonLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdGtOrEq struct {
	Id string
}

func (c InvoiceIdGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdGtOrEq struct {
	UserId string
}

func (c InvoiceUserIdGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipGtOrEq struct {
	Membership *Membership
}

func (c InvoiceMembershipGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressGtOrEq struct {
	Address *Address
}

func (c InvoiceAddressGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileGtOrEq struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartGtOrEq struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndGtOrEq struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnGtOrEq struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidGtOrEq struct {
	Paid bool
}

func (c InvoicePaidGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedGtOrEq struct {
	Failed bool
}

func (c InvoiceFailedGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonGtOrEq struct {
	FailedReason string
}

func (c InvoiceFailedReasonGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdLtOrEq struct {
	Id string
}

func (c InvoiceIdLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdLtOrEq struct {
	UserId string
}

func (c InvoiceUserIdLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceMembershipLtOrEq struct {
	Membership *Membership
}

func (c InvoiceMembershipLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "membership", Value: c.Membership, Operator: d, Descriptor: &Invoice{}, FieldMask: "membership", RootDescriptor: &Invoice{}}
}

type InvoiceAddressLtOrEq struct {
	Address *Address
}

func (c InvoiceAddressLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Invoice{}, FieldMask: "address", RootDescriptor: &Invoice{}}
}

type InvoiceUserProfileLtOrEq struct {
	UserProfile *UserProfile
}

func (c InvoiceUserProfileLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "user_profile", Value: c.UserProfile, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_profile", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartLtOrEq struct {
	InvoiceStart *timestamp.Timestamp
}

func (c InvoiceInvoiceStartLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndLtOrEq struct {
	InvoiceEnd *timestamp.Timestamp
}

func (c InvoiceInvoiceEndLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnLtOrEq struct {
	GeneratedOn *timestamp.Timestamp
}

func (c InvoiceGeneratedOnLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidLtOrEq struct {
	Paid bool
}

func (c InvoicePaidLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedLtOrEq struct {
	Failed bool
}

func (c InvoiceFailedLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonLtOrEq struct {
	FailedReason string
}

func (c InvoiceFailedReasonLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceDeleted struct{}

func (c InvoiceDeleted) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceNotDeleted struct{}

func (c InvoiceNotDeleted) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByEq struct {
	By string
}

func (c InvoiceCreatedByEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByNotEq struct {
	By string
}

func (c InvoiceCreatedByNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByGt struct {
	By string
}

func (c InvoiceCreatedByGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByLt struct {
	By string
}

func (c InvoiceCreatedByLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByGtOrEq struct {
	By string
}

func (c InvoiceCreatedByGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedByLtOrEq struct {
	By string
}

func (c InvoiceCreatedByLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceCreatedOnLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByEq struct {
	By string
}

func (c InvoiceUpdatedByEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByNotEq struct {
	By string
}

func (c InvoiceUpdatedByNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByGt struct {
	By string
}

func (c InvoiceUpdatedByGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByLt struct {
	By string
}

func (c InvoiceUpdatedByLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByGtOrEq struct {
	By string
}

func (c InvoiceUpdatedByGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedByLtOrEq struct {
	By string
}

func (c InvoiceUpdatedByLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceUpdatedOnLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByEq struct {
	By string
}

func (c InvoiceDeletedByEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByNotEq struct {
	By string
}

func (c InvoiceDeletedByNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnNotEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByGt struct {
	By string
}

func (c InvoiceDeletedByGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnGt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByLt struct {
	By string
}

func (c InvoiceDeletedByLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnLt) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByGtOrEq struct {
	By string
}

func (c InvoiceDeletedByGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnGtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedByLtOrEq struct {
	By string
}

func (c InvoiceDeletedByLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c InvoiceDeletedOnLtOrEq) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Invoice{}, RootDescriptor: &Invoice{}}
}

type InvoiceIdIn struct {
	Id []string
}

func (c InvoiceIdIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdIn struct {
	UserId []string
}

func (c InvoiceUserIdIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartIn struct {
	InvoiceStart []*timestamp.Timestamp
}

func (c InvoiceInvoiceStartIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndIn struct {
	InvoiceEnd []*timestamp.Timestamp
}

func (c InvoiceInvoiceEndIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnIn struct {
	GeneratedOn []*timestamp.Timestamp
}

func (c InvoiceGeneratedOnIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidIn struct {
	Paid []bool
}

func (c InvoicePaidIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedIn struct {
	Failed []bool
}

func (c InvoiceFailedIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonIn struct {
	FailedReason []string
}

func (c InvoiceFailedReasonIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

type InvoiceIdNotIn struct {
	Id []string
}

func (c InvoiceIdNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Invoice{}, FieldMask: "id", RootDescriptor: &Invoice{}}
}

type InvoiceUserIdNotIn struct {
	UserId []string
}

func (c InvoiceUserIdNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Invoice{}, FieldMask: "user_id", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceStartNotIn struct {
	InvoiceStart []*timestamp.Timestamp
}

func (c InvoiceInvoiceStartNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "invoice_start", Value: c.InvoiceStart, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_start", RootDescriptor: &Invoice{}}
}

type InvoiceInvoiceEndNotIn struct {
	InvoiceEnd []*timestamp.Timestamp
}

func (c InvoiceInvoiceEndNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "invoice_end", Value: c.InvoiceEnd, Operator: d, Descriptor: &Invoice{}, FieldMask: "invoice_end", RootDescriptor: &Invoice{}}
}

type InvoiceGeneratedOnNotIn struct {
	GeneratedOn []*timestamp.Timestamp
}

func (c InvoiceGeneratedOnNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "generated_on", Value: c.GeneratedOn, Operator: d, Descriptor: &Invoice{}, FieldMask: "generated_on", RootDescriptor: &Invoice{}}
}

type InvoicePaidNotIn struct {
	Paid []bool
}

func (c InvoicePaidNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "paid", Value: c.Paid, Operator: d, Descriptor: &Invoice{}, FieldMask: "paid", RootDescriptor: &Invoice{}}
}

type InvoiceFailedNotIn struct {
	Failed []bool
}

func (c InvoiceFailedNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "failed", Value: c.Failed, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed", RootDescriptor: &Invoice{}}
}

type InvoiceFailedReasonNotIn struct {
	FailedReason []string
}

func (c InvoiceFailedReasonNotIn) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "failed_reason", Value: c.FailedReason, Operator: d, Descriptor: &Invoice{}, FieldMask: "failed_reason", RootDescriptor: &Invoice{}}
}

func (c TrueCondition) invoiceCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type invoiceMapperObject struct {
	id           string
	userId       string
	membership   *Membership
	address      *Address
	userProfile  *UserProfile
	invoiceStart *timestamp.Timestamp
	invoiceEnd   *timestamp.Timestamp
	generatedOn  *timestamp.Timestamp
	paid         bool
	failed       bool
	failedReason string
}

func (s *invoiceMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperInvoice(rows []*Invoice) []*Invoice {

	combinedInvoiceMappers := map[string]*invoiceMapperObject{}

	for _, rw := range rows {

		tempInvoice := &invoiceMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempInvoice.id = rw.Id
		tempInvoice.userId = rw.UserId
		tempInvoice.membership = rw.Membership
		tempInvoice.address = rw.Address
		tempInvoice.userProfile = rw.UserProfile
		tempInvoice.invoiceStart = rw.InvoiceStart
		tempInvoice.invoiceEnd = rw.InvoiceEnd
		tempInvoice.generatedOn = rw.GeneratedOn
		tempInvoice.paid = rw.Paid
		tempInvoice.failed = rw.Failed
		tempInvoice.failedReason = rw.FailedReason

		if combinedInvoiceMappers[tempInvoice.GetUniqueIdentifier()] == nil {
			combinedInvoiceMappers[tempInvoice.GetUniqueIdentifier()] = tempInvoice
		}
	}

	combinedInvoices := []*Invoice{}

	for _, invoice := range combinedInvoiceMappers {
		tempInvoice := &Invoice{}
		tempInvoice.Id = invoice.id
		tempInvoice.UserId = invoice.userId
		tempInvoice.Membership = invoice.membership
		tempInvoice.Address = invoice.address
		tempInvoice.UserProfile = invoice.userProfile
		tempInvoice.InvoiceStart = invoice.invoiceStart
		tempInvoice.InvoiceEnd = invoice.invoiceEnd
		tempInvoice.GeneratedOn = invoice.generatedOn
		tempInvoice.Paid = invoice.paid
		tempInvoice.Failed = invoice.failed
		tempInvoice.FailedReason = invoice.failedReason

		if tempInvoice.Id == "" {
			continue
		}

		combinedInvoices = append(combinedInvoices, tempInvoice)

	}
	return combinedInvoices
}
