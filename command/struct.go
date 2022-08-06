package command

type CMD struct{
  Name string
  Man string
  Description string
  Handle func(...string)
}
