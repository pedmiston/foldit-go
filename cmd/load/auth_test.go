package main

import "testing"

func TestReadSecrets(t *testing.T) {
	authenticate("test.yml")
	if auth.DOAccessKey != "do_access_key" {
		t.Errorf("creds not loaded in auth")
	}
}
