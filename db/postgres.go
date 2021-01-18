package db

import (
	"context"
	"fmt"
	ogconfig "og/const"
	req "og/reqeuest"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-pg/pg/v10"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	b, _ := q.FormattedQuery()
	fmt.Println(string(b))
	return nil
}

type PgSQL struct {
	// db
	Conn *pg.DB
}

func New() *PgSQL {
	// opt, err := pg.ParseURL("postgres://postgres:88888888@192.168.100.121:5432/postgres?sslmode=disable")
	opt, err := pg.ParseURL("postgres://postgres:88888888@172.17.0.5:5432/postgres?sslmode=disable")
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
			Set("update_date=EXCLUDED.update_date").
			Set("insert_date=EXCLUDED.insert_date").
			Set("fresh_life=EXCLUDED.fresh_life").
			Set("log=EXCLUDED.log").
			Insert()
		if err != nil {
			return err
		}

		// fmt.Println(rst)
		// fmt.Println(err)
	} else {
		_, err := self.Conn.Model(r).
			OnConflict("(uuid) DO NOTHING").
			Insert()
		if err != nil {
			return err
		}

	}
	return nil
}

// upsert == true ,存在更新，不存在插入
// upser == false, 存在丢弃，不存在插入
func (self *PgSQL) MustUpdate(r req.Request, upsert bool) {
	// self.Conn.AddQueryHook(dbLogger{})
	err := self.Update(&r, upsert)
	if err != nil {

		panic(err)

	}
}

func (self *PgSQL) Insert(map[string]interface{}) {

}

func (self *PgSQL) Select(i int) []*req.Request {
	var requests []*req.Request
	self.Conn.Model(&requests).Limit(i).Select()
	return requests
}

func (self *PgSQL) SelectExpired() []*req.Request {
	var requests []*req.Request

	self.Conn.Model(&requests).Where("update_date + concat(to_char(fresh_life, '9999999999999999999'), ' seconds')::INTERVAL<?0", time.Now()).Select()
	return requests
}

func (self *PgSQL) GroupStatus() []struct {
	Host        string
	Status      string
	StatusCount int
} {
	var requests []*req.Request
	var res []struct {
		Host        string
		Status      string
		StatusCount int
	}
	self.Conn.Model(&requests).
		Column("host").
		Column("status").
		ColumnExpr("count(*) as status_count").
		Group("host").
		Group("status").
		Select(&res)
	return res
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
	m := self.Conn.Model(&rst).TableExpr(tablename).OnConflict("(" + ogconfig.CrawlerRstKey + ")" + " DO UPDATE")
	for key, _ := range rst {
		m.Set(key + "=EXCLUDED." + key)
	}
	_, err := m.Insert()
	if err != nil {
		panic(err)
	}
}
