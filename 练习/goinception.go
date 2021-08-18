package main

import (
	"fmt"
	"time"
)

func main() {
	//currentTime:=time.Now().Format("2006-01-02 15:04:05")
	socketFile := fmt.Sprintf("/tmp/gh-ost.%s.%d.%s.%s.%d.sock", "opt.Host", 3306,
		"data_base", "table_name", time.Now().Unix())
	fmt.Println(socketFile)

	//socketFile := fmt.Sprintf("/tmp/gh-ost.%s.%d.%s.%s.%d.sock", s.opt.Host, s.opt.Port,
	//	r.TableInfo.Schema, r.TableInfo.Name, )

}
