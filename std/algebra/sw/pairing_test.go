/*
Copyright © 2020 ConsenSys

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

package sw

import (
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	bls12377 "github.com/consensys/gnark-crypto/ecc/bls12-377"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/backend/groth16"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/algebra/fields"
)

type lineEvalBLS377 struct {
	Q, R G2Jac
	P    G1Jac `gnark:",public"`
}

func (circuit *lineEvalBLS377) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	var expected LineEvalRes
	LineEvalBLS377(cs, circuit.Q, circuit.R, circuit.P, &expected, fields.GetBLS377ExtensionFp12(cs))
	cs.AssertIsEqual(expected.r0.A0, "220291599185938038585565774521033812062947190299680306664648725201730830885666933651848261361463591330567860207241")
	cs.AssertIsEqual(expected.r0.A1, "232134458700276476669584229661634543747068594368664068937164975724095736595288995356706959089579876199020312643174")
	cs.AssertIsEqual(expected.r1.A0, "74241662856820718491669277383162555524896537826488558937227282983357670568906847284642533051528779250776935382660")
	cs.AssertIsEqual(expected.r1.A1, "9787836945036920457066634104342154603142239983688979247440278426242314457905122599227144555989168817796094251258")
	cs.AssertIsEqual(expected.r2.A0, "85129589817387660717039592198118788807152207633847410148299763250229022303850156734979397272700502238285752744807")
	cs.AssertIsEqual(expected.r2.A1, "245761211327131018855579902758747359135620549826797077633679496719449586668701082009536667506317412690997533857875")

	return nil
}

func TestLineEvalBLS377(t *testing.T) {

	// create the cs
	var circuit, witness lineEvalBLS377
	r1cs, err := frontend.Compile(ecc.BW6_761, backend.GROTH16, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	var Q, R bls12377.G2Jac
	var P bls12377.G1Jac

	Q.X.A0.SetString("11467063222684898633036104763692544506257812867640109164430855414494851760297509943081481005947955008078272733624")
	Q.X.A1.SetString("153924906120314059329163510034379429156688480181182668999642334674073859906019623717844462092443710331558842221198")
	Q.Y.A0.SetString("217426664443013466493849511677243421913435679616098405782168799962712362374085608530270502677771125796970144049342")
	Q.Y.A1.SetString("220113305559851867470055261956775835250492241909876276448085325823827669499391027597256026508256704101389743638320")
	Q.Z.A0.SetOne()

	R.X.A0.SetString("38348804106969641131654336618231918247608720362924380120333996440589719997236048709530218561145001033408367199467")
	R.X.A1.SetString("208837221672103828632878568310047865523715993428626260492233587961023171407529159232705047544612759994485307437530")
	R.Y.A0.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	R.Y.A1.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")
	R.Z.A0.SetOne()

	P.X.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	P.Y.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")
	P.Z.SetOne()

	witness.Q.Assign(&Q)
	witness.R.Assign(&R)
	witness.P.Assign(&P)

	assert := groth16.NewAssert(t)
	assert.SolvingSucceeded(r1cs, &witness)
}

type lineEvalAffineBLS377 struct {
	Q, R G2Affine
	P    G1Affine `gnark:",public"`
}

func (circuit *lineEvalAffineBLS377) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {
	var expected LineEvalRes
	LineEvalAffineBLS377(cs, circuit.Q, circuit.R, circuit.P, &expected, fields.GetBLS377ExtensionFp12(cs))
	cs.AssertIsEqual(expected.r0.A0, "220291599185938038585565774521033812062947190299680306664648725201730830885666933651848261361463591330567860207241")
	cs.AssertIsEqual(expected.r0.A1, "232134458700276476669584229661634543747068594368664068937164975724095736595288995356706959089579876199020312643174")
	cs.AssertIsEqual(expected.r1.A0, "74241662856820718491669277383162555524896537826488558937227282983357670568906847284642533051528779250776935382660")
	cs.AssertIsEqual(expected.r1.A1, "9787836945036920457066634104342154603142239983688979247440278426242314457905122599227144555989168817796094251258")
	cs.AssertIsEqual(expected.r2.A0, "85129589817387660717039592198118788807152207633847410148299763250229022303850156734979397272700502238285752744807")
	cs.AssertIsEqual(expected.r2.A1, "245761211327131018855579902758747359135620549826797077633679496719449586668701082009536667506317412690997533857875")

	return nil
}

func TestLineEvalAffineBLS377(t *testing.T) {

	// create the cs
	var circuit, witness lineEvalAffineBLS377
	r1cs, err := frontend.Compile(ecc.BW6_761, backend.GROTH16, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	var Q, R bls12377.G2Affine
	var P bls12377.G1Affine

	Q.X.A0.SetString("11467063222684898633036104763692544506257812867640109164430855414494851760297509943081481005947955008078272733624")
	Q.X.A1.SetString("153924906120314059329163510034379429156688480181182668999642334674073859906019623717844462092443710331558842221198")
	Q.Y.A0.SetString("217426664443013466493849511677243421913435679616098405782168799962712362374085608530270502677771125796970144049342")
	Q.Y.A1.SetString("220113305559851867470055261956775835250492241909876276448085325823827669499391027597256026508256704101389743638320")

	R.X.A0.SetString("38348804106969641131654336618231918247608720362924380120333996440589719997236048709530218561145001033408367199467")
	R.X.A1.SetString("208837221672103828632878568310047865523715993428626260492233587961023171407529159232705047544612759994485307437530")
	R.Y.A0.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	R.Y.A1.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")

	P.X.SetString("219129261975485221488302932474367447253380009436652290437731529751224807932621384667224625634955419310221362804739")
	P.Y.SetString("62857965187173987050461294586432573826521562230975685098398439555961148392353952895313161290735015726193379258321")

	witness.Q.Assign(&Q)
	witness.R.Assign(&R)
	witness.P.Assign(&P)

	assert := groth16.NewAssert(t)
	assert.SolvingSucceeded(r1cs, &witness)
}

type pairingAffineBLS377 struct {
	Q          G2Affine
	P          G1Affine `gnark:",public"`
	pairingRes bls12377.GT
}

func (circuit *pairingAffineBLS377) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {

	ateLoop := uint64(9586122913090633729)
	ext := fields.GetBLS377ExtensionFp12(cs)
	pairingInfo := PairingContext{AteLoop: ateLoop, Extension: ext}

	milRes := fields.E12{}
	pairingRes := fields.E12{}

	MillerLoopAffine(cs, circuit.P, circuit.Q, &milRes, pairingInfo)
	pairingRes.FinalExponentiation(cs, &milRes, ateLoop, ext)

	mustbeEq(cs, pairingRes, &circuit.pairingRes)

	return nil
}

func TestPairingAffineBLS377(t *testing.T) {

	P, Q, pairingRes := pairingData()

	// create cs
	var circuit, witness pairingAffineBLS377
	circuit.pairingRes = pairingRes
	r1cs, err := frontend.Compile(ecc.BW6_761, backend.GROTH16, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	// set the cs
	witness.P.Assign(&P)
	witness.Q.Assign(&Q)

	assert := groth16.NewAssert(t)
	assert.SolvingSucceeded(r1cs, &witness)

}

type pairingBLS377 struct {
	P          G1Affine `gnark:",public"`
	Q          G2Affine
	pairingRes bls12377.GT
}

func (circuit *pairingBLS377) Define(curveID ecc.ID, cs *frontend.ConstraintSystem) error {

	ateLoop := uint64(9586122913090633729)
	ext := fields.GetBLS377ExtensionFp12(cs)
	pairingInfo := PairingContext{AteLoop: ateLoop, Extension: ext}
	pairingInfo.BTwistCoeff.A0 = cs.Constant(0)
	pairingInfo.BTwistCoeff.A1 = cs.Constant("155198655607781456406391640216936120121836107652948796323930557600032281009004493664981332883744016074664192874906")

	milRes := fields.E12{}
	MillerLoop(cs, circuit.P, circuit.Q, &milRes, pairingInfo)

	pairingRes := fields.E12{}
	pairingRes.FinalExponentiation(cs, &milRes, ateLoop, ext)

	mustbeEq(cs, pairingRes, &circuit.pairingRes)

	return nil
}

func TestPairingBLS377(t *testing.T) {

	// pairing test data
	P, Q, pairingRes := pairingData()

	// create cs
	var circuit, witness pairingBLS377
	circuit.pairingRes = pairingRes
	r1cs, err := frontend.Compile(ecc.BW6_761, backend.GROTH16, &circuit)
	if err != nil {
		t.Fatal(err)
	}

	// assign values to witness
	witness.P.Assign(&P)
	witness.Q.Assign(&Q)

	assert := groth16.NewAssert(t)
	assert.SolvingSucceeded(r1cs, &witness)

}

func pairingData() (P bls12377.G1Affine, Q bls12377.G2Affine, pairingRes bls12377.GT) {
	_, _, P, Q = bls12377.Generators()
	milRes, _ := bls12377.MillerLoop([]bls12377.G1Affine{P}, []bls12377.G2Affine{Q})
	pairingRes = bls12377.FinalExponentiation(&milRes)
	return
}

func mustbeEq(cs *frontend.ConstraintSystem, fp12 fields.E12, e12 *bls12377.GT) {
	cs.AssertIsEqual(fp12.C0.B0.A0, e12.C0.B0.A0)
	cs.AssertIsEqual(fp12.C0.B0.A1, e12.C0.B0.A1)
	cs.AssertIsEqual(fp12.C0.B1.A0, e12.C0.B1.A0)
	cs.AssertIsEqual(fp12.C0.B1.A1, e12.C0.B1.A1)
	cs.AssertIsEqual(fp12.C0.B2.A0, e12.C0.B2.A0)
	cs.AssertIsEqual(fp12.C0.B2.A1, e12.C0.B2.A1)
	cs.AssertIsEqual(fp12.C1.B0.A0, e12.C1.B0.A0)
	cs.AssertIsEqual(fp12.C1.B0.A1, e12.C1.B0.A1)
	cs.AssertIsEqual(fp12.C1.B1.A0, e12.C1.B1.A0)
	cs.AssertIsEqual(fp12.C1.B1.A1, e12.C1.B1.A1)
	cs.AssertIsEqual(fp12.C1.B2.A0, e12.C1.B2.A0)
	cs.AssertIsEqual(fp12.C1.B2.A1, e12.C1.B2.A1)
}
