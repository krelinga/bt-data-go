package mech

import "github.com/krelinga/bg-data-go/equip"

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