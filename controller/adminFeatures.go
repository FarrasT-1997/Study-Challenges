package controller

import (
	"SC/auth"
	"SC/config"
	"SC/database"
	"SC/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//AUTHORIZED AND ROLE AS ADMIN
func AuthorizedAdmin(c echo.Context) (bool, models.User) {
	loggedInAdminId, role := auth.ExtractTokenUserId(c)
	adminList, _ := database.GetAdminid(loggedInAdminId)

	if loggedInAdminId != int(adminList.ID) || role != "admin" {
		return false, adminList
	}
	return true, adminList
}

//ADMIN FEATURES: EDIT QUESTION
func EditQuestion(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid problem id",
		})
	}
	oneProblem, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "cannot find soal based on id")
	}
	editedSoal, err := database.EditSoal(oneProblem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot edit soal")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"Data":    editedSoal,
	})
}

//ADMIN FEATURES: UPDATING APPROVAL STATUS (ACCEPT, REJECT, AND NOT YET)
func EditSubmitQuestion(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid problem id",
		})
	}
	oneProblem, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cannot find data",
		})
	}
	approvalStatus, err := database.EditStatusApproval(oneProblem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Cannot Change Status Approval",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"Data":    approvalStatus,
	})
}

//ADMIN FEATURES: GET ALL QUESTION BASED ON CATEGORY
func GetQuestionByCategory(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	categoryId, err := strconv.Atoi(c.Param("MataPelajaranId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid category id",
		})
	}
	soalByCategoryList, err := database.GetAllSoalInSpecifiedCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find problem based on the category id",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    soalByCategoryList,
	})

}

//ADMIN FEATURES: GET QUESTION BASED ON ID
func GetQuestionById(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid problem id",
		})
	}
	soal, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Cannot find the problem",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    soal,
	})
}

//ADMIN FEATURES: DELETE QUESTION BASED ON ID
func DeleteQuestion(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	id, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	soalToDelete, err := database.DeleteOneSoalSpecifiedId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete soal",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    soalToDelete,
	})

}

//ADMIN FEATURES: SHOW ALL PROBLEM THAT HAS NOT BEEN REVIEWED -- BY CATEGORY
func ShowSubmittedQuestion(c echo.Context) error {
	auth, adminList := AuthorizedAdmin(c)
	if auth == false || adminList.Role != "admin" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	categoryId, err := strconv.Atoi("kategori_materi_pelajaran_id")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot use the id",
		})
	}
	soalByCategoryList, err := database.GetAllSoalInSpecifiedCategory(categoryId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot find problem based on the category id",
		})
	}
	if err := config.DB.Find(&soalByCategoryList, "approval=?", "not yet").Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot find the problem that needs approval")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"Data":    soalByCategoryList,
	})
}
