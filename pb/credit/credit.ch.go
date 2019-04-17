package credit

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
	"credit": {
		"id":        "credit",
		"credits":   "credit",
		"expire_on": "credit",
	},
}

func (m *Credit) PackageName() string {
	return "go_emailer_com"
}

func (m *Credit) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Credit) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Credit) ObjectName() string {
	return "credit"
}

func (m *Credit) Fields() []string {
	return []string{
		"id", "credits", "expire_on",
	}
}

func (m *Credit) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Credit) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Credit) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "credits":
		return m.Credits, nil
	case "expire_on":
		return m.ExpireOn, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Credit) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "credits":
		return &m.Credits, nil
	case "expire_on":
		return &m.ExpireOn, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Credit) New(field string) error {
	switch field {
	case "id":
		return nil
	case "credits":
		return nil
	case "expire_on":
		if m.ExpireOn == nil {
			m.ExpireOn = &timestamp.Timestamp{}
		}
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Credit) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "credits":
		return "int64"
	case "expire_on":
		return "timestamp"
	default:
		return ""
	}
}

func (_ *Credit) GetEmptyObject() (m *Credit) {
	m = &Credit{}
	return
}

func (m *Credit) GetPrefix() string {
	return "ced"
}

func (m *Credit) GetID() string {
	return m.Id
}

func (m *Credit) SetID(id string) {
	m.Id = id
}

func (m *Credit) IsRoot() bool {
	return true
}

func (m *Credit) IsFlatObject(f string) bool {
	return false
}

type CreditStore struct {
	d driver.Driver
}

func NewCreditStore(d driver.Driver) CreditStore {
	return CreditStore{d: d}
}

func NewPostgresCreditStore(db *x.DB, usr driver.IUserInfo) CreditStore {
	return CreditStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s CreditStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s CreditStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s CreditStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s CreditStore) CreateCredits(ctx context.Context, list ...*Credit) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Credit{}, &Credit{}, "")
}

func (s CreditStore) DeleteCredit(ctx context.Context, cond creditCondition) error {
	return s.d.Delete(ctx, cond.creditCondToDriverCond(s.d), &Credit{}, &Credit{})
}

func (s CreditStore) UpdateCredit(ctx context.Context, req *Credit, cond creditCondition, fields []string) error {
	return s.d.Update(ctx, cond.creditCondToDriverCond(s.d), req, &Credit{}, fields...)
}

func (s CreditStore) ListCredits(ctx context.Context, fields []string, cond creditCondition) ([]*Credit, error) {
	if len(fields) == 0 {
		fields = (&Credit{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.creditCondToDriverCond(s.d), &Credit{}, &Credit{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Credit, 0, 1000)
	for res.Next(ctx) {
		obj := &Credit{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperCredit(list), nil
}

func (s CreditStore) GetCredit(ctx context.Context, fields []string, cond creditCondition) (*Credit, error) {
	if len(fields) == 0 {
		fields = (&Credit{}).Fields()
	}
	objList, err := s.ListCredits(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type creditCondition interface {
	creditCondToDriverCond(d driver.Driver) driver.Conditioner
}

type CreditAnd []creditCondition

func (p CreditAnd) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.creditCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type CreditOr []creditCondition

func (p CreditOr) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.creditCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type CreditParentEq struct {
	Parent string
}

func (c CreditParentEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditIdEq struct {
	Id string
}

func (c CreditIdEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsEq struct {
	Credits int64
}

func (c CreditCreditsEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnEq struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdNotEq struct {
	Id string
}

func (c CreditIdNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsNotEq struct {
	Credits int64
}

func (c CreditCreditsNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnNotEq struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdGt struct {
	Id string
}

func (c CreditIdGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsGt struct {
	Credits int64
}

func (c CreditCreditsGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnGt struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdLt struct {
	Id string
}

func (c CreditIdLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsLt struct {
	Credits int64
}

func (c CreditCreditsLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnLt struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdGtOrEq struct {
	Id string
}

func (c CreditIdGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsGtOrEq struct {
	Credits int64
}

func (c CreditCreditsGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnGtOrEq struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdLtOrEq struct {
	Id string
}

func (c CreditIdLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsLtOrEq struct {
	Credits int64
}

func (c CreditCreditsLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnLtOrEq struct {
	ExpireOn *timestamp.Timestamp
}

func (c CreditExpireOnLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditDeleted struct{}

func (c CreditDeleted) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditNotDeleted struct{}

func (c CreditNotDeleted) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByEq struct {
	By string
}

func (c CreditCreatedByEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByNotEq struct {
	By string
}

func (c CreditCreatedByNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByGt struct {
	By string
}

func (c CreditCreatedByGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByLt struct {
	By string
}

func (c CreditCreatedByLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByGtOrEq struct {
	By string
}

func (c CreditCreatedByGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedByLtOrEq struct {
	By string
}

func (c CreditCreatedByLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditCreatedOnLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByEq struct {
	By string
}

func (c CreditUpdatedByEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByNotEq struct {
	By string
}

func (c CreditUpdatedByNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByGt struct {
	By string
}

func (c CreditUpdatedByGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByLt struct {
	By string
}

func (c CreditUpdatedByLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByGtOrEq struct {
	By string
}

func (c CreditUpdatedByGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedByLtOrEq struct {
	By string
}

func (c CreditUpdatedByLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditUpdatedOnLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByEq struct {
	By string
}

func (c CreditDeletedByEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByNotEq struct {
	By string
}

func (c CreditDeletedByNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnNotEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByGt struct {
	By string
}

func (c CreditDeletedByGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnGt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByLt struct {
	By string
}

func (c CreditDeletedByLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnLt) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByGtOrEq struct {
	By string
}

func (c CreditDeletedByGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnGtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedByLtOrEq struct {
	By string
}

func (c CreditDeletedByLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CreditDeletedOnLtOrEq) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Credit{}, RootDescriptor: &Credit{}}
}

type CreditIdIn struct {
	Id []string
}

func (c CreditIdIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsIn struct {
	Credits []int64
}

func (c CreditCreditsIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnIn struct {
	ExpireOn []*timestamp.Timestamp
}

func (c CreditExpireOnIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

type CreditIdNotIn struct {
	Id []string
}

func (c CreditIdNotIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Credit{}, FieldMask: "id", RootDescriptor: &Credit{}}
}

type CreditCreditsNotIn struct {
	Credits []int64
}

func (c CreditCreditsNotIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "credits", Value: c.Credits, Operator: d, Descriptor: &Credit{}, FieldMask: "credits", RootDescriptor: &Credit{}}
}

type CreditExpireOnNotIn struct {
	ExpireOn []*timestamp.Timestamp
}

func (c CreditExpireOnNotIn) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Credit{}, FieldMask: "expire_on", RootDescriptor: &Credit{}}
}

func (c TrueCondition) creditCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type creditMapperObject struct {
	id       string
	credits  int64
	expireOn *timestamp.Timestamp
}

func (s *creditMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperCredit(rows []*Credit) []*Credit {

	combinedCreditMappers := map[string]*creditMapperObject{}

	for _, rw := range rows {

		tempCredit := &creditMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempCredit.id = rw.Id
		tempCredit.credits = rw.Credits
		tempCredit.expireOn = rw.ExpireOn

		if combinedCreditMappers[tempCredit.GetUniqueIdentifier()] == nil {
			combinedCreditMappers[tempCredit.GetUniqueIdentifier()] = tempCredit
		}
	}

	combinedCredits := []*Credit{}

	for _, credit := range combinedCreditMappers {
		tempCredit := &Credit{}
		tempCredit.Id = credit.id
		tempCredit.Credits = credit.credits
		tempCredit.ExpireOn = credit.expireOn

		if tempCredit.Id == "" {
			continue
		}

		combinedCredits = append(combinedCredits, tempCredit)

	}
	return combinedCredits
}
