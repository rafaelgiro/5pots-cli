package champion

type DownloadedData struct {
	Character map[string]interface{} `json:"-"`
	Tooltips  map[string]interface{} `json:"tooltips"`
}

type JSONData struct {
	Character map[string]SpellObject `mapstructure:"character"`
	Tooltips  map[string]string      `mapstructure:"tooltips"`
}

type SpellObject struct {
	ScriptName string            `mapstructure:"mScriptName"`
	Spell      SpellDataResource `mapstructure:"mSpell"`
}

type SpellDataResource struct {
	// Type                                string                     `json:"__type"`
	// AlwaysSnapFacing                    bool                       `json:"alwaysSnapFacing"`
	CastRange    []float64 `mapstructure:"castRange"`
	CooldownTime []float64 `mapstructure:"cooldownTime"`
	// DelayCastOffsetPercent              float64                    `json:"delayCastOffsetPercent"`
	// DelayTotalTimePercent               float64                    `json:"delayTotalTimePercent"`
	// LuaOnMissileUpdateDistanceInterval  float64                    `json:"luaOnMissileUpdateDistanceInterval"`
	// AffectsTypeFlags                    int                        `json:"mAffectsTypeFlags"`
	// AnimationName                       string                     `json:"mAnimationName"`
	// CantCancelWhileWindingUp            bool                       `json:"mCantCancelWhileWindingUp"`
	CastTime float64 `mapstructure:"mCastTime"`
	// CastType                            int                        `json:"mCastType"`
	// ClientData                          SpellDataResourceClient    `json:"mClientData"`
	DataValues []SpellDataValue `mapstructure:"mDataValues"`
	// ImgIconName                         []string                   `json:"mImgIconName"`
	// MissileSpec                         MissileSpecification       `json:"mMissileSpec"`
	// SpellCalculations                   map[string]GameCalculation `json:"mSpellCalculations"`
	// SpellCooldownOrSealedQueueThreshold float64                    `json:"mSpellCooldownOrSealedQueueThreshold"`
	// TargetingTypeData                   TargetingTypeData          `json:"mTargetingTypeData"`
	Mana []float64 `mapstructure:"mana"`
	// UseAnimatorFramerate                bool                       `json:"useAnimatorFramerate"`
	// UnknownField                        bool                       `json:"{00f7e5bc}"`
}

type SpellDataValue struct {
	Type   string        `mapstructure:"__type"`
	Name   string        `mapstructure:"mName"`
	Values []interface{} `mapstructure:"mValues"`
}

type SpellDataResourceClient struct {
	Type                string               `json:"__type"`
	TargeterDefinitions []TargeterDefinition `json:"mTargeterDefinitions"`
	TooltipData         TooltipInstanceSpell `json:"mTooltipData"`
}

type TooltipInstanceSpell struct {
	Type       string              `json:"__type"`
	Format     string              `json:"mFormat"`
	Lists      TooltipInstanceList `json:"mLists"`
	LocKeys    TooltipLocKeys      `json:"mLocKeys"`
	ObjectName string              `json:"mObjectName"`
}

type TooltipInstanceList struct {
	Type       string                       `json:"__type"`
	Elements   []TooltipInstanceListElement `json:"elements"`
	LevelCount int                          `json:"levelCount"`
}

type TooltipInstanceListElement struct {
	NameOverride string `json:"nameOverride"`
	Type         string `json:"type"`
}

type TooltipLocKeys struct {
	KeyName                     string `json:"keyName"`
	KeySummary                  string `json:"keySummary"`
	KeyTooltip                  string `json:"keyTooltip"`
	KeyTooltipExtendedBelowLine string `json:"keyTooltipExtendedBelowLine"`
}

type TargeterDefinition struct {
	Type                   string                   `json:"__type"`
	EndLocator             *DrawablePositionLocator `json:"endLocator,omitempty"`
	LineStopsAtEndPosition bool                     `json:"lineStopsAtEndPosition,omitempty"`
	LineWidth              *FloatPerSpellLevel      `json:"lineWidth,omitempty"`
	OverrideBaseRange      *FloatPerSpellLevel      `json:"overrideBaseRange,omitempty"`
	HideWithLineIndicator  bool                     `json:"hideWithLineIndicator,omitempty"`
}

type DrawablePositionLocator struct {
	Type         string `json:"__type"`
	BasePosition int    `json:"basePosition"`
}

type FloatPerSpellLevel struct {
	Type           string    `json:"__type"`
	PerLevelValues []float64 `json:"mPerLevelValues"`
	ValueType      int       `json:"mValueType"`
}

type MissileSpecification struct {
	Type              string             `json:"__type"`
	Behaviors         []MissileBehavior  `json:"behaviors"`
	MissileWidth      float64            `json:"mMissileWidth"`
	MovementComponent FixedSpeedMovement `json:"movementComponent"`
	VerticalFacing    VerticalFacing     `json:"verticalFacing"`
}

type MissileBehavior struct {
	Type string `json:"__type"`
}

type FixedSpeedMovement struct {
	Type                     string  `json:"__type"`
	ProjectTargetToCastRange bool    `json:"mProjectTargetToCastRange"`
	Speed                    float64 `json:"mSpeed"`
	StartBoneName            string  `json:"mStartBoneName"`
	TargetBoneName           string  `json:"mTargetBoneName"`
	TracksTarget             bool    `json:"mTracksTarget"`
	UseHeightOffsetAtEnd     bool    `json:"mUseHeightOffsetAtEnd"`
}

type VerticalFacing struct {
	Type string `json:"__type"`
}

type GameCalculation struct {
	Type         string        `json:"__type"`
	FormulaParts []FormulaPart `json:"mFormulaParts"`
}

type FormulaPart struct {
	Type        string   `json:"__type"`
	DataValue   *string  `json:"mDataValue,omitempty"`
	Coefficient *float64 `json:"mCoefficient,omitempty"`
	EndValue    *float64 `json:"mEndValue,omitempty"`
	StartValue  *float64 `json:"mStartValue,omitempty"`
}

type TargetingTypeData struct {
	Type string `json:"__type"`
}