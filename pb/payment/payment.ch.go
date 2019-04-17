package payment

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
	"subscription": {
		"id":                      "subscription",
		"user_id":                 "subscription",
		"gateway_id":              "subscription",
		"gateway_customer_id":     "subscription",
		"gateway_subscription_id": "subscription",
		"membership_id":           "subscription",
		"is_recurring":            "subscription",
		"card_id":                 "subscription",
		"valid_till":              "subscription",
		"subscription_start_date": "subscription",
		"current_period_start":    "subscription",
		"current_period_end":      "subscription",
	},
}

func (m *Subscription) PackageName() string {
	return "go_emailer_com"
}

func (m *Subscription) TableOfObject(f, s string) string {
	return objectTableMap[f][s]
}

func (m *Subscription) GetDescriptorsOf(f string) (driver.Descriptor, error) {
	switch f {
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Subscription) ObjectName() string {
	return "subscription"
}

func (m *Subscription) Fields() []string {
	return []string{
		"id", "user_id", "gateway_id", "gateway_customer_id", "gateway_subscription_id", "membership_id", "is_recurring", "card_id", "valid_till", "subscription_start_date", "current_period_start", "current_period_end",
	}
}

func (m *Subscription) IsObject(field string) bool {
	switch field {
	default:
		return false
	}
}

func (m *Subscription) ValuerSlice(field string) ([]driver.Descriptor, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	default:
		return []driver.Descriptor{}, errors.ErrInvalidField
	}
}

func (m *Subscription) Valuer(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return m.Id, nil
	case "user_id":
		return m.UserId, nil
	case "gateway_id":
		return m.GatewayId, nil
	case "gateway_customer_id":
		return m.GatewayCustomerId, nil
	case "gateway_subscription_id":
		return m.GatewaySubscriptionId, nil
	case "membership_id":
		return m.MembershipId, nil
	case "is_recurring":
		return m.IsRecurring, nil
	case "card_id":
		return m.CardId, nil
	case "valid_till":
		return m.ValidTill, nil
	case "subscription_start_date":
		return m.SubscriptionStartDate, nil
	case "current_period_start":
		return m.CurrentPeriodStart, nil
	case "current_period_end":
		return m.CurrentPeriodEnd, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Subscription) Addresser(field string) (interface{}, error) {
	if m == nil {
		return nil, nil
	}
	switch field {
	case "id":
		return &m.Id, nil
	case "user_id":
		return &m.UserId, nil
	case "gateway_id":
		return &m.GatewayId, nil
	case "gateway_customer_id":
		return &m.GatewayCustomerId, nil
	case "gateway_subscription_id":
		return &m.GatewaySubscriptionId, nil
	case "membership_id":
		return &m.MembershipId, nil
	case "is_recurring":
		return &m.IsRecurring, nil
	case "card_id":
		return &m.CardId, nil
	case "valid_till":
		return &m.ValidTill, nil
	case "subscription_start_date":
		return &m.SubscriptionStartDate, nil
	case "current_period_start":
		return &m.CurrentPeriodStart, nil
	case "current_period_end":
		return &m.CurrentPeriodEnd, nil
	default:
		return nil, errors.ErrInvalidField
	}
}

func (m *Subscription) New(field string) error {
	switch field {
	case "id":
		return nil
	case "user_id":
		return nil
	case "gateway_id":
		return nil
	case "gateway_customer_id":
		return nil
	case "gateway_subscription_id":
		return nil
	case "membership_id":
		return nil
	case "is_recurring":
		return nil
	case "card_id":
		return nil
	case "valid_till":
		if m.ValidTill == nil {
			m.ValidTill = &timestamp.Timestamp{}
		}
		return nil
	case "subscription_start_date":
		if m.SubscriptionStartDate == nil {
			m.SubscriptionStartDate = &timestamp.Timestamp{}
		}
		return nil
	case "current_period_start":
		if m.CurrentPeriodStart == nil {
			m.CurrentPeriodStart = &timestamp.Timestamp{}
		}
		return nil
	case "current_period_end":
		if m.CurrentPeriodEnd == nil {
			m.CurrentPeriodEnd = &timestamp.Timestamp{}
		}
		return nil
	default:
		return errors.ErrInvalidField
	}
}

func (m *Subscription) Type(field string) string {
	switch field {
	case "id":
		return "string"
	case "user_id":
		return "string"
	case "gateway_id":
		return "string"
	case "gateway_customer_id":
		return "string"
	case "gateway_subscription_id":
		return "string"
	case "membership_id":
		return "string"
	case "is_recurring":
		return "bool"
	case "card_id":
		return "string"
	case "valid_till":
		return "timestamp"
	case "subscription_start_date":
		return "timestamp"
	case "current_period_start":
		return "timestamp"
	case "current_period_end":
		return "timestamp"
	default:
		return ""
	}
}

func (_ *Subscription) GetEmptyObject() (m *Subscription) {
	m = &Subscription{}
	return
}

func (m *Subscription) GetPrefix() string {
	return "sub"
}

func (m *Subscription) GetID() string {
	return m.Id
}

func (m *Subscription) SetID(id string) {
	m.Id = id
}

func (m *Subscription) IsRoot() bool {
	return true
}

func (m *Subscription) IsFlatObject(f string) bool {
	return false
}

type SubscriptionStore struct {
	d driver.Driver
}

func NewSubscriptionStore(d driver.Driver) SubscriptionStore {
	return SubscriptionStore{d: d}
}

func NewPostgresSubscriptionStore(db *x.DB, usr driver.IUserInfo) SubscriptionStore {
	return SubscriptionStore{
		&sql.Sql{DB: db, UserInfo: usr, Placeholder: sqrl.Dollar},
	}
}

func (s SubscriptionStore) StartTransaction(ctx context.Context) error {
	return s.d.StartTransaction(ctx)
}

func (s SubscriptionStore) CommitTransaction(ctx context.Context) error {
	return s.d.CommitTransaction(ctx)
}

func (s SubscriptionStore) RollBackTransaction(ctx context.Context) error {
	return s.d.RollBackTransaction(ctx)
}

func (s SubscriptionStore) CreateSubscriptions(ctx context.Context, list ...*Subscription) ([]string, error) {
	vv := make([]driver.Descriptor, len(list))
	for i := range list {
		vv[i] = list[i]
	}
	return s.d.Insert(ctx, vv, &Subscription{}, &Subscription{}, "")
}

func (s SubscriptionStore) DeleteSubscription(ctx context.Context, cond subscriptionCondition) error {
	return s.d.Delete(ctx, cond.subscriptionCondToDriverCond(s.d), &Subscription{}, &Subscription{})
}

func (s SubscriptionStore) UpdateSubscription(ctx context.Context, req *Subscription, cond subscriptionCondition, fields []string) error {
	return s.d.Update(ctx, cond.subscriptionCondToDriverCond(s.d), req, &Subscription{}, fields...)
}

func (s SubscriptionStore) ListSubscriptions(ctx context.Context, fields []string, cond subscriptionCondition) ([]*Subscription, error) {
	if len(fields) == 0 {
		fields = (&Subscription{}).Fields()
	}
	res, err := s.d.Get(ctx, cond.subscriptionCondToDriverCond(s.d), &Subscription{}, &Subscription{}, fields...)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	list := make([]*Subscription, 0, 1000)
	for res.Next(ctx) {
		obj := &Subscription{}
		if err := res.Scan(obj); err != nil {
			return nil, err
		}
		list = append(list, obj)
	}

	if err := res.Close(); err != nil {
		return nil, err
	}

	return MapperSubscription(list), nil
}

func (s SubscriptionStore) GetSubscription(ctx context.Context, fields []string, cond subscriptionCondition) (*Subscription, error) {
	if len(fields) == 0 {
		fields = (&Subscription{}).Fields()
	}
	objList, err := s.ListSubscriptions(ctx, fields, cond)
	if len(objList) == 0 && err == nil {
		err = errors.ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return objList[0], nil
}

type TrueCondition struct{}

type subscriptionCondition interface {
	subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner
}

type SubscriptionAnd []subscriptionCondition

func (p SubscriptionAnd) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.subscriptionCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type SubscriptionOr []subscriptionCondition

func (p SubscriptionOr) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	dc := make([]driver.Conditioner, 0, len(p))
	for _, c := range p {
		dc = append(dc, c.subscriptionCondToDriverCond(d))
	}
	return driver.And{Conditioners: dc, Operator: d}
}

type SubscriptionParentEq struct {
	Parent string
}

func (c SubscriptionParentEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "parent", Value: c.Parent, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionIdEq struct {
	Id string
}

func (c SubscriptionIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdEq struct {
	UserId string
}

func (c SubscriptionUserIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdEq struct {
	GatewayId string
}

func (c SubscriptionGatewayIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdEq struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdEq struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdEq struct {
	MembershipId string
}

func (c SubscriptionMembershipIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringEq struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdEq struct {
	CardId string
}

func (c SubscriptionCardIdEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillEq struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateEq struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartEq struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndEq struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdNotEq struct {
	Id string
}

func (c SubscriptionIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdNotEq struct {
	UserId string
}

func (c SubscriptionUserIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdNotEq struct {
	GatewayId string
}

func (c SubscriptionGatewayIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdNotEq struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdNotEq struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdNotEq struct {
	MembershipId string
}

func (c SubscriptionMembershipIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringNotEq struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdNotEq struct {
	CardId string
}

func (c SubscriptionCardIdNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillNotEq struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateNotEq struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartNotEq struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndNotEq struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdGt struct {
	Id string
}

func (c SubscriptionIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdGt struct {
	UserId string
}

func (c SubscriptionUserIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdGt struct {
	GatewayId string
}

func (c SubscriptionGatewayIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdGt struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdGt struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdGt struct {
	MembershipId string
}

func (c SubscriptionMembershipIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringGt struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdGt struct {
	CardId string
}

func (c SubscriptionCardIdGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillGt struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateGt struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartGt struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndGt struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdLt struct {
	Id string
}

func (c SubscriptionIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdLt struct {
	UserId string
}

func (c SubscriptionUserIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdLt struct {
	GatewayId string
}

func (c SubscriptionGatewayIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdLt struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdLt struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdLt struct {
	MembershipId string
}

func (c SubscriptionMembershipIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringLt struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdLt struct {
	CardId string
}

func (c SubscriptionCardIdLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillLt struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateLt struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartLt struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndLt struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdGtOrEq struct {
	Id string
}

func (c SubscriptionIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdGtOrEq struct {
	UserId string
}

func (c SubscriptionUserIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdGtOrEq struct {
	GatewayId string
}

func (c SubscriptionGatewayIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdGtOrEq struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdGtOrEq struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdGtOrEq struct {
	MembershipId string
}

func (c SubscriptionMembershipIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringGtOrEq struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdGtOrEq struct {
	CardId string
}

func (c SubscriptionCardIdGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillGtOrEq struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateGtOrEq struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartGtOrEq struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndGtOrEq struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdLtOrEq struct {
	Id string
}

func (c SubscriptionIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdLtOrEq struct {
	UserId string
}

func (c SubscriptionUserIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdLtOrEq struct {
	GatewayId string
}

func (c SubscriptionGatewayIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdLtOrEq struct {
	GatewayCustomerId string
}

func (c SubscriptionGatewayCustomerIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdLtOrEq struct {
	GatewaySubscriptionId string
}

func (c SubscriptionGatewaySubscriptionIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdLtOrEq struct {
	MembershipId string
}

func (c SubscriptionMembershipIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringLtOrEq struct {
	IsRecurring bool
}

func (c SubscriptionIsRecurringLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdLtOrEq struct {
	CardId string
}

func (c SubscriptionCardIdLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillLtOrEq struct {
	ValidTill *timestamp.Timestamp
}

func (c SubscriptionValidTillLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateLtOrEq struct {
	SubscriptionStartDate *timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartLtOrEq struct {
	CurrentPeriodStart *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndLtOrEq struct {
	CurrentPeriodEnd *timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionDeleted struct{}

func (c SubscriptionDeleted) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: true, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionNotDeleted struct{}

func (c SubscriptionNotDeleted) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "is_deleted", Value: false, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByEq struct {
	By string
}

func (c SubscriptionCreatedByEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByNotEq struct {
	By string
}

func (c SubscriptionCreatedByNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByGt struct {
	By string
}

func (c SubscriptionCreatedByGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnGt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByLt struct {
	By string
}

func (c SubscriptionCreatedByLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnLt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByGtOrEq struct {
	By string
}

func (c SubscriptionCreatedByGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedByLtOrEq struct {
	By string
}

func (c SubscriptionCreatedByLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionCreatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionCreatedOnLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "created_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByEq struct {
	By string
}

func (c SubscriptionUpdatedByEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByNotEq struct {
	By string
}

func (c SubscriptionUpdatedByNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByGt struct {
	By string
}

func (c SubscriptionUpdatedByGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnGt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByLt struct {
	By string
}

func (c SubscriptionUpdatedByLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnLt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByGtOrEq struct {
	By string
}

func (c SubscriptionUpdatedByGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedByLtOrEq struct {
	By string
}

func (c SubscriptionUpdatedByLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionUpdatedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionUpdatedOnLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "updated_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByEq struct {
	By string
}

func (c SubscriptionDeletedByEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Eq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByNotEq struct {
	By string
}

func (c SubscriptionDeletedByNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnNotEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnNotEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByGt struct {
	By string
}

func (c SubscriptionDeletedByGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnGt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnGt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Gt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByLt struct {
	By string
}

func (c SubscriptionDeletedByLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnLt struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnLt) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.Lt{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByGtOrEq struct {
	By string
}

func (c SubscriptionDeletedByGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnGtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnGtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.GtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedByLtOrEq struct {
	By string
}

func (c SubscriptionDeletedByLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_by", Value: c.By, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionDeletedOnLtOrEq struct {
	On *timestamp.Timestamp
}

func (c SubscriptionDeletedOnLtOrEq) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.LtOrEq{Key: "deleted_on", Value: c.On, Operator: d, Descriptor: &Subscription{}, RootDescriptor: &Subscription{}}
}

type SubscriptionIdIn struct {
	Id []string
}

func (c SubscriptionIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdIn struct {
	UserId []string
}

func (c SubscriptionUserIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdIn struct {
	GatewayId []string
}

func (c SubscriptionGatewayIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdIn struct {
	GatewayCustomerId []string
}

func (c SubscriptionGatewayCustomerIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdIn struct {
	GatewaySubscriptionId []string
}

func (c SubscriptionGatewaySubscriptionIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdIn struct {
	MembershipId []string
}

func (c SubscriptionMembershipIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringIn struct {
	IsRecurring []bool
}

func (c SubscriptionIsRecurringIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdIn struct {
	CardId []string
}

func (c SubscriptionCardIdIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillIn struct {
	ValidTill []*timestamp.Timestamp
}

func (c SubscriptionValidTillIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateIn struct {
	SubscriptionStartDate []*timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartIn struct {
	CurrentPeriodStart []*timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndIn struct {
	CurrentPeriodEnd []*timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.In{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

type SubscriptionIdNotIn struct {
	Id []string
}

func (c SubscriptionIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "id", Value: c.Id, Operator: d, Descriptor: &Subscription{}, FieldMask: "id", RootDescriptor: &Subscription{}}
}

type SubscriptionUserIdNotIn struct {
	UserId []string
}

func (c SubscriptionUserIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "user_id", Value: c.UserId, Operator: d, Descriptor: &Subscription{}, FieldMask: "user_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayIdNotIn struct {
	GatewayId []string
}

func (c SubscriptionGatewayIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "gateway_id", Value: c.GatewayId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewayCustomerIdNotIn struct {
	GatewayCustomerId []string
}

func (c SubscriptionGatewayCustomerIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "gateway_customer_id", Value: c.GatewayCustomerId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_customer_id", RootDescriptor: &Subscription{}}
}

type SubscriptionGatewaySubscriptionIdNotIn struct {
	GatewaySubscriptionId []string
}

func (c SubscriptionGatewaySubscriptionIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "gateway_subscription_id", Value: c.GatewaySubscriptionId, Operator: d, Descriptor: &Subscription{}, FieldMask: "gateway_subscription_id", RootDescriptor: &Subscription{}}
}

type SubscriptionMembershipIdNotIn struct {
	MembershipId []string
}

func (c SubscriptionMembershipIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "membership_id", Value: c.MembershipId, Operator: d, Descriptor: &Subscription{}, FieldMask: "membership_id", RootDescriptor: &Subscription{}}
}

type SubscriptionIsRecurringNotIn struct {
	IsRecurring []bool
}

func (c SubscriptionIsRecurringNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "is_recurring", Value: c.IsRecurring, Operator: d, Descriptor: &Subscription{}, FieldMask: "is_recurring", RootDescriptor: &Subscription{}}
}

type SubscriptionCardIdNotIn struct {
	CardId []string
}

func (c SubscriptionCardIdNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "card_id", Value: c.CardId, Operator: d, Descriptor: &Subscription{}, FieldMask: "card_id", RootDescriptor: &Subscription{}}
}

type SubscriptionValidTillNotIn struct {
	ValidTill []*timestamp.Timestamp
}

func (c SubscriptionValidTillNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "valid_till", Value: c.ValidTill, Operator: d, Descriptor: &Subscription{}, FieldMask: "valid_till", RootDescriptor: &Subscription{}}
}

type SubscriptionSubscriptionStartDateNotIn struct {
	SubscriptionStartDate []*timestamp.Timestamp
}

func (c SubscriptionSubscriptionStartDateNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "subscription_start_date", Value: c.SubscriptionStartDate, Operator: d, Descriptor: &Subscription{}, FieldMask: "subscription_start_date", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodStartNotIn struct {
	CurrentPeriodStart []*timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodStartNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "current_period_start", Value: c.CurrentPeriodStart, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_start", RootDescriptor: &Subscription{}}
}

type SubscriptionCurrentPeriodEndNotIn struct {
	CurrentPeriodEnd []*timestamp.Timestamp
}

func (c SubscriptionCurrentPeriodEndNotIn) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.NotIn{Key: "current_period_end", Value: c.CurrentPeriodEnd, Operator: d, Descriptor: &Subscription{}, FieldMask: "current_period_end", RootDescriptor: &Subscription{}}
}

func (c TrueCondition) subscriptionCondToDriverCond(d driver.Driver) driver.Conditioner {
	return driver.TrueCondition{Operator: d}
}

type subscriptionMapperObject struct {
	id                    string
	userId                string
	gatewayId             string
	gatewayCustomerId     string
	gatewaySubscriptionId string
	membershipId          string
	isRecurring           bool
	cardId                string
	validTill             *timestamp.Timestamp
	subscriptionStartDate *timestamp.Timestamp
	currentPeriodStart    *timestamp.Timestamp
	currentPeriodEnd      *timestamp.Timestamp
}

func (s *subscriptionMapperObject) GetUniqueIdentifier() string {
	return s.id
}

func MapperSubscription(rows []*Subscription) []*Subscription {

	combinedSubscriptionMappers := map[string]*subscriptionMapperObject{}

	for _, rw := range rows {

		tempSubscription := &subscriptionMapperObject{}

		if rw == nil {
			rw = rw.GetEmptyObject()
		}
		tempSubscription.id = rw.Id
		tempSubscription.userId = rw.UserId
		tempSubscription.gatewayId = rw.GatewayId
		tempSubscription.gatewayCustomerId = rw.GatewayCustomerId
		tempSubscription.gatewaySubscriptionId = rw.GatewaySubscriptionId
		tempSubscription.membershipId = rw.MembershipId
		tempSubscription.isRecurring = rw.IsRecurring
		tempSubscription.cardId = rw.CardId
		tempSubscription.validTill = rw.ValidTill
		tempSubscription.subscriptionStartDate = rw.SubscriptionStartDate
		tempSubscription.currentPeriodStart = rw.CurrentPeriodStart
		tempSubscription.currentPeriodEnd = rw.CurrentPeriodEnd

		if combinedSubscriptionMappers[tempSubscription.GetUniqueIdentifier()] == nil {
			combinedSubscriptionMappers[tempSubscription.GetUniqueIdentifier()] = tempSubscription
		}
	}

	combinedSubscriptions := []*Subscription{}

	for _, subscription := range combinedSubscriptionMappers {
		tempSubscription := &Subscription{}
		tempSubscription.Id = subscription.id
		tempSubscription.UserId = subscription.userId
		tempSubscription.GatewayId = subscription.gatewayId
		tempSubscription.GatewayCustomerId = subscription.gatewayCustomerId
		tempSubscription.GatewaySubscriptionId = subscription.gatewaySubscriptionId
		tempSubscription.MembershipId = subscription.membershipId
		tempSubscription.IsRecurring = subscription.isRecurring
		tempSubscription.CardId = subscription.cardId
		tempSubscription.ValidTill = subscription.validTill
		tempSubscription.SubscriptionStartDate = subscription.subscriptionStartDate
		tempSubscription.CurrentPeriodStart = subscription.currentPeriodStart
		tempSubscription.CurrentPeriodEnd = subscription.currentPeriodEnd

		if tempSubscription.Id == "" {
			continue
		}

		combinedSubscriptions = append(combinedSubscriptions, tempSubscription)

	}
	return combinedSubscriptions
}
