package database

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/stretchr/testify/mock"
)

// MockDB is a mock implementation of PgxIface using testify/mock
type MockDB struct {
	mock.Mock
}

// Query mocks the Query method
func (m *MockDB) Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error) {
	callArgs := m.Called(ctx, sql, args)
	if callArgs.Get(0) == nil {
		return nil, callArgs.Error(1)
	}
	return callArgs.Get(0).(pgx.Rows), callArgs.Error(1)
}

// QueryRow mocks the QueryRow method
func (m *MockDB) QueryRow(ctx context.Context, query string, args ...any) pgx.Row {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Get(0).(pgx.Row)
}

// Exec mocks the Exec method
func (m *MockDB) Exec(ctx context.Context, query string, args ...any) (pgconn.CommandTag, error) {
	callArgs := m.Called(ctx, query, args)
	return callArgs.Get(0).(pgconn.CommandTag), callArgs.Error(1)
}

// Close mocks the Close method
func (m *MockDB) Close(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

// MockRow is a mock implementation of pgx.Row
type MockRow struct {
	mock.Mock
}

// Scan mocks the Scan method
func (m *MockRow) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

// MockRows is a mock implementation of pgx.Rows
type MockRows struct {
	mock.Mock
	CurrentIndex int
	Data         [][]any
}

// NewMockRows creates a new MockRows with data
func NewMockRows(data [][]any) *MockRows {
	return &MockRows{
		Data:         data,
		CurrentIndex: -1,
	}
}

// Close mocks the Close method
func (m *MockRows) Close() {
	m.Called()
}

// Err mocks the Err method
func (m *MockRows) Err() error {
	args := m.Called()
	return args.Error(0)
}

// CommandTag mocks the CommandTag method
func (m *MockRows) CommandTag() pgconn.CommandTag {
	args := m.Called()
	return args.Get(0).(pgconn.CommandTag)
}

// FieldDescriptions mocks the FieldDescriptions method
func (m *MockRows) FieldDescriptions() []pgconn.FieldDescription {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([]pgconn.FieldDescription)
}

// Next mocks the Next method - returns true if there are more rows
func (m *MockRows) Next() bool {
	m.CurrentIndex++
	return m.CurrentIndex < len(m.Data)
}

// Scan mocks the Scan method - copies data to dest
func (m *MockRows) Scan(dest ...any) error {
	args := m.Called(dest)
	return args.Error(0)
}

// Values mocks the Values method
func (m *MockRows) Values() ([]any, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]any), args.Error(1)
}

// RawValues mocks the RawValues method
func (m *MockRows) RawValues() [][]byte {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).([][]byte)
}

// Conn mocks the Conn method
func (m *MockRows) Conn() *pgx.Conn {
	args := m.Called()
	if args.Get(0) == nil {
		return nil
	}
	return args.Get(0).(*pgx.Conn)
}
