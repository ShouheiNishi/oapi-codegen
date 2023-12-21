package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	"golang.org/x/tools/imports"
	"gopkg.in/yaml.v3"
)

func init() {
	imports.LocalPrefix = "github.com/deepmap/oapi-codegen/v2"
}

func main() {
	fTestGo := &bytes.Buffer{}
	fmt.Fprintf(fTestGo, `// Code generated by generator/generate.go DO NOT EDIT.

package pkg1_test

import (
	"github.com/stretchr/testify/assert"

	"github.com/deepmap/oapi-codegen/v2/internal/test/strict-server-matrix/pkg1"
    "github.com/deepmap/oapi-codegen/v2/internal/test/strict-server-matrix/pkg2"
)

type strictServerInterface struct{}

	`)

	paths := map[string]any{}
	responses := map[int]map[string]any{
		1: {},
		2: {},
	}

	for _, ref := range []bool{false, true} {
		for _, extRef := range []bool{false, true} {
			nameSlice := []string{"test"}
			var pkgId int = 1
			if ref {
				if extRef {
					continue
					nameSlice = append(nameSlice, "Ext")
					pkgId = 2
				} else {
					nameSlice = append(nameSlice, "Ref")
				}
			} else {
				if extRef {
					continue
				}
			}
			for _, header := range []bool{false, true} {
				nameSlice := nameSlice
				if header {
					nameSlice = append(nameSlice, "Header")
				}
				for _, fixedStatusCode := range []bool{false, true} {
					nameSlice := nameSlice
					if fixedStatusCode {
						nameSlice = append(nameSlice, "Fixed")
					}
					for _, content := range []struct {
						name    []string
						content string
						tag     string
					}{
						{name: []string{"JSON"}, content: "application/json", tag: "JSON"},
						{name: []string{"Special", "JSON"}, content: "application/test+json", tag: "ApplicationTestPlusJSON"},
						{name: []string{"Formdata"}, content: "application/x-www-form-urlencoded", tag: "Formdata"},
						{name: []string{"Multipart"}, content: "multipart/test", tag: "Multipart"},
						// issue B
						// {name: []string{"Wildcard", "Multipart"}, content: "multipart/*", tag: "Multipart"},
						{name: []string{"Text"}, content: "text/plain", tag: "Text"},
						{name: []string{"Other"}, content: "application/test", tag: "Applicationtest"},
						{name: []string{"Wildcard"}, content: "application/*", tag: "Application"},
						{name: []string{"NoContent"}},
					} {
						if content.content == "text/plain" && (header || !fixedStatusCode || ref) {
							// issue A
							continue
						}

						if !(content.content == "application/x-www-form-urlencoded" && ref) {
							// issue #1402
							continue
						}

						if strings.Contains(content.content, "json") && extRef && !fixedStatusCode {
							// issue #1202
							continue
						}

						nameSlice := append(nameSlice, content.name...)

						response := map[string]any{}

						if content.content != "" {
							response["content"] = map[string]any{
								content.content: map[string]any{
									"schema": map[string]any{
										"$ref": "#/components/schemas/TestSchema",
									},
								},
							}
						}

						if header {
							response["headers"] = map[string]any{
								"header1": map[string]any{
									"schema": map[string]any{
										"type": "string",
									},
								},
								"header2": map[string]any{
									"schema": map[string]any{
										"type": "integer",
									},
								},
							}
						}

						var statusCode string
						if fixedStatusCode {
							if content.content == "" {
								statusCode = "204"
							} else {
								statusCode = "200"
							}
						} else {
							statusCode = "default"
						}

						var responseInResponses map[string]any
						if ref {
							responseName := "testResp" + strings.Join(nameSlice[1:], "")
							responses[pkgId][responseName] = response
							var ref string
							if pkgId == 1 {
								ref = fmt.Sprintf("#/components/responses/%s", responseName)
							} else {
								ref = fmt.Sprintf("pkg%d.yaml#/components/responses/%s", pkgId, responseName)
							}
							responseInResponses = map[string]any{
								"$ref": ref,
							}
						} else {
							responseInResponses = response
						}

						pathSlice := make([]string, 0, len(nameSlice))
						for _, s := range nameSlice {
							pathSlice = append(pathSlice, strings.ToLower(s))
						}
						path := "/" + strings.Join(pathSlice, "-")
						paths[path] = map[string]any{
							"get": map[string]any{
								"operationId": strings.Join(nameSlice, ""),
								"responses": map[string]any{
									statusCode: responseInResponses,
								},
							},
						}

						method := "Test" + strings.Join(nameSlice[1:], "")
						body := ""
						var statusCodeRes int
						if content.content == "" {
							statusCodeRes = 204
						} else {
							statusCodeRes = 200
						}
						switch content.content {
						case "":
						case "text/plain":
							body = "\"bar\""
						case "multipart/test", "multipart/*":
							body = fmt.Sprintf(`func(writer *multipart.Writer) error {
						if p, err := writer.CreatePart(textproto.MIMEHeader{"Content-Type": []string{"application/json"}}); err != nil {
							return err
						} else {
							return json.NewEncoder(p).Encode(pkg%d.TestSchema{
								Field1: "bar",
								Field2: 456,
							})
						}
					}`, pkgId)
						case "application/test", "application/*":
							body = "bytes.NewReader(buf)"
						default:
							body = fmt.Sprintf(`pkg%d.TestSchema{
						Field1: "bar",
						Field2: 456,
					}`, pkgId)
						}
						contentRes := content.content
						if strings.HasSuffix(contentRes, "/*") {
							contentRes = strings.TrimSuffix(contentRes, "/*") + "/baz"
						}
						fmt.Fprintf(fTestGo, "func (s strictServerInterface) %s(ctx context.Context, request pkg1.%sRequestObject) (pkg1.%sResponseObject, error) {\n", method, method, method)
						switch content.content {
						case "application/test", "application/*":
							fmt.Fprintf(fTestGo, "buf := []byte(\"bar\")\n")
						}
						var resRet string
						if ref && fixedStatusCode {
							resRet = fmt.Sprintf("pkg%d.TestResp%s%sResponse", pkgId, strings.Join(nameSlice[1:], ""), content.tag)
						} else {
							resRet = fmt.Sprintf("pkg1.%s%s%sResponse", method, statusCode, content.tag)
						}

						if header || !fixedStatusCode ||
							content.content == "application/test" || content.content == "application/*" ||
							content.content == "multipart/*" || content.content == "" {
							resRet += " {\n"
							if body != "" {
								resRet += fmt.Sprintf("Body: %s,\n", body)
							}
							if header {
								var headerType string
								if ref {
									headerType = fmt.Sprintf("pkg%d.TestResp%sResponseHeaders", pkgId, strings.Join(nameSlice[1:], ""))
								} else {
									headerType = fmt.Sprintf("pkg1.%s%sResponseHeaders", method, statusCode)
								}
								resRet += fmt.Sprintf(`Headers: %s {
									Header1: "foo",
									Header2: 123,
								},
								`, headerType)
							}
							if !fixedStatusCode {
								resRet += fmt.Sprintf("StatusCode: %d,\n", statusCodeRes)
							}
							if strings.HasSuffix(content.content, "/*") {
								resRet += fmt.Sprintf("ContentType: \"%s\",\n", contentRes)
							}
							switch content.content {
							case "application/test", "application/*":
								resRet += "ContentLength: int64(len(buf)),\n"
							}
							resRet += "}"
						} else {
							resRet += "(" + body + ")"
						}
						if ref && fixedStatusCode {
							if (strings.HasPrefix(content.content, "multipart/") && !header) ||
								(content.content == "" && extRef) {
								resRet = fmt.Sprintf("pkg1.%s%s%sResponse(%s)", method, statusCode, content.tag, resRet)
							} else if content.content != "" {
								resRet = fmt.Sprintf("pkg1.%s%s%sResponse{%s}", method, statusCode, content.tag, resRet)
							}
						}
						fmt.Fprintf(fTestGo, "return %s, nil\n", resRet)
						fmt.Fprintf(fTestGo, "}\n")
						fmt.Fprintf(fTestGo, "\n")

						fmt.Fprintf(fTestGo, "func %s(t *testing.T) {\n", method)
						fmt.Fprintf(fTestGo, "hh := pkg1.Handler(pkg1.NewStrictHandler(strictServerInterface{}, nil))\n")
						fmt.Fprintf(fTestGo, "\n")
						fmt.Fprintf(fTestGo, "ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {\n")
						fmt.Fprintf(fTestGo, "if !assert.Equal(t, \"%s\", r.URL.Path) {\n", path)
						fmt.Fprintf(fTestGo, "w.WriteHeader(http.StatusInternalServerError)\n")
						fmt.Fprintf(fTestGo, "return\n")
						fmt.Fprintf(fTestGo, "}\n")
						fmt.Fprintf(fTestGo, "hh.ServeHTTP(w, r)\n")
						fmt.Fprintf(fTestGo, "}))\n")
						fmt.Fprintf(fTestGo, "defer ts.Close()\n")
						fmt.Fprintf(fTestGo, "\n")
						fmt.Fprintf(fTestGo, "c, err := pkg1.NewClientWithResponses(ts.URL)\n")
						fmt.Fprintf(fTestGo, "assert.NoError(t, err)\n")
						fmt.Fprintf(fTestGo, "res, err := c.%sWithResponse(context.TODO())\n", method)
						fmt.Fprintf(fTestGo, "assert.NoError(t, err)\n")
						fmt.Fprintf(fTestGo, "assert.Equal(t, %d, res.StatusCode())\n", statusCodeRes)
						if header {
							fmt.Fprintf(fTestGo, "assert.Equal(t, \"foo\", res.HTTPResponse.Header.Get(\"header1\"))\n")
							fmt.Fprintf(fTestGo, "assert.Equal(t, \"123\", res.HTTPResponse.Header.Get(\"header2\"))\n")
						}
						if !strings.HasPrefix(contentRes, "multipart/") {
							fmt.Fprintf(fTestGo, "assert.Equal(t, \"%s\", res.HTTPResponse.Header.Get(\"Content-Type\"))\n", contentRes)
						}
						switch content.content {
						case "":
							fmt.Fprintf(fTestGo, "assert.Equal(t, []byte{}, res.Body)\n")
						case "multipart/test", "multipart/*":
							fmt.Fprintf(fTestGo, `mediaType, params, err := mime.ParseMediaType(res.HTTPResponse.Header.Get("Content-Type"))
					if assert.NoError(t, err) {
						assert.Equal(t, "%s", mediaType)
						assert.NotEmpty(t, params["boundary"])
						reader := multipart.NewReader(bytes.NewReader(res.Body), params["boundary"])
						jsonExist := false
						for {
							if p, err := reader.NextPart(); err == io.EOF {
								break
							} else {
								assert.NoError(t, err)
								switch p.Header.Get("Content-Type") {
								case "application/json":
									var j pkg%d.TestSchema
									err := json.NewDecoder(p).Decode(&j)
									assert.NoError(t, err)
									assert.Equal(t, pkg%d.TestSchema{
										Field1: "bar",
										Field2: 456,
									}, j)
									jsonExist = true
								default:
									assert.Fail(t, "Bad Content-Type: %%s", p.Header.Get("Content-Type"))
								}
							}
						}
						assert.True(t, jsonExist)
					}
					`, contentRes, pkgId, pkgId)
						case "application/x-www-form-urlencoded":
							fmt.Fprintf(fTestGo, `form, err := url.ParseQuery(string(res.Body))
					assert.NoError(t, err)
					assert.Equal(t, url.Values{
						"field1": []string{"bar"},
						"field2": []string{"456"},
					}, form)
					`)
						case "application/json":
							fmt.Fprintf(fTestGo, "assert.Equal(t, &pkg%d.TestSchema{\n", pkgId)
							fmt.Fprintf(fTestGo, "Field1: \"bar\",\n")
							fmt.Fprintf(fTestGo, "Field2: 456,\n")
							fmt.Fprintf(fTestGo, "}, res.JSON%s%s)\n", strings.ToUpper(statusCode[:1]), statusCode[1:])
						case "application/test+json":
							fmt.Fprintf(fTestGo, "assert.Equal(t, &pkg%d.TestSchema{\n", pkgId)
							fmt.Fprintf(fTestGo, "Field1: \"bar\",\n")
							fmt.Fprintf(fTestGo, "Field2: 456,\n")
							fmt.Fprintf(fTestGo, "}, res.ApplicationtestJSON%s%s)\n", strings.ToUpper(statusCode[:1]), statusCode[1:])
						default:
							fmt.Fprintf(fTestGo, "assert.Equal(t, []byte(\"bar\"), res.Body)\n")
						}
						fmt.Fprintf(fTestGo, "}\n")
						fmt.Fprintf(fTestGo, "\n")
					}
				}
			}
		}
	}

	specs := map[int]map[string]any{
		1: {
			"openapi": "3.0.1",
			"components": map[string]any{
				"schemas": map[string]any{
					"TestSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"field1": map[string]any{
								"type": "string",
							},
							"field2": map[string]any{
								"type": "integer",
							},
						},
						"required": []any{
							"field1",
							"field2",
						},
					},
				},
				"responses": responses[1],
			},
			"paths": paths,
		},
		2: {
			"openapi": "3.0.1",
			"components": map[string]any{
				"schemas": map[string]any{
					"TestSchema": map[string]any{
						"type": "object",
						"properties": map[string]any{
							"field1": map[string]any{
								"type": "string",
							},
							"field2": map[string]any{
								"type": "integer",
							},
						},
						"required": []any{
							"field1",
							"field2",
						},
					},
				},
				"responses": responses[2],
			},
		},
	}

	for pkgId, spec := range specs {
		if buf, err := yaml.Marshal(spec); err != nil {
			panic(err)
		} else if fYAML, err := os.Create(fmt.Sprintf("pkg%d.yaml", pkgId)); err != nil {
			panic(err)
		} else {
			defer func() {
				if err := fYAML.Close(); err != nil {
					fmt.Fprintln(os.Stderr, err)
				}
			}()
			fmt.Fprintln(fYAML, "# Code generated by generator/generate.go DO NOT EDIT.")
			if _, err := fYAML.Write(buf); err != nil {
				panic(err)
			}
		}
	}

	testGoName := "pkg1/pkg1_test.go"
	if buf, err := imports.Process(testGoName, fTestGo.Bytes(), nil); err != nil {
		panic(err)
	} else if fTestGoOut, err := os.Create(testGoName); err != nil {
		panic(err)
	} else {
		defer func() {
			if err := fTestGoOut.Close(); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}()
		if _, err := fTestGoOut.Write(buf); err != nil {
			panic(err)
		}
	}
}
