package models

import (
	"bubble/dao"
)

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func CreateTodo(todo *Todo) (err error) {
	err =dao.DB.Create(todo).Error
	return
}

func GetAllTodo() (todoList []*Todo , err error){
	//引用传递
	if err := dao.DB.Find(&todoList).Error;err!=nil{
		return nil,err
	}
	return
}

func GetTheTodo(id string)(todo *Todo,err error){
	todo = new(Todo)//初始化结构体，返回指针
	if err=dao.DB.Where("id=?",id).First(todo).Error;err!=nil{
		return nil,err
	}
	return
}

func UpdateTodo(todo *Todo) (err error){
	err =dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error){
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}