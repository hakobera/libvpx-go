// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 06 Jan 2016 08:36:22 MSK.
// By http://git.io/cgogen. DO NOT EDIT.

package vpx

/*
#cgo pkg-config: vpx
#include <vpx/vp8.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// RefFrameType as declared in vpx/vp8.h:99
type RefFrameType uint

// RefFrameType enumeration from vpx/vp8.h:99
const (
	Vp8LastFrame RefFrameType = C.VP8_LAST_FRAME
	Vp8GoldFrame RefFrameType = C.VP8_GOLD_FRAME
	Vp8AltrFrame RefFrameType = C.VP8_ALTR_FRAME
)

// CodecErr as declared in vpx/vpx_codec.h:142
type CodecErr uint

// CodecErr enumeration from vpx/vpx_codec.h:142
const (
	CodecOk             CodecErr = 0
	CodecError          CodecErr = C.VPX_CODEC_ERROR
	CodecMemError       CodecErr = C.VPX_CODEC_MEM_ERROR
	CodecABIMismatch    CodecErr = C.VPX_CODEC_ABI_MISMATCH
	CodecIncapable      CodecErr = C.VPX_CODEC_INCAPABLE
	CodecUnsupBitstream CodecErr = C.VPX_CODEC_UNSUP_BITSTREAM
	CodecUnsupFeature   CodecErr = C.VPX_CODEC_UNSUP_FEATURE
	CodecCorruptFrame   CodecErr = C.VPX_CODEC_CORRUPT_FRAME
	CodecInvalidParam   CodecErr = C.VPX_CODEC_INVALID_PARAM
	CodecListEnd        CodecErr = C.VPX_CODEC_LIST_END
)

// BitDepth as declared in vpx/vpx_codec.h:223
type BitDepth uint

// BitDepth enumeration from vpx/vpx_codec.h:223
const (
	Bits8  BitDepth = C.VPX_BITS_8
	Bits10 BitDepth = C.VPX_BITS_10
	Bits12 BitDepth = C.VPX_BITS_12
)

// ImageFormat as declared in vpx/vpx_image.h:67
type ImageFormat uint

// ImageFormat enumeration from vpx/vpx_image.h:67
const (
	ImageFormatNone     ImageFormat = 0
	ImageFormatRgb24    ImageFormat = C.VPX_IMG_FMT_RGB24
	ImageFormatRgb32    ImageFormat = C.VPX_IMG_FMT_RGB32
	ImageFormatRgb565   ImageFormat = C.VPX_IMG_FMT_RGB565
	ImageFormatRgb555   ImageFormat = C.VPX_IMG_FMT_RGB555
	ImageFormatUyvy     ImageFormat = C.VPX_IMG_FMT_UYVY
	ImageFormatYuy2     ImageFormat = C.VPX_IMG_FMT_YUY2
	ImageFormatYvyu     ImageFormat = C.VPX_IMG_FMT_YVYU
	ImageFormatBgr24    ImageFormat = C.VPX_IMG_FMT_BGR24
	ImageFormatRgb32Le  ImageFormat = C.VPX_IMG_FMT_RGB32_LE
	ImageFormatArgb     ImageFormat = C.VPX_IMG_FMT_ARGB
	ImageFormatArgbLe   ImageFormat = C.VPX_IMG_FMT_ARGB_LE
	ImageFormatRgb565Le ImageFormat = C.VPX_IMG_FMT_RGB565_LE
	ImageFormatRgb555Le ImageFormat = C.VPX_IMG_FMT_RGB555_LE
	ImageFormatYv12     ImageFormat = C.VPX_IMG_FMT_YV12
	ImageFormatI420     ImageFormat = C.VPX_IMG_FMT_I420
	ImageFormatVpxyv12  ImageFormat = C.VPX_IMG_FMT_VPXYV12
	ImageFormatVpxi420  ImageFormat = C.VPX_IMG_FMT_VPXI420
	ImageFormatI422     ImageFormat = C.VPX_IMG_FMT_I422
	ImageFormatI444     ImageFormat = C.VPX_IMG_FMT_I444
	ImageFormatI440     ImageFormat = C.VPX_IMG_FMT_I440
	ImageFormat444a     ImageFormat = C.VPX_IMG_FMT_444A
	ImageFormatI42016   ImageFormat = C.VPX_IMG_FMT_I42016
	ImageFormatI42216   ImageFormat = C.VPX_IMG_FMT_I42216
	ImageFormatI44416   ImageFormat = C.VPX_IMG_FMT_I44416
	ImageFormatI44016   ImageFormat = C.VPX_IMG_FMT_I44016
)

// ColorSpace as declared in vpx/vpx_image.h:79
type ColorSpace uint

// ColorSpace enumeration from vpx/vpx_image.h:79
const (
	ColorSpaceUnknown  ColorSpace = C.VPX_CS_UNKNOWN
	ColorSpaceBt601    ColorSpace = C.VPX_CS_BT_601
	ColorSpaceBt709    ColorSpace = C.VPX_CS_BT_709
	ColorSpaceSmpte170 ColorSpace = C.VPX_CS_SMPTE_170
	ColorSpaceSmpte240 ColorSpace = C.VPX_CS_SMPTE_240
	ColorSpaceBt2020   ColorSpace = C.VPX_CS_BT_2020
	ColorSpaceReserved ColorSpace = C.VPX_CS_RESERVED
	ColorSpaceSrgb     ColorSpace = C.VPX_CS_SRGB
)
