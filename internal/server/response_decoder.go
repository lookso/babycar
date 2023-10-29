package server

import (
	"encoding/json"
	netHttp "net/http"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"
	kjson "github.com/go-kratos/kratos/v2/encoding/json"
	"google.golang.org/protobuf/encoding/protojson"
)

func MyResponseEncoder(w netHttp.ResponseWriter, r *netHttp.Request, i interface{}) error {
	kjson.MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	reply := Response{
		ErrorReason: "success",
		ErrorMsg:    "",
		MetaData:    nil,
		//		TraceId:     r.Header.Get(common.TraceId),
		ServerTime: time.Now().Unix(),
	}
	codec := encoding.GetCodec("json")
	data, err := codec.Marshal(i)
	err = json.Unmarshal(data, &reply.Data)
	if err != nil {
		return err
	}

	data, err = codec.Marshal(reply)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
