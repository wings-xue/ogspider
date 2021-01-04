package db

import (
	"fmt"
	req "og/reqeuest"
	"testing"
)

var db = New()

func TestPgSQL_SelectExpired(t *testing.T) {

	tests := []struct {
		name string
		self *PgSQL
		want []*req.Request
	}{
		// TODO: Add test cases.
		{
			self: db,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.self.SelectExpired()
			fmt.Println(got)
		})
	}
}

func TestPgSQL_GroupStatus(t *testing.T) {
	tests := []struct {
		name string
		self *PgSQL
	}{
		// TODO: Add test cases.
		{self: db},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.self.GroupStatus()
			fmt.Println(got)
		})
	}
}
