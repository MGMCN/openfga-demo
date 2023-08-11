package main

import (
	"context"
	"fmt"
	"github.com/MGMCN/openfga-demo/openfga"
)

func main() {
	ctx := context.Background()
	client := openfga.NewClient("http", "0.0.0.0:8080", "", "")
	if client != nil {
		if stores, err := client.ListStore(ctx); err != nil {
			fmt.Println(err)
		} else {
			for _, store := range *stores {
				if *store.Name == "demo" {
					client.SetStoreId(*store.Id)
					break
				}
			}

		}
		// Add data for gaoshan and normal_user for testing purposes
		if err := client.CreateRelationTuple(ctx, "folder:docs", "owner", "admin:gaoshan"); err != nil {
			fmt.Println(err)
		}
		if err := client.CreateRelationTuple(ctx, "folder:docs", "viewer", "user:normal_user"); err != nil {
			fmt.Println(err)
		}
		if err := client.CreateRelationTuple(ctx, "doc:doc1", "parent", "folder:docs"); err != nil {
			fmt.Println(err)
		}

		// Test whether the authentication of gaoshan and normal_user conforms to our settings
		if permission, err := client.GetCheck(ctx, "doc:doc1", "can_read", "admin:gaoshan"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The permission for gaoshan to read the 'doc1' file is", permission)
		}

		if permission, err := client.GetCheck(ctx, "doc:doc1", "can_write", "admin:gaoshan"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The permission for gaoshan to write the 'doc1' file is", permission)
		}

		if permission, err := client.GetCheck(ctx, "doc:doc1", "can_read", "user:normal_user"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The permission for normal_user to read the 'doc1' file is", permission)
		}

		if permission, err := client.GetCheck(ctx, "doc:doc1", "can_write", "user:normal_user"); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("The permission for normal_user to write the 'doc1' file is", permission)
		}
	} else {
		fmt.Println("Failed to connect to the OpenFGA server.")
	}
}
