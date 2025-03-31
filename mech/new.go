package mech

import "github.com/krelinga/bg-data-go/equip"

func New(opts... Option) Mech {
	return Extend(Mech{}, opts...)
}

func Extend(base Mech, opts... Option) Mech {
	for _, opt := range opts {
		opt(&base)
	}
	return base
}

type Option func(m *Mech)

func WithName(name string) Option {
	return func(m *Mech) {
		m.Name = name
	}
}

func WithClanName(clanName string) Option {
	return func(m *Mech) {
		m.ClanName = clanName
	}
}

func WithTwoLegs(weight int) Option {
	return func(m *Mech) {
		m.Weight = weight
		structureTable(m, false)
	}
}

func WithFourLegs(weight int) Option {
	return func(m *Mech) {
		m.Weight = weight
		structureTable(m, true)
	}
}

func structureTable(m *Mech, quad bool) {
	m.Structure = make(map[Loc]int)
	m.Structure[LocHead] = 3

	ct := func(x int) {
		m.Structure[LocCenterTorso] = x
	}
	st := func(x int) {
		m.Structure[LocLeftTorso] = x
		m.Structure[LocRightTorso] = x
	}
	a := func(x int) {
		if !quad {
			m.Structure[LocLeftArm] = x
			m.Structure[LocRightArm] = x
		}
	}
	l := func(x int) {
		if quad {
			m.Structure[LocLeftFrontLeg] = x
			m.Structure[LocRightFrontLeg] = x
			m.Structure[LocLeftRearLeg] = x
			m.Structure[LocRightRearLeg] = x
		} else {
			m.Structure[LocLeftLeg] = x
			m.Structure[LocRightLeg] = x
		}
	}
	switch m.Weight {
	case 20:
		ct(6)
		st(5)
		a(3)
		l(4)
	case 25:
		ct(8)
		st(6)
		a(4)
		l(6)
	case 30:
		ct(10)
		st(7)
		a(5)
		l(7)
	case 35:
		ct(11)
		st(8)
		a(6)
		l(8)
	case 40:
		ct(12)
		st(10)
		a(6)
		l(10)
	case 45:
		ct(14)
		st(11)
		a(7)
		l(11)
	case 50:
		ct(16)
		st(12)
		a(8)
		l(12)
	case 55:
		ct(18)
		st(13)
		a(9)
		l(13)
	case 60:
		ct(20)
		st(14)
		a(10)
		l(14)
	case 65:
		ct(21)
		st(15)
		a(10)
		l(15)
	case 70:
		ct(22)
		st(15)
		a(11)
		l(15)
	case 75:
		ct(23)
		st(16)
		a(12)
		l(16)
	case 80:
		ct(25)
		st(17)
		a(13)
		l(17)
	case 85:
		ct(27)
		st(18)
		a(14)
		l(18)
	case 90:
		ct(29)
		st(19)
		a(15)
		l(19)
	case 95:
		ct(30)
		st(20)
		a(16)
		l(20)
	case 100:
		ct(31)
		st(21)
		a(17)
		l(21)
	default:
		panic("weight not supported")
	}
}

func WithArmArmor(x int) Option {
	return func(m *Mech) {
		m.Armor[LocLeftArm] = x
		m.Armor[LocRightArm] = x
	}
}

func WithLegArmor(x int) Option {
	return func(m *Mech) {
		m.Armor[LocLeftLeg] = x
		m.Armor[LocRightLeg] = x
	}
}

func WithSideTorsoArmor(f, r int) Option {
	return func(m *Mech) {
		m.Armor[LocLeftTorso] = f
		m.Armor[LocRightTorso] = f
		m.Armor[LocLeftTorsoRear] = r
		m.Armor[LocRightTorsoRear] = r
	}
}

func WithCenterTorsoArmor(f, r int) Option {
	return func(m *Mech) {
		m.Armor[LocCenterTorso] = f
		m.Armor[LocCenterTorsoRear] = r
	}
}

func WithHeadArmor(x int) Option {
	return func(m *Mech) {
		m.Armor[LocHead] = x
	}
}

func WithFrontLegArmor(x int) Option {
	return func(m *Mech) {
		m.Armor[LocLeftFrontLeg] = x
		m.Armor[LocRightFrontLeg] = x
	}
}

func WithRearLegArmor(x int) Option {
	return func(m *Mech) {
		m.Armor[LocLeftRearLeg] = x
		m.Armor[LocRightRearLeg] = x
	}
}

func WithEngine(engine Engine, walkSpeed int) Option {
	return func(m *Mech) {
		m.Engine = engine
		m.WalkSpeed = walkSpeed
	}
}

func WithMount(equip equip.Equip, loc Loc) Option {
	return WithMountN(1, equip, loc)
}

func WithMountN(n int, equip equip.Equip, loc Loc) Option {
	return func(m *Mech) {
		for range n {
			mount(m, equip, loc)
		}
	}
}

func WithSplitMount(equip equip.Equip, locSizes ...LocSize) Option {
	return func(m *Mech) {
		m.SplitMounts = append(m.SplitMounts, SplitMount{Equip: equip, LocSizes: locSizes})
	}
}

func WithArmMount(equip equip.Equip) Option {
	return func(m *Mech) {
		mount(m, equip, LocLeftArm, LocRightArm)
	}
}

func WithSideTorsoMount(equip equip.Equip) Option {
	return func(m *Mech) {
		mount(m, equip, LocLeftTorso, LocRightTorso)
	}
}

func WithLegMount(equip equip.Equip) Option {
	return func(m *Mech) {
		mount(m, equip, LocLeftLeg, LocRightLeg)
	}
}

func mount(m *Mech, equip equip.Equip, loc... Loc) {
	for _, l := range loc {
		m.Mounts = append(m.Mounts, Mount{Equip: equip, Loc: l})
	}
}

func WithVariant(varName string) Option {
	return func(m *Mech) {
		if m.Name == "" {
			panic("variant name set but mech name is empty")
		}
		m.Name = m.Name + " " + varName
		if m.ClanName != "" {
			m.ClanName = m.ClanName + " " + varName
		}
	}
}