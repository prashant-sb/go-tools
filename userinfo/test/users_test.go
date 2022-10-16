package users

import (
	"testing"

	uinfo "github.com/prashant-sb/go-tools/userinfo/users"
)

const (
	testSchema = "usr.json"
	testUser   = "test"
	testSysDB  = "/etc/passwd"
)

func TestAddUser(t *testing.T) {
	ui := uinfo.NewUserOps()
	userName, err := ui.AddUser("/media/common/workspace/go-tools/userinfo/test/usr.json")
	if err != nil {
		t.Errorf("Error in adding user %v", err.Error())
		return
	}

	if userName == testUser {
		t.Logf("AddUser() PASSED, expected: %v got: %v", testUser, userName)
	} else {
		t.Logf("AddUser() FAILED, expected %v got %v", testUser, userName)
	}
}

func TestGetUser(t *testing.T) {
	ui := uinfo.NewUserOps()
	u, err := ui.Get(testUser)
	if err != nil {
		t.Errorf("Get() FAILED for user %v", err.Error())
	} else if u.Username == testUser {
		t.Logf("Get() PASSED expected: %v got: %v", testUser, u.Username)
	} else {
		t.Errorf("Get() FAILED for user %v", u.Username)
	}
}

func TestGetUsers(t *testing.T) {
	ul := uinfo.NewUserList()
	ulist, err := ul.Get()
	if err != nil {
		t.Errorf("Get() FAILED to User list: %v", err.Error())
		return
	}

	udb, err := ul.ReadEtcPasswd(testSysDB)
	if err != nil {
		t.Errorf("Get() FAILED to User list from system: %v", err.Error())
		return
	}
	if len(udb) == len(ulist.Users) {
		t.Logf("Get() PASSED for user list")
	} else {
		t.Errorf("Get() FAILED to User list from system")
	}
}

func TestDeleteUser(t *testing.T) {
	ui := uinfo.NewUserOps()
	if userName, err := ui.DeleteUser(testUser); err != nil {
		t.Errorf("DeleteUser() FAILED, %v", err.Error())
	} else {
		t.Logf("DeleteUser() PASSED, expected: %v got: %v", testUser, userName)
	}
}
