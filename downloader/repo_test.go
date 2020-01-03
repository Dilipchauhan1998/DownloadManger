package downloader

import "testing"

func TestInsertIntoRepository(t *testing.T){
	var N=10
	for i:=0;i<N;i++{
		download :=New(SERIAL)
		insertIntoRepository(&download)
	}

	if len(downloads)!=10{
		t.Fail()
	}
}

