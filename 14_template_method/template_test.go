package templatemethod

import "testing"

func TestWorkerDaily(t *testing.T) {
	worker := NewWorker(&Coder{})
	worker.Daily()
}
