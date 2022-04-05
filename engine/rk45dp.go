package engine

// MODIFIED INMA
import (
	"math"

	"github.com/mumax/3/cuda"
	"github.com/mumax/3/data"
	"github.com/mumax/3/util"
)

// var (
// 	slice_temp *data.Slice
// 	time_temp  float64
// )

type RK45DP struct {
	k1 *data.Slice // torque at end of step is kept for beginning of next step

}

func (rk *RK45DP) Step() {
	m := M.Buffer()
	size := m.Size()

	if TimeEvolution {
		setStatusLock(false)
	}

	// log.Println("Z elem: ", cuda.GetZElem(m))

	if TimeEvolution {
		initMRKArray(size)
	}

	if FixDt != 0 {
		Dt_si = FixDt
	}

	// upon resize: remove wrongly sized k1
	if rk.k1.Size() != m.Size() {
		rk.Free()
	}

	// first step ever: one-time k1 init and eval
	if rk.k1 == nil {
		rk.k1 = cuda.NewSlice(3, size)

		// log.Println("buffer1: ", M.Buffer().Size())
		torqueFn(rk.k1)
		// log.Println("pasa runge-kutta1")
	}

	// FSAL cannot be used with finite temperature
	if !Temp.isZero() {

		torqueFn(rk.k1)
		// log.Println("pasa runge-kutta0")
	}

	t0 := Time
	// backup magnetization
	m0 := cuda.Buffer(3, size)
	defer cuda.Recycle(m0)
	data.Copy(m0, m)

	k2, k3, k4, k5, k6 := cuda.Buffer(3, size), cuda.Buffer(3, size), cuda.Buffer(3, size), cuda.Buffer(3, size), cuda.Buffer(3, size)
	defer cuda.Recycle(k2)
	defer cuda.Recycle(k3)
	defer cuda.Recycle(k4)
	defer cuda.Recycle(k5)
	defer cuda.Recycle(k6)
	// k2 will be re-used as k7

	h := float32(Dt_si * GammaLL) // internal time step = Dt * gammaLL

	// there is no explicit stage 1: k1 from previous step

	// stage 2
	Time = t0 + (1./5.)*Dt_si
	cuda.Madd2(m, m, rk.k1, 1, (1./5.)*h) // m = m*1 + k1*h/5

	M.normalize()
	torqueFn(k2)
	// log.Println("pasa runge-kutta2")

	// stage 3
	Time = t0 + (3./10.)*Dt_si
	cuda.Madd3(m, m0, rk.k1, k2, 1, (3./40.)*h, (9./40.)*h)

	M.normalize()
	torqueFn(k3)
	// log.Println("pasa runge-kutta3")

	// stage 4
	Time = t0 + (4./5.)*Dt_si
	madd4(m, m0, rk.k1, k2, k3, 1, (44./45.)*h, (-56./15.)*h, (32./9.)*h)

	M.normalize()
	torqueFn(k4)
	// log.Println("pasa runge-kutta4")

	// stage 5
	Time = t0 + (8./9.)*Dt_si
	madd5(m, m0, rk.k1, k2, k3, k4, 1, (19372./6561.)*h, (-25360./2187.)*h, (64448./6561.)*h, (-212./729.)*h)

	M.normalize()
	torqueFn(k5)
	// log.Println("pasa runge-kutta5")

	// stage 6
	Time = t0 + (1.)*Dt_si
	madd6(m, m0, rk.k1, k2, k3, k4, k5, 1, (9017./3168.)*h, (-355./33.)*h, (46732./5247.)*h, (49./176.)*h, (-5103./18656.)*h)

	M.normalize()
	torqueFn(k6)
	// log.Println("pasa runge-kutta6")

	// stage 7: 5th order solution
	Time = t0 + (1.)*Dt_si
	// no k2
	madd6(m, m0, rk.k1, k3, k4, k5, k6, 1, (35./384.)*h, (500./1113.)*h, (125./192.)*h, (-2187./6784.)*h, (11./84.)*h) // 5th

	if TimeEvolution {

		// mx_temp := GetCell(m, 0, 0, 0, 0)
		// my_temp := GetCell(m, 1, 0, 0, 0)
		// mz_temp := GetCell(m, 2, 0, 0, 0)
		//
		// log.Println("vvvv rk45:", mx_temp, ", ", my_temp, ", ", mz_temp)

		attachTimeToFormula(m, Time, true)
	}

	M.normalize()
	k7 := k2     // re-use k2
	torqueFn(k7) // next torque if OK
	// log.Println("pasa runge-kutta7")
	setStatusLock(false)
	// error estimate
	Err := cuda.Buffer(3, size) //k3 // re-use k3 as error estimate
	defer cuda.Recycle(Err)
	madd6(Err, rk.k1, k3, k4, k5, k6, k7, (35./384.)-(5179./57600.), (500./1113.)-(7571./16695.), (125./192.)-(393./640.), (-2187./6784.)-(-92097./339200.), (11./84.)-(187./2100.), (0.)-(1./40.))

	// determine error
	err := cuda.MaxVecNorm(Err) * float64(h)

	// adjust next time step
	if err < MaxErr || Dt_si <= MinDt || FixDt != 0 { // mindt check to avoid infinite loop
		// step OK

		// mx_temp := cuda.GetCell(m, 0, 0, 0, 0)
		// my_temp := cuda.GetCell(m, 1, 0, 0, 0)
		// mz_temp := cuda.GetCell(m, 2, 0, 0, 0)
		//
		// log.Println("vvvv1:", mx_temp, ", ", my_temp, ", ", mz_temp)

		// if TimeEvolution {
		// 	// log.Println("time_temp rk45: ", time_temp)
		// 	cuda.LockMExec = true
		//
		// 	mtemps = append(mtemps, slice_temp)
		// 	times_items = append(times_items, time_temp)
		//
		// 	// log.Println("items: ", len(mtemps))
		// 	// log.Println("time_temp: ", Time)
		// 	// cuda.MdataTemp(cuda.M_rk, slice_temp, float32(cuda.Wc_cuda), time_temp)
		//
		// 	// cuda.M_rk --> destino
		// 	// slice_temp --> m_i actual
		// 	// mtemps --> array con todas la m_i de todos los pasos
		// 	// times_items --> array de tiempos asociados a cada mi
		// 	// time_temp --> tiempo actual
		// 	cuda.MdataTemp(cuda.M_rk, slice_temp, mtemps, times_items, cuda.Wc_cuda, time_temp)
		//
		// 	// log.Println("time_temp: ", time_temp)
		// }

		setLastErr(err)
		setMaxTorque(k7)
		NSteps++
		Time = t0 + Dt_si
		adaptDt(math.Pow(MaxErr/err, 1./5.))
		data.Copy(rk.k1, k7) // FSAL
	} else {
		// undo bad step
		//util.Println("Bad step at t=", t0, ", err=", err)
		util.Assert(FixDt == 0)
		Time = t0
		data.Copy(m, m0)
		NUndone++
		adaptDt(math.Pow(MaxErr/err, 1./6.))
	}
}

func (rk *RK45DP) Free() {
	rk.k1.Free()
	rk.k1 = nil
}

// TODO: into cuda
func madd5(dst, src1, src2, src3, src4, src5 *data.Slice, w1, w2, w3, w4, w5 float32) {
	cuda.Madd3(dst, src1, src2, src3, w1, w2, w3)
	cuda.Madd3(dst, dst, src4, src5, 1, w4, w5)
}

func madd6(dst, src1, src2, src3, src4, src5, src6 *data.Slice, w1, w2, w3, w4, w5, w6 float32) {
	madd5(dst, src1, src2, src3, src4, src5, w1, w2, w3, w4, w5)
	cuda.Madd2(dst, dst, src6, 1, w6)
}
