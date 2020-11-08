package context

import (
	req "og/reqeuest"
	"reflect"
	"testing"
)

func TestManager_Pop(t *testing.T) {
	manager := New()
	req1 := req.New("")
	req1.URL = "http://www.baidu.com"
	manager.Push(req1)

	req2 := req.New("")
	req2.URL = "http://www.baidu.com"
	manager.Push(req2)
	manager.Push(req2)
	manager.Push(req2)

	tests := []struct {
		name    string
		self    *Manager
		want    *req.Request
		wantLen int
	}{
		// TODO: Add test cases.
		{
			self:    manager,
			want:    req1,
			wantLen: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.self.Pop(); !reflect.DeepEqual(got, tt.want) || !reflect.DeepEqual(tt.self.Len(), tt.wantLen) {
				t.Errorf("Manager.Pop() = %v, want %v", got, tt.want)
				t.Errorf("wantLen: %d, want: %v\n", tt.self.Len(), tt.self.Queue)
			}
		})
	}
}
