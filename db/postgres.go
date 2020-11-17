package db

import (
	"fmt"
	req "og/reqeuest"
	"og/setting"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-pg/pg/v10"
)

type PgSQL struct {
	// db
	Conn *pg.DB
}

func New() *PgSQL {
	opt, err := pg.ParseURL("postgres://postgres:xws09040@121.36.224.198:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}
	db := pg.Connect(opt)
	return &PgSQL{
		Conn: db,
	}

}

// upsert == true ,存在更新，不存在插入
// upser == false, 存在丢弃，不存在插入
func (self *PgSQL) Update(r *req.Request, upsert bool) error {
	if upsert {
		_, err := self.Conn.Model(r).
			OnConflict("(uuid) DO UPDATE").
			Set("download=EXCLUDED.download").
			Set("datas=EXCLUDED.datas").
			Set("status=EXCLUDED.status").
			Set("retry=EXCLUDED.retry").
			Insert()
		if err != nil {
			return err
		}

		// fmt.Println(rst)
		// fmt.Println(err)
	} else {
		_, err := self.Conn.Model(r).OnConflict("DO NOTHING").SelectOrInsert()
		if err != nil {
			return err
		}

	}
	return nil
}

// upsert == true ,存在更新，不存在插入
// upser == false, 存在丢弃，不存在插入
func (self *PgSQL) MustUpdate(r *req.Request, upsert bool) {
	err := self.Update(r, upsert)
	if err != nil {
		// panic(err)
	}
}

func (self *PgSQL) Insert(map[string]interface{}) {

}

func (self *PgSQL) Select(i int) []*req.Request {
	var requests []*req.Request
	self.Conn.Model(&requests).Limit(i).Select()
	return requests
}

func (self *PgSQL) Close() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("close database")
		fmt.Println("close engine")
		self.Conn.Close()
		os.Exit(0)
	}()
}

func (self *PgSQL) Save(tablename string, rst map[string]interface{}) {
	m := self.Conn.Model(&rst).TableExpr(tablename).OnConflict("(" + setting.CrawlerRstKey + ")" + " DO UPDATE")
	for key, _ := range rst {
		m.Set(key + "=EXCLUDED." + key)
	}
	_, err := m.Insert()
	if err != nil {
		panic(err)
	}
}
