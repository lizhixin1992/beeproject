package main

import (
	_ "beeproject/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/orm_test?charset=utf8", 30)
}

func main() {
	//beego.Run()

	orm.Debug = true
	o := orm.NewOrm()
	//o.Using("default") // 默认使用 default，你可以指定为其他数据库

	//profile := new(models.Profile)
	//profile.Age = 30
	//
	//profileId,_ := o.Insert(profile)
	//fmt.Println(profileId)
	//
	//user := new(models.User)
	//user.ProfileId = int(profileId)
	//user.Name = "slene"
	//fmt.Println(o.Insert(user))

	//var r orm.RawSeter
	//r = o.Raw("select * from User")
	//fmt.Println(r.QueryRow())

	//user := new(models.User)
	//user.Name = "test1"
	//user.ProfileId = 9
	////创建一个实体，返回主键
	//fmt.Println(o.Insert(user))
	//user.Name = "test2"
	//user.Id = 3
	////根据整个实体更新
	//fmt.Println(o.Update(user))
	////根据主键删除
	//fmt.Println(o.Delete(user))

	////根据自己设置的属性去查询
	//user := models.User{Name: "slene"}
	//err := o.Read(&user,"Name")
	//if err == orm.ErrNoRows {
	//	fmt.Println("查询不到")
	//} else if err == orm.ErrMissPK {
	//	fmt.Println("找不到主键")
	//} else {
	//	fmt.Println(user.Id, user.Name, user.ProfileId)
	//}

	////尝试从数据库读取，不存在的话就创建一个
	////默认必须传入一个参数作为条件字段，同时也支持多个参数多个条件字段
	//user := models.User{Name: "slene1"}
	//// 三个返回参数依次为：是否新创建的，对象 Id 值，错误
	//if created, id, err := o.ReadOrCreate(&user, "Name"); err == nil {
	//	if created {
	//		fmt.Println("New Insert an object. Id:", id)
	//	} else {
	//		fmt.Println("Get an object. Id:", id)
	//	}
	//}

	////同时插入多个对象(在一个事务里面，类似于说起来的	insert into table (name, age) values("slene", 28),("astaxie", 30),("unknown", 20))
	////第一个参数 bulk 为并列插入的数量，第二个为对象的slice
	////返回值为成功插入的数量
	////bulk 为 1 时，将会顺序插入 slice 中的数据
	//users := []models.User{
	//	{Name: "slene91"},
	//	{Name: "astaxie1"},
	//	{Name: "unknown1", ProfileId: -1},
	//}
	//successNums, err := o.InsertMulti(3, users)
	//fmt.Println(successNums, err)

	////更新整个实体
	//user := models.User{Id: 2}
	////if o.Read(&user) == nil {
	////	user.Name = "MyName"
	////	if num, err := o.Update(&user); err == nil {
	////		fmt.Println(num)
	////	}
	////}
	////Update 默认更新所有的字段，可以更新指定的字段：
	//user.Name = "dasdfdsf"
	//user.ProfileId = 8787
	//num , err :=o.Update(&user, "Name")
	//fmt.Println(num,err)




	/************************************高级查询*********************************************/



}
