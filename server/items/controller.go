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
	uuid "github.com/satori/go.uuid"

	"github.com/richardlt/the-collector/server/api"
	"github.com/richardlt/the-collector/server/api/errors"
	"github.com/richardlt/the-collector/server/files"
	"github.com/richardlt/the-collector/server/types"
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

	file := &types.File{
		ResourceType: types.ItemResource,
		ResourceID:   item.ID,
		Name:         filename,
		Path:         path,
		Size:         int64(len(data)),
		ContentType:  ct,
	}
	if err := files.Create(context.Background(), file); err != nil {
		return err
	}

	uri, err := file.GenerateURI(config.jwtSecret)
	if err != nil {
		return err
	}
	item.Picture = uri

	return c.JSON(http.StatusOK, item)
}

// HandlePostFile .
func HandlePostFile(c echo.Context) error {
	i := c.Get("item").(*types.Item)

	header, err := c.FormFile("file")
	if err != nil {
		return errors.NewData("invalid given file")
	}

	ct := header.Header.Get("Content-Type")
	if !types.IsImageContentType(ct) {
		return errors.NewData("invalid given file")
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
		types.ItemResource, i.UUID, uuid.NewV4().String(), filename))

	if err := files.SaveFile(data, path); err != nil {
		return err
	}

	oldFile, err := files.Get(context.Background(), files.NewCriteria().
		ResourceType(types.ItemResource).ResourceID(i.ID))
	if err != nil {
		return err
	}
	if oldFile != nil {
		if err := files.DeleteFile(oldFile.Path); err != nil {
			return err
		}
		if err := files.Delete(context.Background(), oldFile); err != nil {
			return err
		}
	}

	newFile := &types.File{
		ResourceType: types.ItemResource,
		ResourceID:   i.ID,
		Name:         filename,
		Path:         path,
		Size:         int64(len(data)),
		ContentType:  ct,
	}
	if err := files.Create(context.Background(), newFile); err != nil {
		return err
	}

	uri, err := newFile.GenerateURI(config.jwtSecret)
	if err != nil {
		return err
	}
	i.Picture = uri

	return c.JSON(http.StatusOK, i)
}

// HandleDelete .
func HandleDelete(c echo.Context) error {
	i := c.Get("item").(*types.Item)

	f, err := files.Get(context.Background(), files.NewCriteria().
		ResourceType(types.ItemResource).ResourceID(i.ID))
	if err != nil {
		return err
	}
	if f != nil {
		if err := files.DeleteFile(f.Path); err != nil {
			return err
		}
		if err := files.Delete(context.Background(), f); err != nil {
			return err
		}
	}

	if err := Delete(context.Background(), i); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, nil)
}
