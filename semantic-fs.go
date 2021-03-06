package main

import (
    "fmt"
    "log"
//    "os"
    "path/filepath"
//    "flag"
	_ "github.com/mattn/go-sqlite3"
	"semantic-fs/database"
	"semantic-fs/schema"
)


type Tagger struct {
	tagDB *database.Database
}

func CreateTagger(tagDB *database.Database) *Tagger {
    return &Tagger { tagDB: tagDB }
}

func (t *Tagger) TagFile(filePath string, tag string) {
    absPath, err := filepath.Abs(filePath)
    if err != nil {
        log.Fatal(err)
    }
    
    dbFile := t.tagDB.FindFile(absPath)
    dbTag := t.tagDB.FindTag(tag)
    fmt.Printf("TagFile:\nfile: %v\ntag: %v\n", dbFile, dbTag)
    t.tagDB.TagFile(dbFile, dbTag)
}

func main() {
    fmt.Println("Semantic filesystem")

    tagDB := database.CreateDatabase("sqlite3", "./semantic-fs-go.db")
	schema.ExecuteSchema(tagDB, schema.SchemaCreateFileTable())
	schema.ExecuteSchema(tagDB, schema.SchemaCreateTagTable())
	schema.ExecuteSchema(tagDB, schema.SchemaCreateFileTagTable())

    absPath, _ := filepath.Abs("semantic-fs-go.db")
    tagDB.InsertFile("semantic-fs-go.db", absPath, false)
    tagDB.InsertTag("test")
    CreateTagger(tagDB).TagFile("semantic-fs-go.db", "test")

    tagDB.Close()
}
