package database

import (
  "database/sql"
)

// TokensDB is the interface for the tokens database.
type TokensDB interface {
  // GetToken gets the authorization token for a social network.
  GetToken(network string) (string, error)
  // InsertToken inserts a new authorization token into the database.
  InsertToken(network, token string) error
}

// PostgresTokensDB implements the TokensDB interface for a PostgreSQL database.
type PostgresTokensDB struct {
  db *sql.DB
}

// NewPostgresTokensDB creates a new instance of PostgresTokensDB.
func NewPostgresTokensDB(db *sql.DB) *PostgresTokensDB {
  return &PostgresTokensDB{db}
}

// GetToken implements the GetToken function of the TokensDB interface.
func (db *PostgresTokensDB) GetToken(network string) (string, error) {
  // Query to get the authorization token.
  query := `SELECT token FROM tokens WHERE network = $1`

  // Execute the query.
  row := db.db.QueryRow(query, network)

  // Get the authorization token from the query result.
  var token string
  err := row.Scan(&token)
  if err != nil {
    return "", err
  }

  return token, nil
}

// InsertToken implements the InsertToken function of the TokensDB interface.
func (db *PostgresTokensDB) InsertToken(network, token string) error {
  // Query to insert a new authorization token.
  query := `INSERT INTO tokens (network, token) VALUES ($1, $2)`

  // Execute the query.
  _, err := db.db.Exec(query, network, token)
  return err
}
