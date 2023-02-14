package shared

import (
	"time"
)

type OrganizationDomainInput struct {
	Domain                 string  `json:"domain" form:"name=domain" multipartForm:"name=domain"`
	JitProvisioningEnabled *bool   `json:"jit_provisioning_enabled,omitempty" form:"name=jit_provisioning_enabled" multipartForm:"name=jit_provisioning_enabled"`
	SamlAcsURL             *string `json:"saml_acs_url,omitempty" form:"name=saml_acs_url" multipartForm:"name=saml_acs_url"`
	SamlEntityID           *string `json:"saml_entity_id,omitempty" form:"name=saml_entity_id" multipartForm:"name=saml_entity_id"`
	SamlX509Cert           *string `json:"saml_x509_cert,omitempty" form:"name=saml_x509_cert" multipartForm:"name=saml_x509_cert"`
	SsoEnforcement         *string `json:"sso_enforcement,omitempty" form:"name=sso_enforcement" multipartForm:"name=sso_enforcement"`
}

type OrganizationDomain struct {
	Domain                 string    `json:"domain"`
	HasSaml                bool      `json:"has_saml"`
	ID                     string    `json:"id"`
	IsVerified             bool      `json:"is_verified"`
	JitProvisioningEnabled *bool     `json:"jit_provisioning_enabled,omitempty"`
	SamlAcsURL             *string   `json:"saml_acs_url,omitempty"`
	SamlEntityID           *string   `json:"saml_entity_id,omitempty"`
	SamlX509Cert           *string   `json:"saml_x509_cert,omitempty"`
	SsoEnforcement         *string   `json:"sso_enforcement,omitempty"`
	VerificationChallenge  string    `json:"verification_challenge"`
	VerifiedAt             time.Time `json:"verified_at"`
}
