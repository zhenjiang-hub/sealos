package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// LaunchTemplateVersionSet is a nested struct in ecs response
type LaunchTemplateVersionSet struct {
	CreateTime         string             `json:"CreateTime" xml:"CreateTime"`
	ModifiedTime       string             `json:"ModifiedTime" xml:"ModifiedTime"`
	LaunchTemplateId   string             `json:"LaunchTemplateId" xml:"LaunchTemplateId"`
	LaunchTemplateName string             `json:"LaunchTemplateName" xml:"LaunchTemplateName"`
	DefaultVersion     bool               `json:"DefaultVersion" xml:"DefaultVersion"`
	VersionNumber      int64              `json:"VersionNumber" xml:"VersionNumber"`
	VersionDescription string             `json:"VersionDescription" xml:"VersionDescription"`
	CreatedBy          string             `json:"CreatedBy" xml:"CreatedBy"`
	LaunchTemplateData LaunchTemplateData `json:"LaunchTemplateData" xml:"LaunchTemplateData"`
}