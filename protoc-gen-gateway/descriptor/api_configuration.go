package descriptor

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/ghodss/yaml"
	"github.com/golang/protobuf/jsonpb"
)




func loadGrpcAPIServiceFromYAML(yamlFileContents []byte, yamlSourceLogName string) (*GrpcAPIService, error) {
	jsonContents, err := yaml.YAMLToJSON(yamlFileContents)
	if err != nil {
		return nil, fmt.Errorf("Failed to convert gRPC API Configuration from YAML in '%v' to JSON: %v", yamlSourceLogName, err)
	}

	// As our GrpcAPIService is incomplete accept unkown fields.
	unmarshaler := jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}

	serviceConfiguration := GrpcAPIService{}
	if err := unmarshaler.Unmarshal(bytes.NewReader(jsonContents), &serviceConfiguration); err != nil {
		return nil, fmt.Errorf("Failed to parse gRPC API Configuration from YAML in '%v': %v", yamlSourceLogName, err)
	}

	return &serviceConfiguration, nil
}





func registerHTTPRulesFromGrpcAPIService(registry *Registry, service *GrpcAPIService, sourceLogName string) error {
	if service.HTTP == nil {
		// Nothing to do
		return nil
	}

	for _, rule := range service.HTTP.GetRules() {
		selector := "." + strings.Trim(rule.GetSelector(), " ")
		if strings.ContainsAny(selector, "*, ") {
			return fmt.Errorf("Selector '%v' in %v must specify a single service method without wildcards", rule.GetSelector(), sourceLogName)
		}

		registry.AddExternalHTTPRule(selector, rule)
	}

	return nil
}
