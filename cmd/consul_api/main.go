// Package consul implements grpclb service discovery via Consul (consul.io).
// Service names may be specified as: svcname,tag1,tag2,tag3
package main

import (
	"fmt"
	"log"

	"github.com/hashicorp/consul/api"
)

func main() {
	cfg := api.DefaultConfig()
	cfg.Address = "consul:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	leader, err := client.Status().Leader()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(leader)

	peers, err := client.Status().Peers()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(peers)

	fmt.Println("------------------------------ Nodes ------------------------------")
	nodes, _, err := client.Health().Node("223c4c806f23", nil)
	if err != nil {
		log.Fatal(err)
	}
	for _, node := range nodes {
		fmt.Println(describeHealthcheck(node))
	}
	fmt.Println("-------------------------------------------------------------------")

	services, _, err := client.Health().Service("addition", "", true, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("------------------------------ Services ------------------------------")
	for _, service := range services {
		fmt.Println(describeNode(service.Node))
	}
	fmt.Println("----------------------------------------------------------------------")
}

func describeHealthcheck(healthCheck *api.HealthCheck) string {
	tmpl := `
	Node        %v
	CheckID     %v
	Name        %v
	Status      %v
	Notes       %v
	Output      %v
	ServiceID   %v
	ServiceName %v
	ServiceTags %v

	Definition %v

	CreateIndex %v
	ModifyIndex %v
	`
	return fmt.Sprintf(
		tmpl,
		healthCheck.Node,
		healthCheck.CheckID,
		healthCheck.Name,
		healthCheck.Status,
		healthCheck.Notes,
		healthCheck.Output,
		healthCheck.ServiceID,
		healthCheck.ServiceName,
		healthCheck.ServiceTags,

		healthCheck.Definition,

		healthCheck.CreateIndex,
		healthCheck.ModifyIndex,
	)
}

func describeNode(node *api.Node) string {
	tmpl := `
	ID:              %v
	Node:            %v
	Address:         %v
	Datacenter:      %v
	TaggedAddresses: %v
	Meta:            %v
	CreateIndex:     %v
	ModifyIndex:     %v
	`
	return fmt.Sprintf(
		tmpl,
		node.ID,
		node.Node,
		node.Address,
		node.Datacenter,
		node.TaggedAddresses,
		node.Meta,
		node.CreateIndex,
		node.ModifyIndex,
	)
}