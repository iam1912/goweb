package model

import (
   "fmt"
)
type Student struct {
   ID int
   Name string
   Major string
   Sex int
   Birthday string
   Socre int
   Note string
}
func NewStu(id int, name string, major string,sex int,
            birthday string, socre int, note string) Student {
   return Student{
       ID: id,
       Name: name,
       Major: major,
       Sex: sex,
       Birthday: birthday,
       Socre: socre,
       Note: note,
   }
}
func (this *Student) Show() string {
     info := fmt.Sprintf("%d\t%s\t%s\t%d\t%s\t%d\t%s",this.ID, this.Name,
             this.Major, this.Sex, this.Birthday, this.Socre, this.Note)
     return info
}
