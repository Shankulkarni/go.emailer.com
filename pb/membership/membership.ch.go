package membership

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
	"membership": {
		"id":                "membership",
		"title":             "membership",
		"description":       "membership",
		"duration_in_month": "membership",
		"cost":              "membership",
		"quantity":          "membership",
	},
}

func (m *Membership) PackageName() string {
	return "go_emailer_com"
}

func (m *Membership) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Membership) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Membership) ObjectName() string {
	return "membership"
}

func (m *Membership) Fields() []string {
	return []string{
		"id", "title", "description", "duration_in_month", "cost", "quantity",
	}
}

func (m *Membership) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Membership) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Membership) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "title":
		return m.Title, nil
	case "description":
		return m.Description, nil
	case "duration_in_month":
		return m.DurationInMonth, nil
	case "cost":
		return m.Cost, nil
	case "quantity":
		return m.Quantity, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Membership) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "title":
		return &m.Title, nil
	case "description":
		return &m.Description, nil
	case "duration_in_month":
		return &m.DurationInMonth, nil
	case "cost":
		return &m.Cost, nil
	case "quantity":
		return &m.Quantity, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Membership) New(field string) error {
	switch field {
	case "id":
		return nil
	case "title":
		return nil
	case "description":
		return nil
	case "duration_in_month":
		return nil
	case "cost":
		return nil
	case "quantity":
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Membership) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "title":
		return "string"
	case "description":
		return "string"
	case "duration_in_month":
		return "int32"
	case "cost":
		return "int32"
	case "quantity":
		return "int64"
	default:
		return ""
	}
}

func (_ *Membership) GetEmptyObject() (m *Membership) {
	m = &Membership{}
	return
}

func (m *Membership) GetPrefix() string {
	return "mem"
}

func (m *Membership) GetID() string {
	return m.Id
}

func (m *Membership) SetID(id string) {
	m.Id = id
}

func (m *Membership) IsRoot() bool {
	return true
}

func (m *Membership) IsFlatObject(f string) bool {
	return false
}

type MembershipStore struct {
	d driver.Driver
}

func NewMembershipStore(d driver.Driver) MembershipStore {
	return MembershipStore{d: d}
}

func NewPostgresMembershipStore(db *x.DB, usr driver.IUserInfo) MembershipStore {
	return MembershipStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s MembershipStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s MembershipStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s MembershipStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s MembershipStore) CreateMemberships(ctx context.Context, list ...*Membership) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Membership{}, &Membership{}, "")
}

func (s MembershipStore) DeleteMembership(ctx context.Context, cond membershipCondition) error {
	return s.d.Delete(ctx, cond.membershipCondToDriverCond(s.d), &Membership{}, &Membership{})
}

func (s MembershipStore) UpdateMembership(ctx context.Context, req *Membership, cond membershipCondition, fields []string) error {
	return s.d.Update(ctx, cond.membershipCondToDriverCond(s.d), req, &Membership{}, fields...)
}

func (s MembershipStore) ListMemberships(ctx context.Context, fields []string, cond membershipCondition) ([]*Membership, error) {
	if len(fields) == 0 {
		fields = (&Membership{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.membershipCondToDriverCond(s.d), &Membership{}, &Membership{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Membership, 0, 1000)
	for res.Next(ctx) {
		obj := &Membership{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperMembership(list), nil
}

func (s MembershipStore) GetMembership(ctx context.Context, fields []string, cond membershipCondition) (*Membership, error) {
	if len(fields) == 0 {
		fields = (&Membership{}).Fields()
	}
	objList, err := s.ListMemberships(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type membershipCondition interface {
	membershipCondToDriverCond(d driver.Driver) driver.Conditioner
}

type MembershipAnd []membershipCondition

func (p MembershipAnd) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.membershipCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type MembershipOr []membershipCondition

func (p MembershipOr) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.membershipCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type MembershipParentEq struct {
	Parent string
}

func (c MembershipParentEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipIdEq struct {
	Id string
}

func (c MembershipIdEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleEq struct {
	Title string
}

func (c MembershipTitleEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionEq struct {
	Description string
}

func (c MembershipDescriptionEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthEq struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostEq struct {
	Cost int32
}

func (c MembershipCostEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityEq struct {
	Quantity int64
}

func (c MembershipQuantityEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdNotEq struct {
	Id string
}

func (c MembershipIdNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleNotEq struct {
	Title string
}

func (c MembershipTitleNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionNotEq struct {
	Description string
}

func (c MembershipDescriptionNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthNotEq struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostNotEq struct {
	Cost int32
}

func (c MembershipCostNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityNotEq struct {
	Quantity int64
}

func (c MembershipQuantityNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdGt struct {
	Id string
}

func (c MembershipIdGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleGt struct {
	Title string
}

func (c MembershipTitleGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionGt struct {
	Description string
}

func (c MembershipDescriptionGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthGt struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostGt struct {
	Cost int32
}

func (c MembershipCostGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityGt struct {
	Quantity int64
}

func (c MembershipQuantityGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdLt struct {
	Id string
}

func (c MembershipIdLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleLt struct {
	Title string
}

func (c MembershipTitleLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionLt struct {
	Description string
}

func (c MembershipDescriptionLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthLt struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostLt struct {
	Cost int32
}

func (c MembershipCostLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityLt struct {
	Quantity int64
}

func (c MembershipQuantityLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdGtOrEq struct {
	Id string
}

func (c MembershipIdGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleGtOrEq struct {
	Title string
}

func (c MembershipTitleGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionGtOrEq struct {
	Description string
}

func (c MembershipDescriptionGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthGtOrEq struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostGtOrEq struct {
	Cost int32
}

func (c MembershipCostGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityGtOrEq struct {
	Quantity int64
}

func (c MembershipQuantityGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdLtOrEq struct {
	Id string
}

func (c MembershipIdLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleLtOrEq struct {
	Title string
}

func (c MembershipTitleLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionLtOrEq struct {
	Description string
}

func (c MembershipDescriptionLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthLtOrEq struct {
	DurationInMonth int32
}

func (c MembershipDurationInMonthLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostLtOrEq struct {
	Cost int32
}

func (c MembershipCostLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityLtOrEq struct {
	Quantity int64
}

func (c MembershipQuantityLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipDeleted struct{}

func (c MembershipDeleted) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipNotDeleted struct{}

func (c MembershipNotDeleted) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByEq struct {
	By string
}

func (c MembershipCreatedByEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByNotEq struct {
	By string
}

func (c MembershipCreatedByNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByGt struct {
	By string
}

func (c MembershipCreatedByGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByLt struct {
	By string
}

func (c MembershipCreatedByLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByGtOrEq struct {
	By string
}

func (c MembershipCreatedByGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedByLtOrEq struct {
	By string
}

func (c MembershipCreatedByLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipCreatedOnLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByEq struct {
	By string
}

func (c MembershipUpdatedByEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByNotEq struct {
	By string
}

func (c MembershipUpdatedByNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByGt struct {
	By string
}

func (c MembershipUpdatedByGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByLt struct {
	By string
}

func (c MembershipUpdatedByLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByGtOrEq struct {
	By string
}

func (c MembershipUpdatedByGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedByLtOrEq struct {
	By string
}

func (c MembershipUpdatedByLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipUpdatedOnLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByEq struct {
	By string
}

func (c MembershipDeletedByEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByNotEq struct {
	By string
}

func (c MembershipDeletedByNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnNotEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByGt struct {
	By string
}

func (c MembershipDeletedByGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnGt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByLt struct {
	By string
}

func (c MembershipDeletedByLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnLt) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByGtOrEq struct {
	By string
}

func (c MembershipDeletedByGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnGtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedByLtOrEq struct {
	By string
}

func (c MembershipDeletedByLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c MembershipDeletedOnLtOrEq) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Membership{}, RootDescriptor: &Membership{}}
}

type MembershipIdIn struct {
	Id []string
}

func (c MembershipIdIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleIn struct {
	Title []string
}

func (c MembershipTitleIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionIn struct {
	Description []string
}

func (c MembershipDescriptionIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthIn struct {
	DurationInMonth []int32
}

func (c MembershipDurationInMonthIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostIn struct {
	Cost []int32
}

func (c MembershipCostIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityIn struct {
	Quantity []int64
}

func (c MembershipQuantityIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

type MembershipIdNotIn struct {
	Id []string
}

func (c MembershipIdNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Membership{}, FieldMask: "id", RootDescriptor: &Membership{}}
}

type MembershipTitleNotIn struct {
	Title []string
}

func (c MembershipTitleNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "title", Value: c.Title, Operator: d, Descriptor: &Membership{}, FieldMask: "title", RootDescriptor: &Membership{}}
}

type MembershipDescriptionNotIn struct {
	Description []string
}

func (c MembershipDescriptionNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "description", Value: c.Description, Operator: d, Descriptor: &Membership{}, FieldMask: "description", RootDescriptor: &Membership{}}
}

type MembershipDurationInMonthNotIn struct {
	DurationInMonth []int32
}

func (c MembershipDurationInMonthNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "duration_in_month", Value: c.DurationInMonth, Operator: d, Descriptor: &Membership{}, FieldMask: "duration_in_month", RootDescriptor: &Membership{}}
}

type MembershipCostNotIn struct {
	Cost []int32
}

func (c MembershipCostNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "cost", Value: c.Cost, Operator: d, Descriptor: &Membership{}, FieldMask: "cost", RootDescriptor: &Membership{}}
}

type MembershipQuantityNotIn struct {
	Quantity []int64
}

func (c MembershipQuantityNotIn) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "quantity", Value: c.Quantity, Operator: d, Descriptor: &Membership{}, FieldMask: "quantity", RootDescriptor: &Membership{}}
}

func (c TrueCondition) membershipCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type membershipMapperObject struct {
	id              string
	title           string
	description     string
	durationInMonth int32
	cost            int32
	quantity        int64
}

func (s *membershipMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperMembership(rows []*Membership) []*Membership {

	combinedMembershipMappers := map[string]*membershipMapperObject{}

	for _, rw := range rows {

		tempMembership := &membershipMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempMembership.id = rw.Id
		tempMembership.title = rw.Title
		tempMembership.description = rw.Description
		tempMembership.durationInMonth = rw.DurationInMonth
		tempMembership.cost = rw.Cost
		tempMembership.quantity = rw.Quantity

		if combinedMembershipMappers[tempMembership.GetUniqueIdentifier()] == nil {
			combinedMembershipMappers[tempMembership.GetUniqueIdentifier()] = tempMembership
		}
	}

	combinedMemberships := []*Membership{}

	for _, membership := range combinedMembershipMappers {
		tempMembership := &Membership{}
		tempMembership.Id = membership.id
		tempMembership.Title = membership.title
		tempMembership.Description = membership.description
		tempMembership.DurationInMonth = membership.durationInMonth
		tempMembership.Cost = membership.cost
		tempMembership.Quantity = membership.quantity

		if tempMembership.Id == "" {
			continue
		}

		combinedMemberships = append(combinedMemberships, tempMembership)

	}
	return combinedMemberships
}
