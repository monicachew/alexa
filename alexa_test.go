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
}
