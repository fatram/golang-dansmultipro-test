package position

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	"github.com/labstack/echo/v4"
)

type PositionService struct {
	logger genlog.Logger
}

func (s *PositionService) Get(ctx context.Context, identifier interface{}) (interface{}, error) {
	id, _ := identifier.(string)
	resp, err := http.Get(fmt.Sprintf("http://dev3.dansmultipro.co.id/api/recruitment/positions/%s", id))
	if err != nil {
		s.logger.Errorf("error when get position: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unable to get position data").SetInternal(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Errorf("error when get position: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unable to get position data").SetInternal(err)
	}

	var position model.Position
	if err = json.Unmarshal(body, &position); err != nil {
		s.logger.Errorf("error when get position: %s", err.Error())
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "unable to get position data").SetInternal(err)
	}

	if position.ID == "" {
		return nil, echo.NewHTTPError(http.StatusNotFound, "position not found")
	}
	return position, nil
}

func (s *PositionService) GetAll(ctx context.Context, filter model.PositionFilter) ([]model.Position, int, error) {
	resp, err := http.Get(fmt.Sprintf("http://dev3.dansmultipro.co.id/api/recruitment/positions.json?description=%s&location=%s", filter.Description, filter.Location))
	if err != nil {
		s.logger.Errorf("error when get multiple positions: %s", err.Error())
		return nil, 0, echo.NewHTTPError(http.StatusInternalServerError, "unable to get multiple position data").SetInternal(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		s.logger.Errorf("error when get multiple positions: %s", err.Error())
		return nil, 0, echo.NewHTTPError(http.StatusInternalServerError, "unable to get multiple position data").SetInternal(err)
	}

	var positions []model.Position
	if err = json.Unmarshal(body, &positions); err != nil {
		s.logger.Errorf("error when get multiple positions: %s", err.Error())
		return nil, 0, echo.NewHTTPError(http.StatusInternalServerError, "unable to get multiple position data").SetInternal(err)
	}

	if filter.FullTime != nil {
		isFullTime := *filter.FullTime
		i := 0
		for _, position := range positions {
			if position.Type == "Full Time" && isFullTime {
				positions[i] = position
				i++
			} else if position.Type != "Full Time" && !isFullTime {
				positions[i] = position
				i++
			}
		}
		positions = positions[:i]
	}

	total := len(positions)

	if filter.PageNumber != nil && filter.PageSize == nil {
		size := 10
		filter.PageSize = &size
	} else if filter.PageNumber == nil && filter.PageSize != nil {
		page := 1
		filter.PageNumber = &page
	}

	if filter.PageNumber != nil && filter.PageSize != nil {
		page := *filter.PageNumber
		limit := *filter.PageSize
		start := (page - 1) * limit
		end := start + limit
		if end > total {
			end = total
		}
		if start > end {
			start = end
		}
		positions = positions[start:end]
	}

	return positions, total, err

}
