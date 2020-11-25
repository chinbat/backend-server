// Code generated by mockery v2.3.0. DO NOT EDIT.

package mocks

import (
	context "context"

	covidshield "github.com/covid-tracing-mongolia/backend-server/pkg/proto/covidshield"
	mock "github.com/stretchr/testify/mock"

	persistence "github.com/covid-tracing-mongolia/backend-server/pkg/persistence"

	time "time"
)

// Conn is an autogenerated mock type for the Conn type
type Conn struct {
	mock.Mock
}

// CheckClaimKeyBan provides a mock function with given fields: _a0
func (_m *Conn) CheckClaimKeyBan(_a0 string) (int, time.Duration, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 time.Duration
	if rf, ok := ret.Get(1).(func(string) time.Duration); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(time.Duration)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClaimKey provides a mock function with given fields: _a0, _a1, _a2
func (_m *Conn) ClaimKey(_a0 string, _a1 []byte, _a2 context.Context) ([]byte, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string, []byte, context.Context) []byte); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, context.Context) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimKeyFailure provides a mock function with given fields: _a0
func (_m *Conn) ClaimKeyFailure(_a0 string) (int, time.Duration, error) {
	ret := _m.Called(_a0)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 time.Duration
	if rf, ok := ret.Get(1).(func(string) time.Duration); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(time.Duration)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ClaimKeySuccess provides a mock function with given fields: _a0
func (_m *Conn) ClaimKeySuccess(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ClearDiagnosisKeys provides a mock function with given fields: _a0
func (_m *Conn) ClearDiagnosisKeys(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *Conn) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountClaimedOneTimeCodes provides a mock function with given fields:
func (_m *Conn) CountClaimedOneTimeCodes() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountDiagnosisKeys provides a mock function with given fields:
func (_m *Conn) CountDiagnosisKeys() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountExhaustedEncryptionKeysByOriginator provides a mock function with given fields:
func (_m *Conn) CountExhaustedEncryptionKeysByOriginator() ([]persistence.CountByOriginator, error) {
	ret := _m.Called()

	var r0 []persistence.CountByOriginator
	if rf, ok := ret.Get(0).(func() []persistence.CountByOriginator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.CountByOriginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountExpiredClaimedEncryptionKeysByOriginator provides a mock function with given fields:
func (_m *Conn) CountExpiredClaimedEncryptionKeysByOriginator() ([]persistence.CountByOriginator, error) {
	ret := _m.Called()

	var r0 []persistence.CountByOriginator
	if rf, ok := ret.Get(0).(func() []persistence.CountByOriginator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.CountByOriginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountExpiredClaimedEncryptionKeysWithNoUploadsByOriginator provides a mock function with given fields:
func (_m *Conn) CountExpiredClaimedEncryptionKeysWithNoUploadsByOriginator() ([]persistence.CountByOriginator, error) {
	ret := _m.Called()

	var r0 []persistence.CountByOriginator
	if rf, ok := ret.Get(0).(func() []persistence.CountByOriginator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.CountByOriginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountUnclaimedEncryptionKeysByOriginator provides a mock function with given fields:
func (_m *Conn) CountUnclaimedEncryptionKeysByOriginator() ([]persistence.CountByOriginator, error) {
	ret := _m.Called()

	var r0 []persistence.CountByOriginator
	if rf, ok := ret.Get(0).(func() []persistence.CountByOriginator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.CountByOriginator)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountUnclaimedOneTimeCodes provides a mock function with given fields:
func (_m *Conn) CountUnclaimedOneTimeCodes() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOldDiagnosisKeys provides a mock function with given fields:
func (_m *Conn) DeleteOldDiagnosisKeys() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOldEncryptionKeys provides a mock function with given fields:
func (_m *Conn) DeleteOldEncryptionKeys() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteOldFailedClaimKeyAttempts provides a mock function with given fields:
func (_m *Conn) DeleteOldFailedClaimKeyAttempts() (int64, error) {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchKeysForHours provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Conn) FetchKeysForHours(_a0 string, _a1 uint32, _a2 uint32, _a3 int32) ([]*covidshield.TemporaryExposureKey, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 []*covidshield.TemporaryExposureKey
	if rf, ok := ret.Get(0).(func(string, uint32, uint32, int32) []*covidshield.TemporaryExposureKey); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*covidshield.TemporaryExposureKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, uint32, uint32, int32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetServerEvents provides a mock function with given fields: startDate
func (_m *Conn) GetServerEvents(startDate string) ([]persistence.Events, error) {
	ret := _m.Called(startDate)

	var r0 []persistence.Events
	if rf, ok := ret.Get(0).(func(string) []persistence.Events); ok {
		r0 = rf(startDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.Events)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(startDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTEKUploads provides a mock function with given fields: startDate
func (_m *Conn) GetTEKUploads(startDate string) ([]persistence.Uploads, error) {
	ret := _m.Called(startDate)

	var r0 []persistence.Uploads
	if rf, ok := ret.Get(0).(func(string) []persistence.Uploads); ok {
		r0 = rf(startDate)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]persistence.Uploads)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(startDate)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewKeyClaim provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *Conn) NewKeyClaim(_a0 context.Context, _a1 string, _a2 string, _a3 string) (string, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) string); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PrivForPub provides a mock function with given fields: _a0
func (_m *Conn) PrivForPub(_a0 []byte) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveEvent provides a mock function with given fields: event
func (_m *Conn) SaveEvent(event persistence.Event) error {
	ret := _m.Called(event)

	var r0 error
	if rf, ok := ret.Get(0).(func(persistence.Event) error); ok {
		r0 = rf(event)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreKeys provides a mock function with given fields: _a0, _a1, _a2
func (_m *Conn) StoreKeys(_a0 *[32]byte, _a1 []*covidshield.TemporaryExposureKey, _a2 context.Context) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(*[32]byte, []*covidshield.TemporaryExposureKey, context.Context) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
