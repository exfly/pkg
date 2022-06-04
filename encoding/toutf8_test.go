package encoding

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewToUTF8Reader(t *testing.T) {
	tests := []struct {
		charset  string
		filename string
	}{
		{"UTF-8", "testdata/utf8.html"},
		{"UTF-8", "testdata/utf8_bom.html"},
		{"GB-18030", "testdata/gb18030.html"},
		{"Big5", "testdata/big5.html"},
		{"EUC-JP", "testdata/euc_jp.html"},
		{"EUC-KR", "testdata/euc_kr.html"},
		// { "", "testdata/8859_1_pt.html"},
		{"Shift_JIS", "testdata/shift_jis.html"},
	}

	for _, args := range tests {
		f, err := os.Open(args.filename)
		defer f.Close()
		require.NoError(t, err)

		reader, err := NewToUTF8Reader(f, DefaultDetectSize)
		require.NoError(t, err)
		require.Equal(t, args.charset, reader.CharsetName())

		buf := make([]byte, 1000)
		n, err := reader.Read(buf)
		require.NoError(t, err)

		got := string(buf[:n])
		t.Log(got)
		require.NotEqual(t, "", got)
	}
}
