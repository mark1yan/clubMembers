package main

import (
	"testing"
	"time"
)

//mockMembersByEmail is mock object
var mockMembersByEmail = []Member{
	{
		1,
		"markiyan.rybak@gmail.com",
		"Markiyan",
		time.Now().Format("02-01-2006"),
	},
}

//TestEmailIsExist test emailIsExist func to check if email exist
func TestEmailIsExist(t *testing.T) {
	got := emailIsExist("markiyan.rybak@gmail.com", &mockMembersByEmail)
	want := true

	if want != got {
		t.Errorf("email must exists")
	}
}

//TestAddMember test addMember func to add new member and avoid duplicates
func TestAddMember(t *testing.T) {
	got := addMember("example@com", "example", &mockMembersByEmail)
	want := true

	gotDuplicate := addMember("example@com", "example", &mockMembersByEmail)
	wantDuplicate := false

	if want != got {
		t.Errorf("cant add new email")
	} else if gotDuplicate != wantDuplicate {
		t.Errorf("make duplicate")
	}
}
