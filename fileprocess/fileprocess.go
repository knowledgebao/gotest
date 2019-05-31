package fileprocess

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

//github.io发布的时候需要特定的格式，这里实现对指定文件
//夹的文件进行处理，使其符合发布条件
//

//DEPTH 遍历的最深层次
var DEPTH = 10

//WalkDir 遍历目录dirpath,如果发现文件，则使用f进行处理
func WalkDir(dirpath string, depth int, f func(string)) {
	if depth > DEPTH { //大于设定的深度
		return
	}
	files, err := ioutil.ReadDir(dirpath) //读取目录下文件
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			WalkDir(dirpath+"/"+file.Name(), depth+1, f)
			continue
		} else {
			f(dirpath + "/" + file.Name())
		}
	}
}

//AddTagData combine func
func AddTagData(absFile string) {
	AddTag(absFile)
	AddDataTime(absFile)
}

//AddDataTime 给文件名添加时间,比如：源文件名叫oldFile.txt,则处理后格式为2019-08-08-oldFile.txt
func AddDataTime(absFile string) {
	// fmt.Println(filepath.Base(absFile)) //获取路径中的文件名test.txt
	paths, fileName := filepath.Split(absFile)
	if CheckSpecialExpr(fileName, "\\d{4}-\\d{2}(-\\d{2})-") {
		return
	}
	newFilename := GetFileModTime(absFile).Format("2006-01-02-") + fileName
	newAbsFile := paths + newFilename
	os.Rename(absFile, newAbsFile)
}

//AddTag 判断文件是否有标题，如果没有添加标题，如果有跳过，标题格式如下：
/*
---
layout: post
title: filename
date: 2016-01-09 11:15:06
description: filename
tag: folderName

---
*/
func AddTag(absFile string) {
	f, err := os.OpenFile(absFile, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	defer f.Close()
	//Read
	br := bufio.NewReader(f)
	a, _, c := br.ReadLine()
	if c == io.EOF {
		//
	} else if (string(a)) == "---" {
		return
	}
	//GetInfo
	fi, err := f.Stat()
	if err != nil {
		fmt.Println("stat fileinfo error")
		return
	}
	data := fi.ModTime().Format("2006-01-02 15:04:05")
	paths, title := filepath.Split(absFile)
	title = strings.TrimSuffix(title, path.Ext(title))
	title = ReplaceSpecialExpr(title, "", "\\d{4}-\\d{2}(-\\d{2})-")
	tag := path.Base(paths)
	//Write
	f.Seek(0, 0)
	buffer := make([]byte, fi.Size())
	if _, err = f.Read(buffer); err != nil {
		log.Fatal("read file fail")
		return
	}
	f.Seek(0, 0)

	bw := bufio.NewWriter(f)
	bw.WriteString("---\nlayout: post")
	bw.WriteString("\ntitle: ")
	bw.WriteString(title)
	bw.WriteString("\ndate: ")
	bw.WriteString(data)
	bw.WriteString("\ndescription: ")
	bw.WriteString(title)
	bw.WriteString("\ntag: ")
	bw.WriteString(tag)
	bw.WriteString("\n\n---\n")
	bw.Write(buffer)
	bw.Flush()
}

//TrimSpecialSuffix 文件删除指定后缀
func TrimSpecialSuffix(absFile string, suffix string) {
	// fmt.Println(filepath.Base(absFile)) //获取路径中的文件名test.txt
	if path.Ext(absFile) == suffix {
		newFileName := strings.TrimSuffix(absFile, suffix)
		os.Rename(absFile, newFileName)
	}
}

//GetFileModTime 获取文件修改时间 返回time.Time对象
func GetFileModTime(path string) time.Time {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("open file error")
		return time.Now()
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		fmt.Println("stat fileinfo error")
		return time.Now()
	}

	return fi.ModTime()
}

//ReplaceSpecialExpr 用repl替换满足正则表达式的内容
func ReplaceSpecialExpr(fileName, repl, expr string) string {
	reg := regexp.MustCompile(expr)
	return reg.ReplaceAllString(fileName, repl)
}

//CheckSpecialExpr 查找filename是否有复合正则要求的内容
func CheckSpecialExpr(fileName, expr string) bool {
	reg := regexp.MustCompile(expr)
	return reg.MatchString(fileName)
}
