// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: Reader,Writer,Closer,ReadWriter,ReadCloser,ReadWriteCloser)

// Package nio is a generated GoMock package.
package nio

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockReader is a mock of Reader interface
type MockReader struct {
	ctrl     *gomock.Controller
	recorder *MockReaderMockRecorder
}

// MockReaderMockRecorder is the mock recorder for MockReader
type MockReaderMockRecorder struct {
	mock *MockReader
}

// NewMockReader creates a new mock instance
func NewMockReader(ctrl *gomock.Controller) *MockReader {
	mock := &MockReader{ctrl: ctrl}
	mock.recorder = &MockReaderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReader) EXPECT() *MockReaderMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReader) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReaderMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReader)(nil).Read), arg0)
}

// MockWriter is a mock of Writer interface
type MockWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWriterMockRecorder
}

// MockWriterMockRecorder is the mock recorder for MockWriter
type MockWriterMockRecorder struct {
	mock *MockWriter
}

// NewMockWriter creates a new mock instance
func NewMockWriter(ctrl *gomock.Controller) *MockWriter {
	mock := &MockWriter{ctrl: ctrl}
	mock.recorder = &MockWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWriter) EXPECT() *MockWriterMockRecorder {
	return m.recorder
}

// Write mocks base method
func (m *MockWriter) Write(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWriter)(nil).Write), arg0)
}

// MockCloser is a mock of Closer interface
type MockCloser struct {
	ctrl     *gomock.Controller
	recorder *MockCloserMockRecorder
}

// MockCloserMockRecorder is the mock recorder for MockCloser
type MockCloserMockRecorder struct {
	mock *MockCloser
}

// NewMockCloser creates a new mock instance
func NewMockCloser(ctrl *gomock.Controller) *MockCloser {
	mock := &MockCloser{ctrl: ctrl}
	mock.recorder = &MockCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCloser) EXPECT() *MockCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockCloser) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockCloserMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloser)(nil).Close))
}

// MockReadWriter is a mock of ReadWriter interface
type MockReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriterMockRecorder
}

// MockReadWriterMockRecorder is the mock recorder for MockReadWriter
type MockReadWriterMockRecorder struct {
	mock *MockReadWriter
}

// NewMockReadWriter creates a new mock instance
func NewMockReadWriter(ctrl *gomock.Controller) *MockReadWriter {
	mock := &MockReadWriter{ctrl: ctrl}
	mock.recorder = &MockReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadWriter) EXPECT() *MockReadWriterMockRecorder {
	return m.recorder
}

// Read mocks base method
func (m *MockReadWriter) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReadWriterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadWriter)(nil).Read), arg0)
}

// Write mocks base method
func (m *MockReadWriter) Write(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockReadWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockReadWriter)(nil).Write), arg0)
}

// MockReadCloser is a mock of ReadCloser interface
type MockReadCloser struct {
	ctrl     *gomock.Controller
	recorder *MockReadCloserMockRecorder
}

// MockReadCloserMockRecorder is the mock recorder for MockReadCloser
type MockReadCloserMockRecorder struct {
	mock *MockReadCloser
}

// NewMockReadCloser creates a new mock instance
func NewMockReadCloser(ctrl *gomock.Controller) *MockReadCloser {
	mock := &MockReadCloser{ctrl: ctrl}
	mock.recorder = &MockReadCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadCloser) EXPECT() *MockReadCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockReadCloser) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReadCloserMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadCloser)(nil).Close))
}

// Read mocks base method
func (m *MockReadCloser) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReadCloserMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadCloser)(nil).Read), arg0)
}

// MockReadWriteCloser is a mock of ReadWriteCloser interface
type MockReadWriteCloser struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriteCloserMockRecorder
}

// MockReadWriteCloserMockRecorder is the mock recorder for MockReadWriteCloser
type MockReadWriteCloserMockRecorder struct {
	mock *MockReadWriteCloser
}

// NewMockReadWriteCloser creates a new mock instance
func NewMockReadWriteCloser(ctrl *gomock.Controller) *MockReadWriteCloser {
	mock := &MockReadWriteCloser{ctrl: ctrl}
	mock.recorder = &MockReadWriteCloserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockReadWriteCloser) EXPECT() *MockReadWriteCloserMockRecorder {
	return m.recorder
}

// Close mocks base method
func (m *MockReadWriteCloser) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockReadWriteCloserMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockReadWriteCloser)(nil).Close))
}

// Read mocks base method
func (m *MockReadWriteCloser) Read(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockReadWriteCloserMockRecorder) Read(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadWriteCloser)(nil).Read), arg0)
}

// Write mocks base method
func (m *MockReadWriteCloser) Write(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockReadWriteCloserMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockReadWriteCloser)(nil).Write), arg0)
}
