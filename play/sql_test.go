package main_test

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	. "ms/libs/go-orma/play"
	"ms/sun/helper"
	"testing"
    "math/rand"
    "io/ioutil"
)

//var DB *sqlx.DB
func load() {
	var err error
	DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms_test?charset=utf8mb4")
	DB.MapperFunc(func(s string) string { return s })
	if err != nil {
		panic("DB")
	}

    //insert data
    sql,err:= ioutil.ReadFile("test_data_tags.sql")
    if err != nil {
        panic("reading sql tag data faild.")
    }
    DB.Exec(string(sql))
}

func TestInsert(t *testing.T) {
	fmt.Println("ME:::::::Tes")
	load()
	//insrtComment(1000)
}

func TestSelectSQLX(t *testing.T) {
	//db, _ := sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms5?charset=utf8mb4")
	var tags []Tag
	err := DB.Select(&tags, "select * from tags")
	if len(tags) == 0 || err != nil {
		t.Error(" TestSelectSQLX faild", len(tags), err)
	} else {
		t.Logf("%v", tags[0])
	}
}

///////////////////////// Selector ////////////////
func TestSelect(t *testing.T) {
	r, err := NewComment_Selector().Id_LT(100).GetRows(DB)
	if len(r) == 0 || err != nil {
		t.Error(" Select faild", len(r), err)
	}
}

func TestSelectGetAll(t *testing.T) {
	r, err := NewTag_Selector().GetRows(DB)
	if len(r) == 0 || err != nil {
		t.Error(" TestSelectGetAll faild", len(r), err)
	} else {
		t.Logf("%v", r[0])
	}
}

func TestSelectGet(t *testing.T) {
	r, err := NewTag_Selector().GetRow(DB)
	if err != nil {
		t.Error("faild", err)
	} else {
		t.Logf("%v", r)
	}
}

func TestSelectGetWithLimit(t *testing.T) {
	r, err := NewTag_Selector().Limit(2).GetRow(DB)
	if err != nil {
		t.Error("faild", err)
	} else {
		t.Logf("%v", r)
	}
}

func TestSelectWithColumnSelect(t *testing.T) {
	r, err := NewTag_Selector().Select_Name().GetStringSlice(DB)
	if err != nil {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("%v", r)
	}
}

func TestSelectWithColumnSelectWithWhere(t *testing.T) {
	r, err := NewTag_Selector().Select_Name().Id_LT(10).GetStringSlice(DB)
	if err != nil {
		t.Error("faild", len(r), err)
	} else {
		t.Logf(" len = %d, %v", len(r), r)
	}
}

func TestSelectOrder(t *testing.T) {
	r, err := NewTag_Selector().OrderBy_Id_Desc().GetRows(DB)
	if err != nil {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectLimit(t *testing.T) {
	r, err := NewTag_Selector().Limit(10).GetRows(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectLimitOrder(t *testing.T) {
	r, err := NewTag_Selector().Limit(10).OrderBy_Id_Asc().GetRows(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectLimitOrder2(t *testing.T) {
	r, err := NewTag_Selector().Limit(10).OrderBy_Count_Desc().GetRows(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectLimitOrderOffset(t *testing.T) {
	r, err := NewTag_Selector().Limit(10).OrderBy_Count_Desc().Offset(10).GetRows(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectColumnLimitOrderOffset(t *testing.T) {
	r, err := NewTag_Selector().Select_Name().Limit(10).OrderBy_Count_Desc().Offset(10).GetRows(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", len(r), err)
	} else {
		t.Logf("len = %d", len(r))
	}
}

func TestSelectColumnWith_Eq_Object(t *testing.T) {
	r, err := NewTag_Selector().Id_EQ(10).GetRow(DB)
	if err != nil || r.Id != 10 {
		t.Error("faild", err)
	} else {
		//t.Logf("len = %d",)
	}
}

func TestSelectColumnWith_Eq_String(t *testing.T) {
	r, err := NewTag_Selector().Select_Name().Id_EQ(10).GetString(DB)
	if err != nil || len(r) == 0 {
		t.Error("faild", err)
	} else {
		//t.Logf("len = %d",)
	}
}

func TestSelectFullPlay1(t *testing.T) {
    ins:= []int{1,2,3,4,5}
    r, err := NewTag_Selector().Id_In(ins).CreatedTime_GE(125).GetRows(DB)
    if err != nil || len(r) == 0 {
        t.Error("faild", err)
    } else {
        t.Logf("len = %d",len(r))
    }
}

func TestSelectInsAndSelectName(t *testing.T) {
    ins:= []int{1,2,3,4,5}
    r, err := NewTag_Selector().Select_Name().Id_In(ins).CreatedTime_GE(125).GetStringSlice(DB)
    if err != nil || len(r) == 0 {
        t.Error("faild", err)
    } else {
        t.Logf("len = %d, ",len(r))
    }
}

///////////////////// Upadter ///////////////////////
func TestUpdaterInsInt(t *testing.T) {
    ins:= []int{1,2,3,4,5}
    rnd:=rand.Intn(100000)
    r, err := NewTag_Updater().Count(rnd).Id_In(ins).CreatedTime_GE(125).Update(DB)
    z,e:=TagById(DB,1)
    if err != nil || r < 0 || e != nil || z.Count != rnd{
        t.Error("faild", err)
    } else {
        t.Logf("update count = %d, %v",r, z)
    }
}

func TestUpdaterInsString(t *testing.T) {
    ins:= []string{"آتش‌سوزی","انسان","شده","MicroTugs"}
    rnd:=rand.Intn(100000)
    r, err := NewTag_Updater().Name_In(ins).Count(rnd).CreatedTime_GE(125).Update(DB)
    z,e:=TagById(DB,1)
    if err != nil || r < 0 || e != nil || z.Count != rnd{
        t.Error("faild", err)
    } else {
        t.Logf("update count = %d",r)
    }
}

func TestUpdaterAll(t *testing.T) {
    r, err := NewTag_Updater().IsBlocked(2).Update(DB)
    z,e:=TagById(DB,1)
    if err != nil || r < 0 || e != nil || z.IsBlocked != 2{
        t.Error("faild", err)
    } else {
        //t.Logf("update count = %d, %v",r, z)
    }
}

///////////////// Deleter ////////////////
// deleter don't support queryies without where for safety

func TestDeleteInsInt(t *testing.T) {
    ins:= []int{11,12,13,14,15}
    r, err := NewTag_Deleter().Id_In(ins).Count_LT(300).Delete(DB)
    z,e:=TagById(DB,13)
    if err != nil || r < 0 || e != nil || z.Id < 1{
        t.Error("faild", err)
    } else {
        //t.Logf("update count = %d, %v",r, z)
    }
}

func TestDeleteId(t *testing.T) {
    r, err := NewTag_Deleter().Id_EQ(20).Delete(DB)
    z,e:=TagById(DB,13)
    if err != nil || r < 0 || e != nil || z.Id < 1{
        t.Error("faild", err)
    } else {
        //t.Logf("update count = %d, %v",r, z)
    }
}

func TestDeleteAll_NO(t *testing.T) {
    r, err := NewTag_Deleter().Delete(DB)
    z,e:=TagById(DB,13)
    if err == nil || r > 0 || e != nil || z.Id < 1{
        t.Error("faild", err)
    } else {
        //t.Logf("update count = %d, %v",r, z)
    }
}

func insrtComment(num int) {
	for i := 1; i < num; i++ {
		c := Comment{}
		c.Id = i
		c.Text = helper.FactRandStrEmoji(100, true)
		c.UserId = i
		c.PostId = 10 * int64(i)
		c.Replace(DB)
	}
}

type Tag2 struct {
	TagId       int
	Name        string
	Count       int
	IsBlocked   int
	CreatedTime int
}
