package shared

type PatchedOrganizationDomainInput struct {
	Domain                 *string `json:"domain,omitempty" form:"name=domain" multipartForm:"name=domain"`
	JitProvisioningEnabled *bool   `json:"jit_provisioning_enabled,omitempty" form:"name=jit_provisioning_enabled" multipartForm:"name=jit_provisioning_enabled"`
	SamlAcsURL             *string `json:"saml_acs_url,omitempty" form:"name=saml_acs_url" multipartForm:"name=saml_acs_url"`
	SamlEntityID           *string `json:"saml_entity_id,omitempty" form:"name=saml_entity_id" multipartForm:"name=saml_entity_id"`
	SamlX509Cert           *string `json:"saml_x509_cert,omitempty" form:"name=saml_x509_cert" multipartForm:"name=saml_x509_cert"`
	SsoEnforcement         *string `json:"sso_enforcement,omitempty" form:"name=sso_enforcement" multipartForm:"name=sso_enforcement"`
}
