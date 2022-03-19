package models

type ListModel struct {
	Name string
	Structure string
	Jsonschema string
}
type ListCatalog struct {
	Catalog string
	Patch string
	Files   []string
}

type GroupRoute struct {
	Name string
	Group []ListRoute
}
type ListRoute struct {
	Name string
	Link string
	Function string
	Folder string
	MetodFunc string
	Origin string
	RawBody string
}
type PostmanResponse struct {
	Name            string `json:"name"`
	OriginalRequest struct {
		Method string `json:"method"`
		Header []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"header"`
		Body struct {
			Mode string `json:"mode"`
			Raw  string `json:"raw"`
			Options struct {
				Raw struct {
					Language string `json:"language"`
				} `json:"raw"`
			} `json:"options"`
		} `json:"body"`
		Url struct {
			Raw  string   `json:"raw"`
			Host []string `json:"host"`
			Path []string `json:"path"`
		} `json:"url"`
	} `json:"originalRequest"`
	Status                 string `json:"status"`
	Code                   int    `json:"code"`
	PostmanPreviewlanguage string `json:"_postman_previewlanguage"`
	Header                 []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"header"`
	Cookie []interface{} `json:"cookie"`
	Body   string        `json:"body"`
}
type PostmanRoute struct {
	Name    string `json:"name"`
	Request struct {
		Method string `json:"method"`
		Header []struct {
			Key   string `json:"key"`
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"header"`
		Body struct {
			Mode string `json:"mode"`
			Raw  string `json:"raw"`
			Options struct {
				Raw struct {
					Language string `json:"language"`
				} `json:"raw"`
			} `json:"options"`
		} `json:"body"`
		Url struct {
			Raw  string   `json:"raw"`
			Host []string `json:"host"`
			Path []string `json:"path"`
		} `json:"url"`
	} `json:"request"`
	Response []PostmanResponse `json:"response"`
}

type PostmanFolders struct {
	Name string `json:"name"`
	Item []PostmanRoute `json:"item"`
}

type Postman struct {
	Info struct {
		PostmanId string `json:"_postman_id"`
		Name      string `json:"name"`
		Schema    string `json:"schema"`
	} `json:"info"`
	Item []PostmanFolders `json:"item"`
	Event []PostmanEventList `json:"event"`
	Variable []PostmanVariables `json:"variable"`
}

type PostmanEventList struct {
		Listen string `json:"listen"`
		Script struct {
			Type string   `json:"type"`
			Exec []string `json:"exec"`
		} `json:"script"`
}
type PostmanVariables struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}