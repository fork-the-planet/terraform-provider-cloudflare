// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package account

import (
	"context"

	"github.com/cloudflare/terraform-provider-cloudflare/internal/customfield"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

var _ datasource.DataSourceWithConfigValidators = (*AccountsDataSource)(nil)

func ListDataSourceSchema(ctx context.Context) schema.Schema {
	return schema.Schema{
		Attributes: map[string]schema.Attribute{
			"direction": schema.StringAttribute{
				Description: "Direction to order results.\nAvailable values: \"asc\", \"desc\".",
				Optional:    true,
				Validators: []validator.String{
					stringvalidator.OneOfCaseInsensitive("asc", "desc"),
				},
			},
			"name": schema.StringAttribute{
				Description: "Name of the account.",
				Optional:    true,
			},
			"max_items": schema.Int64Attribute{
				Description: "Max items to fetch, default: 1000",
				Optional:    true,
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"result": schema.ListNestedAttribute{
				Description: "The items returned by the data source",
				Computed:    true,
				CustomType:  customfield.NewNestedObjectListType[AccountsResultDataSourceModel](ctx),
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: "Identifier",
							Computed:    true,
						},
						"name": schema.StringAttribute{
							Description: "Account name",
							Computed:    true,
						},
						"created_on": schema.StringAttribute{
							Description: "Timestamp for the creation of the account",
							Computed:    true,
							CustomType:  timetypes.RFC3339Type{},
						},
						"settings": schema.SingleNestedAttribute{
							Description: "Account settings",
							Computed:    true,
							CustomType:  customfield.NewNestedObjectType[AccountsSettingsDataSourceModel](ctx),
							Attributes: map[string]schema.Attribute{
								"abuse_contact_email": schema.StringAttribute{
									Description: "Sets an abuse contact email to notify for abuse reports.",
									Computed:    true,
								},
								"enforce_twofactor": schema.BoolAttribute{
									Description: "Indicates whether membership in this account requires that\nTwo-Factor Authentication is enabled",
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *AccountsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = ListDataSourceSchema(ctx)
}

func (d *AccountsDataSource) ConfigValidators(_ context.Context) []datasource.ConfigValidator {
	return []datasource.ConfigValidator{}
}
