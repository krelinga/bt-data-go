package equip

import "iter"

type Equip struct{}

func (e Equip) List() iter.Seq[Equip] {
	return func(yield func(Equip) bool) {
		yield(e)
	}
}

var (
	ClanMachineGun       Equip
	ClanMachineGunAmmo   Equip
	ClanMediumPulseLaser Equip
	ClanLRM20            Equip
	ClanERLargeLaser     Equip
	ClanERMediumLaser    Equip
	ClanDoubleHeatSink   Equip
	ClanLRM20Ammo        Equip
	ClanCASE             Equip
	ClanERSmallLaser     Equip
	ClanStreakSRM6       Equip
	ClanStreakSRM6Ammo   Equip
	ClanERPPC            Equip
)

type Lister interface {
	List() iter.Seq[Equip]
}

func Join(in ...Lister) Lister {
	return join(in)
}

type join []Lister

func (j join) List() iter.Seq[Equip] {
	return func(yield func(Equip) bool) {
		for _, l := range j {
			for equip := range l.List() {
				if !yield(equip) {
					return
				}
			}
		}
	}
}

func Mult(n int, in Lister) Lister {
	return mult{n, in}
}

type mult struct {
	n  int
	in Lister
}

func (m mult) List() iter.Seq[Equip] {
	return func(yield func(Equip) bool) {
		for range m.n {
			for equip := range m.in.List() {
				if !yield(equip) {
					return
				}
			}
		}
	}
}
