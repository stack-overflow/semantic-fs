package schema

import (
    "log"
    "os"
    "semantic-fs/database"
)

func SchemaCreateTagTable() string {
    return "CREATE TABLE IF NOT EXISTS tag (id INTEGER PRIMARY KEY, name TEXT);"
}

func SchemaCreateFileTable() string {
    return "CREATE TABLE IF NOT EXISTS file (id INTEGER PRIMARY KEY, path TEXT, is_dir INTEGER);"
}

func SchemaCreateFileTagTable() string {
    return `CREATE TABLE IF NOT EXISTS file_tag (
		file_id INTEGER NOT NULL,
		tag_id INTEGET NOT NULL,
        PRIMARY KEY(file_id, tag_id),
		FOREIGN KEY(file_id) REFERENCES file(id),
		FOREIGN KEY(tag_id) REFERENCES tag(id)
	);`
}

func ExecuteSchema(db *database.Database, sqlStmt string) {
    status, err := db.Exec(sqlStmt)

    if err != nil {
        log.Printf("DB error: status: %v, error: %q, query: %s\n", status, err, sqlStmt)
        os.Exit(1)
    }
}

