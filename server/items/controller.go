package items

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	errorsP "github.com/pkg/errors"
	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/files"
	"github.com/richardlt/the-collector/server/types"
	uuid "github.com/satori/go.uuid"
)

// HandleGetAllForCollection .
func HandleGetAllForCollection(c echo.Context) error {
	co := c.Get("collection").(*types.Collection)

	is, err := GetAll(context.Background(), newCriteria().
		CollectionID(co.ID))
	if err != nil {
		return err
	}

	ids := []objectid.ObjectID{}
	for _, i := range is {
		ids = append(ids, i.ID)
	}

	fs, err := files.GetAll(context.Background(), files.NewCriteria().
		ResourceType(types.ItemResource).ResourceID(ids...))
	if err != nil {
		return err
	}

	m := map[objectid.ObjectID]*types.File{}
	for _, f := range fs {
		m[f.ResourceID] = f
	}

	for _, i := range is {
		if f, ok := m[i.ID]; ok {
			uri, err := f.GenerateURI(config.jwtSecret)
			if err != nil {
				return err
			}
			i.Picture = uri
		}
	}

	return c.JSON(http.StatusOK, is)
}

// HandleGet .
func HandleGet(c echo.Context) error {
	i := c.Get("item").(*types.Item)

	f, err := files.Get(context.Background(), files.NewCriteria().
		ResourceType(types.ItemResource).ResourceID(i.ID))
	if err != nil {
		return err
	}
	if f != nil {
		uri, err := f.GenerateURI(config.jwtSecret)
		if err != nil {
			return err
		}
		i.Picture = uri
	}

	return c.JSON(http.StatusOK, i)
}

// HandlePost .
func HandlePost(c echo.Context) error {
	co := c.Get("collection").(*types.Collection)

	header, err := c.FormFile("file")
	if err != nil {
		return errors.NewData("invalid given file")
	}

	ct := header.Header.Get("Content-Type")
	if !types.IsImageContentType(ct) {
		return errors.NewData("invalid given file")
	}

	item := &types.Item{CollectionID: co.ID}
	if err := Create(context.Background(), item); err != nil {
		return err
	}

	f, err := header.Open()
	if err != nil {
		return errorsP.WithStack(err)
	}

	var data []byte
	if ct == types.ImageJpeg {
		data, err = files.FixJpegImageRotation(f)
	} else {
		data, err = ioutil.ReadAll(f)
	}
	if err != nil {
		return errorsP.WithStack(err)
	}

	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s%s",
		api.Slugify(strings.TrimSuffix(header.Filename, ext)), ext)
	path := filepath.Clean(fmt.Sprintf("%s/%s/%s/%s",
		types.ItemResource, item.UUID, uuid.NewV4().String(), filename))

	if err := files.SaveFile(data, path); err != nil {
		return err
	}

	if err := files.Create(context.Background(), &types.File{
		ResourceType: types.ItemResource,
		ResourceID:   item.ID,
		Name:         filename,
		Path:         path,
		Size:         int64(len(data)),
		ContentType:  ct,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, item)
}

// HandleDelete .
func HandleDelete(c echo.Context) error {
	i := c.Get("item").(*types.Item)

	if err := Delete(context.Background(), i); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
