package models

import (
	"os"
	"io"
	"time"
	"fmt"
)

func FileMv(account,ID string){
	t := time.Now().Format("2006-01-02 15")
	err := os.MkdirAll("/tmp/compare/"+account+"/"+t, 0700)
	if err!=nil{
	    fmt.Println(err)
	}
	CopyFile("/tmp/compare/"+account+"/"+t+"/"+account+"_"+ID+"_person.jpg","/tmp/compare/"+account+"/base/"+account+"_person.jpg")
	CopyFile("/tmp/compare/"+account+"/"+t+"/"+account+"_"+ID+"_own.jpg","/tmp/compare/"+account+"/base/"+account+"_own.jpg")


}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
