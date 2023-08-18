package controllers

import (
	"errors"
	"net/http"
	"strconv"
	"time"
	"unjuk_keterampilan/config"
	"unjuk_keterampilan/models"
	"unjuk_keterampilan/repository"

	"github.com/labstack/echo/v4"
)

type CtrlProduct struct {
}

func AddProduct(c echo.Context) error {
	productReq := models.Product{}
	err := c.Bind(&productReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		return err
	}
	//validasi req
	err = config.Validate.Struct(productReq)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	TimeNow := time.Now()
	productReq.CreatedAt = &TimeNow
	res, err := repository.RepoAddProduct(productReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Responses{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Responses{
		Code:    http.StatusCreated,
		Message: "Success Created Product",
		Data:    res,
	})
}

func ShowAllProduct(c echo.Context) error {
	res, err := repository.RepoShowAllProduct()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Responses{
			Code:    http.StatusCreated,
			Message: err.Error(),
			Data:    res,
		})
	}

	return c.JSON(http.StatusAccepted, models.Responses{
		Code:    http.StatusAccepted,
		Message: "Success Get All data product",
		Data:    res,
	})
}

func FindById(c echo.Context) error {
	id := c.Param("id")
	idres, err := parseStringToInt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	res, err := repository.RepoFindById(idres)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, models.Responses{
		Code:    http.StatusAccepted,
		Message: "Success get data product by id",
		Data:    res,
	})
}

func Updateproduct(c echo.Context) error {
	reqproduct := models.Product{}
	err := c.Bind(&reqproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	//validasi req
	err = config.Validate.Struct(reqproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	//get param id and convert id req to uint
	id := c.Param("id")
	idValue, err := parseStringToInt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	// cek id is  found
	_, err = repository.RepoFindById(idValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	reqproduct.Id = idValue
	timeNow := time.Now()
	reqproduct.UpdatedAt = &timeNow
	res, err := repository.RepoUpdateById(reqproduct)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.Responses{
		Code:    http.StatusOK,
		Message: "success update product by id",
		Data:    res,
	})

}

func DeleteProduct(c echo.Context) error {
	// get param id and convert
	id := c.Param("id")
	idValue, err := parseStringToInt(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	// cek id is found?
	_, err = repository.RepoFindById(idValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	err = repository.RepoSoftDeleteByid(idValue)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Responses{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Responses{
		Code:    http.StatusOK,
		Message: "success delete product by id",
	})

}
func parseStringToInt(id string) (idRes uint64, err error) {

	idRes, err = strconv.ParseUint(id, 10, 64)
	if err != nil {
		err = errors.New("fail to convert id")

		return idRes, err
	}
	return

}
