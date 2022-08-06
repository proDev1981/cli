package command

import "fmt"

var Nada CMD = CMD{
  Name : "nada",
  Man : "nada <arguments...>",
  Description : "description nada command",
  Handle : func(args...string){
    fmt.Println("in nada command")
  },
}
