package mech

import (
	"iter"

	"github.com/krelinga/bg-data-go/equip"
)

type Mech struct {
	Name, ClanName string
	Weight int
	Structure map[Loc]int
	Armor map[Loc]int
	Engine Engine
	WalkSpeed int
	Mounts []Mount
	SplitMounts []SplitMount
}

type Engine int

const (
	EngineNil Engine = iota
	EngineStandard
	EngineXL
	EngineClanXL
	EngineMax
)

type Loc int

func (l Loc) LocList() iter.Seq[Loc] {
	return func(yield func(Loc) bool) {
		yield(l)
	}
}

const (
	LocNil Loc = iota
	LocHead
	LocCenterTorso
	LocCenterTorsoRear
	LocLeftTorso
	LocLeftTorsoRear
	LocRightTorso
	LocRightTorsoRear
	LocLeftArm
	LocRightArm
	LocLeftLeg
	LocRightLeg
	LocLeftFrontLeg
	LocRightFrontLeg
	LocLeftRearLeg
	LocRightRearLeg
	LocMax
)

type Mount struct {
	Equip equip.Equip
	Loc Loc
}

type SplitMount struct {
	Equip equip.Equip
	LocSizes []LocSize
}

type LocSize struct {
	Loc Loc
	Size int
}

type LocLister interface {
	LocList() iter.Seq[Loc]
}

func LocJoin(in... LocLister) LocLister {
	return locJoin(in)
}

type locJoin []LocLister

func (lj locJoin) LocList() iter.Seq[Loc] {
	return func(yield func(Loc) bool) {
		for _, ll := range lj {
			for loc := range ll.LocList() {
				if !yield(loc) {
					return
				}
			}
		}
	}
}

var (
	Arms = LocJoin(
		LocLeftArm,
		LocRightArm,
	)
	BipedLegs = LocJoin(
		LocLeftLeg,
		LocRightLeg,
	)
	SideTorsos = LocJoin(
		LocLeftTorso,
		LocRightTorso,
	)
	SideTorsoRears = LocJoin(
		LocLeftTorsoRear,
		LocRightTorsoRear,
	)
)