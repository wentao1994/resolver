package resolver

import (
	"fmt"
	"testing"
)

func TestResolve1(t *testing.T) {
	r := &simpleResolver{}
	updates, watcher, err := r.Resolve("test.com:1234")
	if err != nil {
		t.Fatal(err)
	}
	if watcher != nil {
		t.FailNow()
	}
	if len(updates) != 1 {
		t.Fatal(updates)
	}
	update := updates[0]
	if update.Addr != "test.com:1234" ||
		update.Id != "test.com:1234" ||
		!update.Master ||
		update.Op != OP_Put ||
		update.Password != "" ||
		update.User != "" ||
		update.Weight != 0 {
		t.Fatal(update)
	}
}

func TestResolve2(t *testing.T) {
	r := &simpleResolver{}
	updates, watcher, err := r.Resolve("pwd@test.com:1234")
	if err != nil {
		t.Fatal(err)
	}
	if watcher != nil {
		t.FailNow()
	}
	if len(updates) != 1 {
		t.Fatal(updates)
	}
	update := updates[0]
	if update.Addr != "test.com:1234" ||
		update.Id != "test.com:1234" ||
		!update.Master ||
		update.Op != OP_Put ||
		update.Password != "pwd" ||
		update.User != "" ||
		update.Weight != 0 {
		t.Fatal(update)
	}
}

func TestResolve3(t *testing.T) {
	r := &simpleResolver{}
	updates, watcher, err := r.Resolve("pwd@test1.com:1234,pwd@test2.com:1234")
	if err != nil {
		t.Fatal(err)
	}
	if watcher != nil {
		t.FailNow()
	}
	if len(updates) != 2 {
		t.Fatal(updates)
	}
	for k, update := range updates {
		addr := fmt.Sprintf("test%d.com:1234", k+1)
		master := k == 0
		if update.Addr != addr ||
			update.Id != addr ||
			update.Master != master ||
			update.Op != OP_Put ||
			update.Password != "pwd" ||
			update.User != "" ||
			update.Weight != 0 {
			t.Fatal(update)
		}
	}
}

func TestResolve4(t *testing.T) {
	r := &simpleResolver{}
	updates, watcher, err := r.Resolve("user:pwd@test.com:1234")
	if err != nil {
		t.Fatal(err)
	}
	if watcher != nil {
		t.FailNow()
	}
	if len(updates) != 1 {
		t.Fatal(updates)
	}
	update := updates[0]
	if update.Addr != "test.com:1234" ||
		update.Id != "test.com:1234" ||
		!update.Master ||
		update.Op != OP_Put ||
		update.Password != "pwd" ||
		update.User != "user" ||
		update.Weight != 0 {
		t.Fatal(update)
	}
}

func TestResolve5(t *testing.T) {
	r := &simpleResolver{}
	updates, watcher, err := r.Resolve("1:pwd@test1.com:1234,2:pwd@test2.com:1234")
	if err != nil {
		t.Fatal(err)
	}
	if watcher != nil {
		t.FailNow()
	}
	if len(updates) != 2 {
		t.Fatal(updates)
	}
	for k, update := range updates {
		addr := fmt.Sprintf("test%d.com:1234", k+1)
		master := k == 0
		if update.Addr != addr ||
			update.Id != addr ||
			update.Master != master ||
			update.Op != OP_Put ||
			update.Password != "pwd" ||
			update.User != fmt.Sprintf("%d", k+1) ||
			update.Weight != 0 {
			t.Fatal(update)
		}
	}
}
