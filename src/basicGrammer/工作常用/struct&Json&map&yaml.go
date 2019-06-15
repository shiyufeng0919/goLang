package 工作常用

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"gopkg.in/yaml.v2"
)

func ExecStructJsonMapYaml(){

	json2Map() //json->map and  map->json

	fmt.Println("json->struct and struct->json")

	json2struct1() //json->struct and struct->json

	json2struct2()

	json2struct3()

	struct2json() //struct - >json
}

//struct to json

var str = `{"kind":"ConfigMap","apiVersion":"v1","metadata":{"name":"prom-prometheus-server","namespace":"prometheus","selfLink":"/api/v1/namespaces/prometheus/configmaps/prom-prometheus-server","uid":"dc2299f3-446f-11e9-96a1-40f2e9cde822","resourceVersion":"311529374","creationTimestamp":"2019-03-12T02:38:03Z","labels":{"app":"prometheus","chart":"prometheus-7.4.4","component":"server","heritage":"Tiller","release":"prom"}},"data":{"alerts":"{}\n","prometheus.yml":"global:\n  evaluation_interval: 1m\n  scrape_interval: 1m\n  scrape_timeout: 10s\n  \nrule_files:\n- /etc/config/rules\n- /etc/config/alerts\nscrape_configs:\n- job_name: prometheus\n  static_configs:\n  - targets:\n    - localhost:9090\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\n  job_name: kubernetes-apiservers\n  kubernetes_sd_configs:\n  - role: endpoints\n  relabel_configs:\n  - action: keep\n    regex: default;kubernetes;https\n    source_labels:\n    - __meta_kubernetes_namespace\n    - __meta_kubernetes_service_name\n    - __meta_kubernetes_endpoint_port_name\n  scheme: https\n  tls_config:\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\n    insecure_skip_verify: true\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\n  job_name: kubernetes-nodes\n  kubernetes_sd_configs:\n  - role: node\n  relabel_configs:\n  - action: labelmap\n    regex: __meta_kubernetes_node_label_(.+)\n  - replacement: kubernetes.default.svc:443\n    target_label: __address__\n  - regex: (.+)\n    replacement: /api/v1/nodes/${1}/proxy/metrics\n    source_labels:\n    - __meta_kubernetes_node_name\n    target_label: __metrics_path__\n  scheme: https\n  tls_config:\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\n    insecure_skip_verify: true\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\n  job_name: kubernetes-nodes-cadvisor\n  kubernetes_sd_configs:\n  - role: node\n  relabel_configs:\n  - action: labelmap\n    regex: __meta_kubernetes_node_label_(.+)\n  - replacement: kubernetes.default.svc:443\n    target_label: __address__\n  - regex: (.+)\n    replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor\n    source_labels:\n    - __meta_kubernetes_node_name\n    target_label: __metrics_path__\n  scheme: https\n  tls_config:\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\n    insecure_skip_verify: true\n- job_name: kubernetes-service-endpoints\n  kubernetes_sd_configs:\n  - role: endpoints\n  relabel_configs:\n  - action: keep\n    regex: true\n    source_labels:\n    - __meta_kubernetes_service_annotation_prometheus_io_scrape\n  - action: replace\n    regex: (https?)\n    source_labels:\n    - __meta_kubernetes_service_annotation_prometheus_io_scheme\n    target_label: __scheme__\n  - action: replace\n    regex: (.+)\n    source_labels:\n    - __meta_kubernetes_service_annotation_prometheus_io_path\n    target_label: __metrics_path__\n  - action: replace\n    regex: ([^:]+)(?::\\d+)?;(\\d+)\n    replacement: $1:$2\n    source_labels:\n    - __address__\n    - __meta_kubernetes_service_annotation_prometheus_io_port\n    target_label: __address__\n  - action: labelmap\n    regex: __meta_kubernetes_service_label_(.+)\n  - action: replace\n    source_labels:\n    - __meta_kubernetes_namespace\n    target_label: kubernetes_namespace\n  - action: replace\n    source_labels:\n    - __meta_kubernetes_service_name\n    target_label: kubernetes_name\n  - action: replace\n    source_labels:\n    - __meta_kubernetes_pod_node_name\n    target_label: kubernetes_node\n- honor_labels: true\n  job_name: prometheus-pushgateway\n  kubernetes_sd_configs:\n  - role: service\n  relabel_configs:\n  - action: keep\n    regex: pushgateway\n    source_labels:\n    - __meta_kubernetes_service_annotation_prometheus_io_probe\n- job_name: kubernetes-services\n  kubernetes_sd_configs:\n  - role: service\n  metrics_path: /probe\n  params:\n    module:\n    - http_2xx\n  relabel_configs:\n  - action: keep\n    regex: true\n    source_labels:\n    - __meta_kubernetes_service_annotation_prometheus_io_probe\n  - source_labels:\n    - __address__\n    target_label: __param_target\n  - replacement: blackbox\n    target_label: __address__\n  - source_labels:\n    - __param_target\n    target_label: instance\n  - action: labelmap\n    regex: __meta_kubernetes_service_label_(.+)\n  - source_labels:\n    - __meta_kubernetes_namespace\n    target_label: kubernetes_namespace\n  - source_labels:\n    - __meta_kubernetes_service_name\n    target_label: kubernetes_name\n- job_name: kubernetes-pods\n  kubernetes_sd_configs:\n  - role: pod\n  relabel_configs:\n  - action: keep\n    regex: true\n    source_labels:\n    - __meta_kubernetes_pod_annotation_prometheus_io_scrape\n  - action: replace\n    regex: (.+)\n    source_labels:\n    - __meta_kubernetes_pod_annotation_prometheus_io_path\n    target_label: __metrics_path__\n  - action: replace\n    regex: ([^:]+)(?::\\d+)?;(\\d+)\n    replacement: $1:$2\n    source_labels:\n    - __address__\n    - __meta_kubernetes_pod_annotation_prometheus_io_port\n    target_label: __address__\n  - action: labelmap\n    regex: __meta_kubernetes_pod_label_(.+)\n  - action: replace\n    source_labels:\n    - __meta_kubernetes_namespace\n    target_label: kubernetes_namespace\n  - action: replace\n    source_labels:\n    - __meta_kubernetes_pod_name\n    target_label: kubernetes_pod_name\n- job_name: node\n  static_configs:\n  - labels:\n      cluster: 3.cn\n    targets:\n    - 192.168.169.139:9100\n    - 192.168.170.138:9100\n    - 192.168.170.137:9100\n    - 192.168.177.103:9100\n    - 192.168.177.104:9100\n    - 192.168.182.42:9100\n    - 192.168.182.43:9100\n","rules":"{}\n"}}`

//json to struct
func json2struct1(){

	fmt.Println("=======json2struct1========")

	body:=[]byte(str)

	promeYaml := PromeYaml{}

	err := yaml.Unmarshal(body, &promeYaml)

	if err != nil {
		logs.Error("cannot unmarshal data: %v", err)
		return
	}

	fmt.Println("promeYaml1:",promeYaml.PromeData.PrometheusYaml)
}

//json to struct
func json2struct2(){

	fmt.Println("=======json2struct2========")

	body:=[]byte(str)

	promeYaml:=PromeYaml{}

	err:=json.Unmarshal(body,&promeYaml)

	if err != nil {
		logs.Error("cannot unmarshal data: %v", err)
		return
	}

	fmt.Println("promeYaml2:",promeYaml.PromeData.PrometheusYaml)
}

//json to struct
func json2struct3(){

	fmt.Println("=======json2struct3========")

	strJson,err:=json.Marshal(str)

	if err != nil {
		logs.Error("cannot marshal data: %v", err)
		return
	}

	fmt.Println("strJson:",string(strJson))  //此时json中会带有/符号

	str:=string(strJson)

	fmt.Println("str:",str)

	proyaml := PromeYaml{}

	str = "{\"kind\":\"ConfigMap\",\"apiVersion\":\"v1\",\"metadata\":{\"name\":\"prom-prometheus-server\",\"namespace\":\"prometheus\",\"selfLink\":\"/api/v1/namespaces/prometheus/configmaps/prom-prometheus-server\",\"uid\":\"dc2299f3-446f-11e9-96a1-40f2e9cde822\",\"resourceVersion\":\"311529374\",\"creationTimestamp\":\"2019-03-12T02:38:03Z\",\"labels\":{\"app\":\"prometheus\",\"chart\":\"prometheus-7.4.4\",\"component\":\"server\",\"heritage\":\"Tiller\",\"release\":\"prom\"}},\"data\":{\"alerts\":\"{}\\n\",\"prometheus.yml\":\"global:\\n  evaluation_interval: 1m\\n  scrape_interval: 1m\\n  scrape_timeout: 10s\\n  \\nrule_files:\\n- /etc/config/rules\\n- /etc/config/alerts\\nscrape_configs:\\n- job_name: prometheus\\n  static_configs:\\n  - targets:\\n    - localhost:9090\\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\\n  job_name: kubernetes-apiservers\\n  kubernetes_sd_configs:\\n  - role: endpoints\\n  relabel_configs:\\n  - action: keep\\n    regex: default;kubernetes;https\\n    source_labels:\\n    - __meta_kubernetes_namespace\\n    - __meta_kubernetes_service_name\\n    - __meta_kubernetes_endpoint_port_name\\n  scheme: https\\n  tls_config:\\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\\n    insecure_skip_verify: true\\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\\n  job_name: kubernetes-nodes\\n  kubernetes_sd_configs:\\n  - role: node\\n  relabel_configs:\\n  - action: labelmap\\n    regex: __meta_kubernetes_node_label_(.+)\\n  - replacement: kubernetes.default.svc:443\\n    target_label: __address__\\n  - regex: (.+)\\n    replacement: /api/v1/nodes/${1}/proxy/metrics\\n    source_labels:\\n    - __meta_kubernetes_node_name\\n    target_label: __metrics_path__\\n  scheme: https\\n  tls_config:\\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\\n    insecure_skip_verify: true\\n- bearer_token_file: /var/run/secrets/kubernetes.io/serviceaccount/token\\n  job_name: kubernetes-nodes-cadvisor\\n  kubernetes_sd_configs:\\n  - role: node\\n  relabel_configs:\\n  - action: labelmap\\n    regex: __meta_kubernetes_node_label_(.+)\\n  - replacement: kubernetes.default.svc:443\\n    target_label: __address__\\n  - regex: (.+)\\n    replacement: /api/v1/nodes/${1}/proxy/metrics/cadvisor\\n    source_labels:\\n    - __meta_kubernetes_node_name\\n    target_label: __metrics_path__\\n  scheme: https\\n  tls_config:\\n    ca_file: /var/run/secrets/kubernetes.io/serviceaccount/ca.crt\\n    insecure_skip_verify: true\\n- job_name: kubernetes-service-endpoints\\n  kubernetes_sd_configs:\\n  - role: endpoints\\n  relabel_configs:\\n  - action: keep\\n    regex: true\\n    source_labels:\\n    - __meta_kubernetes_service_annotation_prometheus_io_scrape\\n  - action: replace\\n    regex: (https?)\\n    source_labels:\\n    - __meta_kubernetes_service_annotation_prometheus_io_scheme\\n    target_label: __scheme__\\n  - action: replace\\n    regex: (.+)\\n    source_labels:\\n    - __meta_kubernetes_service_annotation_prometheus_io_path\\n    target_label: __metrics_path__\\n  - action: replace\\n    regex: ([^:]+)(?::\\\\d+)?;(\\\\d+)\\n    replacement: $1:$2\\n    source_labels:\\n    - __address__\\n    - __meta_kubernetes_service_annotation_prometheus_io_port\\n    target_label: __address__\\n  - action: labelmap\\n    regex: __meta_kubernetes_service_label_(.+)\\n  - action: replace\\n    source_labels:\\n    - __meta_kubernetes_namespace\\n    target_label: kubernetes_namespace\\n  - action: replace\\n    source_labels:\\n    - __meta_kubernetes_service_name\\n    target_label: kubernetes_name\\n  - action: replace\\n    source_labels:\\n    - __meta_kubernetes_pod_node_name\\n    target_label: kubernetes_node\\n- honor_labels: true\\n  job_name: prometheus-pushgateway\\n  kubernetes_sd_configs:\\n  - role: service\\n  relabel_configs:\\n  - action: keep\\n    regex: pushgateway\\n    source_labels:\\n    - __meta_kubernetes_service_annotation_prometheus_io_probe\\n- job_name: kubernetes-services\\n  kubernetes_sd_configs:\\n  - role: service\\n  metrics_path: /probe\\n  params:\\n    module:\\n    - http_2xx\\n  relabel_configs:\\n  - action: keep\\n    regex: true\\n    source_labels:\\n    - __meta_kubernetes_service_annotation_prometheus_io_probe\\n  - source_labels:\\n    - __address__\\n    target_label: __param_target\\n  - replacement: blackbox\\n    target_label: __address__\\n  - source_labels:\\n    - __param_target\\n    target_label: instance\\n  - action: labelmap\\n    regex: __meta_kubernetes_service_label_(.+)\\n  - source_labels:\\n    - __meta_kubernetes_namespace\\n    target_label: kubernetes_namespace\\n  - source_labels:\\n    - __meta_kubernetes_service_name\\n    target_label: kubernetes_name\\n- job_name: kubernetes-pods\\n  kubernetes_sd_configs:\\n  - role: pod\\n  relabel_configs:\\n  - action: keep\\n    regex: true\\n    source_labels:\\n    - __meta_kubernetes_pod_annotation_prometheus_io_scrape\\n  - action: replace\\n    regex: (.+)\\n    source_labels:\\n    - __meta_kubernetes_pod_annotation_prometheus_io_path\\n    target_label: __metrics_path__\\n  - action: replace\\n    regex: ([^:]+)(?::\\\\d+)?;(\\\\d+)\\n    replacement: $1:$2\\n    source_labels:\\n    - __address__\\n    - __meta_kubernetes_pod_annotation_prometheus_io_port\\n    target_label: __address__\\n  - action: labelmap\\n    regex: __meta_kubernetes_pod_label_(.+)\\n  - action: replace\\n    source_labels:\\n    - __meta_kubernetes_namespace\\n    target_label: kubernetes_namespace\\n  - action: replace\\n    source_labels:\\n    - __meta_kubernetes_pod_name\\n    target_label: kubernetes_pod_name\\n- job_name: node\\n  static_configs:\\n  - labels:\\n      cluster: 3.cn\\n    targets:\\n    - 192.168.169.139:9100\\n    - 192.168.170.138:9100\\n    - 192.168.170.137:9100\\n    - 192.168.177.103:9100\\n    - 192.168.177.104:9100\\n    - 192.168.182.42:9100\\n    - 192.168.182.43:9100\\n\",\"rules\":\"{}\\n\"}}"

	//问题:str直接赋值可以，但是经过上述str:=string(strJson)转会报错
	if err := json.Unmarshal([]byte(str), &proyaml); err == nil {
		fmt.Println(proyaml.PromeData.PrometheusYaml)
	} else {
		fmt.Println(err)
	}

}


func struct2json(){

	fmt.Println("=========struct2json=================")

	strJson,err:=json.Marshal([]byte(str))

	fmt.Println("strJson:",string(strJson))

	proYaml := PromeYaml{}

	//struct -> json
	proJson,err:=json.Marshal(proYaml)

	if err !=nil {
		logs.Error("struct->json错误",err.Error())
		fmt.Errorf("struct->json错误",err.Error())
		return
	}

	fmt.Printf("struct->json: %s",string(proJson))

	//json -> struct
	err=json.Unmarshal(proJson,&proYaml)

	if err != nil{
		logs.Error("json->struct异常: %S",err.Error())
		fmt.Errorf("json->struct错误",err.Error())
		return
	}

	fmt.Println("json->struct: %s",proYaml.PromeData.PrometheusYaml)
}

//json to map And map to json
func json2Map(){

	fmt.Println("=======json2Map========")


	jsonStr:=`{"kind":"Deployment","apiVersion":"apps/v1","metadata":{"name":"prom-prometheus-server","namespace":"prometheus","selfLink":"/apis/apps/v1/namespaces/prometheus/deployments/prom-prometheus-server","uid":"4fca3310-8c1a-11e9-a1f8-40f2e9ce08c2","resourceVersion":"6865","generation":6,"creationTimestamp":"2019-06-11T07:27:04Z","labels":{"app":"prometheus","chart":"prometheus-7.4.4","component":"server","heritage":"Tiller","release":"prom"},"annotations":{"deployment.kubernetes.io/revision":"1","field.cattle.io/publicEndpoints":"[{\"addresses\":[\"192.168.169.139\"],\"port\":30100,\"protocol\":\"TCP\",\"serviceName\":\"prometheus:prom-prometheus-server\",\"allNodes\":true}]"}},"spec":{"replicas":1,"selector":{"matchLabels":{"app":"prometheus","component":"server","release":"prom"}},"template":{"metadata":{"creationTimestamp":null,"labels":{"app":"prometheus","component":"server","release":"prom"}},"spec":{"volumes":[{"name":"config-volume","configMap":{"name":"prom-prometheus-server","defaultMode":420}},{"name":"storage-volume","persistentVolumeClaim":{"claimName":"prom-prometheus-server"}}],"containers":[{"name":"prometheus-server-configmap-reload","image":"jimmidyson/configmap-reload:v0.2.2","args":["--volume-dir=/etc/config","--webhook-url=http://127.0.0.1:9090/-/reload"],"resources":{},"volumeMounts":[{"name":"config-volume","readOnly":true,"mountPath":"/etc/config"}],"terminationMessagePath":"/dev/termination-log","terminationMessagePolicy":"File","imagePullPolicy":"IfNotPresent"},{"name":"prometheus-server","image":"prom/prometheus:v2.5.0","args":["--config.file=/etc/config/prometheus.yml","--storage.tsdb.path=/data","--web.console.libraries=/etc/prometheus/console_libraries","--web.console.templates=/etc/prometheus/consoles","--web.enable-lifecycle"],"ports":[{"containerPort":9090,"protocol":"TCP"}],"resources":{},"volumeMounts":[{"name":"config-volume","mountPath":"/etc/config"},{"name":"storage-volume","mountPath":"/data"}],"livenessProbe":{"httpGet":{"path":"/-/healthy","port":9090,"scheme":"HTTP"},"initialDelaySeconds":30,"timeoutSeconds":30,"periodSeconds":10,"successThreshold":1,"failureThreshold":3},"readinessProbe":{"httpGet":{"path":"/-/ready","port":9090,"scheme":"HTTP"},"initialDelaySeconds":30,"timeoutSeconds":30,"periodSeconds":10,"successThreshold":1,"failureThreshold":3},"terminationMessagePath":"/dev/termination-log","terminationMessagePolicy":"File","imagePullPolicy":"IfNotPresent"}],"restartPolicy":"Always","terminationGracePeriodSeconds":300,"dnsPolicy":"ClusterFirst","serviceAccountName":"prom-prometheus-server","serviceAccount":"prom-prometheus-server","securityContext":{"runAsUser":0},"schedulerName":"default-scheduler"}},"strategy":{"type":"RollingUpdate","rollingUpdate":{"maxUnavailable":1,"maxSurge":1}},"revisionHistoryLimit":10,"progressDeadlineSeconds":600},"status":{"observedGeneration":6,"replicas":1,"updatedReplicas":1,"readyReplicas":1,"availableReplicas":1,"conditions":[{"type":"Available","status":"True","lastUpdateTime":"2019-06-11T07:27:04Z","lastTransitionTime":"2019-06-11T07:27:04Z","reason":"MinimumReplicasAvailable","message":"Deployment has minimum availability."},{"type":"Progressing","status":"True","lastUpdateTime":"2019-06-11T07:27:46Z","lastTransitionTime":"2019-06-11T07:27:04Z","reason":"NewReplicaSetAvailable","message":"ReplicaSet \"prom-prometheus-server-8498ccdf84\" has successfully progressed."}]}}`

	//引入第三方包可快速要找json中的key值
	//name := gojsonq.New().JSONString(jsonStr).Find("spec.replicas")
	//
	//println(int(name.(float64)))
	//
	//println(strconv.FormatFloat(name.(float64),'f',0,64))

	var mapResult map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &mapResult) //json转map
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}

	for key,value:=range mapResult { //循环处理map数据

		if key == "spec" {

			val:=value.(map[string]interface{})

			for k,_:=range val{

				if k=="replicas" {

					mapResult[k]=2  //修改map中key的值

					break

				}

			}
		}

		if key == "metadata" {

			val:=value.(map[string]interface{})

			for k,_:=range val{

				if k=="uid" || k=="resourceVersion" || k=="creationTimestamp" {
					delete(val,k)  //删除map中的key
				}

			}

		}

		if key == "status" {
			delete(mapResult,key) //删除map中的key
		}

	}

	json, err := json.Marshal(mapResult)  //map转json

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(json))
}

//configmaps配置文件
type PromeYaml struct {
	Pkind              string `yaml:"kind" json:"kind"`
	PApiVersion        string `yaml:"apiVersion" json:"apiVersion"`
	PromeMetadata `yaml:"metadata" json:"metadata"`
	PromeData     `yaml:"data" json:"data"`
}
type PromeMetadata struct {
	Name      string `yaml:"name" json:"name"`
	Namespace string `yaml:"namespace" json:"namespace"`
	SelfLink  string `yaml:"selfLink" json:"selfLink"`
	Uid               string `yaml:"uid" json:"uid"`
	ResourceVersion   string `yaml:"resourceVersion" json:"resourceVersion"`
	CreationTimestamp string `yaml:"creationTimestamp" json:"creationTimestamp"`
	PromeLabels `yaml:"labels" json:"labels"`
}

type PromeLabels struct {
	App       string `yaml:"app" json:"app"`
	Chart     string `yaml:"chart" json:"chart"`
	Component string `yaml:"component" json:"component"`
	Heritage  string `yaml:"heritage" json:"heritage"`
	Release   string `yaml:"release" json:"release"`
}
type PromeData struct {
	Alerts         string `yaml:"alerts" json:"alerts"`
	PrometheusYaml string `yaml:"prometheus.yml" json:"prometheus.yml"`
	Rules          string `yaml:"rules" json:"rules"`
}



