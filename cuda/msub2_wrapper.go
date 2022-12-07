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

// CUDA handle for msub2 kernel
var msub2_code cu.Function

// Stores the arguments for msub2 kernel invocation
type msub2_args_t struct{
	 arg_dst unsafe.Pointer
	 arg_src1 unsafe.Pointer
	 arg_fac1 float32
	 arg_src2 unsafe.Pointer
	 arg_fac2 float32
	 arg_N int
	 argptr [6]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for msub2 kernel invocation
var msub2_args msub2_args_t

func init(){
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	 msub2_args.argptr[0] = unsafe.Pointer(&msub2_args.arg_dst)
	 msub2_args.argptr[1] = unsafe.Pointer(&msub2_args.arg_src1)
	 msub2_args.argptr[2] = unsafe.Pointer(&msub2_args.arg_fac1)
	 msub2_args.argptr[3] = unsafe.Pointer(&msub2_args.arg_src2)
	 msub2_args.argptr[4] = unsafe.Pointer(&msub2_args.arg_fac2)
	 msub2_args.argptr[5] = unsafe.Pointer(&msub2_args.arg_N)
	 }

// Wrapper for msub2 CUDA kernel, asynchronous.
func k_msub2_async ( dst unsafe.Pointer, src1 unsafe.Pointer, fac1 float32, src2 unsafe.Pointer, fac2 float32, N int,  cfg *config) {
	if Synchronous{ // debug
		Sync()
		timer.Start("msub2")
	}

	msub2_args.Lock()
	defer msub2_args.Unlock()

	if msub2_code == 0{
		msub2_code = fatbinLoad(msub2_map, "msub2")
	}

	 msub2_args.arg_dst = dst
	 msub2_args.arg_src1 = src1
	 msub2_args.arg_fac1 = fac1
	 msub2_args.arg_src2 = src2
	 msub2_args.arg_fac2 = fac2
	 msub2_args.arg_N = N
	

	args := msub2_args.argptr[:]
	cu.LaunchKernel(msub2_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous{ // debug
		Sync()
		timer.Stop("msub2")
	}
}

// maps compute capability on PTX code for msub2 kernel.
var msub2_map = map[int]string{ 0: "" ,
30: msub2_ptx_30 ,
35: msub2_ptx_35 ,
37: msub2_ptx_37 ,
50: msub2_ptx_50 ,
52: msub2_ptx_52 ,
53: msub2_ptx_53 ,
60: msub2_ptx_60 ,
61: msub2_ptx_61 ,
70: msub2_ptx_70 ,
75: msub2_ptx_75  }

// msub2 PTX code for various compute capabilities.
const(
  msub2_ptx_30 = `
.version 6.4
.target sm_30
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_35 = `
.version 6.4
.target sm_35
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_37 = `
.version 6.4
.target sm_37
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_50 = `
.version 6.4
.target sm_50
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_52 = `
.version 6.4
.target sm_52
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_53 = `
.version 6.4
.target sm_53
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_60 = `
.version 6.4
.target sm_60
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_61 = `
.version 6.4
.target sm_61
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_70 = `
.version 6.4
.target sm_70
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
   msub2_ptx_75 = `
.version 6.4
.target sm_75
.address_size 64

	// .globl	msub2

.visible .entry msub2(
	.param .u64 msub2_param_0,
	.param .u64 msub2_param_1,
	.param .f32 msub2_param_2,
	.param .u64 msub2_param_3,
	.param .f32 msub2_param_4,
	.param .u32 msub2_param_5
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<8>;
	.reg .b32 	%r<9>;
	.reg .b64 	%rd<11>;


	ld.param.u64 	%rd1, [msub2_param_0];
	ld.param.u64 	%rd2, [msub2_param_1];
	ld.param.f32 	%f1, [msub2_param_2];
	ld.param.u64 	%rd3, [msub2_param_3];
	ld.param.f32 	%f2, [msub2_param_4];
	ld.param.u32 	%r2, [msub2_param_5];
	mov.u32 	%r3, %ctaid.y;
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.x;
	mad.lo.s32 	%r6, %r4, %r3, %r5;
	mov.u32 	%r7, %ntid.x;
	mov.u32 	%r8, %tid.x;
	mad.lo.s32 	%r1, %r6, %r7, %r8;
	setp.ge.s32	%p1, %r1, %r2;
	@%p1 bra 	BB0_2;

	cvta.to.global.u64 	%rd4, %rd2;
	mul.wide.s32 	%rd5, %r1, 4;
	add.s64 	%rd6, %rd4, %rd5;
	ld.global.nc.f32 	%f3, [%rd6];
	mul.f32 	%f4, %f3, %f1;
	cvta.to.global.u64 	%rd7, %rd3;
	add.s64 	%rd8, %rd7, %rd5;
	ld.global.nc.f32 	%f5, [%rd8];
	mul.f32 	%f6, %f5, %f2;
	sub.f32 	%f7, %f4, %f6;
	cvta.to.global.u64 	%rd9, %rd1;
	add.s64 	%rd10, %rd9, %rd5;
	st.global.f32 	[%rd10], %f7;

BB0_2:
	ret;
}


`
 )
