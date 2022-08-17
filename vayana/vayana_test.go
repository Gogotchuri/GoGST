package vayana

import "testing"

func TestClient_Ping(t *testing.T) {
	client, _ := NewDefaultClient(false, "5dbe13f8-c60b-48a6-8705-d734b8e134e5")
	err := client.Ping()
	if err != nil {
		t.Error(err)
	}
}

func TestClient_Authenticate(t *testing.T) {
	client, _ := NewDefaultClient(false, "5dbe13f8-c60b-48a6-8705-d734b8e134e5")
	err := client.Authenticate("tech+vanaya@kernel.finance", "Strawhats16!")
	if err != nil {
		t.Error(err)
	}
	err = client.Logout()
	if err != nil {
		t.Error("error logging out", err.Error())
	}
}

func TestClient_GetGSTINDetails(t *testing.T) {
	client, _ := NewDefaultClient(false, "5dbe13f8-c60b-48a6-8705-d734b8e134e5")
	err := client.Authenticate("tech+vanaya@kernel.finance", "Strawhats16!")
	if err != nil {
		t.Error(err)
	}
	resp, err := client.GetGSTINDetails("27AHQPG5192G1Z")
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	err = client.Logout()
	if err != nil {
		t.Error("error logging out", err.Error())
	}
}
