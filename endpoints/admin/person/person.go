package person

import (
	"github.com/Team-Podo/podoting-server/database"
	"github.com/Team-Podo/podoting-server/models"
	"github.com/Team-Podo/podoting-server/repository"
	"github.com/Team-Podo/podoting-server/response/admin/person_find"
	"github.com/Team-Podo/podoting-server/response/admin/person_get"
	"github.com/Team-Podo/podoting-server/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var repositories Repository

type request struct {
	Name  string `json:"name" binding:"required"`
	Birth string `json:"birth"`
}

type Repository struct {
	person models.PersonRepository
}

func init() {
	repositories = Repository{
		person: &repository.PersonRepository{DB: database.Gorm},
	}
}

func Get(c *gin.Context) {
	people, err := repositories.person.FindAll()
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"people": person_get.ParseResponseForm(people),
	})
}

func Find(c *gin.Context) {
	personID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(person) id should be Integer")
		return
	}

	person, err := repositories.person.FindByID(personID)
	if err != nil {

		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, person_find.ParseResponseForm(person))
}

func Create(c *gin.Context) {
	var req request
	err := c.ShouldBindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, "name is required")
		return
	}

	var person repository.Person
	person.Name = req.Name

	if req.Birth != "" {
		err = person.SetBirth(req.Birth)
		if err != nil {
			c.JSON(http.StatusBadRequest, "birth should be in YYYY-MM-DD format")
			return
		}
	}

	err = repositories.person.Create(&person)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, person.ID)
}

func Update(c *gin.Context) {
	personID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(person) id should be Integer")
		return
	}

	var req request
	err = c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, "name is required")
		return
	}

	var person *repository.Person
	person, err = repositories.person.FindByID(personID)
	if err != nil {
		c.JSON(http.StatusNotFound, "person not found")
		return
	}

	person.Name = req.Name

	if req.Birth != "" {
		err = person.SetBirth(req.Birth)
		if err != nil {
			c.JSON(http.StatusBadRequest, "birth should be in YYYY-MM-DD format")
			return
		}
	}

	err = repositories.person.Update(person)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, person.ID)
}

func Delete(c *gin.Context) {
	personID, err := utils.ParseUint(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "(person) id should be Integer")
		return
	}

	err = repositories.person.Delete(personID)
	if err != nil {
		c.JSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, nil)
}
