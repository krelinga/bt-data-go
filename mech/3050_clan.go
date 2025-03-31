package mech

import "github.com/krelinga/bg-data-go/equip"

var madCatBase = New(
	WithName("Mad Cat"),
	WithClanName("Timber Wolf"),
	WithWeight(75),
	WithTwoLegs(),
	WithEngine(EngineClanXL, 5),
	WithArmor(Arms, 24),
	WithArmor(BipedLegs, 32),
	WithArmor(SideTorsos, 25),
	WithArmor(SideTorsoRears, 7),
	WithArmor(LocCenterTorso, 36),
	WithArmor(LocCenterTorsoRear, 9),
	WithArmor(LocHead, 9),
)

var (
	MadCatPrime = Extend(madCatBase,
		WithVariant("Prime"),
		WithMount(equip.ClanMachineGun, LocJoin(LocCenterTorso, LocRightTorso)),
		WithMount(equip.ClanMachineGunAmmo, LocRightTorso),
		WithMount(equip.ClanMediumPulseLaser, LocLeftTorso),
		WithMount(
			equip.Join(equip.ClanLRM20, equip.ClanLRM20Ammo, equip.ClanCASE),
			LocJoin(LocLeftTorso, LocRightTorso),
		),
		WithMount(
			equip.Join(equip.ClanDoubleHeatSink, equip.ClanERLargeLaser, equip.ClanERMediumLaser),
			LocJoin(LocLeftArm, LocRightArm),
		),
	)
	MadCatA = Extend(madCatBase,
		WithVariant("A"),
		WithMount(equip.ClanERSmallLaser, LocCenterTorso),
		WithMount(
			equip.Join(equip.ClanStreakSRM6, equip.ClanStreakSRM6Ammo, equip.ClanCASE),
			LocRightTorso),
		WithMount(equip.Mult(3, equip.ClanMediumPulseLaser), LocLeftTorso),
		WithMount(
			equip.Join(equip.ClanERPPC, equip.Mult(2, equip.ClanDoubleHeatSink)),
			LocJoin(LocLeftArm, LocRightArm),
		),
		WithMount(equip.ClanDoubleHeatSink, LocRightArm),
	)
)
