package xerrors

import (
	"bytes"
	"embed"
	"io"

	"github.com/gocarina/gocsv"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
)

//go:embed *.csv
var errorsFS embed.FS

func LoadErrors() {
	csvFiles, err := errorsFS.ReadDir(".")
	if err != nil {
		zap.L().Fatal("failed to read error files", zap.Error(err))
	}

	var (
		csvFileNames []string
	)
	for _, file := range csvFiles {
		if file.IsDir() {
			continue
		}

		csvFileNames = append(csvFileNames, file.Name())
	}

	for _, file := range csvFileNames {
		byteData, errFs := errorsFS.ReadFile(file)
		if errFs != nil {
			zap.L().Fatal("failed to read error file", zap.Error(err))
		}

		reader := bytes.NewReader(byteData)
		err = parseAndLoadErrorFile(reader)
		if err != nil {
			zap.L().Fatal("failed to parse error file", zap.Error(err))
		}
	}
}

type csvErrorStructure struct {
	Method    string `csv:"method"`
	ErrorKey  string `csv:"error_key"`
	GrpcCode  uint32 `csv:"grpc_code"`
	ErrorCode int32  `csv:"error_code"`
	Uz        string `csv:"uz"`
	Ru        string `csv:"ru"`
	En        string `csv:"en"`
}

func parseAndLoadErrorFile(reader io.Reader) error {
	var (
		messages []*csvErrorStructure
	)

	if err := gocsv.Unmarshal(reader, &messages); err != nil {
		return err
	}

	for _, message := range messages {
		if _, ok := GlobalErrors[message.Method]; !ok {
			GlobalErrors[message.Method] = map[string]Message{}
		}

		GlobalErrors[message.Method][message.ErrorKey] = Message{
			Message:   message.ErrorKey,
			Code:      codes.Code(message.GrpcCode),
			ErrorCode: message.ErrorCode,
			Labels: Labels{
				Uz: message.Uz,
				Ru: message.Ru,
				En: message.En,
			},
		}
	}

	return nil
}
