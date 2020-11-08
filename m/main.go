package main

import (
	"fmt"
	req "og/reqeuest"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type OgUser struct {
	tableName struct{} `pg:"oguser,alias:oguser"`
	Id        int64
	Name      string
	Emails    []string
}

func (u OgUser) String() string {
	return fmt.Sprintf("User<%d %s %v>", u.Id, u.Name, u.Emails)
}

type Story struct {
	Id       int64
	Title    string
	AuthorId int64
	Author   *OgUser
}

func (s Story) String() string {
	return fmt.Sprintf("Story<%d %s %s>", s.Id, s.Title, s.Author)
}

func xx() {
	opt, _ := pg.ParseURL("postgres://postgres:xws09040@121.36.224.198:5432/postgres?sslmode=disable")

	db := pg.Connect(opt)
	defer db.Close()
	res, _ := db.Exec(`SELECT   tablename   FROM   pg_tables
    WHERE   tablename   NOT   LIKE   'pg%'
	AND tablename NOT LIKE 'sql_%' 
	ORDER   BY   tablename;`)
	fmt.Println(res)
	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user1 := &OgUser{
		Name:   "admin",
		Emails: []string{"admin1@admin", "admin2@admin"},
	}
	_, err = db.Model(user1).Insert()
	if err != nil {
		panic(err)
	}

	_, err = db.Model(&OgUser{
		Name:   "root",
		Emails: []string{"root1@root", "root2@root"},
	}).Insert()

	if err != nil {
		panic(err)
	}

	story1 := &Story{
		Title:    "Cool story",
		AuthorId: user1.Id,
	}
	_, err = db.Model(story1).Insert()
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	user := &OgUser{Id: user1.Id}
	err = db.Model(user).WherePK().Select()
	if err != nil {
		panic(err)
	}

	// Select all users.
	var users []OgUser
	err = db.Model(&users).Select()
	if err != nil {
		panic(err)
	}

	// Select story and associated author in one query.
	story := new(Story)
	err = db.Model(story).
		Relation("Author").
		Where("story.id = ?", story1.Id).
		Select()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
	fmt.Println(users)
	fmt.Println(story)
	// Output: User<1 admin [admin1@admin admin2@admin]>
	// [User<1 admin [admin1@admin admin2@admin]> User<2 root [root1@root root2@root]>]
	// Story<1 Cool story User<1 admin [admin1@admin admin2@admin]>>
}

// createSchema creates database schema for User and Story models.
func createSchema(db *pg.DB) error {
	models := []interface{}{
		(*OgUser)(nil),
		(*Story)(nil),
	}

	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			// Temp: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {

	opt, err := pg.ParseURL("postgres://postgres:xws09040@121.36.224.198:5432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)
	var requests []*req.Request
	err = db.Model(&requests).Limit(10).Select()
	fmt.Println(err)
	fmt.Println(requests)

}
