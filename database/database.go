package database

import (
    "log"
    "database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
    db *sql.DB
    connString string
}

func CreateDatabase(driver string, connString string) *Database {
    db, err := sql.Open(driver, connString)
    if err != nil {
        log.Fatal(err)
    }
    return &Database{ db, connString }
}

func (db *Database) Exec(sqlStmt string) (sql.Result, error) {
	return db.db.Exec(sqlStmt)
}

func (db *Database) Close() {
    db.db.Close()
}

func (db *Database) DB() *sql.DB {
    return db.db;
}

func execSelectRow(db *sql.DB, selectSqlStmt string, args ...interface{}) *sql.Row {
    stmt, err := db.Prepare(selectSqlStmt)
    if err != nil {
        log.Fatal(err)
    }
    return stmt.QueryRow(args...)
}

func (db *Database) FindFile(filePath string) *DbFile {
    var id int
    var path string
    var isDir bool
    err := execSelectRow(
        db.db,
        "SELECT * FROM file WHERE path = ?",
        filePath).Scan(&id, &path, &isDir)

    if err != nil {
        return nil
    }
    
    return &DbFile { id, path, isDir }
}

func (db *Database) FindTag(tag string) *DbTag {
    var id int
    var name string
    err := execSelectRow(
        db.db,
        "SELECT * FROM tag WHERE name = ?",
        tag).Scan(&id, &name)
    
    if err != nil {
        return nil
    }
    
    return &DbTag { id, name }
}

func execInsert(db *sql.DB, insertSqlStmt string, args ...interface{}) int64 {
    stmt, err := db.Prepare(insertSqlStmt)
    if err != nil {
        log.Fatal(err)
    }
    insert, err := stmt.Exec(args...)
    if err != nil {
        log.Fatal(err)
    }
    id, err := insert.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }
    return id
}

func (db *Database) InsertFile(path string, isDir bool) *DbFile {
    file := db.FindFile(path)
    if file == nil {
        id := execInsert(
            db.db,
            "INSERT INTO file (path, is_dir) VALUES (?, ?)",
            path, isDir)

        return &DbFile{int(id), path, isDir}
    }
    return file
}

func (db *Database) InsertTag(tagName string) *DbTag {
    tag := db.FindTag(tagName)
    if tag == nil {
        id := execInsert(
            db.db,
            "INSERT INTO tag (name) VALUES (?)",
            tagName)

        return &DbTag { int(id), tagName }
    }
    return tag
}

func (db *Database) TagFile(f *DbFile, t *DbTag) {
    stmt, err := db.db.Prepare("INSERT OR IGNORE INTO file_tag (file_id, tag_id) VALUES (?, ?)")
    if err != nil {
        log.Fatal(err)
    }
    _, err = stmt.Exec(f.ID(), t.ID())
    if err != nil {
        log.Fatal(err)
    }
}