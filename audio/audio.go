package audio

import (
	"io"
	"os"
)

type Audio interface {
	ReadAudio() ([]byte, error)
	ConvertAudio(data []byte, format string) ([]byte, error)
	SaveAudio(data []byte, path string) error
}

type FLAC struct {
	path string
}

func (f FLAC) ReadAudio() ([]byte, error) {
	// 打开FLAC文件
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 读取文件头部信息，一般为 128 字节
	header := make([]byte, 128)
	if _, err = file.Read(header); err != nil {
		return nil, err
	}

	// 读取音频数据
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (f FLAC) ConvertAudio(data []byte, format string) ([]byte, error) {
	// 实现转换逻辑，这里仅做示例
	convertedData := []byte("FLAC to " + format + " conversion")

	return convertedData, nil
}

func (f FLAC) SaveAudio(data []byte, path string) error {
	// 创建保存文件
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	// 将数据写入文件
	if _, err = file.Write(data); err != nil {
		return err
	}

	return nil
}
