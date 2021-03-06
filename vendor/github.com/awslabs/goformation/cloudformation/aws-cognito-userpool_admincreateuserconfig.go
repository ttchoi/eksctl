package cloudformation

import (
	"encoding/json"
)

// AWSCognitoUserPool_AdminCreateUserConfig AWS CloudFormation Resource (AWS::Cognito::UserPool.AdminCreateUserConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
type AWSCognitoUserPool_AdminCreateUserConfig struct {

	// AllowAdminCreateUserOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	AllowAdminCreateUserOnly *Value `json:"AllowAdminCreateUserOnly,omitempty"`

	// InviteMessageTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	InviteMessageTemplate *AWSCognitoUserPool_InviteMessageTemplate `json:"InviteMessageTemplate,omitempty"`

	// UnusedAccountValidityDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	UnusedAccountValidityDays *Value `json:"UnusedAccountValidityDays,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_AdminCreateUserConfig) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.AdminCreateUserConfig"
}

func (r *AWSCognitoUserPool_AdminCreateUserConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
