/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// ************************************************************
// DO NOT EDIT.
// THIS FILE IS AUTO-GENERATED BY codecgen.
// ************************************************************

package v1alpha1

import (
	"errors"
	"fmt"
	codec1978 "github.com/ugorji/go/codec"
	pkg4_resource "k8s.io/apimachinery/pkg/api/resource"
	pkg1_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	pkg3_types "k8s.io/apimachinery/pkg/types"
	pkg2_v1 "k8s.io/client-go/pkg/api/v1"
	"reflect"
	"runtime"
	time "time"
)

const (
	// ----- content types ----
	codecSelferC_UTF81234 = 1
	codecSelferC_RAW1234  = 0
	// ----- value types used ----
	codecSelferValueTypeArray1234 = 10
	codecSelferValueTypeMap1234   = 9
	// ----- containerStateValues ----
	codecSelfer_containerMapKey1234    = 2
	codecSelfer_containerMapValue1234  = 3
	codecSelfer_containerMapEnd1234    = 4
	codecSelfer_containerArrayElem1234 = 6
	codecSelfer_containerArrayEnd1234  = 7
)

var (
	codecSelferBitsize1234                         = uint8(reflect.TypeOf(uint(0)).Bits())
	codecSelferOnlyMapOrArrayEncodeToStructErr1234 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1234 struct{}

func init() {
	if codec1978.GenVersion != 5 {
		_, file, _, _ := runtime.Caller(0)
		err := fmt.Errorf("codecgen version mismatch: current: %v, need %v. Re-generate file: %v",
			5, codec1978.GenVersion, file)
		panic(err)
	}
	if false { // reference the types, but skip this branch at build/run time
		var v0 pkg4_resource.Quantity
		var v1 pkg1_v1.TypeMeta
		var v2 pkg3_types.UID
		var v3 pkg2_v1.ObjectReference
		var v4 time.Time
		_, _, _, _, _ = v0, v1, v2, v3, v4
	}
}

func (x *MetricValueList) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [4]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[2] = true
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(4)
			} else {
				yynn2 = 1
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[2] {
					yy10 := &x.ListMeta
					yym11 := z.EncBinary()
					_ = yym11
					if false {
					} else if z.HasExtensions() && z.EncExt(yy10) {
					} else {
						z.EncFallback(yy10)
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[2] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("metadata"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yy12 := &x.ListMeta
					yym13 := z.EncBinary()
					_ = yym13
					if false {
					} else if z.HasExtensions() && z.EncExt(yy12) {
					} else {
						z.EncFallback(yy12)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym15 := z.EncBinary()
					_ = yym15
					if false {
					} else {
						h.encSliceMetricValue(([]MetricValue)(x.Items), e)
					}
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("items"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				if x.Items == nil {
					r.EncodeNil()
				} else {
					yym16 := z.EncBinary()
					_ = yym16
					if false {
					} else {
						h.encSliceMetricValue(([]MetricValue)(x.Items), e)
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *MetricValueList) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *MetricValueList) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "metadata":
			if r.TryDecodeAsNil() {
				x.ListMeta = pkg1_v1.ListMeta{}
			} else {
				yyv8 := &x.ListMeta
				yym9 := z.DecBinary()
				_ = yym9
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv8) {
				} else {
					z.DecFallback(yyv8, false)
				}
			}
		case "items":
			if r.TryDecodeAsNil() {
				x.Items = nil
			} else {
				yyv10 := &x.Items
				yym11 := z.DecBinary()
				_ = yym11
				if false {
				} else {
					h.decSliceMetricValue((*[]MetricValue)(yyv10), d)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *MetricValueList) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj12 int
	var yyb12 bool
	var yyhl12 bool = l >= 0
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv13 := &x.Kind
		yym14 := z.DecBinary()
		_ = yym14
		if false {
		} else {
			*((*string)(yyv13)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv15 := &x.APIVersion
		yym16 := z.DecBinary()
		_ = yym16
		if false {
		} else {
			*((*string)(yyv15)) = r.DecodeString()
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.ListMeta = pkg1_v1.ListMeta{}
	} else {
		yyv17 := &x.ListMeta
		yym18 := z.DecBinary()
		_ = yym18
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv17) {
		} else {
			z.DecFallback(yyv17, false)
		}
	}
	yyj12++
	if yyhl12 {
		yyb12 = yyj12 > l
	} else {
		yyb12 = r.CheckBreak()
	}
	if yyb12 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Items = nil
	} else {
		yyv19 := &x.Items
		yym20 := z.DecBinary()
		_ = yym20
		if false {
		} else {
			h.decSliceMetricValue((*[]MetricValue)(yyv19), d)
		}
	}
	for {
		yyj12++
		if yyhl12 {
			yyb12 = yyj12 > l
		} else {
			yyb12 = r.CheckBreak()
		}
		if yyb12 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj12-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x *MetricValue) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		yym1 := z.EncBinary()
		_ = yym1
		if false {
		} else if z.HasExtensions() && z.EncExt(x) {
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			var yyq2 [7]bool
			_, _, _ = yysep2, yyq2, yy2arr2
			const yyr2 bool = false
			yyq2[0] = x.Kind != ""
			yyq2[1] = x.APIVersion != ""
			yyq2[5] = x.WindowSeconds != nil
			var yynn2 int
			if yyr2 || yy2arr2 {
				r.EncodeArrayStart(7)
			} else {
				yynn2 = 4
				for _, b := range yyq2 {
					if b {
						yynn2++
					}
				}
				r.EncodeMapStart(yynn2)
				yynn2 = 0
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[0] {
					yym4 := z.EncBinary()
					_ = yym4
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[0] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("kind"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym5 := z.EncBinary()
					_ = yym5
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.Kind))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[1] {
					yym7 := z.EncBinary()
					_ = yym7
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				} else {
					r.EncodeString(codecSelferC_UTF81234, "")
				}
			} else {
				if yyq2[1] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("apiVersion"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					yym8 := z.EncBinary()
					_ = yym8
					if false {
					} else {
						r.EncodeString(codecSelferC_UTF81234, string(x.APIVersion))
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy10 := &x.DescribedObject
				yy10.CodecEncodeSelf(e)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("describedObject"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy12 := &x.DescribedObject
				yy12.CodecEncodeSelf(e)
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yym15 := z.EncBinary()
				_ = yym15
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.MetricName))
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("metricName"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yym16 := z.EncBinary()
				_ = yym16
				if false {
				} else {
					r.EncodeString(codecSelferC_UTF81234, string(x.MetricName))
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy18 := &x.Timestamp
				yym19 := z.EncBinary()
				_ = yym19
				if false {
				} else if z.HasExtensions() && z.EncExt(yy18) {
				} else if yym19 {
					z.EncBinaryMarshal(yy18)
				} else if !yym19 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy18)
				} else {
					z.EncFallback(yy18)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("timestamp"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy20 := &x.Timestamp
				yym21 := z.EncBinary()
				_ = yym21
				if false {
				} else if z.HasExtensions() && z.EncExt(yy20) {
				} else if yym21 {
					z.EncBinaryMarshal(yy20)
				} else if !yym21 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy20)
				} else {
					z.EncFallback(yy20)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				if yyq2[5] {
					if x.WindowSeconds == nil {
						r.EncodeNil()
					} else {
						yy23 := *x.WindowSeconds
						yym24 := z.EncBinary()
						_ = yym24
						if false {
						} else {
							r.EncodeInt(int64(yy23))
						}
					}
				} else {
					r.EncodeNil()
				}
			} else {
				if yyq2[5] {
					z.EncSendContainerState(codecSelfer_containerMapKey1234)
					r.EncodeString(codecSelferC_UTF81234, string("window"))
					z.EncSendContainerState(codecSelfer_containerMapValue1234)
					if x.WindowSeconds == nil {
						r.EncodeNil()
					} else {
						yy25 := *x.WindowSeconds
						yym26 := z.EncBinary()
						_ = yym26
						if false {
						} else {
							r.EncodeInt(int64(yy25))
						}
					}
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayElem1234)
				yy28 := &x.Value
				yym29 := z.EncBinary()
				_ = yym29
				if false {
				} else if z.HasExtensions() && z.EncExt(yy28) {
				} else if !yym29 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy28)
				} else {
					z.EncFallback(yy28)
				}
			} else {
				z.EncSendContainerState(codecSelfer_containerMapKey1234)
				r.EncodeString(codecSelferC_UTF81234, string("value"))
				z.EncSendContainerState(codecSelfer_containerMapValue1234)
				yy30 := &x.Value
				yym31 := z.EncBinary()
				_ = yym31
				if false {
				} else if z.HasExtensions() && z.EncExt(yy30) {
				} else if !yym31 && z.IsJSONHandle() {
					z.EncJSONMarshal(yy30)
				} else {
					z.EncFallback(yy30)
				}
			}
			if yyr2 || yy2arr2 {
				z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				z.EncSendContainerState(codecSelfer_containerMapEnd1234)
			}
		}
	}
}

func (x *MetricValue) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	yym1 := z.DecBinary()
	_ = yym1
	if false {
	} else if z.HasExtensions() && z.DecExt(x) {
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeMap1234 {
			yyl2 := r.ReadMapStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerMapEnd1234)
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
		} else if yyct2 == codecSelferValueTypeArray1234 {
			yyl2 := r.ReadArrayStart()
			if yyl2 == 0 {
				z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
			} else {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
		} else {
			panic(codecSelferOnlyMapOrArrayEncodeToStructErr1234)
		}
	}
}

func (x *MetricValue) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yys3Slc = z.DecScratchBuffer() // default slice to decode into
	_ = yys3Slc
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if r.CheckBreak() {
				break
			}
		}
		z.DecSendContainerState(codecSelfer_containerMapKey1234)
		yys3Slc = r.DecodeBytes(yys3Slc, true, true)
		yys3 := string(yys3Slc)
		z.DecSendContainerState(codecSelfer_containerMapValue1234)
		switch yys3 {
		case "kind":
			if r.TryDecodeAsNil() {
				x.Kind = ""
			} else {
				yyv4 := &x.Kind
				yym5 := z.DecBinary()
				_ = yym5
				if false {
				} else {
					*((*string)(yyv4)) = r.DecodeString()
				}
			}
		case "apiVersion":
			if r.TryDecodeAsNil() {
				x.APIVersion = ""
			} else {
				yyv6 := &x.APIVersion
				yym7 := z.DecBinary()
				_ = yym7
				if false {
				} else {
					*((*string)(yyv6)) = r.DecodeString()
				}
			}
		case "describedObject":
			if r.TryDecodeAsNil() {
				x.DescribedObject = pkg2_v1.ObjectReference{}
			} else {
				yyv8 := &x.DescribedObject
				yyv8.CodecDecodeSelf(d)
			}
		case "metricName":
			if r.TryDecodeAsNil() {
				x.MetricName = ""
			} else {
				yyv9 := &x.MetricName
				yym10 := z.DecBinary()
				_ = yym10
				if false {
				} else {
					*((*string)(yyv9)) = r.DecodeString()
				}
			}
		case "timestamp":
			if r.TryDecodeAsNil() {
				x.Timestamp = pkg1_v1.Time{}
			} else {
				yyv11 := &x.Timestamp
				yym12 := z.DecBinary()
				_ = yym12
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv11) {
				} else if yym12 {
					z.DecBinaryUnmarshal(yyv11)
				} else if !yym12 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv11)
				} else {
					z.DecFallback(yyv11, false)
				}
			}
		case "window":
			if r.TryDecodeAsNil() {
				if x.WindowSeconds != nil {
					x.WindowSeconds = nil
				}
			} else {
				if x.WindowSeconds == nil {
					x.WindowSeconds = new(int64)
				}
				yym14 := z.DecBinary()
				_ = yym14
				if false {
				} else {
					*((*int64)(x.WindowSeconds)) = int64(r.DecodeInt(64))
				}
			}
		case "value":
			if r.TryDecodeAsNil() {
				x.Value = pkg4_resource.Quantity{}
			} else {
				yyv15 := &x.Value
				yym16 := z.DecBinary()
				_ = yym16
				if false {
				} else if z.HasExtensions() && z.DecExt(yyv15) {
				} else if !yym16 && z.IsJSONHandle() {
					z.DecJSONUnmarshal(yyv15)
				} else {
					z.DecFallback(yyv15, false)
				}
			}
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
	z.DecSendContainerState(codecSelfer_containerMapEnd1234)
}

func (x *MetricValue) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj17 int
	var yyb17 bool
	var yyhl17 bool = l >= 0
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Kind = ""
	} else {
		yyv18 := &x.Kind
		yym19 := z.DecBinary()
		_ = yym19
		if false {
		} else {
			*((*string)(yyv18)) = r.DecodeString()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.APIVersion = ""
	} else {
		yyv20 := &x.APIVersion
		yym21 := z.DecBinary()
		_ = yym21
		if false {
		} else {
			*((*string)(yyv20)) = r.DecodeString()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.DescribedObject = pkg2_v1.ObjectReference{}
	} else {
		yyv22 := &x.DescribedObject
		yyv22.CodecDecodeSelf(d)
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.MetricName = ""
	} else {
		yyv23 := &x.MetricName
		yym24 := z.DecBinary()
		_ = yym24
		if false {
		} else {
			*((*string)(yyv23)) = r.DecodeString()
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Timestamp = pkg1_v1.Time{}
	} else {
		yyv25 := &x.Timestamp
		yym26 := z.DecBinary()
		_ = yym26
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv25) {
		} else if yym26 {
			z.DecBinaryUnmarshal(yyv25)
		} else if !yym26 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv25)
		} else {
			z.DecFallback(yyv25, false)
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		if x.WindowSeconds != nil {
			x.WindowSeconds = nil
		}
	} else {
		if x.WindowSeconds == nil {
			x.WindowSeconds = new(int64)
		}
		yym28 := z.DecBinary()
		_ = yym28
		if false {
		} else {
			*((*int64)(x.WindowSeconds)) = int64(r.DecodeInt(64))
		}
	}
	yyj17++
	if yyhl17 {
		yyb17 = yyj17 > l
	} else {
		yyb17 = r.CheckBreak()
	}
	if yyb17 {
		z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
		return
	}
	z.DecSendContainerState(codecSelfer_containerArrayElem1234)
	if r.TryDecodeAsNil() {
		x.Value = pkg4_resource.Quantity{}
	} else {
		yyv29 := &x.Value
		yym30 := z.DecBinary()
		_ = yym30
		if false {
		} else if z.HasExtensions() && z.DecExt(yyv29) {
		} else if !yym30 && z.IsJSONHandle() {
			z.DecJSONUnmarshal(yyv29)
		} else {
			z.DecFallback(yyv29, false)
		}
	}
	for {
		yyj17++
		if yyhl17 {
			yyb17 = yyj17 > l
		} else {
			yyb17 = r.CheckBreak()
		}
		if yyb17 {
			break
		}
		z.DecSendContainerState(codecSelfer_containerArrayElem1234)
		z.DecStructFieldNotFound(yyj17-1, "")
	}
	z.DecSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) encSliceMetricValue(v []MetricValue, e *codec1978.Encoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	r.EncodeArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncSendContainerState(codecSelfer_containerArrayElem1234)
		yy2 := &yyv1
		yy2.CodecEncodeSelf(e)
	}
	z.EncSendContainerState(codecSelfer_containerArrayEnd1234)
}

func (x codecSelfer1234) decSliceMetricValue(v *[]MetricValue, d *codec1978.Decoder) {
	var h codecSelfer1234
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []MetricValue{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else if yyl1 > 0 {
		var yyrr1, yyrl1 int
		var yyrt1 bool
		_, _ = yyrl1, yyrt1
		yyrr1 = yyl1 // len(yyv1)
		if yyl1 > cap(yyv1) {

			yyrg1 := len(yyv1) > 0
			yyv21 := yyv1
			yyrl1, yyrt1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 248)
			if yyrt1 {
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]MetricValue, yyrl1)
				}
			} else {
				yyv1 = make([]MetricValue, yyrl1)
			}
			yyc1 = true
			yyrr1 = len(yyv1)
			if yyrg1 {
				copy(yyv1, yyv21)
			}
		} else if yyl1 != len(yyv1) {
			yyv1 = yyv1[:yyl1]
			yyc1 = true
		}
		yyj1 := 0
		for ; yyj1 < yyrr1; yyj1++ {
			yyh1.ElemContainerState(yyj1)
			if r.TryDecodeAsNil() {
				yyv1[yyj1] = MetricValue{}
			} else {
				yyv2 := &yyv1[yyj1]
				yyv2.CodecDecodeSelf(d)
			}

		}
		if yyrt1 {
			for ; yyj1 < yyl1; yyj1++ {
				yyv1 = append(yyv1, MetricValue{})
				yyh1.ElemContainerState(yyj1)
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = MetricValue{}
				} else {
					yyv3 := &yyv1[yyj1]
					yyv3.CodecDecodeSelf(d)
				}

			}
		}

	} else {
		yyj1 := 0
		for ; !r.CheckBreak(); yyj1++ {

			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, MetricValue{}) // var yyz1 MetricValue
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			if yyj1 < len(yyv1) {
				if r.TryDecodeAsNil() {
					yyv1[yyj1] = MetricValue{}
				} else {
					yyv4 := &yyv1[yyj1]
					yyv4.CodecDecodeSelf(d)
				}

			} else {
				z.DecSwallow()
			}

		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = []MetricValue{}
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}
