package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func main() {
	var err error
	DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms_test?charset=utf8mb4")
	//DB, err = sqlx.Connect("mysql", "root:123456@tcp(localhost:3307)/ms5?charset=utf8mb4")
	DB.MapperFunc(func(s string) string { return s })
	if err != nil {
		panic("DB")
	}
	fmt.Println("hjb")

	play()
	update()
	_ = DB

}

func play() {
	for i := 0; i < 100; i++ {
		p := Post{}
		p.HasTag = []byte("as")
		p.LikesCount = i

		err := p.Insert(DB)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i := 0; i < 100; i++ {
		p := Comment{}
		p.Text = ("as")
		p.CreatedTime = i

		err := p.Insert(DB)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func update() {
	for i := 0; i < 10; i++ {
		/*		cnt, err := NewComment_Updater().CreatedTime_GT(10).Text("ssss").Update(DB)
						fmt.Println("ttt: ", cnt)
						if err != nil {
							fmt.Println(err)
						}

						cnt, err = NewComment_Updater().UserId(12).Text("uuuuu").Text_EQ("ssss").CreatedTime_LE(100).Update(DB)
						fmt.Println("uuu: ", cnt)
						if err != nil {
							fmt.Println(err)
						}

						NewComment_Updater().Text_Like("%s").Text("LIKE").Update(DB)
				        NewComment_Selector().Select_Id().Id_
				        u,_:=NewUser_Selector().Get()
				        u.Save(DB)
				        NewRecommendUser_Selecter().UserId_EQ(12).OrderBy_Id_Desc().GetAll()*/
	}
}

func m() {
	w := whereClause{}
	w.condition = "ddd"
}
