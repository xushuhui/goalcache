package geek

import "strconv"

func String2Int(s string, defaultval int) int {
	a, err := strconv.Atoi(s)
	if err != nil {
		return defaultval
	}

	return a
}

// 转换默认值
func String2Int32(s string, defaultval int32) int32 {
	a, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return defaultval
	}

	return int32(a)
}

func String2Int64(s string, defaultval int64) int64 {
	a, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return defaultval
	}

	return a
}

func String2Uint64(s string, defaultval uint64) uint64 {
	a, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return defaultval
	}

	return a
}

func String2Int8(s string, defaultval int8) int8 {
	a, err := strconv.Atoi(s)
	if err != nil {
		return defaultval
	}

	return int8(a)
}

func String2Uint8(s string, defaultval uint8) uint8 {
	a, err := strconv.Atoi(s)
	if err != nil {
		return defaultval
	}

	return uint8(a)
}

func String2Float(s string, defaultval float64) float64 {
	a, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return defaultval
	}

	return float64(a)
}

func Int2String(a int) string {
	s := strconv.Itoa(a)
	return s
}

func Int642String(a int64) string {
	s := strconv.FormatInt(a, 10)
	return s
}

func Uint642String(a uint64) string {
	s := strconv.FormatUint(a, 10)
	return s
}

func Int82String(a int8) string {
	s := strconv.Itoa(int(a))
	return s
}

func Float2String(a float64) string {
	s := strconv.FormatFloat(a, 'f', 6, 64)
	return s
}

func Float2StringmINLenth(a float64) string {
	s := strconv.FormatFloat(a, 'f', -1, 64)
	return s
}
