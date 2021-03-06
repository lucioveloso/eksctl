package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSEC2VPNConnectionRoute AWS CloudFormation Resource (AWS::EC2::VPNConnectionRoute)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
type AWSEC2VPNConnectionRoute struct {

	// DestinationCidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-cidrblock
	DestinationCidrBlock *Value `json:"DestinationCidrBlock,omitempty"`

	// VpnConnectionId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-connectionid
	VpnConnectionId *Value `json:"VpnConnectionId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnectionRoute) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnectionRoute"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSEC2VPNConnectionRoute) MarshalJSON() ([]byte, error) {
	type Properties AWSEC2VPNConnectionRoute
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
func (r *AWSEC2VPNConnectionRoute) UnmarshalJSON(b []byte) error {
	type Properties AWSEC2VPNConnectionRoute
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
		*r = AWSEC2VPNConnectionRoute(*res.Properties)
	}

	return nil
}

// GetAllAWSEC2VPNConnectionRouteResources retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
func (t *Template) GetAllAWSEC2VPNConnectionRouteResources() map[string]AWSEC2VPNConnectionRoute {
	results := map[string]AWSEC2VPNConnectionRoute{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSEC2VPNConnectionRoute:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPNConnectionRoute" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2VPNConnectionRoute{}
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

// GetAWSEC2VPNConnectionRouteWithName retrieves all AWSEC2VPNConnectionRoute items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEC2VPNConnectionRouteWithName(name string) (AWSEC2VPNConnectionRoute, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSEC2VPNConnectionRoute:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::EC2::VPNConnectionRoute" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSEC2VPNConnectionRoute{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSEC2VPNConnectionRoute{}, errors.New("resource not found")
}
