package model
import (
	"time"
	"ShortVideoVerifier/sql"
)
type Fake struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	Name   string    `xorm:"comment('姓名') VARCHAR(512)"`
	Age   int    `xorm:"comment('年龄') INT(11)"`
	Sex string `xorm:"comment('性别') VARCHAR(512)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (f *Fake)InsertOne() (int64,error){
	return sql.GetMysql().InsertOne(f)
}

func (f *Fake) FindByPage(offset, limit int) ([]Fake, error) {
	var fakes []Fake
	err := sql.GetMysql().Limit(limit, offset).Find(&fakes)
	return fakes, err
}
