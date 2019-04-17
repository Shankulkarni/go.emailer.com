package card

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
	"user_card_information": {
		"id":          "user_card_information",
		"gateway_id":  "user_card_information",
		"customer_id": "user_card_information",
		"cards":       "card",
	},
	"card": {
		"id":         "card",
		"card_id":    "card",
		"card_type":  "card",
		"last_four":  "card",
		"name":       "card",
		"email":      "card",
		"expire_on":  "card",
		"address":    "card",
		"is_default": "card",
		"metadata":   "card",
	},
}

func (m *UserCardInformation) PackageName() string {
	return "go_emailer_com"
}

func (m *UserCardInformation) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *UserCardInformation) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	case "cards":
		if len(m.Cards) == 0 {
			m.Cards = append(m.Cards, &Card{})
		}
		return m.Cards[0], nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserCardInformation) ObjectName() string {
	return "user_card_information"
}

func (m *UserCardInformation) Fields() []string {
	return []string{
		"id", "gateway_id", "customer_id", "cards",
	}
}

func (m *UserCardInformation) IsObject(field string) bool {
	switch field {
	case "cards":
		return true
	default:
		return false
	}
}

func (m *UserCardInformation) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "cards":
		sli := make([]driver.Descriptor, 0, len(m.Cards))
		for _, c := range m.Cards {
			sli = append(sli, c)
		}
		return sli, nil
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *UserCardInformation) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "gateway_id":
		return m.GatewayId, nil
	case "customer_id":
		return m.CustomerId, nil
	case "cards":
		return m.Cards, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserCardInformation) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "gateway_id":
		return &m.GatewayId, nil
	case "customer_id":
		return &m.CustomerId, nil
	case "cards":
		return &m.Cards, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *UserCardInformation) New(field string) error {
	switch field {
	case "id":
		return nil
	case "gateway_id":
		return nil
	case "customer_id":
		return nil
	case "cards":
		if m.Cards == nil {
			m.Cards = make([]*Card, 0)
		}
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *UserCardInformation) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "gateway_id":
		return "int32"
	case "customer_id":
		return "string"
	case "cards":
		return "repeated"
	default:
		return ""
	}
}

func (_ *UserCardInformation) GetEmptyObject() (m *UserCardInformation) {
	m = &UserCardInformation{}
	_ = m.New("cards")
	m.Cards = append(m.Cards, &Card{})
	m.Cards[0].GetEmptyObject()
	return
}

func (m *UserCardInformation) GetPrefix() string {
	return "crd"
}

func (m *UserCardInformation) GetID() string {
	return m.Id
}

func (m *UserCardInformation) SetID(id string) {
	m.Id = id
}

func (m *UserCardInformation) IsRoot() bool {
	return true
}

func (m *UserCardInformation) IsFlatObject(f string) bool {
	return false
}

func (m *Card) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Card) ObjectName() string {
	return "card"
}

func (m *Card) Fields() []string {
	return []string{
		"id", "card_id", "card_type", "last_four", "name", "email", "expire_on", "address", "is_default", "metadata",
	}
}

func (m *Card) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Card) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Card) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "card_id":
		return m.CardId, nil
	case "card_type":
		return m.CardType, nil
	case "last_four":
		return m.LastFour, nil
	case "name":
		return m.Name, nil
	case "email":
		return m.Email, nil
	case "expire_on":
		return m.ExpireOn, nil
	case "address":
		return m.Address, nil
	case "is_default":
		return m.IsDefault, nil
	case "metadata":
		return m.Metadata, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Card) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "card_id":
		return &m.CardId, nil
	case "card_type":
		return &m.CardType, nil
	case "last_four":
		return &m.LastFour, nil
	case "name":
		return &m.Name, nil
	case "email":
		return &m.Email, nil
	case "expire_on":
		return &m.ExpireOn, nil
	case "address":
		return &m.Address, nil
	case "is_default":
		return &m.IsDefault, nil
	case "metadata":
		return &m.Metadata, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Card) New(field string) error {
	switch field {
	case "id":
		return nil
	case "card_id":
		return nil
	case "card_type":
		return nil
	case "last_four":
		return nil
	case "name":
		return nil
	case "email":
		return nil
	case "expire_on":
		return nil
	case "address":
		return nil
	case "is_default":
		return nil
	case "metadata":
		if m.Metadata == nil {
			m.Metadata = map[string]string{}
		}
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Card) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "card_id":
		return "string"
	case "card_type":
		return "string"
	case "last_four":
		return "string"
	case "name":
		return "string"
	case "email":
		return "string"
	case "expire_on":
		return "string"
	case "address":
		return "json"
	case "is_default":
		return "bool"
	case "metadata":
		return "map"
	default:
		return ""
	}
}

func (_ *Card) GetEmptyObject() (m *Card) {
	m = &Card{}
	return
}

func (m *Card) GetPrefix() string {
	return "cdu"
}

func (m *Card) GetID() string {
	return m.Id
}

func (m *Card) SetID(id string) {
	m.Id = id
}

func (m *Card) IsRoot() bool {
	return false
}

func (m *Card) IsFlatObject(f string) bool {
	return false
}

type UserCardInformationStore struct {
	d driver.Driver
}

func NewUserCardInformationStore(d driver.Driver) UserCardInformationStore {
	return UserCardInformationStore{d: d}
}

func NewPostgresUserCardInformationStore(db *x.DB, usr driver.IUserInfo) UserCardInformationStore {
	return UserCardInformationStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s UserCardInformationStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s UserCardInformationStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s UserCardInformationStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s UserCardInformationStore) CreateUserCardInformations(ctx context.Context, list ...*UserCardInformation) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &UserCardInformation{}, &UserCardInformation{}, "")
}

func (s UserCardInformationStore) DeleteUserCardInformation(ctx context.Context, cond userCardInformationCondition) error {
	return s.d.Delete(ctx, cond.userCardInformationCondToDriverCond(s.d), &UserCardInformation{}, &UserCardInformation{})
}

func (s UserCardInformationStore) UpdateUserCardInformation(ctx context.Context, req *UserCardInformation, cond userCardInformationCondition, fields []string) error {
	return s.d.Update(ctx, cond.userCardInformationCondToDriverCond(s.d), req, &UserCardInformation{}, fields...)
}

func (s UserCardInformationStore) ListUserCardInformations(ctx context.Context, fields []string, cond userCardInformationCondition) ([]*UserCardInformation, error) {
	if len(fields) == 0 {
		fields = (&UserCardInformation{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.userCardInformationCondToDriverCond(s.d), &UserCardInformation{}, &UserCardInformation{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*UserCardInformation, 0, 1000)
	for res.Next(ctx) {
		obj := &UserCardInformation{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperUserCardInformation(list), nil
}

func (s UserCardInformationStore) GetUserCardInformation(ctx context.Context, fields []string, cond userCardInformationCondition) (*UserCardInformation, error) {
	if len(fields) == 0 {
		fields = (&UserCardInformation{}).Fields()
	}
	objList, err := s.ListUserCardInformations(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type userCardInformationCondition interface {
	userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner
}

type UserCardInformationAnd []userCardInformationCondition

func (p UserCardInformationAnd) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.userCardInformationCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type UserCardInformationOr []userCardInformationCondition

func (p UserCardInformationOr) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.userCardInformationCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type UserCardInformationParentEq struct {
	Parent string
}

func (c UserCardInformationParentEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdEq struct {
	Id string
}

func (c UserCardInformationIdEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdEq struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdEq struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdEq struct {
	Id string
}

func (c UserCardInformationCardIdEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdEq struct {
	CardId string
}

func (c UserCardInformationCardCardIdEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeEq struct {
	CardType string
}

func (c UserCardInformationCardCardTypeEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourEq struct {
	LastFour string
}

func (c UserCardInformationCardLastFourEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameEq struct {
	Name string
}

func (c UserCardInformationCardNameEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailEq struct {
	Email string
}

func (c UserCardInformationCardEmailEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnEq struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressEq struct {
	Address *Address
}

func (c UserCardInformationCardAddressEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultEq struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataEq struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdNotEq struct {
	Id string
}

func (c UserCardInformationIdNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdNotEq struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdNotEq struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdNotEq struct {
	Id string
}

func (c UserCardInformationCardIdNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdNotEq struct {
	CardId string
}

func (c UserCardInformationCardCardIdNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeNotEq struct {
	CardType string
}

func (c UserCardInformationCardCardTypeNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourNotEq struct {
	LastFour string
}

func (c UserCardInformationCardLastFourNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameNotEq struct {
	Name string
}

func (c UserCardInformationCardNameNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailNotEq struct {
	Email string
}

func (c UserCardInformationCardEmailNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnNotEq struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressNotEq struct {
	Address *Address
}

func (c UserCardInformationCardAddressNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultNotEq struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataNotEq struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdGt struct {
	Id string
}

func (c UserCardInformationIdGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdGt struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdGt struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdGt struct {
	Id string
}

func (c UserCardInformationCardIdGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdGt struct {
	CardId string
}

func (c UserCardInformationCardCardIdGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeGt struct {
	CardType string
}

func (c UserCardInformationCardCardTypeGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourGt struct {
	LastFour string
}

func (c UserCardInformationCardLastFourGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameGt struct {
	Name string
}

func (c UserCardInformationCardNameGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailGt struct {
	Email string
}

func (c UserCardInformationCardEmailGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnGt struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressGt struct {
	Address *Address
}

func (c UserCardInformationCardAddressGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultGt struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataGt struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdLt struct {
	Id string
}

func (c UserCardInformationIdLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdLt struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdLt struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdLt struct {
	Id string
}

func (c UserCardInformationCardIdLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdLt struct {
	CardId string
}

func (c UserCardInformationCardCardIdLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeLt struct {
	CardType string
}

func (c UserCardInformationCardCardTypeLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourLt struct {
	LastFour string
}

func (c UserCardInformationCardLastFourLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameLt struct {
	Name string
}

func (c UserCardInformationCardNameLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailLt struct {
	Email string
}

func (c UserCardInformationCardEmailLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnLt struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressLt struct {
	Address *Address
}

func (c UserCardInformationCardAddressLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultLt struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataLt struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdGtOrEq struct {
	Id string
}

func (c UserCardInformationIdGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdGtOrEq struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdGtOrEq struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdGtOrEq struct {
	Id string
}

func (c UserCardInformationCardIdGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdGtOrEq struct {
	CardId string
}

func (c UserCardInformationCardCardIdGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeGtOrEq struct {
	CardType string
}

func (c UserCardInformationCardCardTypeGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourGtOrEq struct {
	LastFour string
}

func (c UserCardInformationCardLastFourGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameGtOrEq struct {
	Name string
}

func (c UserCardInformationCardNameGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailGtOrEq struct {
	Email string
}

func (c UserCardInformationCardEmailGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnGtOrEq struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressGtOrEq struct {
	Address *Address
}

func (c UserCardInformationCardAddressGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultGtOrEq struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataGtOrEq struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdLtOrEq struct {
	Id string
}

func (c UserCardInformationIdLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdLtOrEq struct {
	GatewayId int32
}

func (c UserCardInformationGatewayIdLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdLtOrEq struct {
	CustomerId string
}

func (c UserCardInformationCustomerIdLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdLtOrEq struct {
	Id string
}

func (c UserCardInformationCardIdLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdLtOrEq struct {
	CardId string
}

func (c UserCardInformationCardCardIdLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeLtOrEq struct {
	CardType string
}

func (c UserCardInformationCardCardTypeLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourLtOrEq struct {
	LastFour string
}

func (c UserCardInformationCardLastFourLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameLtOrEq struct {
	Name string
}

func (c UserCardInformationCardNameLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailLtOrEq struct {
	Email string
}

func (c UserCardInformationCardEmailLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnLtOrEq struct {
	ExpireOn string
}

func (c UserCardInformationCardExpireOnLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardAddressLtOrEq struct {
	Address *Address
}

func (c UserCardInformationCardAddressLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "cards.address", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultLtOrEq struct {
	IsDefault bool
}

func (c UserCardInformationCardIsDefaultLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardMetadataLtOrEq struct {
	Metadata map[string]string
}

func (c UserCardInformationCardMetadataLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "cards.metadata", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeleted struct{}

func (c UserCardInformationDeleted) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationNotDeleted struct{}

func (c UserCardInformationNotDeleted) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByEq struct {
	By string
}

func (c UserCardInformationCreatedByEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByNotEq struct {
	By string
}

func (c UserCardInformationCreatedByNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByGt struct {
	By string
}

func (c UserCardInformationCreatedByGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByLt struct {
	By string
}

func (c UserCardInformationCreatedByLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByGtOrEq struct {
	By string
}

func (c UserCardInformationCreatedByGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedByLtOrEq struct {
	By string
}

func (c UserCardInformationCreatedByLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationCreatedOnLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByEq struct {
	By string
}

func (c UserCardInformationUpdatedByEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByNotEq struct {
	By string
}

func (c UserCardInformationUpdatedByNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByGt struct {
	By string
}

func (c UserCardInformationUpdatedByGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByLt struct {
	By string
}

func (c UserCardInformationUpdatedByLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByGtOrEq struct {
	By string
}

func (c UserCardInformationUpdatedByGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedByLtOrEq struct {
	By string
}

func (c UserCardInformationUpdatedByLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationUpdatedOnLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByEq struct {
	By string
}

func (c UserCardInformationDeletedByEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByNotEq struct {
	By string
}

func (c UserCardInformationDeletedByNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnNotEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByGt struct {
	By string
}

func (c UserCardInformationDeletedByGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnGt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByLt struct {
	By string
}

func (c UserCardInformationDeletedByLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnLt) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByGtOrEq struct {
	By string
}

func (c UserCardInformationDeletedByGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnGtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedByLtOrEq struct {
	By string
}

func (c UserCardInformationDeletedByLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c UserCardInformationDeletedOnLtOrEq) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &UserCardInformation{}, RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdIn struct {
	Id []string
}

func (c UserCardInformationIdIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdIn struct {
	GatewayId []int32
}

func (c UserCardInformationGatewayIdIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdIn struct {
	CustomerId []string
}

func (c UserCardInformationCustomerIdIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdIn struct {
	Id []string
}

func (c UserCardInformationCardIdIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdIn struct {
	CardId []string
}

func (c UserCardInformationCardCardIdIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeIn struct {
	CardType []string
}

func (c UserCardInformationCardCardTypeIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourIn struct {
	LastFour []string
}

func (c UserCardInformationCardLastFourIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameIn struct {
	Name []string
}

func (c UserCardInformationCardNameIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailIn struct {
	Email []string
}

func (c UserCardInformationCardEmailIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnIn struct {
	ExpireOn []string
}

func (c UserCardInformationCardExpireOnIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultIn struct {
	IsDefault []bool
}

func (c UserCardInformationCardIsDefaultIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationIdNotIn struct {
	Id []string
}

func (c UserCardInformationIdNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationGatewayIdNotIn struct {
	GatewayId []int32
}

func (c UserCardInformationGatewayIdNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "gateway_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCustomerIdNotIn struct {
	CustomerId []string
}

func (c UserCardInformationCustomerIdNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "customer_id", Value: c.CustomerId, Operator: d, Descriptor: &UserCardInformation{}, FieldMask: "customer_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIdNotIn struct {
	Id []string
}

func (c UserCardInformationCardIdNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "cards.id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardIdNotIn struct {
	CardId []string
}

func (c UserCardInformationCardCardIdNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_id", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardCardTypeNotIn struct {
	CardType []string
}

func (c UserCardInformationCardCardTypeNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "cards.card_type", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardLastFourNotIn struct {
	LastFour []string
}

func (c UserCardInformationCardLastFourNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "cards.last_four", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardNameNotIn struct {
	Name []string
}

func (c UserCardInformationCardNameNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "cards.name", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardEmailNotIn struct {
	Email []string
}

func (c UserCardInformationCardEmailNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "cards.email", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardExpireOnNotIn struct {
	ExpireOn []string
}

func (c UserCardInformationCardExpireOnNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "cards.expire_on", RootDescriptor: &UserCardInformation{}}
}

type UserCardInformationCardIsDefaultNotIn struct {
	IsDefault []bool
}

func (c UserCardInformationCardIsDefaultNotIn) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "cards.is_default", RootDescriptor: &UserCardInformation{}}
}

func (c TrueCondition) userCardInformationCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type userCardInformationMapperObject struct {
	id         string
	gatewayId  int32
	customerId string
	cards      map[string]*userCardInformationCardMapperObject
}

func (s *userCardInformationMapperObject) GetUniqueIdentifier() string {
	return s.id
}

type userCardInformationCardMapperObject struct {
	id        string
	cardId    string
	cardType  string
	lastFour  string
	name      string
	email     string
	expireOn  string
	address   *Address
	isDefault bool
	metadata  map[string]string
}

func (s *userCardInformationCardMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperUserCardInformation(rows []*UserCardInformation) []*UserCardInformation {

	combinedUserCardInformationMappers := map[string]*userCardInformationMapperObject{}

	for _, rw := range rows {

		tempUserCardInformation := &userCardInformationMapperObject{}
		tempUserCardInformationCard := &userCardInformationCardMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		if len(rw.Cards) == 0 {
			_ = rw.New("cards")
			rw.Cards = append(rw.Cards, &Card{})
			rw.Cards[0] = rw.Cards[0].GetEmptyObject()
		}
		tempUserCardInformationCard.id = rw.Cards[0].Id
		tempUserCardInformationCard.cardId = rw.Cards[0].CardId
		tempUserCardInformationCard.cardType = rw.Cards[0].CardType
		tempUserCardInformationCard.lastFour = rw.Cards[0].LastFour
		tempUserCardInformationCard.name = rw.Cards[0].Name
		tempUserCardInformationCard.email = rw.Cards[0].Email
		tempUserCardInformationCard.expireOn = rw.Cards[0].ExpireOn
		tempUserCardInformationCard.address = rw.Cards[0].Address
		tempUserCardInformationCard.isDefault = rw.Cards[0].IsDefault
		tempUserCardInformationCard.metadata = rw.Cards[0].Metadata

		tempUserCardInformation.id = rw.Id
		tempUserCardInformation.gatewayId = rw.GatewayId
		tempUserCardInformation.customerId = rw.CustomerId
		tempUserCardInformation.cards = map[string]*userCardInformationCardMapperObject{
			tempUserCardInformationCard.GetUniqueIdentifier(): tempUserCardInformationCard,
		}

		if combinedUserCardInformationMappers[tempUserCardInformation.GetUniqueIdentifier()] == nil {
			combinedUserCardInformationMappers[tempUserCardInformation.GetUniqueIdentifier()] = tempUserCardInformation
		} else {

			userCardInformationMapper := combinedUserCardInformationMappers[tempUserCardInformation.GetUniqueIdentifier()]

			if userCardInformationMapper.cards[tempUserCardInformationCard.GetUniqueIdentifier()] == nil {
				userCardInformationMapper.cards[tempUserCardInformationCard.GetUniqueIdentifier()] = tempUserCardInformationCard
			}
			combinedUserCardInformationMappers[tempUserCardInformation.GetUniqueIdentifier()] = userCardInformationMapper
		}

	}

	combinedUserCardInformations := []*UserCardInformation{}

	for _, userCardInformation := range combinedUserCardInformationMappers {
		tempUserCardInformation := &UserCardInformation{}
		tempUserCardInformation.Id = userCardInformation.id
		tempUserCardInformation.GatewayId = userCardInformation.gatewayId
		tempUserCardInformation.CustomerId = userCardInformation.customerId

		combinedUserCardInformationCards := []*Card{}

		for _, userCardInformationCard := range userCardInformation.cards {
			tempUserCardInformationCard := &Card{}
			tempUserCardInformationCard.Id = userCardInformationCard.id
			tempUserCardInformationCard.CardId = userCardInformationCard.cardId
			tempUserCardInformationCard.CardType = userCardInformationCard.cardType
			tempUserCardInformationCard.LastFour = userCardInformationCard.lastFour
			tempUserCardInformationCard.Name = userCardInformationCard.name
			tempUserCardInformationCard.Email = userCardInformationCard.email
			tempUserCardInformationCard.ExpireOn = userCardInformationCard.expireOn
			tempUserCardInformationCard.Address = userCardInformationCard.address
			tempUserCardInformationCard.IsDefault = userCardInformationCard.isDefault
			tempUserCardInformationCard.Metadata = userCardInformationCard.metadata

			if tempUserCardInformationCard.Id == "" {
				continue
			}

			combinedUserCardInformationCards = append(combinedUserCardInformationCards, tempUserCardInformationCard)

		}
		tempUserCardInformation.Cards = combinedUserCardInformationCards

		if tempUserCardInformation.Id == "" {
			continue
		}

		combinedUserCardInformations = append(combinedUserCardInformations, tempUserCardInformation)

	}
	return combinedUserCardInformations
}

func (s UserCardInformationStore) CreateCards(ctx context.Context, pid string, list ...*Card) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Card{}, &UserCardInformation{}, pid)
}

func (s UserCardInformationStore) DeleteCard(ctx context.Context, cond cardCondition) error {
	return s.d.Delete(ctx, cond.cardCondToDriverCond(s.d), &Card{}, &UserCardInformation{})
}

func (s UserCardInformationStore) UpdateCard(ctx context.Context, req *Card, cond cardCondition, fields []string) error {
	return s.d.Update(ctx, cond.cardCondToDriverCond(s.d), req, &UserCardInformation{}, fields...)
}

func (s UserCardInformationStore) ListCards(ctx context.Context, fields []string, cond cardCondition) ([]*Card, error) {
	if len(fields) == 0 {
		fields = (&Card{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.cardCondToDriverCond(s.d), &Card{}, &UserCardInformation{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Card, 0, 1000)
	for res.Next(ctx) {
		obj := &Card{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperCard(list), nil
}

func (s UserCardInformationStore) GetCard(ctx context.Context, fields []string, cond cardCondition) (*Card, error) {
	if len(fields) == 0 {
		fields = (&Card{}).Fields()
	}
	objList, err := s.ListCards(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type cardCondition interface {
	cardCondToDriverCond(d driver.Driver) driver.Conditioner
}

type CardAnd []cardCondition

func (p CardAnd) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.cardCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type CardOr []cardCondition

func (p CardOr) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.cardCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type CardParentEq struct {
	Pid string
}

func (c CardParentEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "pid", Value: c.Pid, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardIdEq struct {
	Id string
}

func (c CardIdEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdEq struct {
	CardId string
}

func (c CardCardIdEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeEq struct {
	CardType string
}

func (c CardCardTypeEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourEq struct {
	LastFour string
}

func (c CardLastFourEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameEq struct {
	Name string
}

func (c CardNameEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailEq struct {
	Email string
}

func (c CardEmailEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnEq struct {
	ExpireOn string
}

func (c CardExpireOnEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressEq struct {
	Address *Address
}

func (c CardAddressEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultEq struct {
	IsDefault bool
}

func (c CardIsDefaultEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataEq struct {
	Metadata map[string]string
}

func (c CardMetadataEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardIdNotEq struct {
	Id string
}

func (c CardIdNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdNotEq struct {
	CardId string
}

func (c CardCardIdNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeNotEq struct {
	CardType string
}

func (c CardCardTypeNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourNotEq struct {
	LastFour string
}

func (c CardLastFourNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameNotEq struct {
	Name string
}

func (c CardNameNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailNotEq struct {
	Email string
}

func (c CardEmailNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnNotEq struct {
	ExpireOn string
}

func (c CardExpireOnNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressNotEq struct {
	Address *Address
}

func (c CardAddressNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultNotEq struct {
	IsDefault bool
}

func (c CardIsDefaultNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataNotEq struct {
	Metadata map[string]string
}

func (c CardMetadataNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardIdGt struct {
	Id string
}

func (c CardIdGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdGt struct {
	CardId string
}

func (c CardCardIdGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeGt struct {
	CardType string
}

func (c CardCardTypeGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourGt struct {
	LastFour string
}

func (c CardLastFourGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameGt struct {
	Name string
}

func (c CardNameGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailGt struct {
	Email string
}

func (c CardEmailGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnGt struct {
	ExpireOn string
}

func (c CardExpireOnGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressGt struct {
	Address *Address
}

func (c CardAddressGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultGt struct {
	IsDefault bool
}

func (c CardIsDefaultGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataGt struct {
	Metadata map[string]string
}

func (c CardMetadataGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardIdLt struct {
	Id string
}

func (c CardIdLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdLt struct {
	CardId string
}

func (c CardCardIdLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeLt struct {
	CardType string
}

func (c CardCardTypeLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourLt struct {
	LastFour string
}

func (c CardLastFourLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameLt struct {
	Name string
}

func (c CardNameLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailLt struct {
	Email string
}

func (c CardEmailLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnLt struct {
	ExpireOn string
}

func (c CardExpireOnLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressLt struct {
	Address *Address
}

func (c CardAddressLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultLt struct {
	IsDefault bool
}

func (c CardIsDefaultLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataLt struct {
	Metadata map[string]string
}

func (c CardMetadataLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardIdGtOrEq struct {
	Id string
}

func (c CardIdGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdGtOrEq struct {
	CardId string
}

func (c CardCardIdGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeGtOrEq struct {
	CardType string
}

func (c CardCardTypeGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourGtOrEq struct {
	LastFour string
}

func (c CardLastFourGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameGtOrEq struct {
	Name string
}

func (c CardNameGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailGtOrEq struct {
	Email string
}

func (c CardEmailGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnGtOrEq struct {
	ExpireOn string
}

func (c CardExpireOnGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressGtOrEq struct {
	Address *Address
}

func (c CardAddressGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultGtOrEq struct {
	IsDefault bool
}

func (c CardIsDefaultGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataGtOrEq struct {
	Metadata map[string]string
}

func (c CardMetadataGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardIdLtOrEq struct {
	Id string
}

func (c CardIdLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdLtOrEq struct {
	CardId string
}

func (c CardCardIdLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeLtOrEq struct {
	CardType string
}

func (c CardCardTypeLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourLtOrEq struct {
	LastFour string
}

func (c CardLastFourLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameLtOrEq struct {
	Name string
}

func (c CardNameLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailLtOrEq struct {
	Email string
}

func (c CardEmailLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnLtOrEq struct {
	ExpireOn string
}

func (c CardExpireOnLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardAddressLtOrEq struct {
	Address *Address
}

func (c CardAddressLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "address", Value: c.Address, Operator: d, Descriptor: &Card{}, FieldMask: "address", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultLtOrEq struct {
	IsDefault bool
}

func (c CardIsDefaultLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardMetadataLtOrEq struct {
	Metadata map[string]string
}

func (c CardMetadataLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "metadata", Value: c.Metadata, Operator: d, Descriptor: &Card{}, FieldMask: "metadata", RootDescriptor: &UserCardInformation{}}
}

type CardDeleted struct{}

func (c CardDeleted) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardNotDeleted struct{}

func (c CardNotDeleted) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByEq struct {
	By string
}

func (c CardCreatedByEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByNotEq struct {
	By string
}

func (c CardCreatedByNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByGt struct {
	By string
}

func (c CardCreatedByGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByLt struct {
	By string
}

func (c CardCreatedByLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByGtOrEq struct {
	By string
}

func (c CardCreatedByGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedByLtOrEq struct {
	By string
}

func (c CardCreatedByLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardCreatedOnLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByEq struct {
	By string
}

func (c CardUpdatedByEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByNotEq struct {
	By string
}

func (c CardUpdatedByNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByGt struct {
	By string
}

func (c CardUpdatedByGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByLt struct {
	By string
}

func (c CardUpdatedByLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByGtOrEq struct {
	By string
}

func (c CardUpdatedByGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedByLtOrEq struct {
	By string
}

func (c CardUpdatedByLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardUpdatedOnLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByEq struct {
	By string
}

func (c CardDeletedByEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByNotEq struct {
	By string
}

func (c CardDeletedByNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnNotEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByGt struct {
	By string
}

func (c CardDeletedByGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnGt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByLt struct {
	By string
}

func (c CardDeletedByLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnLt) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByGtOrEq struct {
	By string
}

func (c CardDeletedByGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnGtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedByLtOrEq struct {
	By string
}

func (c CardDeletedByLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c CardDeletedOnLtOrEq) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Card{}, RootDescriptor: &UserCardInformation{}}
}

type CardIdIn struct {
	Id []string
}

func (c CardIdIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdIn struct {
	CardId []string
}

func (c CardCardIdIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeIn struct {
	CardType []string
}

func (c CardCardTypeIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourIn struct {
	LastFour []string
}

func (c CardLastFourIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameIn struct {
	Name []string
}

func (c CardNameIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailIn struct {
	Email []string
}

func (c CardEmailIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnIn struct {
	ExpireOn []string
}

func (c CardExpireOnIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultIn struct {
	IsDefault []bool
}

func (c CardIsDefaultIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

type CardIdNotIn struct {
	Id []string
}

func (c CardIdNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Card{}, FieldMask: "id", RootDescriptor: &UserCardInformation{}}
}

type CardCardIdNotIn struct {
	CardId []string
}

func (c CardCardIdNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Card{}, FieldMask: "card_id", RootDescriptor: &UserCardInformation{}}
}

type CardCardTypeNotIn struct {
	CardType []string
}

func (c CardCardTypeNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "card_type", Value: c.CardType, Operator: d, Descriptor: &Card{}, FieldMask: "card_type", RootDescriptor: &UserCardInformation{}}
}

type CardLastFourNotIn struct {
	LastFour []string
}

func (c CardLastFourNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "last_four", Value: c.LastFour, Operator: d, Descriptor: &Card{}, FieldMask: "last_four", RootDescriptor: &UserCardInformation{}}
}

type CardNameNotIn struct {
	Name []string
}

func (c CardNameNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "name", Value: c.Name, Operator: d, Descriptor: &Card{}, FieldMask: "name", RootDescriptor: &UserCardInformation{}}
}

type CardEmailNotIn struct {
	Email []string
}

func (c CardEmailNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "email", Value: c.Email, Operator: d, Descriptor: &Card{}, FieldMask: "email", RootDescriptor: &UserCardInformation{}}
}

type CardExpireOnNotIn struct {
	ExpireOn []string
}

func (c CardExpireOnNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "expire_on", Value: c.ExpireOn, Operator: d, Descriptor: &Card{}, FieldMask: "expire_on", RootDescriptor: &UserCardInformation{}}
}

type CardIsDefaultNotIn struct {
	IsDefault []bool
}

func (c CardIsDefaultNotIn) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "is_default", Value: c.IsDefault, Operator: d, Descriptor: &Card{}, FieldMask: "is_default", RootDescriptor: &UserCardInformation{}}
}

func (c TrueCondition) cardCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type cardMapperObject struct {
	id        string
	cardId    string
	cardType  string
	lastFour  string
	name      string
	email     string
	expireOn  string
	address   *Address
	isDefault bool
	metadata  map[string]string
}

func (s *cardMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperCard(rows []*Card) []*Card {

	combinedCardMappers := map[string]*cardMapperObject{}

	for _, rw := range rows {

		tempCard := &cardMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempCard.id = rw.Id
		tempCard.cardId = rw.CardId
		tempCard.cardType = rw.CardType
		tempCard.lastFour = rw.LastFour
		tempCard.name = rw.Name
		tempCard.email = rw.Email
		tempCard.expireOn = rw.ExpireOn
		tempCard.address = rw.Address
		tempCard.isDefault = rw.IsDefault
		tempCard.metadata = rw.Metadata

		if combinedCardMappers[tempCard.GetUniqueIdentifier()] == nil {
			combinedCardMappers[tempCard.GetUniqueIdentifier()] = tempCard
		}
	}

	combinedCards := []*Card{}

	for _, card := range combinedCardMappers {
		tempCard := &Card{}
		tempCard.Id = card.id
		tempCard.CardId = card.cardId
		tempCard.CardType = card.cardType
		tempCard.LastFour = card.lastFour
		tempCard.Name = card.name
		tempCard.Email = card.email
		tempCard.ExpireOn = card.expireOn
		tempCard.Address = card.address
		tempCard.IsDefault = card.isDefault
		tempCard.Metadata = card.metadata

		if tempCard.Id == "" {
			continue
		}

		combinedCards = append(combinedCards, tempCard)

	}
	return combinedCards
}
