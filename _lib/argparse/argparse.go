// A SIMPLE ARG PARSER

package argparse

import (
	"reflect"
	"strconv"
	"strings"
)

func ParseArgs(strct any, args []string) {
	argMap := make(map[string]string)

	for i := 0; i < len(args); i++ {
		if strings.HasPrefix(args[i], "-") {
			b, a, r := strings.Cut(args[i], "=")
			if r {
				argMap[b] = a
			} else {
				argMap[b] = "True"
			}
		}
	}

	v := reflect.ValueOf(strct).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("arg")
		if tag == "" {
			continue
		}

		b, a, r := strings.Cut(tag, ",")
		var (
			raw string
			ok  bool
		)
		if r {
			raw, ok = argMap[a]
			if !ok {
				raw, _ = argMap[b]
			}
		} else {
			raw, _ = argMap[a]
		}

		switch field.Type.Kind() {
		case reflect.String:
			v.Field(i).SetString(raw)
		case reflect.Int:
			n, err := strconv.Atoi(raw)
			if err == nil {
				v.Field(i).SetInt(int64(n))
			}
		case reflect.Bool:
			boo, err := strconv.ParseBool(raw)
			if err == nil {
				v.Field(i).SetBool(boo)
			}
		default:
		}
	}
}
