// Package api handles communication with the SpaceTraders API.
package api

/**
* Note that this package deliberately only generates models from the OpenAPI spec:
* - the commonly used github.com/oapi-codegen/oapi-codegen produces invalid code (gofmt throws errors) when generating client code
* - the generator at https://openapi-generator.tech produces invalid code (e.g. unused imports) that would require manual cleanup
* - the online generator at swagger.io doesn't allow for proper customization
* - the offline generator at swagger.io (swagger-codegen-cli) does not work with the current spec
* - no other client code generators are known at the moment
*
* Obviously, some of these tools would work with some manual labour, but I decided against this.
* Generated code should be used as-is, without any manual changes.
**/

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.3.0 -config oapi-cfg.yaml https://stoplight.io/api/v1/projects/spacetraders/spacetraders/nodes/reference/SpaceTraders.json?fromExportButton=true&snapshotType=http_service&deref=optimizedBundle
