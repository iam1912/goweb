package servce

import (
   "github.com/jinzhu/gorm"
   _"github.com/jinzhu/gorm/dialects/mysql"
   "github.com/XIE_5/model"
)
type StudentSlice struct {
   Students []model.Student
}
var (
  db *gorm.DB
  err error
)
func init() {
    db, err = gorm.Open("mysql",
            "root:15219331409@/stu?charset=utf8&parseTime=True&loc=Local")
    db.AutoMigrate(&model.Student{})
    checkError(err)
}
func NewStuSlice() *StudentSlice {
    return &StudentSlice{
    }
}
func (this *StudentSlice) List() []model.Student {
    db.Find(&this.Students)
    return this.Students
}
func (this *StudentSlice) Sort() []model.Student {
    db.Order("Socre").Find(&this.Students)
    return this.Students
}
func (this *StudentSlice) FindIndex(id int) bool {
    for i :=0;i<len(this.Students);i++ {
        if this.Students[i].ID == id {
           return true
           break
        }
    }
    return false
}
func (this *StudentSlice) Find(id int) []model.Student{
    db.Where("ID = ?", id).Find(&this.Students)
    return this.Students
}
func (this *StudentSlice) Add(stu model.Student) {
    db.Create(&stu)
}
func (this *StudentSlice) Delete(id int) {
    db.Where("ID = ?",id).Delete(&model.Student{})
}
func (this *StudentSlice) Modify(id int, name string, major string,
                      sex int, birthday string, socre int, note string){
    db.Model(model.Student{}).Where(" ID = ?",id).Updates(model.Student{Name: name,Major:major,Sex: sex,Socre: socre,Note: note})
}
func checkError(err error) {
    if err != nil {
         panic(err)
    }
}
