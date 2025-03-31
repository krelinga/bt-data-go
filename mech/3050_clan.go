package mech

import "github.com/krelinga/bg-data-go/equip"

var madCatBase = New(
	WithName("Mad Cat"),
	WithClanName("Timber Wolf"),
	WithTwoLegs(75),
	WithEngine(EngineClanXL, 5),
	WithArmArmor(24),
	WithLegArmor(32),
	WithSideTorsoArmor(25, 7),
	WithCenterTorsoArmor(36, 9),
	WithHeadArmor(9),
)

var MadCatPrime = Extend(madCatBase,
	WithVariant("Prime"),
	WithMount(equip.ClanMachineGun, LocCenterTorso),
	WithMount(equip.ClanMachineGun, LocRightTorso),
	WithMount(equip.ClanMediumPulseLaser, LocLeftTorso),
	WithSideTorsoMount(equip.ClanLRM20),
	WithArmMount(equip.ClanERLargeLaser),
	WithArmMount(equip.ClanERMediumLaser),
	WithArmMount(equip.ClanDoubleHeatSink),
	WithSideTorsoMount(equip.ClanLRM20Ammo),
	WithMount(equip.ClanMachineGunAmmo, LocRightTorso),
)