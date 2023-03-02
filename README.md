# swagger-go

## 1. OpenAPI
#####  https://github.com/getkin/kin-openapi
#####  https://github.com/OAI/OpenAPI-Specification/blob/main/versions/2.0.md


To generate Swagger documentation using go-openapi, you can use the swagger generate command. Here are the steps to generate Swagger documentation for your Go API using go-openapi:

1. Install go-swagger: First, you need to install the go-swagger tool, which provides a set of command-line tools to work with Swagger specifications. You can install go-swagger by running the following command:
	go get -u github.com/go-swagger/go-swagger/cmd/swagger
	
	One way that worked for me on Windows: (https://stackoverflow.com/questions/56413951/setting-up-go-swagger-for-first-time)
		git clone https://github.com/go-swagger/go-swagger
		cd go-swagger
		go install ./cmd/swagger
	
2. Add Swagger annotations to your Go code: In order to generate Swagger documentation, you need to add annotations to your Go code using go-openapi. These annotations provide information about your API, such as the API path, request and response formats, and authentication methods. Here is an example of how to add a Swagger annotation to a simple API endpoint in your Go code:
	// swagger:route GET /hello hello-world
	//
	// Returns a hello world message.
	//
	// This will return a hello world message.
	//
	//     Responses:
	//       200: helloResponse
	//
	func helloWorld(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	}

	// swagger:response helloResponse
	type helloResponse struct {
		// in: body
		Body struct {
			// The hello world message
			Message string `json:"message"`
		}
	}
This adds Swagger annotations to the helloWorld function and defines a response schema for the API endpoint.

3. Generate Swagger documentation: Once you have added the Swagger annotations to your Go code, you can use the swagger generate command to generate Swagger documentation. Run the following command in the root directory of your project:
	swagger generate spec -o ./swagger.yaml --scan-models
	
	This generates a Swagger specification file named swagger.yaml in the current directory. The --scan-models flag tells swagger to include model definitions in the generated Swagger specification.

4. Serve Swagger documentation: You can serve the generated Swagger documentation using a web server or Swagger UI. Here is an example of how to serve the Swagger documentation using a simple web server:
	1. You can run command : swagger serve -F=swagger swagger.yaml
		OR
	2. You can write code as below:
	 // Serve Swagger documentation
    router.PathPrefix("/docs/").Handler(httpSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Serve API on port 8080
    log.Fatal(http.ListenAndServe(":8080", router))
	
	This serves the Swagger documentation at the /docs/ endpoint using the http-swagger package. You can access the Swagger documentation by navigating to http://localhost:8080/docs/ in your web browser.

That's it! You have generated Swagger documentation for your Go API using go-openapi.



##### $ go get -u github.com/go-swagger/go-swagger/cmd/swagger
##### $ swagger generate spec -o ./swagger.yaml –scan-models
