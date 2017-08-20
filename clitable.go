package clitable

import (
	"io"
)

type Row interface {
	GetValue(int) string
}

func writeChar(out io.Writer, c byte, times int) {
	for i := 0; i < times; i++ {
		out.Write([]byte{c})
	}
}

func Write(out io.Writer, head []string, rows []Row) {
	cols := len(head)
	cw := make([]int, cols)
	{ // calculation of column widths, maybe extra function?
		for col, h := range head {
			n := len([]rune(h))
			if n > cw[col] {
				cw[col] = n
			}
		}
		for _, row := range rows {
			for col := 0; col < cols; col++ {
				v := row.GetValue(col)
				n := len([]rune(v))
				if n > cw[col] {
					cw[col] = n
				}
			}
		}
	}
	for col, h := range head {
		d := cw[col] - len(h) + 1
		writeChar(out, ' ', d)
		io.WriteString(out, h)
		if col < cols-1 {
			out.Write([]byte{'|'})
		}
	}
	out.Write([]byte{'\n'})
	for i, w := range cw {
		writeChar(out, '-', w+1)
		if i < cols-1 {
			out.Write([]byte{'+'})
		}
	}
	out.Write([]byte{'\n'})
	for _, row := range rows {
		for col := 0; col < cols; col++ {
			v := row.GetValue(col)
			d := cw[col] - len([]rune(v)) + 1
			writeChar(out, ' ', d)
			io.WriteString(out, v)
			if col < cols-1 {
				out.Write([]byte{'|'})
			}
		}
		out.Write([]byte{'\n'})
	}
}
