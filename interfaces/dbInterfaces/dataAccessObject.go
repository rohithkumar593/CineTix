// The above function initializes and retrieves a database client by name from a map.
package dbInterface

type DatabaseAccessObject interface {
	ConnectDB(connStr string) error
	CloseConnection()
	// QueryDB(query string, args []any) (any, error) //CRUD operations
	// QueryRow(query string, args []any) *sql.Row
}
