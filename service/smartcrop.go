package service

import (
	"errors"
	"fmt"
	"image"
	"net/http"
	"os"
	"path"

	"github.com/go-chi/chi"
	"github.com/muesli/smartcrop"
	"github.com/muesli/smartcrop/nfnt"
	"github.com/titpetric/factory/resputil"
)

type cropSize struct {
	x, y int
}

func smartCrop(filename string) (map[string]image.Rectangle, error) {
	fullPath := path.Join(config.sourcePath, filename)

	fi, err := os.Open(fullPath)
	if err != nil {
		// don't leak local file path information
		if os.IsNotExist(err) {
			return nil, errors.New("The requested file doesn't exist")
		}
		return nil, err
	}

	img, _, err := image.Decode(fi)
	if err != nil {
		return nil, err
	}

	cropSizes := []cropSize{
		cropSize{1, 1},
		cropSize{4, 3},
		cropSize{3, 4},
		cropSize{16, 9},
	}

	results := map[string]image.Rectangle{}
	analyzer := smartcrop.NewAnalyzer(nfnt.NewDefaultResizer())
	for _, size := range cropSizes {
		crop, err := analyzer.FindBestCrop(img, size.x, size.y)
		if err != nil {
			return nil, err
		}
		results[fmt.Sprintf("%d:%d", size.x, size.y)] = crop
	}
	return results, nil
}

func smartCropHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	filename := chi.URLParam(r, "*")
	resputil.JSON(w, func() (interface{}, error) {
		return smartCrop(filename)
	})
}
