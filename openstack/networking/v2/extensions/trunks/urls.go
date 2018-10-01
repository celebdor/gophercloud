package trunks

import "github.com/gophercloud/gophercloud"

const resourcePath = "trunks"

func resourceURL(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id)
}

func rootURL(c *gophercloud.ServiceClient) string {
	return c.ServiceURL(resourcePath)
}

func listURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func getURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func createURL(c *gophercloud.ServiceClient) string {
	return rootURL(c)
}

func updateURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func deleteURL(c *gophercloud.ServiceClient, id string) string {
	return resourceURL(c, id)
}

func addSubports(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "add_subports")
}

func removeSubports(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "remove_subports")
}

func getSubports(c *gophercloud.ServiceClient, id string) string {
	return c.ServiceURL(resourcePath, id, "get_subports")
}
