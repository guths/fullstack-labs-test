package utilstests

import (
	"battle-of-monsters/app/router"
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
)

var errReq = errors.New("error creating multipart request")

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	r := router.Router()
	nr := httptest.NewRecorder()

	r.ServeHTTP(nr, req)

	return nr
}

func CreateMultiPartRequestByFilename(filename string) (*http.Request, error) {
	body := new(bytes.Buffer)

	multipartWriter := multipart.NewWriter(body)

	fileHeader := make(textproto.MIMEHeader)

	cd := fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", "monsters-empty-monster.csv")

	fileHeader.Set("Content-Disposition", cd)
	fileHeader.Set("Content-Type", "text/plain")
	writer, _ := multipartWriter.CreatePart(fileHeader)

	file, err := os.Open(fmt.Sprintf("data/%s", filename))

	if err != nil {
		return nil, errReq
	}

	_, err = io.Copy(writer, file)

	if err != nil {
		return nil, errReq
	}

	multipartWriter.Close()

	req, _ := http.NewRequestWithContext(context.TODO(), http.MethodPost, "/monsters/import", body)
	req.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	return req, nil
}
