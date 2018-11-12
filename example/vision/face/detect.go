package main

import (
	"fmt"

	"github.com/deanzhang/baidu-ai-go-sdk/vision"
	"github.com/deanzhang/baidu-ai-go-sdk/vision/face/v3"
)

func DetectAndAnalysis() {
	client := v3.NewFaceClient(APIKEY, APISECRET)
	options := map[string]interface{}{
		"max_face_num": 10,
		"face_fields":  "age,beauty,expression,faceshape,gender,glasses,landmark,race,qualities",
		"image_type": "BASE64",
		"face_type": "LIVE",
	}
	rs, err := client.DetectAndAnalysis(
		vision.MustFromFile("face.jpg"),
		options,
	)
	if err != nil {
		panic(err)
	}
	fmt.Println(rs.ToString())
}
