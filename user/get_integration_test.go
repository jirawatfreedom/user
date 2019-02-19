// +build integration

package user_test

import (
	"fmt"
	"testing"

	"github.com/jirawatfreedom/user/user"
)

func TestIntegrationGetUser(t *testing.T) {
	u, err := user.Get()

	if err != nil {
		t.Error(err)
	}
	fmt.Println(u)
}
