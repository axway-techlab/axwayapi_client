package axwayapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"reflect"
	"strings"
)

func assertIsPointer(object interface{}) error {
	rv := reflect.ValueOf(object)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return fmt.Errorf("should provide a pointer as 'object', got a %s", rv.Kind())
	}
	return nil
}

func (c *Client) post(object interface{}, url string, expected ...int) error {
	if err := assertIsPointer(object); err != nil {
		return err
	}
	rb, err := json.Marshal(object)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.HostURL+"/"+url, strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, expected...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) get(object interface{}, url string, expected ...int) error {
	if err := assertIsPointer(object); err != nil {
		return err
	}

	req, err := http.NewRequest("GET", c.HostURL+"/"+url, nil)
	if err != nil {
		return err
	}

	body, err := c.doRequest(req, expected...)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) delete(url string, expected ...int) error {
	req, err := http.NewRequest("DELETE", c.HostURL+"/"+url, nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, expected...)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) put(object interface{}, url string, expected ...int) error {
	err := assertIsPointer(object)
	if err != nil {
		return err
	}
	rb, err := json.Marshal(object)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PUT", c.HostURL+"/"+url, strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, object)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) sendFiles(url string, files ...filePart) ([]byte, error) {
	enc, err := newMultiPart().addFiles(files...).writeMultiPart()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.HostURL+"/"+url, enc.body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", enc.formDataContentType)
	return c.doRequest(req)
}

func (c *Client) sendParts(url string, mpart multiPart) ([]byte, error) {
	enc, err := mpart.writeMultiPart()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", c.HostURL+"/"+url, enc.body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", enc.formDataContentType)
	return c.doRequest(req)
}

type WithId interface {
	GetId() string
}

func (c *Client) UpdateImageFor(t WithId, image string) error {
	var path string
	switch t.(type) {
	case *Frontend:
		path = "proxies"
	case *Application:
		path = "applications"
	case *User:
		path = "users"
	case *Org:
		path = "organizations"
	}
	return c.updateImage(fmt.Sprintf("%s/%s/image/", path, t.GetId()), image)
}

func (c *Client) updateImage(url, imageB64 string) error {
	img, err := base64.StdEncoding.DecodeString(imageB64)
	if err != nil {
		return err
	}
	_, err = c.sendFiles(url,
		filePart{
			data:        img,
			contentType: "image/jpeg",
			fileName:    "image.jpeg",
		})
	if err != nil {
		return err
	}
	return nil
}

// multipart specifics

type multiPart struct {
	files  []filePart
	fields map[string]string
}

func newMultiPart() multiPart {
	return multiPart{}
}
func (mpart multiPart) addField(name string, value string) multiPart {
	if nil == mpart.fields {
		mpart.fields = map[string]string{}
	}
	mpart.fields[name] = value
	return mpart
}
func (mpart multiPart) addFile(file filePart) multiPart {
	mpart.files = append(mpart.files, file)
	return mpart
}
func (mpart multiPart) addFiles(file ...filePart) multiPart {
	mpart.files = append(mpart.files, file...)
	return mpart
}

type multiPartEncoded struct {
	formDataContentType string
	body                *bytes.Buffer
}

func (mpart multiPart) writeMultiPart() (enc *multiPartEncoded, err error) {
	body := &bytes.Buffer{}
	enc = &multiPartEncoded{}
	writer := multipart.NewWriter(body)
	for _, f := range mpart.files {
		err = f.writePart(writer)
		if err != nil {
			return nil, err
		}
	}
	for n, v := range mpart.fields {
		err = writer.WriteField(n, v)
		if err != nil {
			return nil, err
		}
	}
	err = writer.Close()
	if err != nil {
		return nil, err
	}
	enc.body = body
	enc.formDataContentType = writer.FormDataContentType()
	return enc, nil
}

type filePart struct {
	data        []byte
	contentType string
	fileName    string
}

func (part *filePart) writePart(writer *multipart.Writer) error {
	partHeader := textproto.MIMEHeader{}
	partHeader.Add("Content-Disposition", "form-data; name=\"file\"; filename=\""+part.fileName+"\"")
	partHeader.Add("Content-Type", part.contentType)
	p, err := writer.CreatePart(partHeader)
	if err != nil {
		return err
	}
	_, err = p.Write(part.data)
	if err != nil {
		return err
	}
	return nil
}
