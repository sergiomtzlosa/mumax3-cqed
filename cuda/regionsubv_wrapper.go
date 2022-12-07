package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import(
	"unsafe"
	"github.com/mumax/3/cuda/cu"
	"github.com/mumax/3/timer"
	"sync"
)

// CUDA handle for regionsubv kernel
var regionsubv_code cu.Function

// Stores the arguments for regionsubv kernel invocation
type regionsubv_args_t struct{
	 arg_dstx unsafe.Pointer
	 arg_dsty unsafe.Pointer
	 arg_dstz unsafe.Pointer
	 arg_LUTx unsafe.Pointer
	 arg_LUTy unsafe.Pointer
	 arg_LUTz unsafe.Pointer
	 arg_regions unsafe.Pointer
	 arg_N int
	 argptr [8]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for regionsubv kernel invocation
var regionsubv_args regionsubv_args_t

func init(){
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	 regionsubv_args.argptr[0] = unsafe.Pointer(&regionsubv_args.arg_dstx)
	 regionsubv_args.argptr[1] = unsafe.Pointer(&regionsubv_args.arg_dsty)
	 regionsubv_args.argptr[2] = unsafe.Pointer(&regionsubv_args.arg_dstz)
	 regionsubv_args.argptr[3] = unsafe.Pointer(&regionsubv_args.arg_LUTx)
	 regionsubv_args.argptr[4] = unsafe.Pointer(&regionsubv_args.arg_LUTy)
	 regionsubv_args.argptr[5] = unsafe.Pointer(&regionsubv_args.arg_LUTz)
	 regionsubv_args.argptr[6] = unsafe.Pointer(&regionsubv_args.arg_regions)
	 regionsubv_args.argptr[7] = unsafe.Pointer(&regionsubv_args.arg_N)
	 }

// Wrapper for regionsubv CUDA kernel, asynchronous.
func k_regionsubv_async ( dstx unsafe.Pointer, dsty unsafe.Pointer, dstz unsafe.Pointer, LUTx unsafe.Pointer, LUTy unsafe.Pointer, LUTz unsafe.Pointer, regions unsafe.Pointer, N int,  cfg *config) {
	if Synchronous{ // debug
		Sync()
		timer.Start("regionsubv")
	}

	regionsubv_args.Lock()
	defer regionsubv_args.Unlock()

	if regionsubv_code == 0{
		regionsubv_code = fatbinLoad(regionsubv_map, "regionsubv")
	}

	 regionsubv_args.arg_dstx = dstx
	 regionsubv_args.arg_dsty = dsty
	 regionsubv_args.arg_dstz = dstz
	 regionsubv_args.arg_LUTx = LUTx
	 regionsubv_args.arg_LUTy = LUTy
	 regionsubv_args.arg_LUTz = LUTz
	 regionsubv_args.arg_regions = regions
	 regionsubv_args.arg_N = N
	

	args := regionsubv_args.argptr[:]
	cu.LaunchKernel(regionsubv_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous{ // debug
		Sync()
		timer.Stop("regionsubv")
	}
}

// maps compute capability on PTX code for regionsubv kernel.
var regionsubv_map = map[int]string{ 0: "" ,
30: regionsubv_ptx_30 ,
35: regionsubv_ptx_35 ,
37: regionsubv_ptx_37 ,
50: regionsubv_ptx_50 ,
52: regionsubv_ptx_52 ,
53: regionsubv_ptx_53 ,
60: regionsubv_ptx_60 ,
61: regionsubv_ptx_61 ,
70: regionsubv_ptx_70 ,
75: regionsubv_ptx_75  }

// regionsubv PTX code for various compute capabilities.
const(
  regionsubv_ptx_30 = `
.version 6.4
.target sm_30
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<10>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	cvta.to.global.u64 	%rd11, %rd4;
	ld.global.u8 	%r9, [%rd10];
	mul.wide.u32 	%rd12, %r9, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_35 = `
.version 6.4
.target sm_35
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_37 = `
.version 6.4
.target sm_37
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_50 = `
.version 6.4
.target sm_50
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_52 = `
.version 6.4
.target sm_52
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_53 = `
.version 6.4
.target sm_53
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_60 = `
.version 6.4
.target sm_60
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_61 = `
.version 6.4
.target sm_61
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_70 = `
.version 6.4
.target sm_70
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
   regionsubv_ptx_75 = `
.version 6.4
.target sm_75
.address_size 64

	// .globl	regionsubv

.visible .entry regionsubv(
	.param .u64 regionsubv_param_0,
	.param .u64 regionsubv_param_1,
	.param .u64 regionsubv_param_2,
	.param .u64 regionsubv_param_3,
	.param .u64 regionsubv_param_4,
	.param .u64 regionsubv_param_5,
	.param .u64 regionsubv_param_6,
	.param .u32 regionsubv_param_7
)
{
	.reg .pred 	%p<2>;
	.reg .b16 	%rs<2>;
	.reg .f32 	%f<10>;
	.reg .b32 	%r<11>;
	.reg .b64 	%rd<25>;


	ld.param.u64 	%rd1, [regionsubv_param_0];
	ld.param.u64 	%rd2, [regionsubv_param_1];
	ld.param.u64 	%rd3, [regionsubv_param_2];
	ld.param.u64 	%rd4, [regionsubv_param_3];
	ld.param.u64 	%rd5, [regionsubv_param_4];
	ld.param.u64 	%rd6, [regionsubv_param_5];
	ld.param.u64 	%rd7, [regionsubv_param_6];
	ld.param.u32 	%r2, [regionsubv_param_7];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd8, %rd7;
	cvt.s64.s32	%rd9, %r1;
	add.s64 	%rd10, %rd8, %rd9;
	ld.global.nc.u8 	%rs1, [%rd10];
	cvta.to.global.u64 	%rd11, %rd4;
	cvt.u32.u16	%r9, %rs1;
	and.b32  	%r10, %r9, 255;
	mul.wide.u32 	%rd12, %r10, 4;
	add.s64 	%rd13, %rd11, %rd12;
	cvta.to.global.u64 	%rd14, %rd1;
	mul.wide.s32 	%rd15, %r1, 4;
	add.s64 	%rd16, %rd14, %rd15;
	ld.global.f32 	%f1, [%rd16];
	ld.global.nc.f32 	%f2, [%rd13];
	sub.f32 	%f3, %f1, %f2;
	st.global.f32 	[%rd16], %f3;
	cvta.to.global.u64 	%rd17, %rd5;
	add.s64 	%rd18, %rd17, %rd12;
	cvta.to.global.u64 	%rd19, %rd2;
	add.s64 	%rd20, %rd19, %rd15;
	ld.global.f32 	%f4, [%rd20];
	ld.global.nc.f32 	%f5, [%rd18];
	sub.f32 	%f6, %f4, %f5;
	st.global.f32 	[%rd20], %f6;
	cvta.to.global.u64 	%rd21, %rd6;
	add.s64 	%rd22, %rd21, %rd12;
	cvta.to.global.u64 	%rd23, %rd3;
	add.s64 	%rd24, %rd23, %rd15;
	ld.global.f32 	%f7, [%rd24];
	ld.global.nc.f32 	%f8, [%rd22];
	sub.f32 	%f9, %f7, %f8;
	st.global.f32 	[%rd24], %f9;

BB0_2:
	ret;
}


`
 )
