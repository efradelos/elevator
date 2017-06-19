package elevator

import (
  "bytes"
)

type Floor struct {
  UpRequest bool
  DownRequest bool
  SummonsRequest bool
}

func (f *Floor) HasRequest() bool {
  return f.UpRequest || f.DownRequest || f.SummonsRequest
}

func (f *Floor) String() string {
  var buffer bytes.Buffer
  if (f.UpRequest) {
    buffer.WriteString("↑ ")
  } else {
    buffer.WriteString("  ")
  }
  if (f.DownRequest) {
    buffer.WriteString("↓ ")
  } else {
    buffer.WriteString("  ")
  }
  if (f.SummonsRequest) {
    buffer.WriteString("o ")
  } else {
    buffer.WriteString("  ")
  }
  return buffer.String()
}
