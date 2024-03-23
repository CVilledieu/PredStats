package APIData

type PredecessorMatch struct {
	WinningTeam           int64                   `json:"winningTeam"`
	GameDuration          int64                   `json:"gameDuration,omitempty"`
	GameMode              string                  `json:"gameMode,omitempty"`
	MatchID               string                  `json:"matchID,omitempty"`
	Region                string                  `json:"region,omitempty"`
	StartTime             string                  `json:"startTime,omitempty"`
	EndTime               string                  `json:"endTime,omitempty"`
	MatchEndReason        string                  `json:"matchEndReason,omitempty"`
	PlayerData            []*PlayerData           `json:"playerData,omitempty"`
	HeroKills             []*HeroKill             `json:"heroKills,omitempty"`
	StructureDestructions []*StructureDestruction `json:"structureDestructions,omitempty"`
	ObjectiveKills        []*ObjectiveKill        `json:"objectiveKills,omitempty"`
}

type PlayerData struct {
	PlayerID       string           `json:"playerID,omitempty"`
	TeamID         int64            `json:"teamID,omitempty"`
	HeroName       string           `json:"heroName,omitempty"`
	RoleName       string           `json:"roleName,omitempty"`
	PlayerName     string           `json:"playerName,omitempty"`
	MinionData     *MinionData      `json:"minionData,omitempty"`
	CombatData     *CombatData      `json:"combatData,omitempty"`
	DamageHealData *DamageHealData  `json:"damageHealData,omitempty"`
	WardsData      *WardsData       `json:"wardsData,omitempty"`
	IncomeData     *IncomeData      `json:"incomeData,omitempty"`
	AbilityData    []*AbilityData   `json:"abilityData,omitempty"`
	InventoryData  []*InventoryData `json:"inventoryData,omitempty"`
}

type MinionData struct {
	MinionsKilled             int64 `json:"minionsKilled,omitempty"`
	LaneMinionsKilled         int64 `json:"laneMinionsKilled,omitempty"`
	NeutralMinionsKilled      int64 `json:"neutralMinionsKilled,omitempty"`
	NeutralMinionsTeamJungle  int64 `json:"neutralMinionsTeamJungle,omitempty"`
	NeutralMinionsEnemyJungle int64 `json:"neutralMinionsEnemyJungle,omitempty"`
}

type CombatData struct {
	Kills               int64 `json:"kills,omitempty"`
	Deaths              int64 `json:"deaths,omitempty"`
	Assists             int64 `json:"assists,omitempty"`
	LargestKillingSpree int64 `json:"largestKillingSpree,omitempty"`
	LargestMultiKill    int64 `json:"largestMultiKill,omitempty"`
}

type DamageHealData struct {
	TotalDamageDealt              int64 `json:"totalDamageDealt,omitempty"`
	PhysicalDamageDealt           int64 `json:"physicalDamageDealt,omitempty"`
	MagicalDamageDealt            int64 `json:"magicalDamageDealt,omitempty"`
	TrueDamageDealt               int64 `json:"trueDamageDealt,omitempty"`
	LargestCriticalStrike         int64 `json:"largestCriticalStrike,omitempty"`
	TotalDamageDealtToHeroes      int64 `json:"totalDamageDealtToHeroes,omitempty"`
	PhysicalDamageDealtToHeroes   int64 `json:"physicalDamageDealtToHeroes,omitempty"`
	MagicalDamageDealtToHeroes    int64 `json:"magicalDamageDealtToHeroes,omitempty"`
	TrueDamageDealtToHeroes       int64 `json:"trueDamageDealtToHeroes,omitempty"`
	TotalDamageDealtToStructures  int64 `json:"totalDamageDealtToStructures,omitempty"`
	TotalDamageDealtToObjectives  int64 `json:"totalDamageDealtToObjectives,omitempty"`
	TotalDamageTaken              int64 `json:"totalDamageTaken,omitempty"`
	PhysicalDamageTaken           int64 `json:"physicalDamageTaken,omitempty"`
	MagicalDamageTaken            int64 `json:"magicalDamageTaken,omitempty"`
	TrueDamageTaken               int64 `json:"trueDamageTaken,omitempty"`
	TotalDamageTakenFromHeroes    int64 `json:"totalDamageTakenFromHeroes,omitempty"`
	PhysicalDamageTakenFromHeroes int64 `json:"physicalDamageTakenFromHeroes,omitempty"`
	MagicalDamageTakenFromHeroes  int64 `json:"magicalDamageTakenFromHeroes,omitempty"`
	TrueDamageTakenFromHeroes     int64 `json:"trueDamageTakenFromHeroes,omitempty"`
	TotalDamageMitigated          int64 `json:"totalDamageMitigated,omitempty"`
	TotalHealingDone              int64 `json:"totalHealingDone,omitempty"`
	ItemHealingDone               int64 `json:"itemHealingDone,omitempty"`
	CrestHealingDone              int64 `json:"crestHealingDone,omitempty"`
	UtilityHealingDone            int64 `json:"utilityHealingDone,omitempty"`
	TotalShieldingReceived        int64 `json:"totalShieldingReceived,omitempty"`
}

type WardsData struct {
	WardsPlaced      int64       `json:"wardsPlaced,omitempty"`
	WardsDestroyed   int64       `json:"wardsDestroyed,omitempty"`
	WardDestructions []*WardData `json:"wardDestructions,omitempty"`
	WardPlacements   []*WardData `json:"wardPlacements,omitempty"`
}
type WardData struct {
	TypeID   int64     `json:"typeID,omitempty"`
	GameTime int64     `json:"gameTime,omitempty"`
	Location *Location `json:"location,omitempty"`
}

type Location struct {
	X float64 `json:"x,omitempty"`
	Y float64 `json:"y,omitempty"`
	Z float64 `json:"z,omitempty"`
}

type IncomeData struct {
	GoldEarned           int64          `json:"goldEarned,omitempty"`
	GoldSpent            int64          `json:"goldSpent,omitempty"`
	GoldEarnedAtInterval []int64        `json:"goldEarnedAtInterval,omitempty"`
	Transactions         []*Transaction `json:"transactions,omitempty"`
}

type Transaction struct {
	ItemID          int64 `json:"itemID,omitempty"`
	TransactionType int64 `json:"transactionType,omitempty"`
	GameTime        int64 `json:"gameTime,omitempty"`
}

type AbilityData struct {
	AbilityInputTag string `json:"abilityInputTag,omitempty"`
	GameTime        int64  `json:"gameTime,omitempty"`
}

type InventoryData struct {
	ItemSlot int64 `json:"itemSlot,omitempty"`
	ItemID   int64 `json:"itemID,omitempty"`
}

type HeroKill struct {
	KilledPlayerID   string    `json:"killedPlayerID,omitempty"`
	KilledHeroName   string    `json:"killedHeroName,omitempty"`
	KillerPlayerID   string    `json:"killerPlayerID,omitempty"`
	KillerHeroName   string    `json:"killerHeroName,omitempty"`
	KillerEntityType string    `json:"killerEntityType,omitempty"`
	IsFirstBlood     bool      `json:"isFirstBlood,omitempty"`
	Location         *Location `json:"location,omitempty"`
	GameTime         int64     `json:"gameTime,omitempty"`
}

type StructureDestruction struct {
	DestructionPlayerID string    `json:"destructionPlayerID,omitempty"`
	DestructionHeroName string    `json:"destructionHeroName,omitempty"`
	StructureEntityType string    `json:"structureEntityType,omitempty"`
	Location            *Location `json:"location,omitempty"`
	TeamID              int64     `json:"teamID,omitempty"`
	GameTime            int64     `json:"gameTime,omitempty"`
}

type ObjectiveKill struct {
	KilledEntityType string    `json:"killedEntityType,omitempty"`
	KillerPlayerID   string    `json:"killerPlayerID,omitempty"`
	KillerHeroName   string    `json:"killerHeroName,omitempty"`
	Location         *Location `json:"location,omitempty"`
	GameTime         int64     `json:"gameTime,omitempty"`
}
