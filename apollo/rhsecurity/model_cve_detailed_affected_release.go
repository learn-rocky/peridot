/*
 * Red Hat Security Data API
 *
 * Unofficial OpenAPI definitions for Red Hat Security Data API
 *
 * API version: 1.0
 * Contact: mustafa@ctrliq.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package rhsecurity

import (
	"encoding/json"
)

// CVEDetailedAffectedRelease struct for CVEDetailedAffectedRelease
type CVEDetailedAffectedRelease struct {
	ProductName string  `json:"product_name"`
	ReleaseDate string  `json:"release_date"`
	Advisory    string  `json:"advisory"`
	Cpe         string  `json:"cpe"`
	Package     *string `json:"package,omitempty"`
}

// NewCVEDetailedAffectedRelease instantiates a new CVEDetailedAffectedRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCVEDetailedAffectedRelease(productName string, releaseDate string, advisory string, cpe string) *CVEDetailedAffectedRelease {
	this := CVEDetailedAffectedRelease{}
	this.ProductName = productName
	this.ReleaseDate = releaseDate
	this.Advisory = advisory
	this.Cpe = cpe
	return &this
}

// NewCVEDetailedAffectedReleaseWithDefaults instantiates a new CVEDetailedAffectedRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCVEDetailedAffectedReleaseWithDefaults() *CVEDetailedAffectedRelease {
	this := CVEDetailedAffectedRelease{}
	return &this
}

// GetProductName returns the ProductName field value
func (o *CVEDetailedAffectedRelease) GetProductName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductName
}

// GetProductNameOk returns a tuple with the ProductName field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedAffectedRelease) GetProductNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductName, true
}

// SetProductName sets field value
func (o *CVEDetailedAffectedRelease) SetProductName(v string) {
	o.ProductName = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *CVEDetailedAffectedRelease) GetReleaseDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedAffectedRelease) GetReleaseDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *CVEDetailedAffectedRelease) SetReleaseDate(v string) {
	o.ReleaseDate = v
}

// GetAdvisory returns the Advisory field value
func (o *CVEDetailedAffectedRelease) GetAdvisory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Advisory
}

// GetAdvisoryOk returns a tuple with the Advisory field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedAffectedRelease) GetAdvisoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Advisory, true
}

// SetAdvisory sets field value
func (o *CVEDetailedAffectedRelease) SetAdvisory(v string) {
	o.Advisory = v
}

// GetCpe returns the Cpe field value
func (o *CVEDetailedAffectedRelease) GetCpe() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpe
}

// GetCpeOk returns a tuple with the Cpe field value
// and a boolean to check if the value has been set.
func (o *CVEDetailedAffectedRelease) GetCpeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpe, true
}

// SetCpe sets field value
func (o *CVEDetailedAffectedRelease) SetCpe(v string) {
	o.Cpe = v
}

// GetPackage returns the Package field value if set, zero value otherwise.
func (o *CVEDetailedAffectedRelease) GetPackage() string {
	if o == nil || o.Package == nil {
		var ret string
		return ret
	}
	return *o.Package
}

// GetPackageOk returns a tuple with the Package field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CVEDetailedAffectedRelease) GetPackageOk() (*string, bool) {
	if o == nil || o.Package == nil {
		return nil, false
	}
	return o.Package, true
}

// HasPackage returns a boolean if a field has been set.
func (o *CVEDetailedAffectedRelease) HasPackage() bool {
	if o != nil && o.Package != nil {
		return true
	}

	return false
}

// SetPackage gets a reference to the given string and assigns it to the Package field.
func (o *CVEDetailedAffectedRelease) SetPackage(v string) {
	o.Package = &v
}

func (o CVEDetailedAffectedRelease) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["product_name"] = o.ProductName
	}
	if true {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if true {
		toSerialize["advisory"] = o.Advisory
	}
	if true {
		toSerialize["cpe"] = o.Cpe
	}
	if o.Package != nil {
		toSerialize["package"] = o.Package
	}
	return json.Marshal(toSerialize)
}

type NullableCVEDetailedAffectedRelease struct {
	value *CVEDetailedAffectedRelease
	isSet bool
}

func (v NullableCVEDetailedAffectedRelease) Get() *CVEDetailedAffectedRelease {
	return v.value
}

func (v *NullableCVEDetailedAffectedRelease) Set(val *CVEDetailedAffectedRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableCVEDetailedAffectedRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableCVEDetailedAffectedRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCVEDetailedAffectedRelease(val *CVEDetailedAffectedRelease) *NullableCVEDetailedAffectedRelease {
	return &NullableCVEDetailedAffectedRelease{value: val, isSet: true}
}

func (v NullableCVEDetailedAffectedRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCVEDetailedAffectedRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
