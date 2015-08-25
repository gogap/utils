package range_parser

import (
	"errors"
	"strconv"
	"strings"
)

type OrderType int32

const (
	ORDER_TYPE_ASC  OrderType = 0
	ORDER_TYPE_DESC           = 1
)

var (
	DefaultMaxCount int64 = 20
	DefaultOffset   int64 = 1
)

type RangeOptions struct {
	OrderBy   []string
	OrderType OrderType
	Offset    int64
	Max       int64
	Want      []string
}

func Parse(opt string) (rangeOptions *RangeOptions, err error) {
	opt = strings.Trim(opt, " ")

	if opt == "" {
		return
	} else if opt = strings.TrimRight(opt, ";"); opt == "" {
		return
	}

	rangeOpts := new(RangeOptions)
	rangeOpts.Max = DefaultMaxCount
	rangeOpts.Offset = DefaultOffset
	rangeOpts.OrderType = ORDER_TYPE_DESC
	rangeOpts.Want = nil

	optArray := strings.Split(opt, ";")
	for _, option := range optArray {
		optKV := strings.Split(option, "=")
		if len(optKV) != 2 {
			return nil, errors.New("range options params should be key-value format")
		}

		k := strings.Trim(optKV[0], " ")
		v := strings.Trim(optKV[1], " ")

		switch k {
		case "order":
			{
				if v == "desc" {
					rangeOpts.OrderType = ORDER_TYPE_DESC
				} else if v == "asc" {
					rangeOpts.OrderType = ORDER_TYPE_ASC
				} else {
					return nil, errors.New("param order only could be desc or asc")
				}
			}
		case "max":
			{
				iv, e := strconv.ParseInt(v, 10, 64)
				if e != nil {
					return nil, errors.New("param max only could be int")
				}

				if iv <= 0 {
					return nil, errors.New("param max should greater than zero")
				}

				rangeOpts.Max = iv
			}
		case "offset":
			{
				iv, e := strconv.ParseInt(v, 10, 64)
				if e != nil {
					return nil, errors.New("param offset only could be int")
				}

				if iv < 0 {
					return nil, errors.New("param offset should greater or equal than zero")
				}

				rangeOpts.Offset = iv
			}
		case "order_by":
			{
				orderBy := []string{}
				keys := strings.Split(v, " ")
				for _, key := range keys {
					key = strings.Trim(key, " ")
					if key != "" {
						orderBy = append(orderBy, key)
					}
				}

				if len(orderBy) == 0 {
					return nil, errors.New("none order keys exist")
				}

				rangeOpts.OrderBy = orderBy
			}
		case "want":
			{
				want := []string{}
				keys := strings.Split(v, " ")
				for _, key := range keys {
					key = strings.Trim(key, " ")
					if key != "" && key != "..." {
						want = append(want, key)
					}
				}

				if len(want) == 0 {
					want = nil
				}

				rangeOpts.Want = want
			}
		}
	}

	rangeOptions = rangeOpts

	return
}
