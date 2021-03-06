/*
 * Copyright (c) 2019 Entrust Datacard Corporation.
 * All rights reserved.
 */

package main

import (
	"github.com/hashicorp/vault/logical"
	"github.com/hashicorp/vault/logical/framework"
)

func pathConfig(b *backend) *framework.Path {
	ret := &framework.Path{
		Pattern: "config/" + framework.GenericNameRegex("caId"),

		Operations: map[logical.Operation]framework.OperationHandler{
			logical.ReadOperation:   &framework.PathOperation{Callback: b.opReadConfigCA},
			logical.UpdateOperation: &framework.PathOperation{Callback: b.opWriteConfigCA},
		},

		HelpSynopsis:    "CAGW Configuration",
		HelpDescription: "Configures CAGW parameters including client cert and key.",
		Fields:          map[string]*framework.FieldSchema{},
	}

	ret.Fields["pem_bundle"] = &framework.FieldSchema{
		Type:        framework.TypeString,
		Default:     "",
		Description: `PEM encoded client certificate and key.`,
		Required:    true,
	}

	ret.Fields["caId"] = &framework.FieldSchema{
		Type:        framework.TypeString,
		Default:     "",
		Description: `CA identifier`,
		Required:    true,
	}

	ret.Fields["url"] = &framework.FieldSchema{
		Type:        framework.TypeString,
		Default:     "",
		Description: `URL for CAGW including base context path`,
		Required:    true,
	}

	ret.Fields["cacerts"] = &framework.FieldSchema{
		Type:    framework.TypeString,
		Default: "",
		Description: "PEM encoded CA certificate chain. Not needed if the gateway's " +
			"certificate is publicly trusted.",
	}

	return ret
}
