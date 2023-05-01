package position

import "github.com/labstack/echo/v4"

// Listposition godoc
// @Summary			List position
// @Description		Menampilkan daftar position
// @Tags         	position
// @Accept       	json
// @Produce      	json
// @Param			page number			query	int		false	"page"
// @Param			page size			query	int		false	"limit"
// @Success      	200  {object}   model.GetAllModel{data=[]model.Position}
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/positions [get]
func (ctr PositionController) GetAll(c echo.Context) error {
	return ctr.GetAllPositionControllerImpl.GetAll(c)
}

// Getposition godoc
// @Summary			Get one position
// @Description		Menampilkan satu position
// @Tags         	position
// @Accept       	json
// @Produce      	json
// @param 	   		id		path 	string 	true 	"id position"
// @Success      	200  {object}   model.Position
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/positions/{id} [get]
func (ctr PositionController) Get(c echo.Context) error {
	return ctr.GetPositionControllerImpl.Get(c)
}