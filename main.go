package main

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

	tempCtx := context.Background()

	repo := pet.NewPetRepository(db)

	testItem, err := repo.FetchPets(tempCtx, 1, 1)
	if err != nil {
		panic(err)
	}

	now := time.Now()
	if len(testItem) == 0 {
		if err := repo.InsertPet(tempCtx, &pet.Pet{
			Name:     "Aden",
			Nickname: "Taden",
			Images: []pet.Image{
				{
					PetID: 1,
					URL:   "https://www.shutterstock.com/image-photo/beautiful-domestic-cat-resting-light-260nw-1600042501.jpg",
				},
			},
			Birthday: &now,
			Race:     "Scottish Fold",
			PetType:  pet.Cat,
			Address1: "My Home 1",
			Address2: "My Home 2",
		}); err != nil {
			panic(err)
		}

		if err := repo.InsertPet(tempCtx, &pet.Pet{
			Name:     "Loly",
			Nickname: "Lolii",
			Images: []pet.Image{
				{
					PetID: 2,
					URL:   "https://www.shutterstock.com/image-photo/beautiful-domestic-cat-resting-light-260nw-1600042501.jpg",
				},
			},
			Birthday: &now,
			Race:     "Brittish Shorthair",
			PetType:  pet.Cat,
			Address1: "My Home 1",
			Address2: "My Home 2",
		}); err != nil {
			panic(err)
		}
	}

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
		pets, err := repo.FetchPets(c.Request().Context(), 1, 10)
		if err != nil {
			return errs.Wrap(err)
		}

		adoption := adoption.AdoptionList(pets, "/adoptions/list?page=2&page_size=10")
		return templhelper.RenderAssert(c, http.StatusOK, "", adoption)
	})

	e.GET("/adoptions/list", func(c echo.Context) error {
		page, _ := strconv.Atoi(c.QueryParam("page"))
		pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

		pets, err := repo.FetchPets(c.Request().Context(), page, pageSize)
		if err != nil {
			return errs.Wrap(err)
		}
		adoption := adoption.List(pets, "/adoptions/list?page=3&page_size=10")
		return templhelper.RenderAssert(c, http.StatusOK, "", adoption)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.Get(config.APP_PORT))))
}
