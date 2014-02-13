package mysql

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
)

type Field struct {
	Schema       []byte
	Table        []byte
	OrgTable     []byte
	Name         []byte
	OrgName      []byte
	Charset      uint16
	ColumnLength uint32
	Type         uint8
	Flag         uint16
	Decimal      uint8

	//below if command was fieldlist
	DefaultValueLength uint64
	DefaultValue       []byte

	isFieldList bool
}

func (f Field) Dump() []byte {
	l := len(f.Schema) + len(f.Table) + len(f.OrgTable) + len(f.Name) + len(f.OrgName) + len(f.DefaultValue) + 48

	data := make([]byte, 0, l)

	data = append(data, PutLengthEncodedString([]byte("def"))...)

	data = append(data, PutLengthEncodedString(f.Schema)...)

	data = append(data, PutLengthEncodedString(f.Table)...)
	data = append(data, PutLengthEncodedString(f.OrgTable)...)

	data = append(data, PutLengthEncodedString(f.Name)...)
	data = append(data, PutLengthEncodedString(f.OrgName)...)

	data = append(data, 0x0c)

	data = append(data, Uint16ToBytes(f.Charset)...)
	data = append(data, Uint32ToBytes(f.ColumnLength)...)
	data = append(data, f.Type)
	data = append(data, Uint16ToBytes(f.Flag)...)
	data = append(data, f.Decimal)
	data = append(data, 0, 0)

	if f.isFieldList {
		data = append(data, Uint64ToBytes(f.DefaultValueLength)...)
		data = append(data, f.DefaultValue...)
	}

	return data
}

type fieldPacket []byte

func (p fieldPacket) Parse() (f Field, err error) {
	var n int
	pos := 0
	//skip catelog, always def
	n, err = SkipLengthEnodedString(p)
	if err != nil {
		return
	}
	pos += n

	//schema
	f.Schema, _, n, err = LengthEnodedString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//table
	f.Table, _, n, err = LengthEnodedString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//org_table
	f.OrgTable, _, n, err = LengthEnodedString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//name
	f.Name, _, n, err = LengthEnodedString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//org_name
	f.OrgName, _, n, err = LengthEnodedString(p[pos:])
	if err != nil {
		return
	}
	pos += n

	//skip oc
	pos += 1

	//charset
	f.Charset = binary.LittleEndian.Uint16(p[pos:])
	pos += 2

	//column length
	f.ColumnLength = binary.LittleEndian.Uint32(p[pos:])
	pos += 4

	//type
	f.Type = p[pos]
	pos++

	//flag
	f.Flag = binary.LittleEndian.Uint16(p[pos:])
	pos += 2

	//decimals 1
	f.Decimal = p[pos]
	pos++

	//filter [0x00][0x00]
	pos += 2

	//if more data, command was field list
	if pos < len(p) {
		f.isFieldList = true

		//length of default value lenenc-int
		f.DefaultValueLength, _, n = LengthEncodedInt(p[pos:])
		pos += n

		if pos+int(f.DefaultValueLength) > len(p) {
			err = ErrMalformPacket
			return
		}

		//default value string[$len]
		f.DefaultValue = p[pos:(pos + int(f.DefaultValueLength))]
	}

	return
}

type rowPacket []byte

func (p rowPacket) Parse(f []Field, binary bool) ([]interface{}, error) {
	if binary {
		return p.ParseBinary(f)
	} else {
		return p.ParseText(f)
	}
}

func (p rowPacket) ParseText(f []Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	var err error
	var v []byte
	var isNull, isUnsigned bool
	var pos int = 0
	var n int = 0

	for i := range f {
		v, isNull, n, err = LengthEnodedString(p[pos:])
		if err != nil {
			return nil, err
		}

		pos += n

		if isNull {
			data[i] = nil
		} else {
			isUnsigned = (f[i].Flag&UNSIGNED_FLAG > 0)

			switch f[i].Type {
			case MYSQL_TYPE_TINY, MYSQL_TYPE_SHORT, MYSQL_TYPE_INT24,
				MYSQL_TYPE_LONGLONG, MYSQL_TYPE_YEAR:
				if isUnsigned {
					data[i], err = strconv.ParseUint(string(v), 10, 64)
				} else {
					data[i], err = strconv.ParseInt(string(v), 10, 64)
				}
			case MYSQL_TYPE_FLOAT, MYSQL_TYPE_DOUBLE:
				data[i], err = strconv.ParseFloat(string(v), 64)
			default:
				data[i] = v
			}

			if err != nil {
				return nil, err
			}
		}
	}

	return data, nil
}

func (p rowPacket) ParseBinary(f []Field) ([]interface{}, error) {
	data := make([]interface{}, len(f))

	if p[0] != OK_HEADER {
		return nil, ErrMalformPacket
	}

	pos := 1 + ((len(f) + 7 + 2) >> 3)

	nullBitmap := p[1:pos]

	var isUnsigned bool
	var isNull bool
	var n int
	var err error
	var v []byte
	for i := range data {
		if nullBitmap[(i+2)/8]&(1<<(uint(i+2)%8)) > 0 {
			data[i] = nil
			continue
		}

		isUnsigned = f[i].Flag&UNSIGNED_FLAG > 0

		switch f[i].Type {
		case MYSQL_TYPE_NULL:
			data[i] = nil
			continue

		case MYSQL_TYPE_TINY:
			if isUnsigned {
				data[i] = uint64(p[pos])
			} else {
				data[i] = int64(p[pos])
			}
			pos++
			continue

		case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint16(p[pos : pos+2]))
			} else {
				data[i] = int64((binary.LittleEndian.Uint16(p[pos : pos+2])))
			}
			pos += 2
			continue

		case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
			if isUnsigned {
				data[i] = uint64(binary.LittleEndian.Uint32(p[pos : pos+4]))
			} else {
				data[i] = int64(binary.LittleEndian.Uint32(p[pos : pos+4]))
			}
			pos += 4
			continue

		case MYSQL_TYPE_LONGLONG:
			if isUnsigned {
				data[i] = binary.LittleEndian.Uint64(p[pos : pos+8])
			} else {
				data[i] = int64(binary.LittleEndian.Uint64(p[pos : pos+8]))
			}
			pos += 8
			continue

		case MYSQL_TYPE_FLOAT:
			data[i] = float64(math.Float32frombits(binary.LittleEndian.Uint32(p[pos : pos+4])))
			pos += 4
			continue

		case MYSQL_TYPE_DOUBLE:
			data[i] = math.Float64frombits(binary.LittleEndian.Uint64(p[pos : pos+8]))
			pos += 8
			continue

		case MYSQL_TYPE_DECIMAL, MYSQL_TYPE_NEWDECIMAL, MYSQL_TYPE_VARCHAR,
			MYSQL_TYPE_BIT, MYSQL_TYPE_ENUM, MYSQL_TYPE_SET, MYSQL_TYPE_TINY_BLOB,
			MYSQL_TYPE_MEDIUM_BLOB, MYSQL_TYPE_LONG_BLOB, MYSQL_TYPE_BLOB,
			MYSQL_TYPE_VAR_STRING, MYSQL_TYPE_STRING, MYSQL_TYPE_GEOMETRY:
			v, isNull, n, err = LengthEnodedString(p[pos:])
			pos += n
			if err != nil {
				return nil, err
			}

			if !isNull {
				data[i] = v
				continue
			} else {
				data[i] = nil
				continue
			}
		case MYSQL_TYPE_DATE, MYSQL_TYPE_NEWDATE:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDate(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIMESTAMP, MYSQL_TYPE_DATETIME:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryDateTime(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		case MYSQL_TYPE_TIME:
			var num uint64
			num, isNull, n = LengthEncodedInt(p[pos:])

			pos += n

			if isNull {
				data[i] = nil
				continue
			}

			data[i], err = FormatBinaryTime(int(num), p[pos:])
			pos += int(num)

			if err != nil {
				return nil, err
			}

		default:
			return nil, fmt.Errorf("Stmt Unknown FieldType %d %s", f[i].Type, f[i].Name)
		}
	}

	return data, nil
}

type resultsetPacket struct {
	Status uint16
	Fields []fieldPacket
	Rows   []rowPacket
}

func (p *resultsetPacket) Parse(binary bool) (*Resultset, error) {
	r := new(Resultset)

	r.binary = binary
	r.Status = p.Status

	r.Fields = make([]Field, len(p.Fields))
	r.FieldNames = make(map[string]int, len(p.Fields))

	var err error
	for i := range r.Fields {
		r.Fields[i], err = p.Fields[i].Parse()
		if err != nil {
			return nil, err
		}

		r.FieldNames[string(r.Fields[i].Name)] = i
	}

	r.Data = make([][]interface{}, len(p.Rows))
	for i := range r.Data {
		r.Data[i], err = p.Rows[i].Parse(r.Fields, binary)
		if err != nil {
			return nil, err
		}
	}

	return r, nil
}

type Resultset struct {
	binary bool //test resultset is text or binary

	Status uint16 //server status for this query resultset

	Fields     []Field
	FieldNames map[string]int

	Data [][]interface{}
}

func (r *Resultset) GetStatus() uint16 {
	return r.Status
}

func (r *Resultset) RowNumber() int {
	return len(r.Data)
}

func (r *Resultset) ColumnNumber() int {
	return len(r.Fields)
}

func (r *Resultset) GetData(row, column int) (interface{}, error) {
	if row >= len(r.Data) || row < 0 {
		return nil, fmt.Errorf("invalid row index %d", row)
	}

	if column >= len(r.Fields) || column < 0 {
		return nil, fmt.Errorf("invalid column index %d", column)
	}

	return r.Data[row][column], nil
}

func (r *Resultset) GetDataByName(row int, name string) (interface{}, error) {
	if column, ok := r.FieldNames[name]; ok {
		return r.GetData(row, column)
	} else {
		return nil, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) IsNull(row, column int) (bool, error) {
	d, err := r.GetData(row, column)
	if err != nil {
		return false, err
	}

	return d == nil, nil
}

func (r *Resultset) IsNullByName(row int, name string) (bool, error) {
	if column, ok := r.FieldNames[name]; ok {
		return r.IsNull(row, column)
	} else {
		return false, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) GetUint(row, column int) (uint64, error) {
	d, err := r.GetData(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case uint64:
		return v, nil
	case int64:
		return uint64(v), nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetUintByName(row int, name string) (uint64, error) {
	if column, ok := r.FieldNames[name]; ok {
		return r.GetUint(row, column)
	} else {
		return 0, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) GetInt(row, column int) (int64, error) {
	v, err := r.GetUint(row, column)
	if err != nil {
		return 0, err
	}

	return int64(v), nil
}

func (r *Resultset) GetIntByName(row int, name string) (int64, error) {
	v, err := r.GetUintByName(row, name)
	if err != nil {
		return 0, err
	}

	return int64(v), nil
}

func (r *Resultset) GetFloat(row, column int) (float64, error) {
	d, err := r.GetData(row, column)
	if err != nil {
		return 0, err
	}

	switch v := d.(type) {
	case float64:
		return v, nil
	case nil:
		return 0, nil
	default:
		return 0, fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetFloatByName(row int, name string) (float64, error) {
	if column, ok := r.FieldNames[name]; ok {
		return r.GetFloat(row, column)
	} else {
		return 0, fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) GetString(row, column int) (string, error) {
	d, err := r.GetData(row, column)
	if err != nil {
		return "", err
	}

	switch v := d.(type) {
	case string:
		return v, nil
	case []byte:
		return string(v), nil
	case nil:
		return "", nil
	default:
		return "", fmt.Errorf("data type is %T", v)
	}
}

func (r *Resultset) GetStringByName(row int, name string) (string, error) {
	if column, ok := r.FieldNames[name]; ok {
		return r.GetString(row, column)
	} else {
		return "", fmt.Errorf("invalid field name %s", name)
	}
}

func (r *Resultset) DumpRows() ([][]byte, error) {
	if r.binary {
		return r.dumpBinary()
	} else {
		return r.dumpText()
	}
}

func (r *Resultset) dumpBinary() ([][]byte, error) {
	rows := make([][]byte, len(r.Data))

	var err error
	for i := range rows {
		rows[i], err = r.dumpBinaryRow(r.Data[i])
		if err != nil {
			return nil, err
		}
	}

	return rows, nil
}

func (r *Resultset) dumpUint(v uint64, tp byte) ([]byte, error) {
	switch tp {
	case MYSQL_TYPE_TINY:
		return []byte{byte(v)}, nil
	case MYSQL_TYPE_SHORT, MYSQL_TYPE_YEAR:
		return Uint16ToBytes(uint16(v)), nil
	case MYSQL_TYPE_INT24, MYSQL_TYPE_LONG:
		return Uint32ToBytes(uint32(v)), nil
	case MYSQL_TYPE_LONGLONG:
		return Uint64ToBytes(v), nil
	default:
		return nil, fmt.Errorf("invalid field type %d", tp)
	}
}

func (r *Resultset) dumpBinaryRow(data []interface{}) ([]byte, error) {
	//NULL-bitmap, length: (column-count + 7 + 2) / 8
	nullBitmapLenght := ((len(data) + 7 + 2) >> 3)

	row := make([]byte, 1+nullBitmapLenght, 1024)
	row[0] = OK_HEADER

	for p, i := range data {
		tp := r.Fields[p].Type
		switch v := i.(type) {
		case nil:
			row[1+(p+2)/8] |= (1 << (uint(p+2) % 8))
		case int64:
			if d, err := r.dumpUint(uint64(v), tp); err != nil {
				return nil, err
			} else {
				row = append(row, d...)
			}
		case uint64:
			if d, err := r.dumpUint(uint64(v), tp); err != nil {
				return nil, err
			} else {
				row = append(row, d...)
			}
		case float64:
			if tp == MYSQL_TYPE_FLOAT {
				row = append(row, Uint32ToBytes(math.Float32bits(float32(v)))...)
			} else if tp == MYSQL_TYPE_DOUBLE {
				row = append(row, Uint64ToBytes(math.Float64bits(v))...)
			} else {
				return nil, fmt.Errorf("invalid field type %d", tp)
			}
		case []byte:
			row = append(row, PutLengthEncodedString(v)...)
		case string:
			row = append(row, PutLengthEncodedString([]byte(v))...)
		default:
			return nil, fmt.Errorf("invalid type %T", v)
		}
	}
	return row, nil
}

func (r *Resultset) dumpText() ([][]byte, error) {
	rows := make([][]byte, len(r.Data))

	var err error
	for i := range rows {
		rows[i], err = r.dumpTextRow(r.Data[i])
		if err != nil {
			return nil, err
		}
	}

	return rows, nil
}

func (r *Resultset) dumpTextRow(data []interface{}) ([]byte, error) {
	row := make([]byte, 0, 1024)

	for _, i := range data {
		switch v := i.(type) {
		case nil:
			row = append(row, 0xfb)
		case int64:
			s := strconv.FormatInt(v, 10)
			row = append(row, PutLengthEncodedString([]byte(s))...)
		case uint64:
			s := strconv.FormatUint(v, 10)
			row = append(row, PutLengthEncodedString([]byte(s))...)
		case float64:
			s := strconv.FormatFloat(v, 'f', -1, 64)
			row = append(row, PutLengthEncodedString([]byte(s))...)
		case []byte:
			row = append(row, PutLengthEncodedString(v)...)
		case string:
			row = append(row, PutLengthEncodedString([]byte(v))...)
		default:
			return nil, fmt.Errorf("invalid type %T", i)
		}
	}

	return row, nil
}

type Result struct {
	Status       uint16
	InsertId     uint64
	AffectedRows uint64
}

func (r *Result) GetStatus() uint16 {
	return r.Status
}

func (r *Result) LastInsertId() (int64, error) {
	return int64(r.InsertId), nil
}

func (r *Result) RowsAffected() (int64, error) {
	return int64(r.AffectedRows), nil
}
