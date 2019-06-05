package publichandler

import (
	"net/http"

	"html/template"
	"fmt"
)


type HomePage struct {
	Gateway  string
	Service string
	Domain string
	Id string
}

var gatewayTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="/public/assets/swagger-ui/swagger-ui.css" >
  <style>
  .information-container { display:none; }
  </style>
</head>

<body>
<div class="swagger-ui"><div class="wrapper"><h2 class="title">[{{.Service}}]</h2></div></div>
<div id="swagger-ui"></div>

<script src="/public/assets/swagger-ui/swagger-ui-bundle.js"> </script>
<script>
const DisableTryItOutPlugin = function() {
  return {
    statePlugins: {
      spec: {
        wrapSelectors: {
          allowTryItOutFor: () => () => false
        }
      }
    }
  }
}

window.onload = function() {

  const ui = SwaggerUIBundle({
    url: "http://{{.Gateway}}/swagger/{{.Service}}",
    //spec: {"swagger":"2.0","info":{"description":"Api Documentation","version":"1.0","title":"Api Documentation","termsOfService":"urn:tos","contact":{"name":"Contact Email"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0"}},"host":"10.255.130.73:8080","basePath":"/","tags":[{"name":"back-controller","description":"Back Controller"}],"paths":{"/search/get":{"post":{"tags":["back-controller"],"summary":"search","operationId":"searchUsingPOST","consumes":["application/json"],"produces":["*/*"],"parameters":[{"in":"body","name":"model","description":"model","required":false,"schema":{"$ref":"#/definitions/Model"}}],"responses":{"200":{"description":"OK","schema":{"type":"array","items":{"$ref":"#/definitions/Data"}}},"201":{"description":"Created"},"401":{"description":"Unauthorized"},"403":{"description":"Forbidden"},"404":{"description":"Not Found"}}}}},"definitions":{"Data":{"properties":{"id":{"type":"string"},"name":{"type":"string"}}}}},
    dom_id: '#swagger-ui',
    deepLinking: false,
    //plugins: [DisableTryItOutPlugin],
    layout: "BaseLayout",
	requestInterceptor: (req) => {
	    if(req.loadSpec) {
	      var hash = btoa("admin" + ":" + "password")
	      req.headers.Authorization = "Basic " + hash
	    }
	    return req
	  }
  })

  window.ui = ui
}
</script>
</body>

</html>
`


var entryTemplate string = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Swagger UI</title>
  <link rel="stylesheet" type="text/css" href="/public/assets/swagger-ui/swagger-ui.css" >
  <style>
  .information-container { display:none; }
  </style>
</head>

<body>
<!-- <div class="swagger-ui"><div class="wrapper"><h2 class="title">[{{.Domain}}]</h2></div></div> -->
<div id="swagger-ui"></div>

<script src="/public/assets/swagger-ui/swagger-ui-bundle.js"> </script>
<script>
const DisableTryItOutPlugin = function() {
  return {
    statePlugins: {
      spec: {
        wrapSelectors: {
          allowTryItOutFor: () => () => false
        }
      }
    }
  }
}

window.onload = function() {

  const ui = SwaggerUIBundle({
    url: "{{.Domain}}/apigateway/{{.Id}}/swagger",
    //url: "{{.Domain}}/microservices/{{.Id}}/api",
    //spec: {"swagger":"2.0","info":{"description":"Api Documentation","version":"1.0","title":"Api Documentation","termsOfService":"urn:tos","contact":{"name":"Contact Email"},"license":{"name":"Apache 2.0","url":"http://www.apache.org/licenses/LICENSE-2.0"}},"host":"10.255.130.73:8080","basePath":"/","tags":[{"name":"back-controller","description":"Back Controller"}],"paths":{"/search/get":{"post":{"tags":["back-controller"],"summary":"search","operationId":"searchUsingPOST","consumes":["application/json"],"produces":["*/*"],"parameters":[{"in":"body","name":"model","description":"model","required":false,"schema":{"$ref":"#/definitions/Model"}}],"responses":{"200":{"description":"OK","schema":{"type":"array","items":{"$ref":"#/definitions/Data"}}},"201":{"description":"Created"},"401":{"description":"Unauthorized"},"403":{"description":"Forbidden"},"404":{"description":"Not Found"}}}}},"definitions":{"Data":{"properties":{"id":{"type":"string"},"name":{"type":"string"}}}}},
    dom_id: '#swagger-ui',
    deepLinking: false,
    //plugins: [DisableTryItOutPlugin],
    layout: "BaseLayout",
	requestInterceptor: (req) => {
	    if(req.loadSpec) {
	      var hash = btoa("admin" + ":" + "password")
	      req.headers.Authorization = "Basic " + hash
	    }
	    return req
	  }
  })

  window.ui = ui
}
</script>
</body>

</html>
`


func NewSwaggerGatewayHandler() (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		template := template.Must(template.New("gateway").Parse(gatewayTemplate))
		gateway := r.FormValue("gateway")
		service := r.FormValue("service")
		err := template.Execute(w, HomePage{
			Gateway:  gateway,
			Service : service,
		})
		if err != nil {
			fmt.Println(err)
		}
	}), nil
}

func NewSwaggerEntryHandler() (http.Handler, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

		template := template.Must(template.New("entry").Parse(entryTemplate))
		domain := r.FormValue("domain")
		id := r.FormValue("id")
		err := template.Execute(w, HomePage{
			Domain:  domain,
			Id: id,
		})
		if err != nil {
			fmt.Println(err)
		}
	}), nil
}


