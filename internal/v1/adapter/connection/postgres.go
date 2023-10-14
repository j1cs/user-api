package connection

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/j1cs/api-user/internal/v1/adapter/entity/db"
	"github.com/j1cs/api-user/internal/v1/service/config"

	"github.com/rs/zerolog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
	"github.com/j1cs/api-user/internal/v1/service/log"
)

func getWriteDataSource(envs *config.Environment) string {
	dbHost := envs.WriteDatabase.Host
	dbPort := envs.WriteDatabase.Port
	dbName := envs.WriteDatabase.Name
	dbUser := envs.WriteDatabase.User
	dbPassword := envs.WriteDatabase.Password
	dbSchema := envs.WriteDatabase.SchemaName
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?options=-c%%20search_path=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSchema)
}

func getReadDataSource(envs *config.Environment) string {
	dbHost := envs.ReadDatabase.Host
	dbPort := envs.ReadDatabase.Port
	dbName := envs.ReadDatabase.Name
	dbUser := envs.ReadDatabase.User
	dbPassword := envs.ReadDatabase.Password
	dbSchema := envs.ReadDatabase.SchemaName
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?options=-c%%20search_path=%s", dbUser, dbPassword, dbHost, dbPort, dbName, dbSchema)
}

func NewDatabaseConnection(envs *config.Environment, zlog *zerolog.Logger) *gorm.DB {
	gormConf := config.GormConf()
	gdb, err := gorm.Open(postgres.Open(getWriteDataSource(envs)), &gormConf)
	err = gdb.Use(dbresolver.Register(dbresolver.Config{
		TraceResolverMode: true,
		Replicas:          []gorm.Dialector{postgres.Open(getReadDataSource(envs))},
	}).
		SetMaxIdleConns(10).
		SetMaxOpenConns(50).
		SetConnMaxIdleTime(time.Minute).
		SetConnMaxLifetime(4 * time.Hour),
	)

	if err != nil {
		zlog.Fatal().Err(err).Msg("Error connecting to database")
	}
	if os.Getenv("GORM_LOG") == "true" {
		zlog.Info().Msg("gorm log enabled")
		gdb.Logger = log.NewGormLogger(zlog).LogMode(logger.Info)
	}

	err = ApplyMigrations(gdb, zlog)
	if err != nil {
		zlog.Fatal().Err(err).Msg("Error migrating to database")
	}

	return gdb
}

func ApplyMigrations(gdb *gorm.DB, logger *zerolog.Logger) error {
	var wasApplied bool
	var tables []db.Entity
	tables = append(tables, &db.User{})
	var err error

	for _, table := range tables {
		if err = gdb.AutoMigrate(&table); err == nil {
			if err = gdb.First(&table).Error; errors.Is(err, gorm.ErrRecordNotFound) {
				err = seedTable(table, gdb)
				if err != nil {
					return err
				}
				wasApplied = true
			}
		}
	}

	logger.Info().Msgf("Applied migrations?: %t", wasApplied)
	return err
}

func seedTable(entity db.Entity, gdb *gorm.DB) error {
	err := entity.Seed(gdb)
	if err != nil {
		return err
	}
	return err
}
