package workTest

/*
数据处理，包括：结构体，数组，slice，map，json
*/
import (
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"golandProject/goLang/src/common"
	"log"
	"strconv"
)

type Data struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

type Data1 struct {
	Id      []int  `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"comment"`
}

//构建[{15 jingdong 备注15} {16 jingdong 备注16} {17 huahong 备注17} {18 guoyao 备注18}]数据
func Test1() {
	//一。测试数据start
	enterpriseArray := [...]string{"jingdong", "jingdong", "huahong", "guoyao"}
	dealData := make([]Data, 4)
	for index := range enterpriseArray {
		dealData[index].Id = index + 15
		dealData[index].Name = enterpriseArray[index]
		dealData[index].Comment = "备注" + strconv.Itoa(index+15)
	}
	fmt.Println(dealData) //[{15 jingdong 备注15} {16 jingdong 备注16} {17 huahong 备注17} {18 guoyao 备注18}]
	//测试数据end

	//二。处理name重复数据 start
	//1.取出name,处理重复，保证result中存储的name值不重复
	result := []string{} //存放结果
	for i := range dealData {
		flag := true
		for j := range result {
			if dealData[i].Name == result[j] {
				flag = false //存在重复元素，标识为false
				break
			}
		}
		if flag {
			result = append(result, dealData[i].Name)
		}
	}
	//重复数据处理end

	//三。重新组装数据
	data := make([]Data1, len(result))
	for m := range result {
		flg := true
		data[m].Name = result[m]  //先添加name值
		for n := range dealData { //循环处理id值，如果name值相等，则追加id值
			if result[m] == dealData[n].Name {
				if flg {
					data[m].Comment = dealData[n].Comment //取第1次出现的值
					flg = false
				}
				data[m].Id = append(data[m].Id, dealData[n].Id) //追加ID值
			}
		}
	}
	fmt.Println(data) //[{[15 16] jingdong 备注15} {[17] huahong 备注17} {[18] guoyao 备注18}]

	//四。结构体转换成json
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Json Marshaling failed:%s", err)
	}
	fmt.Printf("%s\n", jsonData) //[{"id":[15,16],"name":"jingdong","comment":"备注15"},{"id":[17],"name":"huahong","comment":"备注17"},{"id":[18],"name":"guoyao","comment":"备注18"}]
}

//构建{"code":200,"msg":"success","result":[{"id":"101","name":"yufeng"},{"id":"101","name":"kaixin"}]}数据
func Test2() {
	//拼凑json   body为map数组
	var rbody []map[string]interface{}
	temp := make(map[string]interface{})
	temp["id"] = "101"
	temp["name"] = "yufeng"

	rbody = append(rbody, temp)

	temp2 := make(map[string]interface{})
	temp2["id"] = "101"
	temp2["name"] = "kaixin"

	rbody = append(rbody, temp2)

	cnnJson := make(map[string]interface{})
	cnnJson["code"] = 200
	cnnJson["msg"] = "success"
	cnnJson["result"] = rbody

	b, _ := json.Marshal(cnnJson)
	cnnn := string(b)
	fmt.Println("cnnn:%s", cnnn)
	cn_json, _ := simplejson.NewJson([]byte(cnnn))
	cn_body, _ := cn_json.Get("body").Array()

	for _, di := range cn_body {
		//就在这里对di进行类型判断
		newdi, _ := di.(map[string]interface{})
		id := newdi["id"]
		name := newdi["name"]
		fmt.Println(id, name)
	}
}

//构建{"success":true,"address":"haerbin","alias":"xiaoyu","result":[{"id":15,"name":"yufeng15","comment":"kaixin15"},{"id":16,"name":"yufeng16","comment":"kaixin16"},{"id":17,"name":"yufeng17","comment":"kaixin17"},{"id":18,"name":"yufeng18","comment":"kaixin18"}],"errors":[],"messages":[]}
func Test4() {
	//造test数据
	data := make([]Data, 4)
	for i := 0; i < len(data); i++ {
		data[i].Id = i + 15
		data[i].Name = "yufeng" + strconv.Itoa(i+15)
		data[i].Comment = "kaixin" + strconv.Itoa(i+15)
	}
	fmt.Println(data) //[{15 yufeng15 kaixin15} {16 yufeng16 kaixin16} {17 yufeng17 kaixin17} {18 yufeng18 kaixin18}]

	//响应
	response := common.NewResponse(data)
	fmt.Println(response) //{true [{15 yufeng15 kaixin15} {16 yufeng16 kaixin16} {17 yufeng17 kaixin17} {18 yufeng18 kaixin18}] [] []}

	//结构体 -> json
	jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println("data convert to json error", err)
	}
	//注意打印格式
	fmt.Printf("%s\n", jsonData) //{"success":true,"result":[{"id":15,"name":"yufeng15","comment":"kaixin15"},{"id":16,"name":"yufeng16","comment":"kaixin16"},{"id":17,"name":"yufeng17","comment":"kaixin17"},{"id":18,"name":"yufeng18","comment":"kaixin18"}],"errors":[],"messages":[]}

	data2 := make(map[string]interface{})
	data2["success"] = response.Success
	data2["result"] = response.Result
	data2["errors"] = response.Errors
	data2["messages"] = response.Messages
	data2["alias"] = "xiaoyu"
	data2["address"] = "haerbin"
	fmt.Println(data2) //map[address:haerbin alias:xiaoyu]
	jsonData2, err := json.Marshal(data2)
	if err != nil {
		fmt.Println("data2 convert to json error", err)
	}
	fmt.Println(string(jsonData2)) //{"address":"haerbin","alias":"xiaoyu"}
	fmt.Printf("%s\n", jsonData2)  //{"address":"haerbin","alias":"xiaoyu"}

}

//构建数组(结构体 -> 数组)
func Test5() {
	//造test数据
	data := make([]Data, 4)
	for i := 0; i < len(data); i++ {
		data[i].Id = i + 15
		data[i].Name = "yufeng" + strconv.Itoa(i+15)
		data[i].Comment = "kaixin" + strconv.Itoa(i+15)
	}
	fmt.Println(data) //[{15 yufeng15 kaixin15} {16 yufeng16 kaixin16} {17 yufeng17 kaixin17} {18 yufeng18 kaixin18}]

	nameArray := make([]string, len(data))

	for index := range data {
		nameArray[index] = data[index].Name
	}

	fmt.Println(nameArray) //[yufeng15 yufeng16 yufeng17 yufeng18]
}

//拼接多个数组为一个
func ConcatArray() {
	userArray := [...]string{"A", "B", "C", "D"}
	fmt.Println("userArray:", userArray) //[A B C D]

	tmp := []Data{}
	for _, user := range userArray {
		dataArray := make([]Data, 3)
		for i := 0; i < 3; i++ {
			dataArray[i].Id = i
			dataArray[i].Name = user + strconv.Itoa(i) //数字转string

		}

		fmt.Println("len:", len(dataArray))  //3
		fmt.Println("dataArray:", dataArray) //[{0 yufeng0} {1 yufeng1} {2 yufeng2}]
		tmp = append(tmp, dataArray...)

	}
	fmt.Printf("%v", tmp) //[{0 A0} {1 A1} {2 A2} {0 B0} {1 B1} {2 B2} {0 C0} {1 C1} {2 C2} {0 D0} {1 D1} {2 D2}]

}
