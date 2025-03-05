package zstd

import (
	"io"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	klauszstd "github.com/klauspost/compress/zstd"
)

type zstd struct {
}

var _ promhttp.CompressionEngine = (*zstd)(nil)

func init() {
	err := promhttp.RegisterCompressionEngine("zstd", &zstd{})
	if err != nil {
		panic(err)
	}
}

func (zstd *zstd) NewWriter(rw io.Writer) (io.WriteCloser, error) {
	// TODO(mrueg): Replace klauspost/compress with stdlib implementation once https://github.com/golang/go/issues/62513 is implemented.
	encoder, err := klauszstd.NewWriter(rw, klauszstd.WithEncoderLevel(klauszstd.SpeedFastest))
	if err != nil {
		return nil, err
	}
	return encoder, nil
}

func (zstd *zstd) NewReader(r io.Reader) (io.ReadCloser, error) {
	decoder, err := klauszstd.NewReader(r)
	if err != nil {
		return nil, err
	}
	return decoder.IOReadCloser(), nil
}
