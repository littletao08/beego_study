package utils

import (
	"github.com/qiniu/api.v6/rs"
	"testing"
	"github.com/qiniu/api.v6/io"
	"log"
	"os"
	"github.com/qiniu/api.v6/conf"
)

func init() {
	//https://portal.qiniu.com/setting/key  秘钥
	conf.ACCESS_KEY = "YOUR_APP_ACCESS_KEY"
	conf.SECRET_KEY = "YOUR_APP_SECRET_KEY"
}


type PutExtra struct {
	Params   map[string]string
	MimeType string
	Crc32    uint32
	CheckCrc uint32
}

func TestUpload(t  *testing.T) {
	var err error
	var ret io.PutRet
	var extra = &io.PutExtra{}
    r,err := os.Open("/users/xxxx/Desktop/psu.jpeg")
	upToken := upToken("threeperson")
	err = io.Put(nil, &ret, upToken,"heads/1001/psu.jpeg", r, extra)

	if err != nil {
		log.Print("io.Put failed:", err)
		return
	}

	log.Print("result:",ret.Hash, ret.Key)
}


//bucketName 空间名称
func upToken(bucketName string) string {
	putPolicy := rs.PutPolicy{
		Scope:         bucketName,
	}
	return putPolicy.Token(nil)
}
