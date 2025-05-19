// Package license helps manage commercial licenses and check if they
// are valid for the version of UniOffice used.
package license

import _g "gitee.com/greatmusicians/unioffice/internal/license"

// SetMeteredKey sets the metered License API key required for SaaS operation.
// Document usage is reported periodically for the product to function correctly.
func SetMeteredKey(apiKey string) error { return _g.SetMeteredKey(apiKey) }

// GetMeteredState checks the currently used metered document usage status,
// documents used and credits available.
func GetMeteredState() (_g.MeteredStatus, error) { return _g.GetMeteredState() }

// MakeUnlicensedKey returns a default key.
func MakeUnlicensedKey() *LicenseKey { return _g.MakeUnlicensedKey() }

const (
	LicenseTierUnlicensed = _g.LicenseTierUnlicensed
	LicenseTierCommunity  = _g.LicenseTierCommunity
	LicenseTierIndividual = _g.LicenseTierIndividual
	LicenseTierBusiness   = _g.LicenseTierBusiness
)

// LegacyLicense holds the old-style unioffice license information.
type LegacyLicense = _g.LegacyLicense

// SetLicenseKey sets and validates the license key.
func SetLicenseKey(content string, customerName string) error {
	return _g.SetLicenseKey(content, customerName)
}

// LegacyLicenseType is the type of license
type LegacyLicenseType = _g.LegacyLicenseType

// LicenseKey represents a loaded license key.
type LicenseKey = _g.LicenseKey

// GetLicenseKey returns the currently loaded license key.
func GetLicenseKey() *LicenseKey { return _g.GetLicenseKey() }

// SetLegacyLicenseKey installs a legacy license code. License codes issued prior to June 2019.
// Will be removed at some point in a future major version.
func SetLegacyLicenseKey(s string) error { return _g.SetLegacyLicenseKey(s) }
