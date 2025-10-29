package str

import "testing"

func TestBufferAppend(t *testing.T) {
	buf := NewBuffer()

	buf.Append("start")
	buf.Append([]byte("-bytes"))
	buf.Append('-')
	buf.Append(42)
	buf.Append(int64(7))
	buf.Append(uint(9))
	buf.Append(uint64(11))
	buf.Append(true)

	got := buf.String()
	want := "start-bytes-427911true"

	if got != want {
		t.Fatalf("Buffer.String() = %q, want %q", got, want)
	}
}
