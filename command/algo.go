package command

import "fmt"

var Algo CMD = CMD{
  Name : "algo",
  Man : "algo <argument...>",
  Description : "description algo command",
  Handle : func (args ...string){
    fmt.Println("in algo command!")
  },
} 
