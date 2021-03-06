package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSIoT1ClickPlacement AWS CloudFormation Resource (AWS::IoT1Click::Placement)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html
type AWSIoT1ClickPlacement struct {

	// AssociatedDevices AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-associateddevices
	AssociatedDevices interface{} `json:"AssociatedDevices,omitempty"`

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-attributes
	Attributes interface{} `json:"Attributes,omitempty"`

	// PlacementName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-placementname
	PlacementName *Value `json:"PlacementName,omitempty"`

	// ProjectName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot1click-placement.html#cfn-iot1click-placement-projectname
	ProjectName *Value `json:"ProjectName,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoT1ClickPlacement) AWSCloudFormationType() string {
	return "AWS::IoT1Click::Placement"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSIoT1ClickPlacement) MarshalJSON() ([]byte, error) {
	type Properties AWSIoT1ClickPlacement
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSIoT1ClickPlacement) UnmarshalJSON(b []byte) error {
	type Properties AWSIoT1ClickPlacement
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSIoT1ClickPlacement(*res.Properties)
	}

	return nil
}

// GetAllAWSIoT1ClickPlacementResources retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
func (t *Template) GetAllAWSIoT1ClickPlacementResources() map[string]AWSIoT1ClickPlacement {
	results := map[string]AWSIoT1ClickPlacement{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSIoT1ClickPlacement:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IoT1Click::Placement" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSIoT1ClickPlacement{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSIoT1ClickPlacementWithName retrieves all AWSIoT1ClickPlacement items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoT1ClickPlacementWithName(name string) (AWSIoT1ClickPlacement, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSIoT1ClickPlacement:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::IoT1Click::Placement" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSIoT1ClickPlacement{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSIoT1ClickPlacement{}, errors.New("resource not found")
}
