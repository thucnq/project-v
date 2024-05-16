// Code generated by MockGen. DO NOT EDIT.
// Source: index.go

// Package notification is a generated GoMock package.
package notification

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockINotification is a mock of INotification interface
type MockINotification struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationMockRecorder
}

// MockINotificationMockRecorder is the mock recorder for MockINotification
type MockINotificationMockRecorder struct {
	mock *MockINotification
}

// NewMockINotification creates a new mock instance
func NewMockINotification(ctrl *gomock.Controller) *MockINotification {
	mock := &MockINotification{ctrl: ctrl}
	mock.recorder = &MockINotificationMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINotification) EXPECT() *MockINotificationMockRecorder {
	return m.recorder
}

// GetData mocks base method
func (m *MockINotification) GetData() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetData")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// GetData indicates an expected call of GetData
func (mr *MockINotificationMockRecorder) GetData() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetData", reflect.TypeOf((*MockINotification)(nil).GetData))
}

// GetKey mocks base method
func (m *MockINotification) GetKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetKey indicates an expected call of GetKey
func (mr *MockINotificationMockRecorder) GetKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockINotification)(nil).GetKey))
}

// MockIDataDelivery is a mock of IDataDelivery interface
type MockIDataDelivery struct {
	ctrl     *gomock.Controller
	recorder *MockIDataDeliveryMockRecorder
}

// MockIDataDeliveryMockRecorder is the mock recorder for MockIDataDelivery
type MockIDataDeliveryMockRecorder struct {
	mock *MockIDataDelivery
}

// NewMockIDataDelivery creates a new mock instance
func NewMockIDataDelivery(ctrl *gomock.Controller) *MockIDataDelivery {
	mock := &MockIDataDelivery{ctrl: ctrl}
	mock.recorder = &MockIDataDeliveryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIDataDelivery) EXPECT() *MockIDataDeliveryMockRecorder {
	return m.recorder
}

// GetExchangeName mocks base method
func (m *MockIDataDelivery) GetExchangeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExchangeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetExchangeName indicates an expected call of GetExchangeName
func (mr *MockIDataDeliveryMockRecorder) GetExchangeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExchangeName", reflect.TypeOf((*MockIDataDelivery)(nil).GetExchangeName))
}

// GetRoutingKey mocks base method
func (m *MockIDataDelivery) GetRoutingKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoutingKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetRoutingKey indicates an expected call of GetRoutingKey
func (mr *MockIDataDeliveryMockRecorder) GetRoutingKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoutingKey", reflect.TypeOf((*MockIDataDelivery)(nil).GetRoutingKey))
}

// Marshal mocks base method
func (m *MockIDataDelivery) Marshal() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Marshal")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Marshal indicates an expected call of Marshal
func (mr *MockIDataDeliveryMockRecorder) Marshal() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Marshal", reflect.TypeOf((*MockIDataDelivery)(nil).Marshal))
}

// MockINotificationVisitor is a mock of INotificationVisitor interface
type MockINotificationVisitor struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationVisitorMockRecorder
}

// MockINotificationVisitorMockRecorder is the mock recorder for MockINotificationVisitor
type MockINotificationVisitorMockRecorder struct {
	mock *MockINotificationVisitor
}

// NewMockINotificationVisitor creates a new mock instance
func NewMockINotificationVisitor(ctrl *gomock.Controller) *MockINotificationVisitor {
	mock := &MockINotificationVisitor{ctrl: ctrl}
	mock.recorder = &MockINotificationVisitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINotificationVisitor) EXPECT() *MockINotificationVisitorMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockINotificationVisitor) Send(arg0 INotification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Send", arg0)
}

// Send indicates an expected call of Send
func (mr *MockINotificationVisitorMockRecorder) Send(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockINotificationVisitor)(nil).Send), arg0)
}

// MockIProducer is a mock of IProducer interface
type MockIProducer struct {
	ctrl     *gomock.Controller
	recorder *MockIProducerMockRecorder
}

// MockIProducerMockRecorder is the mock recorder for MockIProducer
type MockIProducerMockRecorder struct {
	mock *MockIProducer
}

// NewMockIProducer creates a new mock instance
func NewMockIProducer(ctrl *gomock.Controller) *MockIProducer {
	mock := &MockIProducer{ctrl: ctrl}
	mock.recorder = &MockIProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIProducer) EXPECT() *MockIProducerMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockIProducer) Publish(data IDataDelivery) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish
func (mr *MockIProducerMockRecorder) Publish(data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIProducer)(nil).Publish), data)
}

// MockIEventFactory is a mock of IEventFactory interface
type MockIEventFactory struct {
	ctrl     *gomock.Controller
	recorder *MockIEventFactoryMockRecorder
}

// MockIEventFactoryMockRecorder is the mock recorder for MockIEventFactory
type MockIEventFactoryMockRecorder struct {
	mock *MockIEventFactory
}

// NewMockIEventFactory creates a new mock instance
func NewMockIEventFactory(ctrl *gomock.Controller) *MockIEventFactory {
	mock := &MockIEventFactory{ctrl: ctrl}
	mock.recorder = &MockIEventFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventFactory) EXPECT() *MockIEventFactoryMockRecorder {
	return m.recorder
}

// Publish mocks base method
func (m *MockIEventFactory) Publish(noti INotification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Publish", noti)
}

// Publish indicates an expected call of Publish
func (mr *MockIEventFactoryMockRecorder) Publish(noti interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockIEventFactory)(nil).Publish), noti)
}



// WithRabbitmq mocks base method
func (m *MockIEventFactory) WithRabbitmq(rabbitmqClient IProducer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithRabbitmq", rabbitmqClient)
}

// WithRabbitmq indicates an expected call of WithRabbitmq
func (mr *MockIEventFactoryMockRecorder) WithRabbitmq(rabbitmqClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithRabbitmq", reflect.TypeOf((*MockIEventFactory)(nil).WithRabbitmq), rabbitmqClient)
}

// WithEventBusRedis mocks base method
func (m *MockIEventFactory) WithEventBusRedis(redisClient IProducer) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "WithEventBusRedis", redisClient)
}

// WithEventBusRedis indicates an expected call of WithEventBusRedis
func (mr *MockIEventFactoryMockRecorder) WithEventBusRedis(redisClient interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithEventBusRedis", reflect.TypeOf((*MockIEventFactory)(nil).WithEventBusRedis), redisClient)
}

// MockIEventDispatcher is a mock of IEventDispatcher interface
type MockIEventDispatcher struct {
	ctrl     *gomock.Controller
	recorder *MockIEventDispatcherMockRecorder
}

// MockIEventDispatcherMockRecorder is the mock recorder for MockIEventDispatcher
type MockIEventDispatcherMockRecorder struct {
	mock *MockIEventDispatcher
}

// NewMockIEventDispatcher creates a new mock instance
func NewMockIEventDispatcher(ctrl *gomock.Controller) *MockIEventDispatcher {
	mock := &MockIEventDispatcher{ctrl: ctrl}
	mock.recorder = &MockIEventDispatcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIEventDispatcher) EXPECT() *MockIEventDispatcherMockRecorder {
	return m.recorder
}

// AddEvent mocks base method
func (m *MockIEventDispatcher) AddEvent(arg0 INotification) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddEvent", arg0)
}

// AddEvent indicates an expected call of AddEvent
func (mr *MockIEventDispatcherMockRecorder) AddEvent(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEvent", reflect.TypeOf((*MockIEventDispatcher)(nil).AddEvent), arg0)
}
