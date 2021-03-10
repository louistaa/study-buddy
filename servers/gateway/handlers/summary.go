package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

//PreviewImage represents a preview image for a page
type PreviewImage struct {
	URL       string `json:"url,omitempty"`
	SecureURL string `json:"secureURL,omitempty"`
	Type      string `json:"type,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	Alt       string `json:"alt,omitempty"`
}

//PageSummary represents summary properties for a web page
type PageSummary struct {
	Type        string          `json:"type,omitempty"`
	URL         string          `json:"url,omitempty"`
	Title       string          `json:"title,omitempty"`
	SiteName    string          `json:"siteName,omitempty"`
	Description string          `json:"description,omitempty"`
	Author      string          `json:"author,omitempty"`
	Keywords    []string        `json:"keywords,omitempty"`
	Icon        *PreviewImage   `json:"icon,omitempty"`
	Images      []*PreviewImage `json:"images,omitempty"`
}

//SummaryHandler handles requests for the page summary API.
//This API expects one query string parameter named `url`,
//which should contain a URL to a web page. It responds with
//a JSON-encoded PageSummary struct containing the page summary
//meta-data.
func SummaryHandler(w http.ResponseWriter, r *http.Request) {
	/*Helpful Links:
	https://golang.org/pkg/net/http/#Request.FormValue
	https://golang.org/pkg/net/http/#Error
	https://golang.org/pkg/encoding/json/#NewEncoder
	*/

	// - Add an HTTP header to the response with the name
	//  `Access-Control-Allow-Origin` and a value of `*`. This will
	//   allow cross-origin AJAX requests to your server.
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// - Get the `url` query string parameter value from the request.
	//   If not supplied, respond with an http.StatusBadRequest error.
	pageURL := r.FormValue("url")
	if len(pageURL) == 0 {
		http.Error(w, "URL parameter not provided", http.StatusBadRequest)
		return
	}

	// - Call fetchHTML() to fetch the requested URL. See comments in that
	//   function for more details.
	htmlStream, err := fetchHTML(pageURL)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// - Call extractSummary() to extract the page summary meta-data,
	//   as directed in the assignment. See comments in that function
	//   for more details
	summary, err := extractSummary(pageURL, htmlStream)
	if err != nil {
		http.Error(w, "Error extracting summary", http.StatusInternalServerError)
	}

	// - Close the response HTML stream so that you don't leak resources.
	defer htmlStream.Close()

	// - Finally, respond with a JSON-encoded version of the PageSummary
	//   struct. That way the client can easily parse the JSON back into
	//   an object. Remember to tell the client that the response content
	//   type is JSON.
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(summary); err != nil {
		http.Error(w, "error encoding JSON", http.StatusInternalServerError)
		return
	}
}

//fetchHTML fetches `pageURL` and returns the body stream or an error.
//Errors are returned if the response status code is an error (>=400),
//or if the content type indicates the URL is not an HTML page.
func fetchHTML(pageURL string) (io.ReadCloser, error) {
	/*TODO: Do an HTTP GET for the page URL. If the response status
	code is >= 400, return a nil stream and an error. If the response
	content type does not indicate that the content is a web page, return
	a nil stream and an error. Otherwise return the response body and
	no (nil) error.

	To test your implementation of this function, run the TestFetchHTML
	test in summary_test.go. You can do that directly in Visual Studio Code,
	or at the command line by running:
		go test -run TestFetchHTML

	Helpful Links:
	https://golang.org/pkg/net/http/#Get
	*/

	resp, err := http.Get(pageURL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode >= 400 {
		return nil, errors.New("Error Response Status")
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "text/html") {
		return nil, errors.New("Error Content is not a Webpage")
	}

	return resp.Body, nil
}

//extractSummary tokenizes the `htmlStream` and populates a PageSummary
//struct with the page's summary meta-data.
func extractSummary(pageURL string, htmlStream io.ReadCloser) (*PageSummary, error) {
	/* tokenize the `htmlStream` and extract the page summary meta-data
	according to the assignment description.

	To test your implementation of this function, run the TestExtractSummary
	test in summary_test.go. You can do that directly in Visual Studio Code,
	or at the command line by running:
		go test -run TestExtractSummary

	Helpful Links:
	https://drstearns.github.io/tutorials/tokenizing/
	http://ogp.me/
	https://developers.facebook.com/docs/reference/opengraph/
	https://golang.org/pkg/net/url/#URL.ResolveReference
	*/

	summary := &PageSummary{}
	tokenizer := html.NewTokenizer(htmlStream)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err == io.EOF {
				break
			}
			return nil, tokenizer.Err()
		}

		if tokenType == html.EndTagToken {
			token := tokenizer.Token()
			if "head" == token.Data {
				return summary, nil
			}
		}

		if tokenType == html.StartTagToken || tokenType == html.SelfClosingTagToken {
			token := tokenizer.Token()
			if "meta" == token.Data {
				property := extractMetaAttr(token.Attr, "property")
				name := extractMetaAttr(token.Attr, "name")
				content := extractMetaAttr(token.Attr, "content")
				if property == "og:type" {
					summary.Type = content
				} else if property == "og:url" {
					summary.URL = content
				} else if property == "og:title" {
					summary.Title = content
				} else if property == "og:site_name" {
					summary.SiteName = content
				} else if property == "og:description" {
					summary.Description = content
				} else if name == "author" {
					summary.Author = content
				} else if (len(summary.Description) == 0) && (name == "description") {
					summary.Description = content
				} else if name == "keywords" {
					words := strings.Split(content, ",")
					keywords := []string{}
					for _, word := range words {
						keywords = append(keywords, strings.TrimSpace(word))
					}
					summary.Keywords = keywords
				} else if strings.HasPrefix(property, "og:image") {
					if property == "og:image" {
						image := &PreviewImage{}
						image.URL = content
						if strings.HasPrefix(image.URL, "/") {
							urlPart := strings.Split(pageURL, "/")
							urlPart = urlPart[:len(urlPart)-1]
							url := strings.Join(urlPart, "/")
							image.URL = url + image.URL
						}
						summary.Images = append(summary.Images, image)
					} else if property == "og:image:secure_url" {
						image := summary.Images[len(summary.Images)-1]
						image.SecureURL = content
						if strings.HasPrefix(image.SecureURL, "/") {
							urlPart := strings.Split(pageURL, "/")
							urlPart = urlPart[:len(urlPart)-1]
							url := strings.Join(urlPart, "/")
							image.SecureURL = url + image.SecureURL
						}
					} else if property == "og:image:type" {
						image := summary.Images[len(summary.Images)-1]
						image.Type = content
					} else if property == "og:image:width" {
						image := summary.Images[len(summary.Images)-1]
						image.Width, _ = strconv.Atoi(content)
					} else if property == "og:image:height" {
						image := summary.Images[len(summary.Images)-1]
						image.Height, _ = strconv.Atoi(content)
					} else if property == "og:image:alt" {
						image := summary.Images[len(summary.Images)-1]
						image.Alt = content
					}
				}
			} else if "title" == token.Data {
				if len(summary.Title) == 0 {
					tokenType = tokenizer.Next()
					if tokenType == html.TextToken {
						summary.Title = tokenizer.Token().Data
					}
				}
			} else if "link" == token.Data {
				rel := extractMetaAttr(token.Attr, "rel")
				if rel == "icon" {
					icon := &PreviewImage{}
					icon.URL = extractMetaAttr(token.Attr, "href")
					if strings.HasPrefix(icon.URL, "/") {
						urlPart := strings.Split(pageURL, "/")
						urlPart = urlPart[:len(urlPart)-1]
						url := strings.Join(urlPart, "/")
						icon.URL = url + icon.URL
					}
					icon.Type = extractMetaAttr(token.Attr, "type")
					sizes := extractMetaAttr(token.Attr, "sizes")
					if sizes != "" {
						if strings.Contains(sizes, "x") {
							slice := strings.Split(sizes, "x")
							height, err := strconv.Atoi(slice[0])
							width, err := strconv.Atoi(slice[1])
							if err != nil {
								break
							}
							icon.Height = height
							icon.Width = width
						}
					}
					summary.Icon = icon
				}
			}
		}
	}

	return summary, nil
}

func extractMetaAttr(attributes []html.Attribute, key string) string {
	for _, attr := range attributes {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
