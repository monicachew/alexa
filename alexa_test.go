package alexa

import "testing"

func TestInit(t *testing.T) {
	var a AlexaRank
	a.Init("head.csv")
	rank, err := a.GetRank("google.com")
	if err != nil {
		t.Errorf("Expected to find rank for google.com: %s", err)
	}
	if rank != 1 {
		t.Error("Expected rank 1 for google.com: %d", rank)
	}
	rank, err = a.GetRank("example.com")
	if err == nil {
		t.Errorf("Expected error for rank of example.com")
	}
	if rank != -1 {
		t.Errorf("Expected -1 for rank of example.com")
	}
	rep, err := a.GetReputation("taobao.com")
	if err != nil {
		t.Errorf("Expected to find reputation for taobao: %s", err)
	}
	if (rep - 0.1) > 0.0000001 {
		t.Errorf("Expected reputation 0.1 for taobao: %f", rep)
	}
	rep, err = a.GetReputation("google.com")
	if rep != 1.0 {
		t.Errorf("Expected reputation 1.0 for google: %f", rep)
	}
}
