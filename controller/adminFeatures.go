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
func AuthorizedAdmin(c echo.Context) bool {
	_, role := auth.ExtractTokenUserId(c)
	//adminList, err := database.GetAdminid(loggedInAdminId)

	if role != "admin" {
		return false
	}
	return true
}

//ADMIN FEATURES: EDIT QUESTION
func EditQuestion(c echo.Context) error {
	//---------------------------
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	//-------------------------
	soalId, err2 := strconv.Atoi(c.Param("soalId"))
	if err2 != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//---------------------------------------------------------
	oneProblem, err := database.GetOneQuestionById(soalId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "cannot find soal based on id")
	}
	c.Bind(&oneProblem)
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
	auth := AuthorizedAdmin(c)
	if auth == false {
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
	c.Bind(&oneProblem)
	_, err1 := database.EditSoal(oneProblem)
	if err1 != nil {
		return c.JSON(http.StatusBadRequest, "Cannot Edit Status Approval")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Updates Approval",
	})
}

//ADMIN FEATURES: GET ALL QUESTION BASED ON CATEGORY
func GetQuestionByCategory(c echo.Context) error {
	//---------------------------------------------------------
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	//---------------------------------------------------------
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
	auth := AuthorizedAdmin(c)
	if auth == false {
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
	mapQuestion := map[string]interface{}{
		"ID":              soal.ID,
		"Soal_pertanyaan": soal.Soal_pertanyaan,
		"PilihanA":        soal.PilihanA,
		"PilihanB":        soal.PilihanB,
		"PilihanC":        soal.PilihanC,
		"PilihanD":        soal.PilihanD,
		"Jawaban":         soal.Jawaban,
		"KesulitanID":     soal.KesulitanID,
		"Solusi":          soal.Solusi,
		"Approval":        soal.Approval,
		"CategoryID":      soal.CategoryID,
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    mapQuestion,
	})
}

//ADMIN FEATURES: DELETE QUESTION BASED ON ID
func DeleteQuestion(c echo.Context) error {
	//---------------------------------------------------------

	//---------------------------------------------------------
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	soalId, err := strconv.Atoi(c.Param("soalId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	//------------------------
	soalToDelete, err := database.DeleteOneSoalSpecifiedId(soalId)
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
	//---------------------------------------------------------
	auth := AuthorizedAdmin(c)
	if auth == false {
		return echo.NewHTTPError(http.StatusUnauthorized, "Cannot access this account")
	}
	categoryId, err := strconv.Atoi(c.Param("kategori_materi_pelajaran_id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot use the id",
		})
	}
	//---------------------------------------------------------

	var soal []models.Soal
	if err := config.DB.Where(map[string]interface{}{"category_id": categoryId, "approval": "not yet"}).Find(&soal).Error; err != nil {
		return c.JSON(http.StatusBadRequest, "Cannot find the problem that needs approval")
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
		"Data":    soal,
	})
}
