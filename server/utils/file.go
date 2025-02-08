package utils

import (
	"bytes"
	"encoding/base64"
	"io"

	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/pocketbase/pocketbase/tools/security"
)

func GetImageBase64(app *core.App, r *core.Record, col string) (*string, error) {
	buf, err := getFileBuffer(app, r, col)
	if err != nil {
		return nil, err
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	return &base64Str, nil
}

func NewFile(base64Str *string) (*filesystem.File, error) {
	buf, err := base64.StdEncoding.DecodeString(*base64Str)
	if err != nil {
		return nil, err
	}

	return filesystem.NewFileFromBytes(buf, security.RandomString(16))
}

func getFileBuffer(app *core.App, r *core.Record, col string) (*bytes.Buffer, error) {
	a := (*app)

	fsys, err := a.NewFilesystem()
	if err != nil {
		return nil, err
	}
	defer fsys.Close()

	key := r.BaseFilesPath() + "/" + r.GetString(col)
	reader, err := fsys.GetFile(key)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	content := new(bytes.Buffer)
	_, err = io.Copy(content, reader)
	if err != nil {
		return nil, err
	}

	return content, err
}
