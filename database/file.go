package database

type DbFile struct {
    id int
    name string
    path string
    isDir bool
}

func (fd *DbFile) ID() int {
    return fd.id
}