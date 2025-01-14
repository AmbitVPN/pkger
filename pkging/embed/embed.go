package embed

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"io"

	"github.com/ambitvpn/pkger/here"
)

func Decode(src []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(src)))
	_, err := hex.Decode(dst, src)
	if err != nil {
		return nil, err
	}

	r, err := gzip.NewReader(bytes.NewReader(dst))
	if err != nil {
		return nil, err
	}

	bb := &bytes.Buffer{}
	if _, err := io.Copy(bb, r); err != nil {
		return nil, err
	}
	return bb.Bytes(), nil
}

func Encode(b []byte) ([]byte, error) {
	bb := &bytes.Buffer{}
	gz := gzip.NewWriter(bb)

	if _, err := gz.Write(b); err != nil {
		return nil, err
	}

	if err := gz.Flush(); err != nil {
		return nil, err
	}

	if err := gz.Close(); err != nil {
		return nil, err
	}

	s := hex.EncodeToString(bb.Bytes())
	return []byte(s), nil
}

type Data struct {
	Infos map[string]here.Info `json:"infos"`
	Files map[string]File      `json:"files"`
	Here  here.Info            `json:"here"`
}
