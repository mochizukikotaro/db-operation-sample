package main

import (
	_ "github.com/lib/pq"
	"github.com/mochizukikotaro/golang-orm-sample/sample"
)

func main() {
	sample.SampleSql()
	sample.SampleGorm()
	sample.SampleDbr()
	sample.SampleGorp()
	sample.SampleSqlx()
}
