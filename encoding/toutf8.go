package encoding

import (
	"bufio"
	"io"
	"strings"

	"github.com/pkg/errors"
	"github.com/saintfish/chardet"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	DefaultDetectSize = 4096 * 10
)

type charsetReader interface {
	io.Reader
	CharsetName() string
}

func newCharsetReader(reader io.Reader, charset string) implCharsetReader {
	return implCharsetReader{Reader: reader, charset: charset}
}

type implCharsetReader struct {
	io.Reader
	charset string
}

func (c implCharsetReader) CharsetName() string {
	return c.charset
}

// NewToUTF8Reader any charset to utf8 reader
func NewToUTF8Reader(r io.Reader, detectSize int) (charsetReader, error) {
	reader := bufio.NewReaderSize(r, detectSize)
	content, _ := reader.Peek(detectSize)

	dt := chardet.NewTextDetector()
	bestCharset, err := dt.DetectBest(content)
	if err != nil {
		return nil, errors.Wrap(err, "detect charset failed")
	}

	var unicodeReader charsetReader = newCharsetReader(reader, "origin")

	switch vv := bestCharset.Charset; strings.ToLower(vv) {
	case "gb-18030":
		unicodeReader = newCharsetReader(
			transform.NewReader(reader, simplifiedchinese.GB18030.NewDecoder()),
			vv,
		)
	case "utf-16le":
		win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
		tf := unicode.BOMOverride(win16be.NewDecoder())
		unicodeReader = newCharsetReader(
			transform.NewReader(reader, tf),
			vv,
		)
	case "big5":
		unicodeReader = newCharsetReader(
			transform.NewReader(
				reader,
				traditionalchinese.Big5.NewDecoder(),
			),
			vv,
		)
	case "euc-jp":
		unicodeReader = newCharsetReader(
			transform.NewReader(
				reader,
				japanese.EUCJP.NewDecoder(),
			),
			vv,
		)
	case "euc-kr":
		unicodeReader = newCharsetReader(
			transform.NewReader(
				reader,
				korean.EUCKR.NewDecoder(),
			),
			vv,
		)
	case "shift_jis":
		unicodeReader = newCharsetReader(
			transform.NewReader(
				reader,
				japanese.ShiftJIS.NewDecoder(),
			),
			vv,
		)
	case "iso-8859-1":
		// skip
		unicodeReader = newCharsetReader(
			unicodeReader,
			vv,
		)
	case "windows-1252":
		unicodeReader = newCharsetReader(
			transform.NewReader(
				reader,
				charmap.Windows1252.NewDecoder(),
			),
			vv,
		)
	case "utf-8":
		// skip
		unicodeReader = newCharsetReader(
			unicodeReader,
			vv,
		)
	default:
		return nil, errors.Errorf("unknown charset %v", vv)
	}

	return unicodeReader, nil
}
