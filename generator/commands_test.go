package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Describe(t *testing.T) {

	short, long := Describe("ConstituentsGet")
	assert.Equal(t, "ConstituentsGet", short)
	assert.Equal(t, "Get the details of a Constituent using id.\n ConstituentId(*): string\t", long)

	short, long = Describe("ConstituentsUpdate")
	assert.Equal(t, "ConstituentsUpdate", short)
	assert.Equal(t, "Update a constituent. Only the information about constituent can be updated. If addresses, electronicAddresses, salutations or phones data are passed, they will be ignored.\n ConstituentId(*): string\t\n Constituent(*): object\t", long)

	short, long = Describe("ConstituentsCreateConstituent")
	assert.Equal(t, "ConstituentsCreateConstituent", short)
	assert.Equal(t, "Create a new constituent with addresses, electronicAddresses, salutations and phones.\n Constituent(*): object\t", long)

	short, long = Describe("ElectronicAddressesGet")
	assert.Equal(t, "ElectronicAddressesGet", short)
	assert.Equal(t, "Get details of an electronic address.\n ElectronicAddressId(*): string\t", long)

	short, long = Describe("ElectronicAddressesUpdate")
	assert.Equal(t, "ElectronicAddressesUpdate", short)
	assert.Equal(t, "Update an electronic address.\n ElectronicAddressId(*): string\t\n ElectronicAddress(*): object\t", long)

	short, long = Describe("ElectronicAddressesCreate")
	assert.Equal(t, "ElectronicAddressesCreate", short)
	assert.Equal(t, "Create a new electronic address.\n ElectronicAddress(*): object\t", long)

	short, long = Describe("AssociationTypesGet")
	assert.Equal(t, "AssociationTypesGet", short)
	assert.Equal(t, "Get the details of an association type by id. To get the resource only if the user has write/edit access, pass filter=\"writeonly\". To get the resources in maintenance mode(ignore Control Groups), pass maintenanceMode=\"true\". Maintenance mode requires users to have access to the reference tables. (Specified in TX_REFTABLE_USERGROUP).\n Id(*): string\t\n Filter: string\tFilter by user access (default: readwrite)\n MaintenanceMode: string\tIgnore control grouping (default: false)", long)

}

func Test_DescribeFunc(t *testing.T) {
	short, long := DescribeFunc("ConstituentsGet")
	assert.Equal(t, "ConstituentsGet", short)
	assert.Equal(t, "gets the details of a constituent using id\n", long)

	short, long = DescribeFunc("ConstituentsUpdate")
	assert.Equal(t, "ConstituentsUpdate", short)
	assert.Equal(t, "updates a constituent only the information about constituent can be updated if addresses electronic addresses salutations or phones data are passed they will be ignored\n", long)

	short, long = DescribeFunc("ConstituentsCreateConstituent")
	assert.Equal(t, "ConstituentsCreateConstituent", short)
	assert.Equal(t, "creates a new constituent with addresses electronic addresses salutations and phones\n", long)

	short, long = DescribeFunc("ElectronicAddressesGet")
	assert.Equal(t, "ElectronicAddressesGet", short)
	assert.Equal(t, "gets details of an electronic address\n", long)

	short, long = DescribeFunc("ElectronicAddressesUpdate")
	assert.Equal(t, "ElectronicAddressesUpdate", short)
	assert.Equal(t, "updates an electronic address\n", long)

	short, long = DescribeFunc("ElectronicAddressesCreate")
	assert.Equal(t, "ElectronicAddressesCreate", short)
	assert.Equal(t, "creates a new electronic address\n", long)
}
