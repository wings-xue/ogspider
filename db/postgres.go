package db

import (
	req "og/reqeuest"

	"github.com/go-pg/pg"
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
func (self *PgSQL) Update(r *req.Request, upsert bool) {
	if upsert {
		self.Conn.Model(r).OnConflict("(uuid) DO UPDATE").
			Set("Download=EXCLUDD.Download").
			Set("Datas=EXCLUDD.Datas").
			Set("Status=EXCLUDD.Status").
			Set("Retry=EXCLUDD.Retry").
			Insert()
	} else {
		self.Conn.Model(r).OnConflict("DO NOTHING").SelectOrInsert()
	}
}

func (self *PgSQL) Select(i int) []*req.Request {
	var requests []*req.Request
	self.Conn.Model(&requests).Limit(i).Select()
	return requests
}
