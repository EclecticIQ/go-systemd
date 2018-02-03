// Copyright 2015 CoreOS, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package login1

import (
	"fmt"
	"testing"
)

// TestNew ensures that New() works without errors.
func TestNew(t *testing.T) {
	_, err := New()

	if err != nil {
		t.Fatal(err)
	}
}

func TestListSessions(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	sessions, err := c.ListSessions()
	if err != nil {
		t.Fatal(err)
	}

	if len(sessions) < 1 {
		t.Fatal(fmt.Errorf("expected at least one active session"))
	}

	for _, session := range sessions {
		fmt.Println(session.ID)
		fmt.Println(session.UID)
		fmt.Println(session.User)
		fmt.Println(session.Seat)
		fmt.Println(session.Path)
	}
}

func TestListUsers(t *testing.T) {
	c, err := New()
	if err != nil {
		t.Fatal(err)
	}

	users, err := c.ListUsers()
	if err != nil {
		t.Fatal(err)
	}

	if len(users) < 1 {
		t.Fatal(fmt.Errorf("expected at least one active user"))
	}

	for _, user := range users {
		fmt.Println(user.UID)
		fmt.Println(user.Name)
		fmt.Println(user.Path)
	}

}
