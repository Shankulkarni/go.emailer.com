package log

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
	"log": {
		"id": "log",
	},
}

func (m *Log) PackageName() string {
	return "go_emailer_com"
}

func (m *Log) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Log) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Log) ObjectName() string {
	return "log"
}

func (m *Log) Fields() []string {
	return []string{
		"id",
	}
}

func (m *Log) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Log) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Log) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Log) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Log) New(field string) error {
	switch field {
	case "id":
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Log) Type(field string) string {
	switch field {
	case "id":
		return "string"
	default:
		return ""
	}
}

func (_ *Log) GetEmptyObject() (m *Log) {
	m = &Log{}
	return
}

func (m *Log) GetPrefix() string {
	return "inv"
}

func (m *Log) GetID() string {
	return m.Id
}

func (m *Log) SetID(id string) {
	m.Id = id
}

func (m *Log) IsRoot() bool {
	return true
}

func (m *Log) IsFlatObject(f string) bool {
	return false
}

type LogStore struct {
	d driver.Driver
}

func NewLogStore(d driver.Driver) LogStore {
	return LogStore{d: d}
}

func NewPostgresLogStore(db *x.DB, usr driver.IUserInfo) LogStore {
	return LogStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s LogStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s LogStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s LogStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s LogStore) CreateLogs(ctx context.Context, list ...*Log) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Log{}, &Log{}, "")
}

func (s LogStore) DeleteLog(ctx context.Context, cond logCondition) error {
	return s.d.Delete(ctx, cond.logCondToDriverCond(s.d), &Log{}, &Log{})
}

func (s LogStore) UpdateLog(ctx context.Context, req *Log, cond logCondition, fields []string) error {
	return s.d.Update(ctx, cond.logCondToDriverCond(s.d), req, &Log{}, fields...)
}

func (s LogStore) ListLogs(ctx context.Context, fields []string, cond logCondition) ([]*Log, error) {
	if len(fields) == 0 {
		fields = (&Log{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.logCondToDriverCond(s.d), &Log{}, &Log{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Log, 0, 1000)
	for res.Next(ctx) {
		obj := &Log{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperLog(list), nil
}

func (s LogStore) GetLog(ctx context.Context, fields []string, cond logCondition) (*Log, error) {
	if len(fields) == 0 {
		fields = (&Log{}).Fields()
	}
	objList, err := s.ListLogs(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type logCondition interface {
	logCondToDriverCond(d driver.Driver) driver.Conditioner
}

type LogAnd []logCondition

func (p LogAnd) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.logCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type LogOr []logCondition

func (p LogOr) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.logCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type LogParentEq struct {
	Parent string
}

func (c LogParentEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogIdEq struct {
	Id string
}

func (c LogIdEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdNotEq struct {
	Id string
}

func (c LogIdNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdGt struct {
	Id string
}

func (c LogIdGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdLt struct {
	Id string
}

func (c LogIdLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdGtOrEq struct {
	Id string
}

func (c LogIdGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdLtOrEq struct {
	Id string
}

func (c LogIdLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogDeleted struct{}

func (c LogDeleted) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogNotDeleted struct{}

func (c LogNotDeleted) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByEq struct {
	By string
}

func (c LogCreatedByEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByNotEq struct {
	By string
}

func (c LogCreatedByNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByGt struct {
	By string
}

func (c LogCreatedByGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByLt struct {
	By string
}

func (c LogCreatedByLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByGtOrEq struct {
	By string
}

func (c LogCreatedByGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedByLtOrEq struct {
	By string
}

func (c LogCreatedByLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogCreatedOnLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByEq struct {
	By string
}

func (c LogUpdatedByEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByNotEq struct {
	By string
}

func (c LogUpdatedByNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByGt struct {
	By string
}

func (c LogUpdatedByGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByLt struct {
	By string
}

func (c LogUpdatedByLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByGtOrEq struct {
	By string
}

func (c LogUpdatedByGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedByLtOrEq struct {
	By string
}

func (c LogUpdatedByLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogUpdatedOnLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByEq struct {
	By string
}

func (c LogDeletedByEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByNotEq struct {
	By string
}

func (c LogDeletedByNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnNotEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByGt struct {
	By string
}

func (c LogDeletedByGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnGt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByLt struct {
	By string
}

func (c LogDeletedByLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnLt) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByGtOrEq struct {
	By string
}

func (c LogDeletedByGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnGtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedByLtOrEq struct {
	By string
}

func (c LogDeletedByLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c LogDeletedOnLtOrEq) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Log{}, RootDescriptor: &Log{}}
}

type LogIdIn struct {
	Id []string
}

func (c LogIdIn) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

type LogIdNotIn struct {
	Id []string
}

func (c LogIdNotIn) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Log{}, FieldMask: "id", RootDescriptor: &Log{}}
}

func (c TrueCondition) logCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type logMapperObject struct {
	id string
}

func (s *logMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperLog(rows []*Log) []*Log {

	combinedLogMappers := map[string]*logMapperObject{}

	for _, rw := range rows {

		tempLog := &logMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempLog.id = rw.Id

		if combinedLogMappers[tempLog.GetUniqueIdentifier()] == nil {
			combinedLogMappers[tempLog.GetUniqueIdentifier()] = tempLog
		}
	}

	combinedLogs := []*Log{}

	for _, log := range combinedLogMappers {
		tempLog := &Log{}
		tempLog.Id = log.id

		if tempLog.Id == "" {
			continue
		}

		combinedLogs = append(combinedLogs, tempLog)

	}
	return combinedLogs
}
