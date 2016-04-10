package database

type DbTag struct {
    id int
    name string
}

func (fd *DbTag) ID() int {
    return fd.id
}
