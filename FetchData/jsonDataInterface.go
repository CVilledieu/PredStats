package APIData

type Data struct {
	matchInfo MatchInfo
}

type MatchInfo struct {
	gameDuration int
	winningTeam  int
}

type PlayerData struct {
	combatData  CombatData
	wardData    WardData
	dmgHealData DmgHealData
	abilityData AbilityData
	itemData    ItemData
}

type CombatData struct {
}
