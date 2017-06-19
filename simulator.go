package elevator

import (
  "time"
  "math/rand"
  tm "github.com/buger/goterm"

)

type Simulator struct {
  Elevator *Elevator
  numRequests int
  requests [][]int
  processing [][]int
}

func printElevator(elevator *Elevator) {
	tm.Clear() // Clear current screen
	tm.MoveCursor(1,1)
	tm.Println(elevator)
  tm.Println(elevator.Floors[elevator.CurrentFloor - 1])
	tm.Flush()
}

func (s *Simulator) Run(done chan int) {
  ch := make(chan int)
  printElevator(s.Elevator)
  s.Elevator.Start(ch)
  s.generateRandomRequests()
  go s.processRequest()
  go s.processSummons(ch, done)
}

func (s *Simulator) generateRandomRequests() {
  s.requests = make([][]int, s.numRequests)
  for i := 0; i < len(s.requests); i++ {
    pickup := rand.Intn(len(s.Elevator.Floors)) + 1
    dropoff := rand.Intn(len(s.Elevator.Floors)) + 1
    s.requests[i] = []int{pickup, dropoff}
  }
}

func (s *Simulator) processSummons(ch chan int, done chan int) {
  for {
    if len(s.requests) == 0 && len(s.processing) == 0 && s.Elevator.Direction == DIRECTION_STOPPED {
      done <- 0
      return
    }
    floor := <- ch
    printElevator(s.Elevator)
    time.Sleep(time.Second)
    var new [][]int
    for i := 0; i < len(s.processing); i++ {
      request := s.processing[i]
      if request[0] == floor {
        s.Elevator.SummonFloor(request[1])
      } else {
        new = append(new, request)
      }
    }
    s.processing = new
    printElevator(s.Elevator)
  }
}

func (s *Simulator) processRequest() {
  if len(s.requests) == 0 {
    return
  }
  request := s.requests[0]
  s.requests = s.requests[1:]
  s.processing = append(s.processing, request)
  if request[0] < request[1] {
    s.Elevator.RequestFloor(request[0], DIRECTION_UP)
  } else {
    s.Elevator.RequestFloor(request[0], DIRECTION_DOWN)
  }

  time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
  s.processRequest()

	// floor := rand.Intn(len(elevator.Floors)) + 1
	// direction := rand.Intn(2) - 1
	// if direction == 0 {
	// 	elevator.SummonFloor(floor)
	// } else {
	// 	elevator.RequestFloor(floor, el.Direction(direction))
	// }
	// randomRequests(elevator)
}

func NewSimulator(elevator *Elevator, numRequests int) *Simulator {
  return &Simulator{Elevator: elevator, numRequests: numRequests}
}
