package main

func main() {
	var testStruct *Ash = &Ash{On: true, Ammo: 100, Power: 1000}
	testStruct.Shoot()
	testStruct.Shoot()
	testStruct.Shoot()
	testStruct.RideBike()
}

type Ash struct {
	On    bool
	Ammo  int
	Power int
}

func (a *Ash) Shoot() bool {
	if !a.On || a.Ammo < 1 {
		return false
	}
	a.Ammo--
	return true
}

func (a *Ash) RideBike() bool {
	if !a.On || a.Power < 1 {
		return false
	}
	a.Power--
	return true
}
