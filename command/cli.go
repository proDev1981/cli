package command

import "fmt"
import "flag"

func MakeCli(cmds ...CMD){
  flag.Parse()
  for _, cmd := range cmds{
    if flag.Arg(0) == cmd.Name{
      if flag.Arg(1) == "--help"{
          parseHelp(cmd)
      }else{
          runCMD(cmd, flag.Args())
      }
      return
    }
  }
  fmt.Println("command no such!!")
}

func parseHelp(cmd CMD){
        fmt.Println("-- MAN --")
        fmt.Println(cmd.Man)
        fmt.Println("-- DESCRIPTION --")
        fmt.Println(cmd.Description)
}

func runCMD(cmd CMD, args []string){
        cmd.Handle(args[1:]...)
}
