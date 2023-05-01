package position

import "github.com/labstack/echo/v4"

// Listposition godoc
// @Summary			List position
// @Description		Menampilkan daftar position
// @Tags         	Position
// @Accept       	json
// @Produce      	json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @Param			page number			query	int		false	"page"
// @Param			page size			query	int		false	"limit"
// @Param			description			query	int		false	"description"
// @Param			location			query	int		false	"location"
// @Param			full time			query	bool		false	"full_time"
// @Success      	200  {object}   []model.Position
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/positions [get]
func (ctr PositionController) GetAll(c echo.Context) error {
	return ctr.GetAllPositionControllerImpl.GetAll(c)
}

// Getposition godoc
// @Summary			Get one position
// @Description		Menampilkan satu position
// @Tags         	Position
// @Accept       	json
// @Produce      	json
// @Param Authorization header string true "Insert your access token" default(Bearer <Add access token here>)
// @param 	   		id		path 	string 	true 	"id position"
// @Success      	200  {object}   model.Position
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/positions/{id} [get]
func (ctr PositionController) Get(c echo.Context) error {
	return ctr.GetPositionControllerImpl.Get(c)
}
