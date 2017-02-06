package controllers

import (
//	"os"
	"fmt"
	"bytes"
	"regexp"
	"os/exec"
	"strconv"
	"errors"
	"time"
//	"path/filepath"
	"path"

	"github.com/astaxie/beego"
	"github.com/crackcomm/nsqueue/producer"
	"github.com/nsqio/go-nsq"
)

// Operations about Compare
type CompareController struct {
	beego.Controller
}

type Result struct {
	Detail string
	Status string
}

type VerifyMsg struct {
	ID string
	User  string
	Files []string
}

// @Title comparePhotos
// @Description compare photos
// @Param	
// @Success 200 {int} confidence
// @Failure 403 body is empty
// @router / [post]
func (c *CompareController) Post() {
	account := c.GetString("account")
	ID := c.GetString("ID")
	if len(ID) == 0 {
		c.Data["json"] = Result{"no ID input", "fail"}
		c.ServeJSON()
		return
	}

	files := make([]string, 0)
	files = append(files, path.Join("/tmp/compare/"+account+"/base", account+"_person.jpg"), path.Join("/tmp/compare/"+account+"/base", account+"_own.jpg"))

/*	filepath.Walk("/tmp/compare",
		func(path string, f os.FileInfo, err error)  error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			files = append(files, path)
			return nil
		})*/

	if len(files) != 2 {
		c.Data["json"] = Result{"can not find 2 photos to compare", "fail"}
		c.ServeJSON()
		return
	}

/*	confidence, err := c.execVerify("/nvr/verify", files)
	if err != nil {
		c.Data["json"] = Result{err.Error(), "fail"}
		c.ServeJSON()
		return
	}
	
	if confidence < 0.5 {
		c.Data["json"] = Result{"confidence too low", "fail"}
	} else {
		c.Data["json"] = Result{strconv.FormatFloat(confidence, 'f', -1, 32), "success"}
	}*/

	err:= c.execVerifyByNSQ(account, files, 5,ID)


	if err != nil {
		c.Data["json"] = Result{err.Error(), "fail"}
		c.ServeJSON()
		return
	}
	
	c.Data["json"] = Result{"Server is verifying these two photos, please wait a minute.", "success"}
	c.ServeJSON()
	return
}

func (c *CompareController) execVerify(cmdName string, cmdArgs []string) (float64, error) {
	// Create an *exec.Cmd
	cmd := exec.Command(cmdName, cmdArgs...)

	// Stdout buffer
	cmdOutput := &bytes.Buffer{}
	// Attach buffer to command
	cmd.Stdout = cmdOutput

	// Execute command
	err := cmd.Run() // will wait for command to return
	if err != nil {
		fmt.Println(err)
		return float64(0), err
	}

	outs := cmdOutput.Bytes()
	if len(outs) > 0 {
		fmt.Println(string(outs))
		pattern := `^finish loading all nets\s+([0-9\.]+)\s+$`
		reg := regexp.MustCompile(pattern)
		result := reg.FindStringSubmatch(string(outs))
		if len(result) == 0 {
			return float64(0), errors.New("invaild outs of verify.exe")
		}

		confidence, err := strconv.ParseFloat(result[1], 32)
		if err != nil {
			return confidence, err 
		}
		
		return confidence, nil
	}
	
	return float64(0), errors.New("failed to exec verify.exe")
}

func (c *CompareController) execVerifyByNSQ(user string, cmdArgs []string, timeout int,ID string) (err error) {
	msg := VerifyMsg{User: user, Files: cmdArgs,ID:ID}
	myProducer := producer.New()
	myProducer.Connect("192.168.2.122:4150")

	doneChan := make(chan *nsq.ProducerTransaction, 1)
	myProducer.PublishJSONAsync("verify", msg, doneChan)

	for {
		select {
			case pack := <-doneChan:
				fmt.Println(pack)
				if pack.Error != nil {
					return pack.Error
				}
				return nil
			case <-time.After(time.Second * time.Duration(timeout)):
				return errors.New("It's really weird to get Nothing!!!")
		}
	}

	myProducer.Stop()
	return nil
}
