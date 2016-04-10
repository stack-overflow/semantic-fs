package database

type DbFile struct {
    id int
    path string
    isDir bool
}

func (fd *DbFile) ID() int {
    return fd.id
}