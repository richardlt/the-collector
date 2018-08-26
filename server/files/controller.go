package files

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/types"
)

// HandleGet .
func HandleGet(c echo.Context) error {
	fileToken := c.Param("fileToken")
	filename := c.Param("filename")
	if fileToken == "" {
		return errors.NewNotFound()
	}

	cs, valid, err := api.ParseJWT(config.jwtSecret, fileToken)
	if err != nil || !valid {
		return errors.NewNotFound()
	}
	if t, ok := cs["type"].(string); !ok || t != "file" {
		return errors.NewNotFound()
	}
	uuid, ok := cs["uuid"].(string)
	if !ok || !api.IsUUID(uuid) {
		return errors.NewNotFound()
	}

	f, err := Get(context.Background(), NewCriteria().
		Name(filename).UUID(uuid))
	if err != nil {
		return err
	}
	if f == nil {
		return errors.NewNotFound()
	}

	size := c.QueryParam("size")
	if !types.IsValidSize(size) {
		return errors.NewData("invalid given image size")
	}

	c.Response().Header().Set("Content-Type", string(f.ContentType))
	c.Response().Header().Set("Content-Disposition", "inline;filename=\""+f.Name+"\"")
	c.Response().Header().Set("X-Frame-Options", "DENY")
	c.Response().Header().Set("Content-Security-Policy", "Frame-ancestors 'none'")

	data, err := ReadImage(f.Path, size)
	if err != nil {
		return err
	}
	if data == nil {
		if size == types.Original {
			return errors.NewNotFound()
		}
		o, err := ReadImage(f.Path, types.Original)
		if err != nil {
			return err
		}
		if o == nil {
			return errors.NewNotFound()
		}
		data, err = ResizeImageFile(f, o, size)
		if err != nil {
			return err
		}
		if err := SaveFile(data, fmt.Sprintf("%s.%s", f.Path, size)); err != nil {
			return err
		}
	}

	return c.Blob(http.StatusOK, string(f.ContentType), data)
}
