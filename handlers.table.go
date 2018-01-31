package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type TablesHandler struct {
	db                  *gorm.DB
	tableRepository     TableRepository
	clubunionRepository ClubUnionRepository
	responseHandler     ResponseHandler
	requestHandler      RequestHandler
}

func InitTablesHandler(app *App) *TablesHandler {
	h := &TablesHandler{
		app.Db(),
		NewTableRepository(app.Db()),
		NewClubUnionRepository(app.Db()),
		app.ResponseHandler(),
		app.requestHandler,
	}

	authMiddleware := NewAuthMiddleware(app)

	v1 := app.engine.Group("/v1")
	{
		v1.Use().GET("/export/:clubid", h.Export)
		v1.Use(authMiddleware).GET("/tables/club/:clubid", h.List)
		v1.Use(authMiddleware).GET("/tables/detail/:tid", h.Get)

	}

	return h
}
func (h *TablesHandler) List(c *gin.Context) {

	startTime, _ := time.Parse("2006-01-02 15:04:05", c.DefaultQuery("startTime", "2016-01-01 00:00:00"))
	endTime, _ := time.Parse("2006-01-02 15:04:05", c.DefaultQuery("endTime", time.Now().Format("2006-01-02 15:04:05")))
	wheres := fmt.Sprintf("`table`.end_time BETWEEN %d and %d", startTime.Unix(), endTime.Unix())
	if roundName, ok := c.GetQuery("roundName"); ok {
		createByid, _ := strconv.Atoi(roundName)
		wheres = wheres + fmt.Sprintf(" and (`table`.`name` like '%s' or `table`.create_user_id = %d)", "%"+roundName+"%", createByid)
	}

	if gameMod, ok := c.GetQuery("gameMod"); ok {
		wheres = wheres + " and `table`.game_mod IN (" + gameMod + ")"
	}

	id, _ := strconv.Atoi(c.Param("clubid"))
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit := 20
	offset := (page * limit) - limit
	tables, err := h.tableRepository.FindByClubId(id, wheres, limit, offset)
	fmt.Println(tables)
	if err != nil {
		h.responseHandler.InternalServerError(c)
		return
	}
	h.responseHandler.JSON(c, http.StatusOK, tables)
}

func (h *TablesHandler) Get(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("tid"))
	playerstable, err := h.tableRepository.FindById(id)
	if err != nil {
		h.responseHandler.NotFound(c)
		return
	}

	h.responseHandler.JSON(c, http.StatusOK, playerstable)
}

func (h *TablesHandler) Export(c *gin.Context) {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var err error
	var wheres string
	file, err = xlsx.OpenFile("template.xlsx")

	if err != nil {
		log.Fatal(err)
	}
	sheet = file.Sheets[0]
	if tid, ok := c.GetQuery("tid"); ok {
		tableid, _ := strconv.Atoi(tid)
		wheres = fmt.Sprintf("`table`.id = %d", tableid)
	} else {
		wheres = "1 = 1"
	}
	clubid, _ := strconv.Atoi(c.Param("clubid"))
	tables, err := h.tableRepository.FindByClubId(clubid, wheres, 10000, 0)
	if err != nil {
		h.responseHandler.NotFound(c)
		return
	}
	for _, table := range tables {
		players, err := h.tableRepository.FindById(table.ID)
		var tableexport TableExport
		if err == nil {
			for _, player := range players {
				tableexport.GameMod = GameMod[table.GameMod]
				tableexport.Name = table.Name
				tableexport.CreateName = table.CreatedBy.Username
				tableexport.BigBlind = fmt.Sprintf("%d/%d", table.BigBlind/2, table.BigBlind)
				tableexport.LimitPlayers = table.LimitPlayers
				tableexport.ExistTime = fmt.Sprintf("%.2fh", float64(table.ExistTime)/3600)
				tableexport.HandCounts = table.HandCounts
				tableexport.Uno = player.User.Uno
				tableexport.Username = player.User.Username
				tableexport.ClubName = player.Club.Name
				tableexport.ClubNo = player.Club.ClubNumber
				tableexport.Spends = player.Spends
				tableexport.GetBack = player.GetBack
				tableexport.InsurancePay = player.InsurancePay
				tableexport.InsuranceBet = player.InsuranceBet
				tableexport.InsuranceNum = player.InsuranceNum
				tableexport.InsurancePool = table.InsurancePool
				tableexport.Win = player.GetBack - player.Spends
				tableexport.EndTime = time.Unix(table.EndTime, 0).Format("2006-01-02 15:04:05")
				row = sheet.AddRow()
				row.WriteStruct(&tableexport, -1)
			}
		}
	}
	filename := time.Now().Format("2006-01-02 15:04:05") + ".xlsx"
	h.responseHandler.downloadFile(c, filename)
	err = file.Write(c.Writer)
	if err != nil {
		fmt.Printf(err.Error())
	}
}
