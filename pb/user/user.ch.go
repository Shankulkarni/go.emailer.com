package user

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
	"user_profile": {
		"id":             "user_profile",
		"email":          "user_profile",
		"name":           "user_profile",
		"last_name":      "user_profile",
		"contact_number": "user_profile",
		"address":        "user_profile",
	},
}

func (m *UserProfile) PackageName() string {
	return "go_emailer_com"
}

func (m *UserProfile) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *UserProfile) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserProfile) ObjectName() string {
	return "user_profile"
}

func (m *UserProfile) Fields() []string {
	return []string{
		"id", "email", "name", "last_name", "contact_number", "address",
	}
}

func (m *UserProfile) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *UserProfile) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *UserProfile) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "email":
		return m.Email, nil
	case "name":
		return m.Name, nil
	case "last_name":
		return m.LastName, nil
	case "contact_number":
		return m.ContactNumber, nil
	case "address":
		return m.Address, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserProfile) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "email":
		return &m.Email, nil
	case "name":
		return &m.Name, nil
	case "last_name":
		return &m.LastName, nil
	case "contact_number":
		return &m.ContactNumber, nil
	case "address":
		return &m.Address, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserProfile) New(field string) error {
	switch field {
	case "id":
		return nil
	case "email":
		return nil
	case "name":
		return nil
	case "last_name":
		return nil
	case "contact_number":
		return nil
	case "address":
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *UserProfile) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "email":
		return "string"
	case "name":
		return "string"
	case "last_name":
		return "string"
	case "contact_number":
		return "string"
	case "address":
		return "json"
	default:
		return ""
	}
}

func (_ *UserProfile) GetEmptyObject() (m *UserProfile) {
	m = &UserProfile{}
	return
}

func (m *UserProfile) GetPrefix() string {
	return "usr"
}

func (m *UserProfile) GetID() string {
	return m.Id
}

func (m *UserProfile) SetID(id string) {
	m.Id = id
}

func (m *UserProfile) IsRoot() bool {
	return true
}

func (m *UserProfile) IsFlatObject(f string) bool {
	return false
}

type UserProfileStore struct {
	d driver.Driver
}

func NewUserProfileStore(d driver.Driver) UserProfileStore {
	return UserProfileStore{d: d}
}

func NewPostgresUserProfileStore(db *x.DB, usr driver.IUserInfo) UserProfileStore {
	return UserProfileStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s UserProfileStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s UserProfileStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s UserProfileStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s UserProfileStore) CreateUserProfiles(ctx context.Context, list ...*UserProfile) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &UserProfile{}, &UserProfile{}, "")
}

func (s UserProfileStore) DeleteUserProfile(ctx context.Context, cond userProfileCondition) error {
	return s.d.Delete(ctx, cond.userProfileCondToDriverCond(s.d), &UserProfile{}, &UserProfile{})
}

func (s UserProfileStore) UpdateUserProfile(ctx context.Context, req *UserProfile, cond userProfileCondition, fields []string) error {
	return s.d.Update(ctx, cond.userProfileCondToDriverCond(s.d), req, &UserProfile{}, fields...)
}

func (s UserProfileStore) ListUserProfiles(ctx context.Context, fields []string, cond userProfileCondition) ([]*UserProfile, error) {
	if len(fields) == 0 {
		fields = (&UserProfile{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.userProfileCondToDriverCond(s.d), &UserProfile{}, &UserProfile{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*UserProfile, 0, 1000)
	for res.Next(ctx) {
		obj := &UserProfile{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperUserProfile(list), nil
}

func (s UserProfileStore) GetUserProfile(ctx context.Context, fields []string, cond userProfileCondition) (*UserProfile, error) {
	if len(fields) == 0 {
		fields = (&UserProfile{}).Fields()
	}
	objList, err := s.ListUserProfiles(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type userProfileCondition interface {
	userProfileCondToDriverCond(d driver.Driver) driver.Conditioner
}

type UserProfileAnd []userProfileCondition

func (p UserProfileAnd) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.userProfileCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type UserProfileOr []userProfileCondition

func (p UserProfileOr) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.userProfileCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type UserProfileParentEq struct {
	Parent string
}

func (c UserProfileParentEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileIdEq struct {
	Id string
}

func (c UserProfileIdEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailEq struct {
	Email string
}

func (c UserProfileEmailEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameEq struct {
	Name string
}

func (c UserProfileNameEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameEq struct {
	LastName string
}

func (c UserProfileLastNameEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberEq struct {
	ContactNumber string
}

func (c UserProfileContactNumberEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressEq struct {
	Address *Address
}

func (c UserProfileAddressEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileIdNotEq struct {
	Id string
}

func (c UserProfileIdNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailNotEq struct {
	Email string
}

func (c UserProfileEmailNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameNotEq struct {
	Name string
}

func (c UserProfileNameNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameNotEq struct {
	LastName string
}

func (c UserProfileLastNameNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberNotEq struct {
	ContactNumber string
}

func (c UserProfileContactNumberNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressNotEq struct {
	Address *Address
}

func (c UserProfileAddressNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileIdGt struct {
	Id string
}

func (c UserProfileIdGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailGt struct {
	Email string
}

func (c UserProfileEmailGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameGt struct {
	Name string
}

func (c UserProfileNameGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameGt struct {
	LastName string
}

func (c UserProfileLastNameGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberGt struct {
	ContactNumber string
}

func (c UserProfileContactNumberGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressGt struct {
	Address *Address
}

func (c UserProfileAddressGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileIdLt struct {
	Id string
}

func (c UserProfileIdLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailLt struct {
	Email string
}

func (c UserProfileEmailLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameLt struct {
	Name string
}

func (c UserProfileNameLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameLt struct {
	LastName string
}

func (c UserProfileLastNameLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberLt struct {
	ContactNumber string
}

func (c UserProfileContactNumberLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressLt struct {
	Address *Address
}

func (c UserProfileAddressLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileIdGtOrEq struct {
	Id string
}

func (c UserProfileIdGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailGtOrEq struct {
	Email string
}

func (c UserProfileEmailGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameGtOrEq struct {
	Name string
}

func (c UserProfileNameGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameGtOrEq struct {
	LastName string
}

func (c UserProfileLastNameGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberGtOrEq struct {
	ContactNumber string
}

func (c UserProfileContactNumberGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressGtOrEq struct {
	Address *Address
}

func (c UserProfileAddressGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileIdLtOrEq struct {
	Id string
}

func (c UserProfileIdLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailLtOrEq struct {
	Email string
}

func (c UserProfileEmailLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameLtOrEq struct {
	Name string
}

func (c UserProfileNameLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameLtOrEq struct {
	LastName string
}

func (c UserProfileLastNameLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberLtOrEq struct {
	ContactNumber string
}

func (c UserProfileContactNumberLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileAddressLtOrEq struct {
	Address *Address
}

func (c UserProfileAddressLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &UserProfile{}, FieldMask: "address", RootDescriptor: &UserProfile{}}
}

type UserProfileDeleted struct{}

func (c UserProfileDeleted) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileNotDeleted struct{}

func (c UserProfileNotDeleted) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByEq struct {
	By string
}

func (c UserProfileCreatedByEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByNotEq struct {
	By string
}

func (c UserProfileCreatedByNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByGt struct {
	By string
}

func (c UserProfileCreatedByGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByLt struct {
	By string
}

func (c UserProfileCreatedByLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByGtOrEq struct {
	By string
}

func (c UserProfileCreatedByGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedByLtOrEq struct {
	By string
}

func (c UserProfileCreatedByLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileCreatedOnLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByEq struct {
	By string
}

func (c UserProfileUpdatedByEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByNotEq struct {
	By string
}

func (c UserProfileUpdatedByNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByGt struct {
	By string
}

func (c UserProfileUpdatedByGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByLt struct {
	By string
}

func (c UserProfileUpdatedByLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByGtOrEq struct {
	By string
}

func (c UserProfileUpdatedByGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedByLtOrEq struct {
	By string
}

func (c UserProfileUpdatedByLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileUpdatedOnLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByEq struct {
	By string
}

func (c UserProfileDeletedByEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByNotEq struct {
	By string
}

func (c UserProfileDeletedByNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnNotEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByGt struct {
	By string
}

func (c UserProfileDeletedByGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnGt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByLt struct {
	By string
}

func (c UserProfileDeletedByLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnLt) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByGtOrEq struct {
	By string
}

func (c UserProfileDeletedByGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnGtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedByLtOrEq struct {
	By string
}

func (c UserProfileDeletedByLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserProfileDeletedOnLtOrEq) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserProfile{}, RootDescriptor: &UserProfile{}}
}

type UserProfileIdIn struct {
	Id []string
}

func (c UserProfileIdIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailIn struct {
	Email []string
}

func (c UserProfileEmailIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameIn struct {
	Name []string
}

func (c UserProfileNameIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameIn struct {
	LastName []string
}

func (c UserProfileLastNameIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberIn struct {
	ContactNumber []string
}

func (c UserProfileContactNumberIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

type UserProfileIdNotIn struct {
	Id []string
}

func (c UserProfileIdNotIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserProfile{}, FieldMask: "id", RootDescriptor: &UserProfile{}}
}

type UserProfileEmailNotIn struct {
	Email []string
}

func (c UserProfileEmailNotIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "email", Value: c.Email, Operator: d, Descriptor: &UserProfile{}, FieldMask: "email", RootDescriptor: &UserProfile{}}
}

type UserProfileNameNotIn struct {
	Name []string
}

func (c UserProfileNameNotIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "name", Value: c.Name, Operator: d, Descriptor: &UserProfile{}, FieldMask: "name", RootDescriptor: &UserProfile{}}
}

type UserProfileLastNameNotIn struct {
	LastName []string
}

func (c UserProfileLastNameNotIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "last_name", Value: c.LastName, Operator: d, Descriptor: &UserProfile{}, FieldMask: "last_name", RootDescriptor: &UserProfile{}}
}

type UserProfileContactNumberNotIn struct {
	ContactNumber []string
}

func (c UserProfileContactNumberNotIn) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "contact_number", Value: c.ContactNumber, Operator: d, Descriptor: &UserProfile{}, FieldMask: "contact_number", RootDescriptor: &UserProfile{}}
}

func (c TrueCondition) userProfileCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type userProfileMapperObject struct {
	id            string
	email         string
	name          string
	lastName      string
	contactNumber string
	address       *Address
}

func (s *userProfileMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperUserProfile(rows []*UserProfile) []*UserProfile {

	combinedUserProfileMappers := map[string]*userProfileMapperObject{}

	for _, rw := range rows {

		tempUserProfile := &userProfileMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempUserProfile.id = rw.Id
		tempUserProfile.email = rw.Email
		tempUserProfile.name = rw.Name
		tempUserProfile.lastName = rw.LastName
		tempUserProfile.contactNumber = rw.ContactNumber
		tempUserProfile.address = rw.Address

		if combinedUserProfileMappers[tempUserProfile.GetUniqueIdentifier()] == nil {
			combinedUserProfileMappers[tempUserProfile.GetUniqueIdentifier()] = tempUserProfile
		}
	}

	combinedUserProfiles := []*UserProfile{}

	for _, userProfile := range combinedUserProfileMappers {
		tempUserProfile := &UserProfile{}
		tempUserProfile.Id = userProfile.id
		tempUserProfile.Email = userProfile.email
		tempUserProfile.Name = userProfile.name
		tempUserProfile.LastName = userProfile.lastName
		tempUserProfile.ContactNumber = userProfile.contactNumber
		tempUserProfile.Address = userProfile.address

		if tempUserProfile.Id == "" {
			continue
		}

		combinedUserProfiles = append(combinedUserProfiles, tempUserProfile)

	}
	return combinedUserProfiles
}
