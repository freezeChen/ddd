package infrastructure

import "github.com/go-xorm/xorm"

var engine *xorm.Engine



func NewDb() *xorm.Engine {
	return engine
}

