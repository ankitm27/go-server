import "database/sql"

type mysqlArticleRepository struct {
	Conn *sql.DB
}
