package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/muhwyndhamhp/gotes-mx/config"
	"github.com/muhwyndhamhp/gotes-mx/db"
	"github.com/muhwyndhamhp/gotes-mx/modules/pet"
	"github.com/muhwyndhamhp/gotes-mx/template"
	adoption "github.com/muhwyndhamhp/gotes-mx/templates/adopotion"
	"github.com/muhwyndhamhp/gotes-mx/templates/index"
	"github.com/muhwyndhamhp/gotes-mx/templates/nextpage"
	"github.com/muhwyndhamhp/gotes-mx/utils/errs"
	"github.com/muhwyndhamhp/gotes-mx/utils/routing"
	"github.com/muhwyndhamhp/gotes-mx/utils/templhelper"
)

func main() {
	e := echo.New()
	routing.SetupRouter(e)

	db := db.GetDB()

	repo := pet.NewPetRepository(db)

	e.Static("/dist", "dist")

	template.NewTemplateRenderer(e)

	e.GET("/", func(c echo.Context) error {
		hello := index.Index("From Params")
		return templhelper.RenderAssert(c, http.StatusOK, "", hello)
	})

	e.GET("/next-page", func(c echo.Context) error {
		nextPage := nextpage.Next()
		return templhelper.RenderAssert(c, http.StatusOK, "", nextPage)
	})

	e.GET("/adoptions", func(c echo.Context) error {
		pets, err := repo.FetchPets(c.Request().Context(), 1, 10, pet.All, "")
		if err != nil {
			return errs.Wrap(err)
		}

		petTypes := []pet.PetTypeSet{
			{
				PetType: pet.All,
				URL:     "/dist/animal.png",
			},
			{
				PetType: pet.Cat,
				URL:     "/dist/cat.png",
			},
			{
				PetType: pet.Dog,
				URL:     "/dist/dog.png",
			},
			{
				PetType: pet.Bird,
				URL:     "/dist/bird.png",
			},
		}

		adoption := adoption.AdoptionList(pets, petTypes, "/adoptions/list?page=2&page_size=10")
		return templhelper.RenderAssert(c, http.StatusOK, "", adoption)
	})

	e.GET("/adoptions/list", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))
		keyword := c.QueryParam("search")
		petType := pet.PetType(c.QueryParam("pet_type"))

		if page == 0 {
			page = 1
		}
		if pageSize == 0 {
			pageSize = 10
		}

		pets, err := repo.FetchPets(c.Request().Context(), page, pageSize, petType, keyword)
		if err != nil {
			return errs.Wrap(err)
		}

		adoption := adoption.List(pets, fmt.Sprintf("/adoptions/list?page=%d&page_size=%d&search=%s", page+1, pageSize, keyword))
		return templhelper.RenderAssert(c, http.StatusOK, "", adoption)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}
