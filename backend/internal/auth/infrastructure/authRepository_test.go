package infrastructure

import (
	"database/sql"
	"errors"
	"testing"
)

type MockRow struct {
	ScanFunc func(dest ...interface{}) error
}

func (mr *MockRow) Scan(dest ...interface{}) error {
	return mr.ScanFunc(dest...)
}

type DBMock struct {
	QueryRowFunc func(query string, args ...interface{}) *MockRow
	ExecFunc     func(query string, args ...interface{}) (sql.Result, error)
}

func (db *DBMock) QueryRow(query string, args ...interface{}) *MockRow {
	return db.QueryRowFunc(query, args...)
}
func (db *DBMock) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.ExecFunc(query, args...)
}

type authSessionRepo struct {
	db *DBMock
}

func (repo *authSessionRepo) InvalidateSession(username string) error {
	_, err := repo.db.Exec("UPDATE SOC_APP_auth_sessions SET logged = false, jti = NULL WHERE user_name = ?", username)
	return err
}
func (repo *authSessionRepo) GetJTIByUsername(username string) (string, error) {
	row := repo.db.QueryRow("SELECT jti FROM SOC_APP_auth_sessions WHERE user_name = $1 AND logged = true ORDER BY timestamp DESC LIMIT 1", username)
	var jti string
	err := row.Scan(&jti)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("No valid sessions for this user")
		}
		return "", err
	}
	return jti, nil
}
func (repo *authSessionRepo) CheckValidSession(username string) (bool, error) {
	row := repo.db.QueryRow("SELECT logged FROM SOC_APP_auth_sessions WHERE user_name = $1 ORDER BY timestamp DESC LIMIT 1", username)
	var logged bool
	err := row.Scan(&logged)
	if err != nil {
		return false, err
	}
	return logged, nil
}

func TestAuthSessionDBRepository_CheckValidSession(t *testing.T) {
	tests := []struct {
		name           string
		username       string
		mockScanFunc   func(dest ...interface{}) error
		expectedLogged bool
		expectError    bool
		expectedErrMsg string
	}{
		{
			name:     "User is logged in",
			username: "validUser",
			mockScanFunc: func(dest ...interface{}) error {
				*(dest[0].(*bool)) = true
				return nil
			},
			expectedLogged: true,
			expectError:    false,
			expectedErrMsg: "",
		},
		{
			name:     "User is not logged in",
			username: "invalidUser",
			mockScanFunc: func(dest ...interface{}) error {
				*(dest[0].(*bool)) = false
				return nil
			},
			expectedLogged: false,
			expectError:    false,
			expectedErrMsg: "",
		},
		{
			name:     "No valid session for user",
			username: "noSessionUser",
			mockScanFunc: func(dest ...interface{}) error {
				return errors.New("No valid sessions for this user")
			},
			expectedLogged: false,
			expectError:    true,
			expectedErrMsg: "No valid sessions for this user",
		},
		{
			name:     "Database error",
			username: "dbErrorUser",
			mockScanFunc: func(dest ...interface{}) error {
				return errors.New("database error")
			},
			expectedLogged: false,
			expectError:    true,
			expectedErrMsg: "database error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &DBMock{
				QueryRowFunc: func(query string, args ...interface{}) *MockRow {
					return &MockRow{
						ScanFunc: tc.mockScanFunc,
					}
				},
			}

			repo := &authSessionRepo{
				db: mockDB,
			}

			logged, err := repo.CheckValidSession(tc.username)

			if tc.expectError {
				if err == nil {
					t.Errorf("%s: expected an error but got none", tc.name)
				} else if err.Error() != tc.expectedErrMsg {
					t.Errorf("%s: expected '%s' error, got '%v'", tc.name, tc.expectedErrMsg, err)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tc.name, err)
				}
				if logged != tc.expectedLogged {
					t.Errorf("%s: expected logged to be %v, got %v", tc.name, tc.expectedLogged, logged)
				}
			}
		})
	}
}
func TestAuthSessionDBRepository_GetJTIByUsername(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		setupMock   func(db *DBMock)
		expectedJTI string
		expectError bool
	}{
		{
			name:     "Success retrieving JTI",
			username: "validUsername",
			setupMock: func(db *DBMock) {
				db.QueryRowFunc = func(query string, args ...interface{}) *MockRow {
					return &MockRow{
						ScanFunc: func(dest ...interface{}) error {
							*(dest[0].(*string)) = "mockedJTI"
							return nil
						},
					}
				}
			},
			expectedJTI: "mockedJTI",
			expectError: false,
		},
		{
			name:     "Unauthorized access",
			username: "invalidUsername",
			setupMock: func(db *DBMock) {
				db.QueryRowFunc = func(query string, args ...interface{}) *MockRow {
					return &MockRow{
						ScanFunc: func(dest ...interface{}) error {
							return sql.ErrNoRows
						},
					}
				}
			},
			expectedJTI: "",
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &DBMock{}
			tc.setupMock(mockDB)

			repo := &authSessionRepo{
				db: mockDB,
			}

			jti, err := repo.GetJTIByUsername(tc.username)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected an error but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got %v", err)
				}
				if jti != tc.expectedJTI {
					t.Errorf("Expected JTI to be '%v', got '%v'", tc.expectedJTI, jti)
				}
			}
		})
	}
}

type MockExecResult struct{}

func (mer MockExecResult) LastInsertId() (int64, error) { return 0, nil }
func (mer MockExecResult) RowsAffected() (int64, error) { return 1, nil }
func TestAuthSessionDBRepository_InvalidateSession(t *testing.T) {
	tests := []struct {
		name        string
		username    string
		mockExec    func(query string, args ...interface{}) (sql.Result, error)
		expectError bool
	}{
		{
			name:     "Successful invalidation",
			username: "userToInvalidate",
			mockExec: func(query string, args ...interface{}) (sql.Result, error) {
				return MockExecResult{}, nil
			},
			expectError: false,
		},
		{
			name:     "Database error on invalidation",
			username: "userWithError",
			mockExec: func(query string, args ...interface{}) (sql.Result, error) {
				return nil, errors.New("database error")
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockDB := &DBMock{
				ExecFunc: tc.mockExec,
			}

			repo := &authSessionRepo{
				db: mockDB,
			}

			err := repo.InvalidateSession(tc.username)

			if tc.expectError {
				if err == nil {
					t.Errorf("%s: expected an error but got none", tc.name)
				}
			} else {
				if err != nil {
					t.Errorf("%s: unexpected error: %v", tc.name, err)
				}
			}
		})
	}
}
