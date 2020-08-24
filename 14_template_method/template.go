package templatemethod

import "fmt"

type WorkInterface interface {
	GetUp()
	Work()
	Sleep()
}

type Worker struct {
	WorkInterface
}

func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

func (w *Worker) Daily() {
	w.GetUp()
	w.Work()
	w.Sleep()
}

type Coder struct {
}

func (c *Coder) GetUp() {
	fmt.Println("Coder Getup")
}

func (c *Coder) Work() {
	fmt.Println("Coding")
}
func (c *Coder) Sleep() {
	fmt.Println("Coder Sleep")
}
