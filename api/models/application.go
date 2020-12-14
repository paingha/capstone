// Copyright 2020 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"bitbucket.com/irb/api/config"
)

//Application - struct for the Application DB model
type Application struct {
	ID                                             int                              `json:"id,omitempty" sql:"primary_key"`
	ProjectTitle                                   string                           `json:"projectTitle"`
	UUID                                           string                           `json:"uuid"`
	PrincipalInvestigatorFirstName                 string                           `json:"principalInvestigatorFirstName"`
	PrincipalInvestigatorLastName                  string                           `json:"principalInvestigatorLastName"`
	PrincipalInvestigatorTitle                     string                           `json:"principalInvestigatorTitle"`
	PrincipalInvestigatorInstitutionalEmailAddress string                           `json:"principalInvestigatorInstitutionalEmailAddress"`
	PrincipalInvestigatorDepartment                string                           `json:"principalInvestigatorDepartment"`
	PrincipalInvestigatorDaytimePhone              string                           `json:"principalInvestigatorDaytimePhone"`
	PrincipalInvestigatorMailingAddress            string                           `json:"principalInvestigatorMailingAddress"`
	PrincipalInvestigatorStatus                    string                           `json:"principalInvestigatorStatus"`
	CoInvestigators                                []CoInvestigator                 `json:"coInvestigators"`
	FundingSources                                 []FundingSources                 `json:"fundingSources"`
	GrantTitle                                     string                           `json:"grantTitle"`
	FundingSource                                  string                           `json:"fundingSource"`
	PIofGrant                                      string                           `json:"piOfGrant"`
	Sponsor                                        string                           `json:"sponsor"`
	GrantNumber                                    string                           `json:"grantNumber"`
	ProposedStartDate                              string                           `json:"proposedStartDate"`
	ProjectType                                    string                           `json:"projectType"`
	MedicalClearance                               bool                             `json:"medicalClearance"`
	MedicalClearanceExplanation                    string                           `json:"medicalClearanceExplanation"`
	MedicalClearanceInstrumentExplanation          string                           `json:"medicalClearanceInstrumentExplanation"`
	StudySites                                     []StudySites                     `json:"studySites"`
	PotentialVulnerablePopulations                 []PotentialVulnerablePopulations `json:"potentiallyVulnerablePopulations"`
	ExternalOversight                              []ExternalOversight              `json:"externalOversight"`
	ConflictOfInterest                             bool                             `json:"conflictOfInterest"`
	ConflictOfInterestExplanation                  string                           `json:"conflictOfInterestExplanation"`
	StudyBackgroundInformation                     string                           `json:"studyBackgroundInformation"`
	StudyResearchDesign                            string                           `json:"studyResearchDesign"`
	PotentialResearchSubjects                      string                           `json:"potentialResearchSubjects"`
	StudyProcedure                                 string                           `json:"studyProcedure"`
	StudyTime                                      string                           `json:"studyTime"`
	StudyDocument                                  string                           `json:"studyDocument"`
	RecruitmentProcedure                           string                           `json:"recruitmentProcedure"`
	RecruitmentMaterial                            string                           `json:"recruitmentMaterial"`
	WillCollectDirectIdentifiers                   bool                             `json:"willCollectDirectIdentifiers"`
	CollectionOfDirectIdentifiersExplanation       string                           `json:"collectionOfDirectIdentifiersExplanation"`
	CodingSystemProtection                         string                           `json:"codingSystemProtection"`
	IdentifierDestructionOrMaintenance             string                           `json:"identifierDesctructionOrMaintenance"`
	LinkBetweenStudyCodeAndIdentifiers             bool                             `json:"linkBetweenStudyCodeAndIdentifiers"`
	LinkBetweenStudyCodeAndIdentifiersExplanation  string                           `json:"linkBetweenStudyCodeAndIdentifiersExplanation"`
	LinkBetweenOutsideAndResearchTeam              bool                             `json:"linkBetweenOutsideAndResearchTeam"`
	LinkBetweenOutsideAndResearchTeamExplanation   string                           `json:"linkBetweenOutsideAndResearchTeamExplanation"`
	DataFormat                                     string                           `json:"dataFormat"`
	AudioVideoPhotos                               bool                             `json:"audioVideoPhotos"`
	AudioVideoPhotosExplanation                    string                           `json:"audioVideoPhotosExplanation"`
	DataStorageAndSecurityPrecautions              string                           `json:"dataStorageAndSecurityPrecautions"`
	DataDestructionTime                            string                           `json:"dataDestructionTime"`
	DataDestructionMethods                         []DataDestructionMethod          `json:"dataDestructionMethods"`
	ConsentFormRecord                              bool                             `json:"consentFormRecord"`
	ConsentFormRecordExplanation                   string                           `json:"consentFormRecordExplanation"`
	RecordAvailability                             bool                             `json:"recordAvailability"`
	RecordAvailabilityExplanation                  string                           `json:"recordAvailabilityExplanation"`
	FederalCertificateOfConfidentiality            bool                             `json:"federalCertificateOfConfidentiality"`
	FederalCertificateOfConfidentialityExplanation string                           `json:"federalCertificateOfConfidentialityExplanation"`
	ConsentFiles                                   []ConsentFiles                   `json:"consentFiles"`
	Risks                                          []Risk                           `json:"risks"`
	NatureOfRiskExplanation                        string                           `json:"natureOfRisk"`
	StepsToMinimizeRisk                            string                           `json:"stepsToMinimizeRisk"`
	UseOfDeceptionExplanation                      string                           `json:"useOfDeceptionExplanation"`
	BenefitsDescriptionIndividual                  string                           `json:"benefitsDescriptionIndividual"`
	BenefitsDescriptionSociety                     string                           `json:"benefitsDescriptionSociety"`
	AnyCompensation                                bool                             `json:"anyCompensation"`
	CompensationExplanation                        string                           `json:"compensationExplanation"`
	Status                                         string                           `json:"status"`
	CreatedAt                                      time.Time                        `json:"created_at"`
	UpdatedAt                                      time.Time                        `json:"updated_at"`
	DeletedAt                                      *time.Time                       `json:"deleted_at"`
}

//CoInvestigator - struct of all Co Investigators
type CoInvestigator struct {
	CoInvestigatorFirstName                 string `json:"coInvestigatorFirstName"`
	CoInvestigatorLastName                  string `json:"coInvestigatorLastName"`
	CoInvestigatorTitle                     string `json:"coInvestigatorTitle"`
	CoInvestigatorInstitutionalEmailAddress string `json:"coInvestigatorInstitutionalEmailAddress"`
	CoInvestigatorDepartment                string `json:"coInvestigatorDepartment"`
	CoInvestigatorDaytimePhone              string `json:"coInvestigatorDaytimePhone"`
	CoInvestigatorMailingAddress            string `json:"coInvestigatorMailingAddress"`
	CoInvestigatorStatus                    string `json:"coInvestigatorStatus"`
}

//ConsentFiles - struct for array of consentfiles
type ConsentFiles struct {
	FileName string `json:"fileName"`
	FileType string `json:"fileType"`
	URL      string `json:"url"`
}

//Risk - struct for array of risk
type Risk struct {
	Content string `json:"content"`
}

//DataDestructionMethod - struct for array of DataDestructionMethod
type DataDestructionMethod struct {
	Content string `json:"content"`
}

//ExternalOversight - struct for array of ExternalOversight
type ExternalOversight struct {
	Content string `json:"content"`
}

//PotentialVulnerablePopulations - struct for array of PotentialVulnerablePopulations
type PotentialVulnerablePopulations struct {
	Content string `json:"content"`
}

//StudySites - struct for array of StudySites
type StudySites struct {
	Content string `json:"content"`
}

//FundingSources - struct for array of FundingSources
type FundingSources struct {
	Content string `json:"content"`
}

//TableName - application model db table name set to applications
func (Application) TableName() string {
	return "applications"
}

//GetAllApplications - fetch all applications at once
func GetAllApplications(application *[]Application, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Application{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(application).Error; err != nil {
		return count, err
	}
	return count, nil
}

//GetMyApplications - fetch my applications
func GetMyApplications(application *[]Application, id int, offset int, limit int) (int, error) {
	var count = 0
	if err := config.DB.Model(&Application{}).Where("id = ?", id).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(application).Error; err != nil {
		return count, err
	}
	return count, nil
}

//CreateApplication - create an application
func CreateApplication(application *Application) (bool, error) {
	if errs := config.DB.Create(application).Error; errs != nil {
		return false, errs
	}
	return true, nil
}

//GetApplication - fetch one application
func GetApplication(application *Application, id int) error {
	if err := config.DB.Where("id = ?", id).First(application).Error; err != nil {
		return err
	}
	return nil
}

//UpdateApplication - update an application
func UpdateApplication(application *Application, id int) error {
	if err := config.DB.Model(&application).Where("id = ?", id).Updates(application).Error; err != nil {
		return err
	}
	return nil
}

//DeleteApplication - delete an application
func DeleteApplication(id int) error {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Application{}).Error; err != nil {
		return err
	}
	return nil
}
