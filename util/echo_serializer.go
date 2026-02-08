package util

import (
	"encoding/json"
	"io"

	"github.com/labstack/echo/v5"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type ProtobufJsonEchoSerializer struct{}

func (s ProtobufJsonEchoSerializer) Serialize(c *echo.Context, target any, indent string) error {
	if pb, ok := target.(proto.Message); ok {
		marshaler := protojson.MarshalOptions{
			UseProtoNames: true,
			Indent:        indent,
		}
		data, err := marshaler.Marshal(pb)
		if err != nil {
			return err
		}
		_, err = c.Response().Write(data)
		return err
	} else {
		enc := json.NewEncoder(c.Response())
		if indent != "" {
			enc.SetIndent("", indent)
		}
		return enc.Encode(target)
	}
}

func (s ProtobufJsonEchoSerializer) Deserialize(c *echo.Context, target any) error {
	if pb, ok := target.(proto.Message); ok {
		unmarshaler := protojson.UnmarshalOptions{
			DiscardUnknown: true,
		}
		data, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return echo.ErrBadRequest.Wrap(err)
		}
		err = unmarshaler.Unmarshal(data, pb)
		if err != nil {
			return echo.ErrBadRequest.Wrap(err)
		}
		return nil
	} else {
		if err := json.NewDecoder(c.Request().Body).Decode(target); err != nil {
			return echo.ErrBadRequest.Wrap(err)
		}
		return nil
	}
}
