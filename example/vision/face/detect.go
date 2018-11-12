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
		"face_field":   "age,beauty,expression,face_shape,gender,glasses,landmark,race,quality,face_type",
		"image_type":   "BASE64",
		"face_type":    "LIVE",
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
