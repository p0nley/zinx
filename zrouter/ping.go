package zrouter

import (
	"encoding/json"
	"fmt"

	"github.com/p0nley/zinx/zhao"
	"github.com/p0nley/zinx/ziface"
	"github.com/p0nley/zinx/zlog"
	"github.com/p0nley/zinx/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter
}

/*
func RecResultJsonToPlain() {
    var recResult string
    var dat map[string]interface{}
    json.Unmarshal([]byte(json_str), &dat)

    if v, ok := dat["ws"]; ok {
        ws := v.([]interface{})
        for i, wsItem := range ws {
            wsMap := wsItem.(map[string]interface{})
            if vCw, ok := wsMap["cw"]; ok {
                cw := vCw.([]interface{})
                for i, cwItem := range cw {
                    cwItemMap := cwItem.(map[string]interface{})
                    if w, ok := cwItemMap["w"]; ok {
                        recResult = recResult + w.(string)
                    }
                }
            }
        }
    }
    fmt.Println(recResult)
}
*/

func getValue(key string, m map[string]interface{}) string {
	nima := ""
	for k, v := range m {
		if k == key {
			nima = v.(string)
			//fmt.Println(v)
			//fmt.Sprintf("%s", v)
			//fmt.Sprintf("%v", v)

			break
		}
		continue
		switch vv := v.(type) {
		case string:
			fmt.Printf("%s is string, value: %s\n", k, vv)
		case int:
			fmt.Printf("%s is int, value: %d\n", k, vv)
		case int64:
			fmt.Printf("%s is int64, value: %d\n", k, vv)
		case bool:
			fmt.Printf("%s is bool, vaule: %v", k, vv)
		default:
			fmt.Printf("%s is unknow type\n", k)
		}
	}

	return nima
}

//Ping Handle
func (this *PingRouter) Handle(request ziface.IRequest) {

	zlog.Debug("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Debug("recv from client : msgId=0", ", data=", string(request.GetData()))

	jsonStr := string(request.GetData())

	//nima := request.GetData()
	var JObject interface{}

	//json.MarshalIndent()
	err := json.Unmarshal([]byte(jsonStr), &JObject)
	if err != nil {
		fmt.Println(err)
	}

	jo := JObject.(map[string]interface{})

	//fmt.Println(getValue("TEST", jo))

	action := getValue("ACTION", jo)
	fmt.Println(action)

	switch action {
	case "LOGIN":
		fmt.Printf("LOGIN")

		err = request.GetConnection().SendBuffMsg([]byte("{\"ACTION\":\"LOGIN\",\"session_id\":\"" + zhao.GetGUID().Hex() + "\"}"))

	case "HEART_BEAT":
		fmt.Printf("HEART_BEAT")
		//fmt.Printf("%s is int, value: %d\n")
		err = request.GetConnection().SendBuffMsg([]byte("{\"ACTION\":\"HEART_BEAT\"}"))
	case "PIPE":
		fmt.Printf("PIPE")
		//fmt.Printf("%s is int64, value: %d\n")
	case "WEB_EXEC":
		fmt.Printf("WEB_EXEC")
		//fmt.Printf("%s is bool, vaule: %v")
	default:
		//fmt.Printf("%s is unknow type\n")
	}

	type User struct {
		ACTION string
		//SESSION_ID string
		session_id string
	}

	//u := User{ACTION: "LOGIN", session_id: "tom"}
	//jsonU, _ := json.Marshal(u)

	//err = request.GetConnection().SendBuffMsg([]byte("ping...ping...ping[FromServer]"))
	if err != nil {
		zlog.Error(err)
	}
}
