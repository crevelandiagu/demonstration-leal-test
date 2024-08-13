package postgres

import (
	"context"
	"demonstration-leal-test/internal/adapter/config"

	"fmt"
	"github.com/Masterminds/squirrel"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

//// migrationsFS is a filesystem that embeds the migrations folder
////
////go:embed migrations/*.sql
//var migrationsFS embed.FS

/**
 * PostgresDB is a wrapper for PostgreSQL database connection
 * that uses pgxpool as database driver.
 * It also holds a reference to squirrel.StatementBuilderType
 * which is used to build SQL queries that compatible with PostgreSQL syntax
 */
type PostgresDB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

// New creates a new PostgreSQL database instance
func New(ctx context.Context, config *config.DB) (*PostgresDB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		config.Connection,
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
	)

	db, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &PostgresDB{
		db,
		&psql,
		url,
	}, nil
}

//// Migrate runs the database migration
//func (db *PostgresDB) Migrate() error {
//	driver, err := iofs.New(migrationsFS, "migrations")
//	if err != nil {
//		return err
//	}
//
//	migrations, err := migrate.NewWithSourceInstance("iofs", driver, db.url)
//	if err != nil {
//		return err
//	}
//
//	err = migrations.Up()
//	if err != nil && err != migrate.ErrNoChange {
//		return err
//	}
//
//	return nil
//}

// ErrorCode returns the error code of the given error
func (db *PostgresDB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}

// Close closes the database connection
func (db *PostgresDB) Close() {
	db.Pool.Close()
}
