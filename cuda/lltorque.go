package cuda

import (
	"github.com/mumax/3/data"
)

// Landau-Lifshitz torque divided by gamma0:
// 	- 1/(1+α²) [ m x B +  α m x (m x B) ]
// 	torque in Tesla
// 	m normalized
// 	B in Tesla
// see lltorque.cu
func LLTorque(torque, m, B *data.Slice, alpha MSlice) {
	N := torque.Len()
	cfg := make1DConf(N)

	// k_lltorque2_async(torque.DevPtr(X), torque.DevPtr(Y), torque.DevPtr(Z),
	// 	m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
	// 	B.DevPtr(X), B.DevPtr(Y), B.DevPtr(Z),
	// 	alpha.DevPtr(0), alpha.Mul(0), N, util.Bextra_vector, cfg)

	// log.Println("Time_cuda:", Time_cuda)
	// log.Println("Fixed_dt_cuda:", Fixed_dt_cuda)
	// log.Println("Stop_time_cuda: ", Stop_time_cuda)

	k_lltorque2_async(torque.DevPtr(X), torque.DevPtr(Y), torque.DevPtr(Z),
		m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
		B.DevPtr(X), B.DevPtr(Y), B.DevPtr(Z),
		alpha.DevPtr(0), alpha.Mul(0), N, Time_cuda.DevPtr(0), Fixed_dt_cuda.DevPtr(0), Stop_time_cuda, Wc_cuda, Brms_cuda[X], Brms_cuda[Y], Brms_cuda[Z], Step_Times.DevPtr(X), cfg)

	// log.Println("Time_cuda: ", GetElem(Time_cuda, 0, 0))
	// log.Println("Fixed_dt_cuda: ", GetElem(Fixed_dt_cuda, 0, 0))
	// log.Println("Bx: ", GetElem(B, 0, 0))
	// log.Println("Step_Times: ", GetElem(Step_Times, 0, 3))
}

// Landau-Lifshitz torque with precession disabled.
// Used by engine.Relax().
func LLNoPrecess(torque, m, B *data.Slice) {
	N := torque.Len()
	cfg := make1DConf(N)

	k_llnoprecess_async(torque.DevPtr(X), torque.DevPtr(Y), torque.DevPtr(Z),
		m.DevPtr(X), m.DevPtr(Y), m.DevPtr(Z),
		B.DevPtr(X), B.DevPtr(Y), B.DevPtr(Z), N, cfg)
}
