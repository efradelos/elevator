package elevator

import (
  "bytes"
  "fmt"
  "time"
)

type Direction int
const (
  DIRECTION_UP Direction = 1
  DIRECTION_STOPPED Direction = 0
  DIRECTION_DOWN Direction= -1
)


type Elevator struct {
  Floors []Floor
  CurrentFloor int
  Direction Direction
  On bool
}

func (e *Elevator) String() string {
  var buffer bytes.Buffer
  for i := len(e.Floors); i > 0; i -= 1 {
    buffer.WriteString(fmt.Sprintf("%2d: ", i))
    if e.CurrentFloor == i {
      buffer.WriteString("x ")
    } else {
      buffer.WriteString("_ ")
    }
    buffer.WriteString(e.Floors[i - 1].String())
    buffer.WriteString("\n")
  }
  return buffer.String()
}

func (e *Elevator) RequestFloor(floor int, direction Direction) {
  if direction == DIRECTION_UP {
    e.Floors[floor - 1].UpRequest = true
  } else {
    e.Floors[floor - 1].DownRequest = true
  }
}

func (e *Elevator) SummonFloor(floor int) {
  e.Floors[floor - 1].SummonsRequest = true
}


func (e *Elevator) Start(ch chan int) {
  e.On = true
  go e.run(ch)
}

func (e *Elevator) Stop() { }

func (e *Elevator) run(ch chan int) {
  for e.On {
    if e.hasRequests(DIRECTION_STOPPED) {
      if e.processFloor() {
        time.Sleep(2 * time.Second)
      }
      e.moveToNextFloor(ch)
      time.Sleep(time.Second)
    }
  }
}

func (e *Elevator) hasRequests(direction Direction) bool {
  for i, floor := range e.Floors {
    inDirectionOfTravel := direction == DIRECTION_STOPPED || int(direction) * (i + 1 - e.CurrentFloor) > 0

    if inDirectionOfTravel && (floor.HasRequest()) {
      return true
    }
  }
  return false
}

func (e *Elevator) moveToNextFloor(ch chan int) {
  if e.Direction == DIRECTION_STOPPED {
    e.Direction = DIRECTION_UP
  }

  if !e.hasRequests(DIRECTION_STOPPED) {
    e.Direction = DIRECTION_STOPPED
    return
  }

  if !e.hasRequests(e.Direction) {
    e.Direction *= -1
    return
  }

  e.CurrentFloor += int(e.Direction)
  ch <- e.CurrentFloor
}

func (e *Elevator) processFloor() bool {
  floor := &e.Floors[e.CurrentFloor - 1]

  hasRequests := floor.HasRequest()
  if hasRequests {
    floor.UpRequest = floor.UpRequest && e.Direction != DIRECTION_UP
    floor.DownRequest = floor.DownRequest && e.Direction != DIRECTION_DOWN
    floor.SummonsRequest = false
  }

  return hasRequests
}

func NewElevator(num int) *Elevator {
  return &Elevator{Floors: make([]Floor, num), CurrentFloor: 2}
}
