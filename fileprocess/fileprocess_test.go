package fileprocess

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func TestProcess(t *testing.T) {
	WalkDir("D:/work/knowledgebao/knowledgebao.github.io/_posts", 0, AddTagData)
}

func TestReplace(t *testing.T) {
	fmt.Println(ReplaceSpecialExpr("abcdef", "", "a"))
	fmt.Println(ReplaceSpecialExpr("abcdef.ä¸‹è½½", "", ".ä¸‹è½½$"))
	var ymd = "^((([0-9]{3}[1-9]|[0-9]{2}[1-9][0-9]{1}|[0-9]{1}[1-9][0-9]{2}|[1-9][0-9]{3})-(((0[13578]|1[02])-(0[1-9]|[12][0-9]|3[01]))|((0[469]|11)-(0[1-9]|[12][0-9]|30))|(02-(0[1-9]|[1][0-9]|2[0-8]))))|((([0-9]{2})(0[48]|[2468][048]|[13579][26])|((0[48]|[2468][048]|[3579][26])00))-02-29))\\s+([0-1]?[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])$"
	ymd = "\\d{4}-\\d{2}(-\\d{2})-"
	fmt.Println(ReplaceSpecialExpr("2018-01-25-test.file", "", ymd))

	fmt.Println(CheckSpecialExpr("2018-01-25-test.file", ymd))
}

func TestGetParentPath(t *testing.T) {
	file := "D:/work/knowledgebao/knowledgebao.github.io/_posts/2019-06-28-待整理列表.md"
	paths, fileName := filepath.Split(file)
	fileSuffix := path.Ext(fileName) //获取文件后缀
	fileName = strings.TrimSuffix(fileName, fileSuffix)
	parentPath := path.Base(paths)
	fmt.Println(parentPath, fileName)
}

func GbkToUtf8(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func TestWrite(t *testing.T) {
	// absFile, _ := GbkToUtf8(StringUtil.String2bytes("D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\2019-06-28-待整理列表.md"))
	// f, err := os.OpenFile(StringUtil.Bytes2String(absFile), os.O_RDWR, 0)
	// absFile := "D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\test.md"
	absFile := "D:\\work\\knowledgebao\\knowledgebao.github.io\\_posts\\2019-06-28-待整理列表.md"
	f, err := os.OpenFile(absFile, os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err.Error())
		return
	}
	defer f.Close()
	bw := bufio.NewWriter(f)
	bw.WriteString("---\nlayout: post")
	bw.WriteString("\ntitle: ")
	bw.WriteString("\n---\n")
	bw.Flush()
}
