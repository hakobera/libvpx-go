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
import "unsafe"

// CodecVersion function as declared in vpx/vpx_codec.h:242
func CodecVersion() int {
	ret := C.vpx_codec_version()
	v := (int)(ret)
	return v
}

// CodecVersionStr function as declared in vpx/vpx_codec.h:264
func CodecVersionStr() string {
	ret := C.vpx_codec_version_str()
	v := packPCharString(ret)
	return v
}

// CodecVersionExtraStr function as declared in vpx/vpx_codec.h:273
func CodecVersionExtraStr() string {
	ret := C.vpx_codec_version_extra_str()
	v := packPCharString(ret)
	return v
}

// CodecBuildConfig function as declared in vpx/vpx_codec.h:282
func CodecBuildConfig() string {
	ret := C.vpx_codec_build_config()
	v := packPCharString(ret)
	return v
}

// CodecIfaceName function as declared in vpx/vpx_codec.h:292
func CodecIfaceName(iface []CodecIface) string {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iface)).Data)), cgoAllocsUnknown
	ret := C.vpx_codec_iface_name(ciface)
	v := packPCharString(ret)
	return v
}

// CodecErrToString function as declared in vpx/vpx_codec.h:305
func CodecErrToString(err CodecErr) string {
	cerr, _ := (C.vpx_codec_err_t)(err), cgoAllocsUnknown
	ret := C.vpx_codec_err_to_string(cerr)
	v := packPCharString(ret)
	return v
}

// CodecGetError function as declared in vpx/vpx_codec.h:318
func CodecGetError(ctx []CodecCtx) string {
	cctx, _ := unpackArgSCodecCtx(ctx)
	ret := C.vpx_codec_error(cctx)
	packSCodecCtx(ctx, cctx)
	v := packPCharString(ret)
	return v
}

// CodecErrorDetail function as declared in vpx/vpx_codec.h:331
func CodecErrorDetail(ctx []CodecCtx) string {
	cctx, _ := unpackArgSCodecCtx(ctx)
	ret := C.vpx_codec_error_detail(cctx)
	packSCodecCtx(ctx, cctx)
	v := packPCharString(ret)
	return v
}

// CodecDestroy function as declared in vpx/vpx_codec.h:351
func CodecDestroy(ctx []CodecCtx) CodecErr {
	cctx, _ := unpackArgSCodecCtx(ctx)
	ret := C.vpx_codec_destroy(cctx)
	packSCodecCtx(ctx, cctx)
	v := (CodecErr)(ret)
	return v
}

// CodecGetCaps function as declared in vpx/vpx_codec.h:361
func CodecGetCaps(iface []CodecIface) CodecCaps {
	ciface, _ := (*C.vpx_codec_iface_t)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&iface)).Data)), cgoAllocsUnknown
	ret := C.vpx_codec_get_caps(ciface)
	v := (CodecCaps)(ret)
	return v
}

// ImageAlloc function as declared in vpx/vpx_image.h:150
func ImageAlloc(img []Image, fmt ImageFormat, dW uint, dH uint, align uint) *Image {
	cimg, _ := unpackArgSImage(img)
	cfmt, _ := (C.vpx_img_fmt_t)(fmt), cgoAllocsUnknown
	cdW, _ := (C.uint)(dW), cgoAllocsUnknown
	cdH, _ := (C.uint)(dH), cgoAllocsUnknown
	calign, _ := (C.uint)(align), cgoAllocsUnknown
	ret := C.vpx_img_alloc(cimg, cfmt, cdW, cdH, calign)
	packSImage(img, cimg)
	v := NewImageRef(ret)
	return v
}

// ImageWrap function as declared in vpx/vpx_image.h:175
func ImageWrap(img []Image, fmt ImageFormat, dW uint, dH uint, align uint, imgData []byte) *Image {
	cimg, _ := unpackArgSImage(img)
	cfmt, _ := (C.vpx_img_fmt_t)(fmt), cgoAllocsUnknown
	cdW, _ := (C.uint)(dW), cgoAllocsUnknown
	cdH, _ := (C.uint)(dH), cgoAllocsUnknown
	calign, _ := (C.uint)(align), cgoAllocsUnknown
	cimgData, _ := (*C.uchar)(unsafe.Pointer((*sliceHeader)(unsafe.Pointer(&imgData)).Data)), cgoAllocsUnknown
	ret := C.vpx_img_wrap(cimg, cfmt, cdW, cdH, calign, cimgData)
	packSImage(img, cimg)
	v := NewImageRef(ret)
	return v
}

// ImageSetRect function as declared in vpx/vpx_image.h:196
func ImageSetRect(img []Image, x uint, y uint, w uint, h uint) int {
	cimg, _ := unpackArgSImage(img)
	cx, _ := (C.uint)(x), cgoAllocsUnknown
	cy, _ := (C.uint)(y), cgoAllocsUnknown
	cw, _ := (C.uint)(w), cgoAllocsUnknown
	ch, _ := (C.uint)(h), cgoAllocsUnknown
	ret := C.vpx_img_set_rect(cimg, cx, cy, cw, ch)
	packSImage(img, cimg)
	v := (int)(ret)
	return v
}

// ImageFlip function as declared in vpx/vpx_image.h:210
func ImageFlip(img []Image) {
	cimg, _ := unpackArgSImage(img)
	C.vpx_img_flip(cimg)
	packSImage(img, cimg)
}

// ImageFree function as declared in vpx/vpx_image.h:218
func ImageFree(img []Image) {
	cimg, _ := unpackArgSImage(img)
	C.vpx_img_free(cimg)
	packSImage(img, cimg)
}
