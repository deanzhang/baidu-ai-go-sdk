package v3

import (
	"github.com/deanzhang/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
)

const (
	faceDetectUrl = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
)

type FaceResponse struct {
	*req.Resp
}

func (fc *FaceClient) DetectAndAnalysis(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {

	if err := fc.Auth(); err != nil {
		return nil, err
	}

	url := faceDetectUrl + "?access_token=" + fc.AccessToken

	base64Str, err := image.Base64Encode()
	if err != nil {
		return nil, err
	}
	options["image"] = base64Str

	header := req.Header{
		"Content-Type": "application/json",
	}

	resp, err := req.Post(url, req.BodyJSON(options), header)
	if err != nil {
		return nil, err
	}

	return &FaceResponse{
		Resp: resp,
	}, nil

}
