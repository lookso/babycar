package server

import (
	netHttp "net/http"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
)

type Response struct {
	ErrorReason string            `json:"error_reason"`
	ErrorMsg    string            `json:"error_msg"`
	MetaData    map[string]string `json:"meta_data"`
	TraceId     string            `json:"trace_id"`
	ServerTime  int64             `json:"server_time"`
	Data        interface{}       `json:"data"`
}

// MyErrorEncoder encodes the error to the HTTP response.
func MyErrorEncoder(w netHttp.ResponseWriter, r *netHttp.Request, err error) {
	se := errors.FromError(err)
	codec, _ := CodecForRequest(r, "Accept")

	key := r.Method + "#" + r.URL.Path
	hiddenMessage := errorHandle.Default
	if errorMessages, ok := errorHandle.Handle[key]; ok {
		for _, errorMessage := range errorMessages.ErrorMessages {
			if errorMessage.ErrorReason == se.Reason {
				hiddenMessage = errorMessage.Message
				break
			}
		}
	}
	if se.Metadata != nil {
		se.Metadata["origin_message"] = se.Message
	} else {
		se.Metadata = map[string]string{
			"origin_message": se.Message,
		}
	}

	response := Response{
		ErrorReason: se.Reason,
		ErrorMsg:    hiddenMessage,
		MetaData:    se.Metadata,
		//		TraceId:     r.Header.Get(common.TraceId),
		ServerTime: time.Now().Unix(),
		Data:       nil,
	}

	body, err := codec.Marshal(response)
	if err != nil {
		w.WriteHeader(netHttp.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(se.Code))
	_, _ = w.Write(body)
}

// CodecForRequest get encoding.Codec via http.Request
func CodecForRequest(r *netHttp.Request, name string) (encoding.Codec, bool) {
	for _, accept := range r.Header[name] {
		codec := encoding.GetCodec(ContentSubtype(accept))
		if codec != nil {
			return codec, true
		}
	}
	return encoding.GetCodec("json"), false
}

func ContentSubtype(contentType string) string {
	left := strings.Index(contentType, "/")
	if left == -1 {
		return ""
	}
	right := strings.Index(contentType, ";")
	if right == -1 {
		right = len(contentType)
	}
	if right < left {
		return ""
	}
	return contentType[left+1 : right]
}
