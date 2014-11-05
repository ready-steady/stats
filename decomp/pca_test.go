package decomp

import (
	"testing"

	"github.com/ready-steady/support/assert"
)

func TestCovPCA(t *testing.T) {
	C := []float64{
		+1.000000000000000e+00,
		+1.154127058177033e-01,
		+3.134709671301593e-01,
		-2.624605475789556e-01,
		-2.675628902376415e-01,
		+1.154127058177033e-01,
		+1.000000000000000e+00,
		+1.487775664601468e-01,
		+8.251722261621278e-01,
		+6.471493243665016e-01,
		+3.134709671301593e-01,
		+1.487775664601468e-01,
		+1.000000000000000e+00,
		+2.188948568766876e-01,
		-4.746506616202846e-01,
		-2.624605475789556e-01,
		+8.251722261621278e-01,
		+2.188948568766876e-01,
		+1.000000000000000e+00,
		+7.047860759340304e-01,
		-2.675628902376415e-01,
		+6.471493243665017e-01,
		-4.746506616202846e-01,
		+7.047860759340304e-01,
		+1.000000000000000e+00,
	}

	expectedU := []float64{
		-1.777008675251867e-01,
		+5.456098595844089e-01,
		-7.679302728419535e-02,
		+5.862167305996994e-01,
		+5.667319106337718e-01,
		+5.274851240801810e-01,
		+3.209998200695767e-01,
		+7.174077977724287e-01,
		+1.998699259327777e-01,
		-2.531731103264220e-01,
		+7.655128668903998e-01,
		+1.936393917712877e-01,
		-4.774829578033630e-01,
		-2.733383132570983e-01,
		+2.716431999752671e-01,
		-2.541030814371850e-01,
		+7.277625876615521e-01,
		-2.402296315689612e-01,
		-2.880993167658167e-01,
		-5.148609014089138e-01,
		+1.990063320205669e-01,
		-1.792607000828338e-01,
		-4.401461482047020e-01,
		+6.772642819885449e-01,
		-5.252109497950679e-01,
	}

	expectedΛ := []float64{
		2.500258665819985e+00,
		1.525542276483301e+00,
		8.324395052148045e-01,
		1.261010747114109e-01,
		1.565847777049709e-02,
	}

	U, Λ, err := CovPCA(C, 5)

	assert.Success(err, t)
	assert.AlmostEqual(Λ, expectedΛ, t)
	assert.AlmostEqual(abs(U), abs(expectedU), t)
}

func TestCovPCAFailure(t *testing.T) {
	_, _, err := CovPCA([]float64{-1, 0, 0, 0, 1, 0, 0, 0, 1}, 3)

	assert.Failure(err, t)
}

func abs(data []float64) []float64 {
	for i := range data {
		if data[i] < 0 {
			data[i] *= -1
		}
	}

	return data
}
