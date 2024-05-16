package mysqltest

import (
	"context"
	"database/sql"
	"fmt"

	"project-v/config"
	"project-v/internal/store/mysql/sqlc"
	"project-v/pkg/l"
	mysqlconnect "project-v/pkg/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	LocalTest     = "local_test"
	ContainerTest = "container_test"
)

const (
	EnvFileCfg = "../../../../"
)

const (
	defaultImage = "mysql:latest"
	exposedPort  = "3306/tcp"
	mappedPort   = "3306"

	databaseName = "workflow_test"

	mysqlUsername = "root"
	mysqlPassword = "password"
)

func NewDBForIntegrationTest(ctx context.Context, typeTest string) (
	db *sql.DB, cleanUpFunc func(), err error,
) {
	switch typeTest {
	case LocalTest:
		return GetDBLocal()
	case ContainerTest:
		return GetDBContainer(ctx)
	}

	return nil, nil, nil
}

func GetDBLocal() (db *sql.DB, cleanUpFunc func(), err error) {
	ll := l.New()
	cfg := config.Load(ll, EnvFileCfg)

	// Mysql
	db = mysqlconnect.NewConnectMysql(&cfg.MySql)

	return db, func() {}, nil
}

func GetDBContainer(ctx context.Context) (db *sql.DB, cleanUpContainer func(), err error) {
	mysqlCtn, err := testcontainers.GenericContainer(
		ctx, testcontainers.GenericContainerRequest{
			ContainerRequest: testcontainers.ContainerRequest{
				Image:        defaultImage,
				ExposedPorts: []string{exposedPort},
				Env: map[string]string{
					"MYSQL_ROOT_PASSWORD": mysqlPassword,
					"MYSQL_DATABASE":      databaseName,
				},
				WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
			},
			Started: true,
		})
	if err != nil {
		panic(err)
	}
	var cleanContainer = func() {
		mysqlCtn.Terminate(ctx)
	}

	// Get MySQL container's IP and port
	ip, err := mysqlCtn.Host(ctx)
	if err != nil {
		panic(err)
	}

	port, err := mysqlCtn.MappedPort(ctx, mappedPort)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysqlUsername, mysqlPassword, ip, port.Port(), databaseName)
	fmt.Println("===> DSN:", dsn)

	// Connect to database
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	// Run migrate
	err = RunMysqlMigrate(db)
	if err != nil {
		panic(err)
	}

	return db, cleanContainer, nil
}

func RunMysqlMigrate(db *sql.DB) error {
	driver, err := mysql.WithInstance(
		db, &mysql.Config{
			NoLock: true,
		},
	)
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://../../../../database/migrate",
		"mysql",
		driver,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

// NOTE: Minh - must remove after has sqlc query
var (
	_ = sqlc.AdditionalRequest{}
	_ = sqlc.OnHoldRequest{}
	_ = sqlc.OnHoldRequestApprover{}
	_ = sqlc.Ticket{}
	_ = sqlc.TicketActivity{}
	_ = sqlc.TicketEdge{}
	_ = sqlc.TicketField{}
	_ = sqlc.TicketNode{}
	_ = sqlc.TicketTag{}
)
