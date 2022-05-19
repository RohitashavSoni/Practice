package main

import (
	"io"
	//"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// Hello world, the web server

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, `{"assets":[{"id":"0","name":"edgeAssetName","parentId":"","hasChildren":true,"assets":[{"id":"5502a47463e5404b9f496632fc8f7f1a","name":"edgeAssetChield1","parentId":"0","hasChildren":true,"assets":[{"id":"0063d36ac618494fa2c8b93c25349c8b","name":"edgeAssetChield1-1","parentId":"5502a47463e5404b9f496632fc8f7f1a","hasChildren":false,"assets":[],"aspects":[{"id":"144325a1154e4b5aa5e306ac855016dc","name":"Aspect","assetId":"0063d36ac618494fa2c8b93c25349c8b","variables":[{"id":"28d326b4872d48d7a3c1a517af17e3db","name":"tag1","dataType":"Bool","assetId":"0063d36ac618494fa2c8b93c25349c8b","aspectId":"144325a1154e4b5aa5e306ac855016dc","aspectName":"Aspect","unit":"","topic":"plc1::tag1::4::4"}]}]}],"aspects":[]},{"id":"3b7853fdf1b64e3a81232ea3008a8145","name":"edgeAssetChield2","parentId":"0","hasChildren":false,"assets":[],"aspects":[]}],"aspects":[{"id":"1655a06970ea496cbea944a7f55f5716","name":"Aspect1","assetId":"0","variables":[{"id":"97b63eb07ff24757bb96799c90db22d9","name":"tag5","dataType":"Float","assetId":"0","aspectId":"1655a06970ea496cbea944a7f55f5716","aspectName":"Aspect1","unit":"","topic":"plc1::tag5::4::4"},{"id":"cb68406aa32f40fc9f15d8ba09576395","name":"tag3","dataType":"UInt8","assetId":"0","aspectId":"1655a06970ea496cbea944a7f55f5716","aspectName":"Aspect1","unit":"","topic":"plc1::tag3::4::4"},{"id":"c2c99fbd12724a8d9601f81fc1314e66","name":"tag1","dataType":"Bool","assetId":"0","aspectId":"1655a06970ea496cbea944a7f55f5716","aspectName":"Aspect1","unit":"","topic":"plc1::tag1::4::4"},{"id":"c9392f3862d048a3a986cad1855d5d04","name":"tag4","dataType":"Int32","assetId":"0","aspectId":"1655a06970ea496cbea944a7f55f5716","aspectName":"Aspect1","unit":"","topic":"plc1::tag4::4::4"},{"id":"e20d916915774c84bc15214f28bf48fa","name":"tag2","dataType":"Int16","assetId":"0","aspectId":"1655a06970ea496cbea944a7f55f5716","aspectName":"Aspect1","unit":"degree","topic":"plc1::tag2::4::4"}]}]}]}`)
	}

	http.HandleFunc("/dataServiceModel", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*
func apicall() []byte{
	response, _ := http.Get("http://ie-dataservice-connector:4201")
	log.Println("--->", response.Status)
	responseData, _ := ioutil.ReadAll(response.Body)
	return responseData

}*/
