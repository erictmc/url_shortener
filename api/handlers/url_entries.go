package handlers

import (
	"encoding/json"
	"errors"
	env "github.com/erictmc/url_shortener/api/environment"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/labstack/echo"
	"gorm.io/gorm"
	"net/http"
	"regexp"
)

type UrlForm struct {
	OriginalUrl string `json:"original_url"`
}

func (u UrlForm) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.OriginalUrl,
			validation.Required,
			is.URL,
			validation.Match(regexp.MustCompile("^https://.+")).
				Error("https:// is required for urls"),
		),
	)
}

// CreateShortUrl takes an url and creates a "short url" version of it.
func CreateShortUrl(appEnv env.AppEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		form := new(UrlForm)

		if err := json.NewDecoder(c.Request().Body).Decode(&form); err != nil {
			appEnv.Logger.Println(err)
			return err
		}

		if err := form.Validate(); err != nil {
			return c.JSON(http.StatusUnprocessableEntity, err)
		}
		entry, err := appEnv.Db.CreateUrlEntry(form.OriginalUrl)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, nil)
		}

		entry.ShortUrl = appEnv.BuildShortUrl(entry.ShortUrl)

		return c.JSON(http.StatusCreated, entry)
	}
}

// RouteToShortUrl takes a short url and routes to the original url.
func RouteToShortUrl(appEnv env.AppEnvironment) echo.HandlerFunc {
	return func(c echo.Context) error {
		shortUrl := c.Param("short_url")
		urlEntry, err := appEnv.Db.FetchUrlEntry(shortUrl)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.JSON(http.StatusNotFound, nil)
			}
			return c.JSON(http.StatusInternalServerError, nil)
		}
		return c.Redirect(http.StatusFound, urlEntry.OriginalUrl)
	}
}