package main

import (
   "net/http"
   "fmt"
   _"os"
   "html/template"
   "github.com/XIE_5/servce"
   "strconv"
   "github.com/XIE_5/model"
)
var (
  t *template.Template
  StuSlice *servce.StudentSlice
)
func init() {
  StuSlice = servce.NewStuSlice()
  t, _ = template.ParseFiles("StuWeb.html","css.html","button.html")
}
func checkError(err error) {
  if err != nil {
     fmt.Println(err)
     return
  }
}
func MainMenu(w http.ResponseWriter) {
     t, _ = template.ParseFiles("StuWeb.html",
                                 "css.html","button.html")
     t.Execute(w, nil)
}
func Showmes(w http.ResponseWriter) {
     stu := StuSlice.List()
     t, _ = template.ParseFiles("StuTable.html",
                                "css.html","button.html")
     t.Execute(w, stu)
}
func Sorts(w http.ResponseWriter) {
     stu := StuSlice.Sort()
     t, _ = template.ParseFiles("StuTable.html",
                                "css.html","button.html")
     t.Execute(w, stu)
}
func Finds(w http.ResponseWriter, r *http.Request) {
     Id := r.Form.Get("ID")
     if len(Id) == 5 {
       if id, err := strconv.Atoi(Id); err == nil {
          if StuSlice.FindIndex(id) {
              stu := StuSlice.Find(id);
              t, _ = template.ParseFiles("StuTable.html",
                                         "css.html","button.html")
              t.Execute(w, stu)
            } else {
              t, _ = template.ParseFiles("StuWeb.html",
                                          "css.html","button.html")
              t.Execute(w, "查询失败")
            }
          } else {
              t, _ = template.ParseFiles("StuWeb.html",
                                         "css.html","button.html")
              t.Execute(w, "查询失败")
              return
          }
       } else {
        t, _ = template.ParseFiles("StuWeb.html",
                                   "css.html","button.html")
        t.Execute(w, "查询失败")
        return
     }
}
func Adds(w http.ResponseWriter, r *http.Request) {
      id, _ := strconv.Atoi(r.Form.Get("Id"))
      name := r.Form.Get("Name")
      major := r.Form.Get("Major")
      sex, _ := strconv.Atoi(r.Form.Get("Sex"))
      birthday := r.Form.Get("Birthday")
      socre, _ := strconv.Atoi(r.Form.Get("Socre"))
      note := r.Form.Get("Note")
      stu := model.NewStu(id,name,major,sex,birthday,socre,note)
      StuSlice.Add(stu)
      t, _ = template.ParseFiles("StuWeb.html",
                                 "css.html","button.html")
      t.Execute(w, "添加成功")
}
func Deletes(w http.ResponseWriter, r *http.Request) {
     t, _ = template.ParseFiles("StuWeb.html",
                                 "css.html","button.html")
     Id := r.Form.Get("IDD")
     if len(Id) == 5 {
       if id, err:= strconv.Atoi(Id); err == nil {
          StuSlice.Delete(id)
          t.Execute(w, "删除成功")
       } else {
          t.Execute(w, "删除失败")
       }
     } else {
          t.Execute(w, "删除失败")
     }
}
func Modifys(w http.ResponseWriter, r *http.Request) {
     t, _ = template.ParseFiles("StuWeb.html",
                                  "css.html","button.html")
     Id := r.Form.Get("IDM")
     if len(Id) == 5 {
       if id, err:= strconv.Atoi(Id); err == nil {
          name := r.Form.Get("Namem")
          major := r.Form.Get("Majorm")
          sex, _ := strconv.Atoi(r.Form.Get("Sexm"))
          birthday := r.Form.Get("Birthdaym")
          socre, _ := strconv.Atoi(r.Form.Get("Socrem"))
          note := r.Form.Get("Note")
          StuSlice.Modify(id,name,major,sex,birthday,socre,note)
          t.Execute(w, "修改成功")
       } else {
          t.Execute(w, "修改失败")
       }
     } else {
          t.Execute(w, "修改失败")
     }
}
func StuWeb(w http.ResponseWriter, r *http.Request) {
   if r.Method == "GET" {
      t.Execute(w, nil)
   } else {
          r.ParseForm()
          switch r.Form.Get("change") {
            case "主页" :
               MainMenu(w)
            case "学生信息列表":
               Showmes(w)
            case "学生信息排序" :
               Sorts(w)
            case "查询" :
               Finds(w, r)
            case "添加" :
               Adds(w, r)
            case "删除" :
               Deletes(w, r)
            case "修改" :
               Modifys(w, r)
          }
          fmt.Println("信息已处理")
   }
}
func main() {
   http.HandleFunc("/StuWeb",StuWeb)
   err := http.ListenAndServe(":9090", nil)
   checkError(err)
}
