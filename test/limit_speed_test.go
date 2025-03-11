package test

import (
	"ShortVideoVerifier/model"
	"ShortVideoVerifier/sql"
	"testing"
	"time"
)
func init() {
	sql.InitMySQLEngine()
	sql.GetMysql().Sync2(model.Fake{})
} 
// go test -v -timeout 10h -run TestInsert500000
func TestInsert500000(t *testing.T) {
	for i:=0;i<500000;i++{
		one:=new(model.Fake)
		one.Name="fake"
		one.Age=i
		if i%2==0{
			one.Sex="female"
		}else{
			one.Sex="male"
		}
		one.InsertOne()
	}
}
func TestOffseSpeed(t *testing.T) {
	a_start:=time.Now()
	new(model.Fake).FindByPage(500000,10)
	a_end:=time.Now()
	a_speed:=a_end.Sub(a_start).Seconds()
	t.Logf("a查询花费了%v秒\n",a_speed)
	b_start:=time.Now()
	new(model.Fake).FindByPage(500000,10)
	b_end:=time.Now()
	b_speed:=b_end.Sub(b_start).Seconds()
	t.Logf("b查询花费了%v秒\n",b_speed)
}

func TestLimitSpeed(t *testing.T) {
	a_start:=time.Now()
	new(model.Fake).FindByPage(0,10)
	a_end:=time.Now()
	a_speed:=a_end.Sub(a_start).Seconds()
	t.Logf("a查询花费了%v秒\n",a_speed)
	b_start:=time.Now()
	new(model.Fake).FindByPage(0,10)
	b_end:=time.Now()
	b_speed:=b_end.Sub(b_start).Seconds()
	t.Logf("b查询花费了%v秒\n",b_speed)
}