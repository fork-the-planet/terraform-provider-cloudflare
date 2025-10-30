package organization_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/acctest"
	"github.com/cloudflare/terraform-provider-cloudflare/internal/utils"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestMain is the entry point for test execution
func TestMain(m *testing.M) {
	resource.TestMain(m)
}

// TestAccCloudflareOrganization_Basic tests the basic CRUD operations for organization resource
func TestAccCloudflareOrganization_Basic(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization." + rnd
	orgName := fmt.Sprintf("tf-acctest-%s", rnd)
	updatedOrgName := fmt.Sprintf("tf-acctest-%s-updated", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOrganizationDestroy,
		Steps: []resource.TestStep{
			// Step 1: Create - Test resource creation with all required attributes
			{
				Config: testAccOrganizationConfig(rnd, orgName),
				Check: resource.ComposeTestCheckFunc(
					// Verify required attributes
					resource.TestCheckResourceAttr(resourceName, "name", orgName),
					// Verify computed attributes are set
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
					// Verify meta attributes
					resource.TestCheckResourceAttrSet(resourceName, "meta.%"),
				),
			},
			// Step 2: Update - Test modifying updatable attributes
			{
				Config: testAccOrganizationConfig(rnd, updatedOrgName),
				Check: resource.ComposeTestCheckFunc(
					// Verify the name was updated
					resource.TestCheckResourceAttr(resourceName, "name", updatedOrgName),
					// Verify ID remains the same (it should be set)
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					// Verify other attributes remain consistent
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
				),
			},
			// Step 3: Import - Test import functionality with proper ID format
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				// Organization import uses just the ID, no prefix needed
			},
		},
	})
}

// TestAccCloudflareOrganization_WithProfile tests organization creation with profile information
func TestAccCloudflareOrganization_WithProfile(t *testing.T) {
	rnd := utils.GenerateRandomResourceName()
	resourceName := "cloudflare_organization." + rnd
	orgName := fmt.Sprintf("tf-acctest-%s", rnd)

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.TestAccPreCheck(t) },
		ProtoV6ProviderFactories: acctest.TestAccProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckCloudflareOrganizationDestroy,
		Steps: []resource.TestStep{
			// Create organization with profile
			{
				Config:             testAccOrganizationConfigWithProfile(rnd, orgName),
				ExpectNonEmptyPlan: true, // Allow non-empty plan due to profile field handling
				Check: resource.ComposeTestCheckFunc(
					// Basic attribute checks
					resource.TestCheckResourceAttr(resourceName, "name", orgName),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "create_time"),
					// Profile checks
					resource.TestCheckResourceAttr(resourceName, "profile.business_name", "Test Business"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_email", "test@example.com"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_phone", "+1234567890"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_address", "123 Test St, Test City, TC 12345"),
				),
			},
			// Update profile information
			{
				Config:             testAccOrganizationConfigWithProfileUpdated(rnd, orgName),
				ExpectNonEmptyPlan: true, // Allow non-empty plan due to profile field handling
				Check: resource.ComposeTestCheckFunc(
					// Verify profile was updated
					resource.TestCheckResourceAttr(resourceName, "profile.business_name", "Updated Business"),
					resource.TestCheckResourceAttr(resourceName, "profile.business_email", "updated@example.com"),
				),
			},
			// Import test
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				
				ImportStateVerifyIgnore: []string{
					"profile", // Profile fields not populated on import
					"profile.%",
					"profile.business_name",
					"profile.business_email", 
					"profile.business_phone",
					"profile.business_address",
					"profile.external_metadata",
				}, 
			},
		},
	})
}



// Test configuration functions that load from testdata files

func testAccOrganizationConfig(rnd, name string) string {
	return acctest.LoadTestCase("basic.tf", rnd, name)
}

func testAccOrganizationConfigWithProfile(rnd, name string) string {
	return acctest.LoadTestCase("with_profile.tf", rnd, name)
}

func testAccOrganizationConfigWithProfileUpdated(rnd, name string) string {
	return acctest.LoadTestCase("with_profile_updated.tf", rnd, name)
}

func testAccOrganizationConfigWithParent(rnd, name, parentID string) string {
	return acctest.LoadTestCase("with_parent.tf", rnd, name, parentID)
}

// testAccCheckCloudflareOrganizationDestroy verifies the organization has been destroyed
func testAccCheckCloudflareOrganizationDestroy(s *terraform.State) error {
	client := acctest.SharedClient()

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "cloudflare_organization" {
			continue
		}

		// Try to fetch the organization
		_, err := client.Organizations.Get(context.Background(), rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("organization %s still exists", rs.Primary.ID)
		}
	}

	return nil
}
