package v3

import (
	"fmt"
	"github.com/deanzhang/baidu-ai-go-sdk/vision"
	"github.com/imroc/req"
)

const (
	faceDetectUrl = "https://aip.baidubce.com/rest/2.0/face/v3/detect"
)

var (
	typeMap map[string]string
)

/*type FaceResponse struct {
	*req.Resp
}*/

type Location struct {
	Left     float64 `json:"left"`
	Top      float64 `json:"top"`
	Width    float64 `json:"width"`
	Height   float64 `json:"height"`
	Rotation int64   `json:"rotation"`
}

func (l Location) String() string {
	return fmt.Sprintf("[Left: %.2f, Top:%.2f, Width:%.2f, Height:%.2f, Rotation:%d]", l.Left, l.Top, l.Width, l.Height, l.Rotation)
}

type Angle struct {
	Yaw   float64 `json:"yaw"`
	Pitch float64 `json:"pitch"`
	Roll  float64 `json:"roll"`
}

func (a Angle) String() string {
	return fmt.Sprintf("[Yaw:%.2f, Pitch:%.2f, Roll:%.2f]", a.Yaw, a.Pitch, a.Roll)
}

type TypeProb struct {
	Type        string  `json:"type"`
	Probability float64 `json:"probability"`
}

func init() {
	typeMap = map[string]string{"smile": "微笑", "laugh": "大笑", "square": "方形", "triangle": "三角形", "oval": "椭圆", "round": "圆形", "male": "男性", "female": "女性", "common": "普通眼镜", "sun": "墨镜", "yellow": "黄种人", "white": "白种人", "black": "黑种人", "arabs": "阿拉伯人", "human": "真实人脸", "cartoon": "卡通人脸", "none": "none"}
}

func (t TypeProb) String() string {
	return fmt.Sprintf("[Type:%s, 可能性:%.2f]", typeMap[t.Type], t.Probability)
}

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func (p Point) String() string {
	return fmt.Sprintf("[X:%.2f, Y:%.2f]", p.X, p.Y)
}

type Occlusion struct {
	LeftEye    float64 `json:"left_eye"`
	RightEye   float64 `json:"right_eye"`
	Nose       float64 `json:"nose"`
	Mouth      float64 `json:"mouth"`
	LeftCheek  float64 `json:"left_cheek"`
	RightCheek float64 `json:"right_cheek"`
	Chin       float64 `json:"chin_contour"`
}

func (o Occlusion) String() string {
	return fmt.Sprintf("[左眼:%.2f, 右眼:%.2f, 鼻子:%.2f, 嘴巴:%.2f, 左脸颊:%.2f,右脸颊:%.2f,下巴:%.2f", o.LeftEye, o.RightEye, o.Nose, o.Mouth, o.LeftCheek, o.RightCheek, o.Chin)
}

type Quality struct {
	Occlusion    Occlusion `json:"occlusion"`
	Blur         float64   `json:"blur"`
	Illumination float64   `json:"illumination"`
	Completeness int64     `json:"completeness"`
}

func (q Quality) String() string {
	return fmt.Sprintf("[人脸各部分遮挡概率:%v], 模糊程度:%.2f, 光照程度:%.2f, 完整度:%d", q.Occlusion, q.Blur, q.Illumination, q.Completeness)
}

type FaceList struct {
	FaceToken       string    `json:"face_token"`
	Location        Location  `json:"location"`
	FaceProbability float64   `json:"face_probability"`
	Angle           Angle     `json:"angle"`
	Age             float64   `json:"age"`
	Beauty          float64   `json:"beauty"`
	Expression      TypeProb  `json:"expression"`
	FaceShape       TypeProb  `json:"face_shape"`
	Gender          TypeProb  `json:"gender"`
	Glasses         TypeProb  `json:"glasses"`
	Race            TypeProb  `json:"race"`
	FaceType        TypeProb  `json:"face_type"`
	Landmark        [4]Point  `json:"landmark"`
	Landmark72      [72]Point `json:"landmark72"`
	Quality         Quality   `json:"quality"`
}

func (f FaceList) String() string {
	return fmt.Sprintf("\n---[Token:%s, 人脸置信度:%.2f, 年龄:%.2f, 颜值:%.2f,\n在图片中的位置:%v,\n人脸旋转角度:%v,\n表情:%v, 脸型:%v, 性别:%v\n眼镜:%v, 人种:%v, 人脸或卡通:%v,\n人脸图片质量:%v,\n4关键点:%v\n72关键点:%v]---\n",
		f.FaceToken, f.FaceProbability, f.Age, f.Beauty, f.Location, f.Angle, f.Expression, f.FaceShape, f.Gender, f.Glasses, f.Race, f.FaceType, f.Quality, f.Landmark, f.Landmark72)
}

type FaceDetectResp struct {
	FaceNum  int        `json:"face_num"`
	FaceList []FaceList `json:"face_list"`
}

func (f FaceDetectResp) String() string {
	return fmt.Sprintf("[人脸数量：%d, 人脸列表:%v]", f.FaceNum, f.FaceList)
}

type FaceResponse struct {
	ErrCode   int            `json:"error_code"`
	ErrMsg    string         `json:"error_msg"`
	LogId     int            `json:"log_id"`
	TimeStamp int            `json:"timestamp"`
	Cached    int            `json:"cached"`
	Result    FaceDetectResp `json:"result"`
}

func (f FaceResponse) String() string {
	return fmt.Sprintf("[错误编码：%d, 识别结果:%s]", f.ErrCode, f.Result.String())
}

func (fc *FaceClient) DetectAndAnalysis(image *vision.Image, options map[string]interface{}) (*FaceResponse, error) {
	var response FaceResponse
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

	resp.ToJSON(&response)

	/*return &FaceResponse{
		Resp: resp,
	}, nil*/
	return &response, nil

}
