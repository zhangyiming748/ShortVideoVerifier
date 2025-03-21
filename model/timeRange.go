package model
import (
	"time"
	"ShortVideoVerifier/sql"
)
type Range struct {
	Id        int64     `xorm:"not null pk autoincr comment('主键id') INT(11)"`
	TimeStamp   int64    `xorm:"comment('时间戳') BIGINT"`
	Datatime   time.Time    `xorm:"comment('时间') DateTime"`
	Sex string `xorm:"comment('性别') VARCHAR(512)"`
	CreatedAt time.Time `xorm:"created"`
	UpdatedAt time.Time `xorm:"updated"`
	DeletedAt time.Time `xorm:"deleted"`
}

func (r *Range)InsertOne() (int64,error){
	return sql.GetMysql().InsertOne(r)
}

func (r *Range) FindByPage(offset, limit int) ([]Range, error) {
	var rs []Range
	err := sql.GetMysql().Limit(limit, offset).Find(&rs)
	return rs, err
}
