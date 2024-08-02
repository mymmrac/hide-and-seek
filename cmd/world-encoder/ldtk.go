package main

type LDtkFile struct {
	Header              Header   `json:"__header__"`
	Iid                 string   `json:"iid"`
	JSONVersion         string   `json:"jsonVersion"`
	AppBuildID          int      `json:"appBuildId"`
	NextUID             int      `json:"nextUid"`
	IdentifierStyle     string   `json:"identifierStyle"`
	Toc                 []any    `json:"toc"`
	WorldLayout         any      `json:"worldLayout"`
	WorldGridWidth      any      `json:"worldGridWidth"`
	WorldGridHeight     any      `json:"worldGridHeight"`
	DefaultLevelWidth   any      `json:"defaultLevelWidth"`
	DefaultLevelHeight  any      `json:"defaultLevelHeight"`
	DefaultPivotX       int      `json:"defaultPivotX"`
	DefaultPivotY       int      `json:"defaultPivotY"`
	DefaultGridSize     int      `json:"defaultGridSize"`
	DefaultEntityWidth  int      `json:"defaultEntityWidth"`
	DefaultEntityHeight int      `json:"defaultEntityHeight"`
	BgColor             string   `json:"bgColor"`
	DefaultLevelBgColor string   `json:"defaultLevelBgColor"`
	MinifyJSON          bool     `json:"minifyJson"`
	ExternalLevels      bool     `json:"externalLevels"`
	ExportTiled         bool     `json:"exportTiled"`
	SimplifiedExport    bool     `json:"simplifiedExport"`
	ImageExportMode     string   `json:"imageExportMode"`
	ExportLevelBg       bool     `json:"exportLevelBg"`
	PngFilePattern      any      `json:"pngFilePattern"`
	BackupOnSave        bool     `json:"backupOnSave"`
	BackupLimit         int      `json:"backupLimit"`
	BackupRelPath       any      `json:"backupRelPath"`
	LevelNamePattern    string   `json:"levelNamePattern"`
	TutorialDesc        any      `json:"tutorialDesc"`
	CustomCommands      []any    `json:"customCommands"`
	Flags               []string `json:"flags"`
	Defs                Defs     `json:"defs"`
	Levels              []Levels `json:"levels"`
	Worlds              []Worlds `json:"worlds"`
	DummyWorldIid       string   `json:"dummyWorldIid"`
}

type Header struct {
	FileType   string `json:"fileType"`
	App        string `json:"app"`
	Doc        string `json:"doc"`
	Schema     string `json:"schema"`
	AppAuthor  string `json:"appAuthor"`
	AppVersion string `json:"appVersion"`
	URL        string `json:"url"`
}

type Layers struct {
	Type                           string         `json:"__type"`
	Identifier                     string         `json:"identifier"`
	Type0                          string         `json:"type"`
	UID                            int            `json:"uid"`
	Doc                            any            `json:"doc"`
	UIColor                        any            `json:"uiColor"`
	GridSize                       int            `json:"gridSize"`
	GuideGridWid                   int            `json:"guideGridWid"`
	GuideGridHei                   int            `json:"guideGridHei"`
	DisplayOpacity                 int            `json:"displayOpacity"`
	InactiveOpacity                float64        `json:"inactiveOpacity"`
	HideInList                     bool           `json:"hideInList"`
	HideFieldsWhenInactive         bool           `json:"hideFieldsWhenInactive"`
	CanSelectWhenInactive          bool           `json:"canSelectWhenInactive"`
	RenderInWorldView              bool           `json:"renderInWorldView"`
	PxOffsetX                      int            `json:"pxOffsetX"`
	PxOffsetY                      int            `json:"pxOffsetY"`
	ParallaxFactorX                int            `json:"parallaxFactorX"`
	ParallaxFactorY                int            `json:"parallaxFactorY"`
	ParallaxScaling                bool           `json:"parallaxScaling"`
	RequiredTags                   []any          `json:"requiredTags"`
	ExcludedTags                   []any          `json:"excludedTags"`
	AutoTilesKilledByOtherLayerUID any            `json:"autoTilesKilledByOtherLayerUid"`
	UIFilterTags                   []any          `json:"uiFilterTags"`
	UseAsyncRender                 bool           `json:"useAsyncRender"`
	IntGridValues                  []IntGridValue `json:"intGridValues"`
	IntGridValuesGroups            []any          `json:"intGridValuesGroups"`
	AutoRuleGroups                 []any          `json:"autoRuleGroups"`
	AutoSourceLayerDefUID          any            `json:"autoSourceLayerDefUid"`
	TilesetDefUID                  any            `json:"tilesetDefUid"`
	TilePivotX                     int            `json:"tilePivotX"`
	TilePivotY                     int            `json:"tilePivotY"`
	BiomeFieldUID                  any            `json:"biomeFieldUid"`
}

type IntGridValue struct {
	Value      int    `json:"value"`
	Identifier string `json:"identifier"`
	Color      string `json:"color"`
	Tile       any    `json:"tile"`
	GroupUid   int    `json:"groupUid"`
}

type TileRect struct {
	TilesetUID int `json:"tilesetUid"`
	X          int `json:"x"`
	Y          int `json:"y"`
	W          int `json:"w"`
	H          int `json:"h"`
}

type Entities struct {
	Identifier       string   `json:"identifier"`
	UID              int      `json:"uid"`
	Tags             []any    `json:"tags"`
	ExportToToc      bool     `json:"exportToToc"`
	AllowOutOfBounds bool     `json:"allowOutOfBounds"`
	Doc              any      `json:"doc"`
	Width            int      `json:"width"`
	Height           int      `json:"height"`
	ResizableX       bool     `json:"resizableX"`
	ResizableY       bool     `json:"resizableY"`
	MinWidth         any      `json:"minWidth"`
	MaxWidth         any      `json:"maxWidth"`
	MinHeight        any      `json:"minHeight"`
	MaxHeight        any      `json:"maxHeight"`
	KeepAspectRatio  bool     `json:"keepAspectRatio"`
	TileOpacity      int      `json:"tileOpacity"`
	FillOpacity      float64  `json:"fillOpacity"`
	LineOpacity      int      `json:"lineOpacity"`
	Hollow           bool     `json:"hollow"`
	Color            string   `json:"color"`
	RenderMode       string   `json:"renderMode"`
	ShowName         bool     `json:"showName"`
	TilesetID        int      `json:"tilesetId"`
	TileRenderMode   string   `json:"tileRenderMode"`
	TileRect         TileRect `json:"tileRect"`
	UITileRect       any      `json:"uiTileRect"`
	NineSliceBorders []any    `json:"nineSliceBorders"`
	MaxCount         int      `json:"maxCount"`
	LimitScope       string   `json:"limitScope"`
	LimitBehavior    string   `json:"limitBehavior"`
	PivotX           int      `json:"pivotX"`
	PivotY           int      `json:"pivotY"`
	FieldDefs        []any    `json:"fieldDefs"`
}

type CachedPixelData struct {
	OpaqueTiles   string `json:"opaqueTiles"`
	AverageColors string `json:"averageColors"`
}

type Tilesets struct {
	CWid              int             `json:"__cWid"`
	CHei              int             `json:"__cHei"`
	Identifier        string          `json:"identifier"`
	UID               int             `json:"uid"`
	RelPath           string          `json:"relPath"`
	EmbedAtlas        any             `json:"embedAtlas"`
	PxWid             int             `json:"pxWid"`
	PxHei             int             `json:"pxHei"`
	TileGridSize      int             `json:"tileGridSize"`
	Spacing           int             `json:"spacing"`
	Padding           int             `json:"padding"`
	Tags              []any           `json:"tags"`
	TagsSourceEnumUID any             `json:"tagsSourceEnumUid"`
	EnumTags          []any           `json:"enumTags"`
	CustomData        []any           `json:"customData"`
	SavedSelections   []any           `json:"savedSelections"`
	CachedPixelData   CachedPixelData `json:"cachedPixelData"`
}

type Defs struct {
	Layers        []Layers   `json:"layers"`
	Entities      []Entities `json:"entities"`
	Tilesets      []Tilesets `json:"tilesets"`
	Enums         []any      `json:"enums"`
	ExternalEnums []any      `json:"externalEnums"`
	LevelFields   []any      `json:"levelFields"`
}

type Tile struct {
	TilesetUID int `json:"tilesetUid"`
	X          int `json:"x"`
	Y          int `json:"y"`
	W          int `json:"w"`
	H          int `json:"h"`
}

type EntityInstances struct {
	Identifier     string `json:"__identifier"`
	Grid           []int  `json:"__grid"`
	Pivot          []int  `json:"__pivot"`
	Tags           []any  `json:"__tags"`
	Tile           Tile   `json:"__tile"`
	SmartColor     string `json:"__smartColor"`
	Iid            string `json:"iid"`
	Width          int    `json:"width"`
	Height         int    `json:"height"`
	DefUID         int    `json:"defUid"`
	Px             []int  `json:"px"`
	FieldInstances []any  `json:"fieldInstances"`
	WorldX         int    `json:"__worldX"`
	WorldY         int    `json:"__worldY"`
}

type LayerInstances struct {
	Identifier         string            `json:"__identifier"`
	Type               string            `json:"__type"`
	CWid               int               `json:"__cWid"`
	CHei               int               `json:"__cHei"`
	GridSize           int               `json:"__gridSize"`
	Opacity            int               `json:"__opacity"`
	PxTotalOffsetX     int               `json:"__pxTotalOffsetX"`
	PxTotalOffsetY     int               `json:"__pxTotalOffsetY"`
	TilesetDefUID      int               `json:"__tilesetDefUid"`
	TilesetRelPath     any               `json:"__tilesetRelPath"`
	Iid                string            `json:"iid"`
	LevelID            int               `json:"levelId"`
	LayerDefUID        int               `json:"layerDefUid"`
	PxOffsetX          int               `json:"pxOffsetX"`
	PxOffsetY          int               `json:"pxOffsetY"`
	Visible            bool              `json:"visible"`
	OptionalRules      []any             `json:"optionalRules"`
	IntGridCsv         []int             `json:"intGridCsv"`
	AutoLayerTiles     []AutoLayerTile   `json:"autoLayerTiles"`
	Seed               int               `json:"seed"`
	OverrideTilesetUID any               `json:"overrideTilesetUid"`
	GridTiles          []any             `json:"gridTiles"`
	EntityInstances    []EntityInstances `json:"entityInstances"`
}

type AutoLayerTile struct {
	Px  []int `json:"px"`
	Src []int `json:"src"`
	F   int   `json:"f"`
	T   int   `json:"t"`
	D   []int `json:"d"`
	A   int   `json:"a"`
}

type Neighbours struct {
	LevelIid string `json:"levelIid"`
	Dir      string `json:"dir"`
}

type Levels struct {
	Identifier        string           `json:"identifier"`
	Iid               string           `json:"iid"`
	UID               int              `json:"uid"`
	WorldX            int              `json:"worldX"`
	WorldY            int              `json:"worldY"`
	WorldDepth        int              `json:"worldDepth"`
	PxWid             int              `json:"pxWid"`
	PxHei             int              `json:"pxHei"`
	BgColor           string           `json:"__bgColor"`
	BgColor0          any              `json:"bgColor"`
	UseAutoIdentifier bool             `json:"useAutoIdentifier"`
	BgRelPath         any              `json:"bgRelPath"`
	BgPos             any              `json:"bgPos"`
	BgPivotX          float64          `json:"bgPivotX"`
	BgPivotY          float64          `json:"bgPivotY"`
	SmartColor        string           `json:"__smartColor"`
	BgPos0            any              `json:"__bgPos"`
	ExternalRelPath   any              `json:"externalRelPath"`
	FieldInstances    []any            `json:"fieldInstances"`
	LayerInstances    []LayerInstances `json:"layerInstances"`
	Neighbours        []Neighbours     `json:"__neighbours"`
}

type Worlds struct {
	Iid                string   `json:"iid"`
	Identifier         string   `json:"identifier"`
	DefaultLevelWidth  int      `json:"defaultLevelWidth"`
	DefaultLevelHeight int      `json:"defaultLevelHeight"`
	WorldGridWidth     int      `json:"worldGridWidth"`
	WorldGridHeight    int      `json:"worldGridHeight"`
	WorldLayout        string   `json:"worldLayout"`
	Levels             []Levels `json:"levels"`
}
