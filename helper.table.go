package main

const (
	GAME_MOD_GENERAL               = "组建牌局-普通"
	GAME_MOD_SNG                   = "组建牌局-sng"
	GAME_MOD_HALL_GENERAL_STANDARD = "大厅-普通牌局"
	GAME_MOD_HALL_SNG              = "大厅-坐满即玩"
	GAME_MOD_HALL_GENERAL_HEADSUP  = "大厅-单挑"
	GAME_MOD_21                    = "俱乐部-普通牌局"
	GAME_MOD_22                    = "俱乐部-sng"
	GAME_MOD_23                    = "组建牌局-普通"
	GAME_MOD_43                    = "组建牌局-普通"
	GAME_MOD_41                    = "联盟-普通牌局"
	GAME_MOD_42                    = "联盟-sng"
	GAME_MOD_MTT_GENERAl           = "组建牌局-MTT"
	GAME_MOD_53                    = "大厅-MTT"
	GAME_MOD_63                    = "本地化-MTT"
)

var GameMod = map[string]string{
	"general": GAME_MOD_GENERAL,
	"sng":     GAME_MOD_SNG,
	"hall_general_standard": GAME_MOD_HALL_GENERAL_STANDARD,
	"hall_sng":              GAME_MOD_HALL_SNG,
	"hall_general_headsup":  GAME_MOD_HALL_GENERAL_HEADSUP,
	"21":          GAME_MOD_21,
	"22":          GAME_MOD_22,
	"23":          GAME_MOD_23,
	"41":          GAME_MOD_41,
	"42":          GAME_MOD_42,
	"43":          GAME_MOD_43,
	"mtt_general": GAME_MOD_MTT_GENERAl,
	"52":          GAME_MOD_53,
	"63":          GAME_MOD_63,
}
