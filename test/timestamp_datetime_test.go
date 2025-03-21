package test

import (
	"ShortVideoVerifier/model"
	"ShortVideoVerifier/sql"
	"testing"
	"time"
)

func init() {
	if _, err := sql.InitMySQLEngine(); err != nil {
		panic(err)
	}
	if err := sql.GetMysql().Sync2(model.Range{}); err != nil {
		panic(err)
	}
}
// https://www.bilibili.com/video/BV12FX5Y9Ehk?t=99.9
func TestSetTime(t *testing.T) {
	zero := "1970-01-01 00:00:00"
	negative := "1000-01-01 00:00:00"
	positive := "2039-01-01 00:00:00"
	max:="9999-12-31 23:59:59"
	d0, t0 := Time2Data(zero)
	r1 := new(model.Range)
	r1.TimeStamp = t0
	r1.Datatime = d0
	r1.InsertOne()

	d1, t1 := Time2Data(negative)
	r2 := new(model.Range)
	r2.TimeStamp = t1
	r2.Datatime = d1.AddDate(-10, 0, 0)
	r2.InsertOne()
	d2, t2 := Time2Data(positive)
	r3 := new(model.Range)
	r3.TimeStamp = t2
	r3.Datatime = d2
	r3.InsertOne()

	d3,t3:=Time2Data(max)
	r4:=new(model.Range)
	r4.TimeStamp=t3
	r4.Datatime=d3
	r4.InsertOne()

	outofrange:=d3.AddDate(10,0,0)
	r5:=new(model.Range)
	r5.Datatime=outofrange
	r5.TimeStamp=t3
	r5.InsertOne()
}

func Time2Data(formatTime string) (time.Time, int64) {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", formatTime, loc)
	timestamp := t.Unix()
	return t, timestamp
}
