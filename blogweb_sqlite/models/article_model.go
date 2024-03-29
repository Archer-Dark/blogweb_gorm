package models

import (
	"../database"
	"fmt"
	"../config"
	"log"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

//---------添加文章-----------
func AddArticle(article Article) error {
	err := insertArticle(article)
	SetArticleRowsNum()
	return err
}

//插入一篇文章
func insertArticle(article Article) error {
	db := database.GetDB()
	db = db.Create(&article)
	return db.Error
}

//-----------查询文章---------

//根据页码查询文章
func FindArticleWithPage(page int) ([]Article, error) {
	page--
	fmt.Println("---------->page", page)
	return QueryArticleWithPage(page,config.NUM)
}

/**
分页查询数据库
limit分页查询语句，
	语法：limit m，n

	m代表从多少位开始获取，与id值无关
	n代表获取多少条数据

注意limit前面咩有where
*/
func QueryArticleWithPage(page, num int) ([]Article, error) {
	//sql := fmt.Sprintf("limit %d,%d", page*num, num)
	//return QueryArticlesWithCon(sql)
	var articles []Article
	db := database.GetDB()
	db = db.Limit(num).Offset(page*num).Find(&articles);if db.Error != nil {
		return nil,db.Error
	}
	return articles,nil
}

//------翻页------

//存储表的行数，只有自己可以更改，当文章新增或者删除时需要更新这个值
var articleRowsNum = 0

//只有首次获取行数的时候采取统计表里的行数
func GetArticleRowsNum() int {
	if articleRowsNum == 0{
		articleRowsNum = QueryArticleRowNum()
	}
	return articleRowsNum
}

//查询文章的总条数
func QueryArticleRowNum() int {
	db := database.GetDB()
	num := 0
	db.Model(&Article{}).Count(&num)
	return num
}

//设置页数
func SetArticleRowsNum()  {
	articleRowsNum = QueryArticleRowNum()
}

//----------查询文章-------------
func QueryArticleWithId(id int) Article {
	art := Article{}
	db := database.GetDB()
	db.First(&art,id)
	return art
}

//----------修改数据----------

func UpdateArticle(article Article) error {
	db := database.GetDB()
	db = db.Model(&article).Updates(article)
	return db.Error
}

//----------删除文章---------
func DeleteArticle(artID int) error {
	db := database.GetDB()
	db = db.Unscoped().Delete(&Article{},"id = ?",artID)
	return db.Error
}

//查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	db := database.GetDB()
	paramList := []string{}
	db = db.Model(&Article{}).Select(param).Find(&paramList);if db.Error != nil {
		log.Println(db.Error)
	}
	return paramList
}

//--------------按照标签查询--------------
/*
通过标签查询首页的数据
有四种情况
	1.左右两边有&符和其他符号
	2.左边有&符号和其他符号，同时右边没有任何符号
	3.右边有&符号和其他符号，同时左边没有任何符号
	4.左右两边都没有符号

通过%去匹配任意多个字符，至少是一个
*/

func QueryArticlesWithTag(tag string) ([]Article,error) {
	var articles []Article
	db := database.GetDB()
	db = db.Where("tags like '%&" + tag + "&%'")
	db = db.Or(" tags like '%&" + tag + "'")
	db = db.Or("tags like '" + tag + "&%'")
	db = db.Or("tags like '" + tag + "'")
	db = db.Find(&articles)
	return articles,db.Error
}