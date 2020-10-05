package main

import (
	"context"
	"fmt"
	"time"

	"github.com/influxdata/influxdb-client-go"
)

func main() {
	userName := "root"
	password := "mashiro"
	// Create a client
	// Supply a string in the form: "username:password" as a token. Set empty value for an unauthenticated server
	client := influxdb2.NewClient("http://116.62.7.230:8086", fmt.Sprintf("%s:%s",userName, password))
	// Get the blocking write client
	// Supply a string in the form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
	// Org name is not used
	writeApi := client.WriteApiBlocking("", "testdb/autogen")
	// create point using full params constructor
	p := influxdb2.NewPointWithMeasurement("stat").
		AddTag("unit", "temperature").
		AddField("avg", 23.2).
		AddField("max", 45).
		SetTime(time.Now())
	// Write data
	err := writeApi.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Write error: %s\n", err.Error())
	}

	// Get query client. Org name is not used
	queryApi := client.QueryApi("")
	// Supply string in a form database/retention-policy as a bucket. Skip retention policy for the default one, use just a database name (without the slash character)
	result, err := queryApi.Query(context.Background(), `from(bucket:"testdb/autogen")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	if err == nil {
		for result.Next() {
			if result.TableChanged() {
				fmt.Printf("table: %s\n", result.TableMetadata().String())
			}
			fmt.Printf("row: %s\n\n", result.Record().String())
		}
		if result.Err() != nil {
			fmt.Printf("Query error: %s\n", result.Err().Error())
		}
	} else {
		fmt.Printf("Query error: %s\n", err.Error())
	}
	// Close client
	client.Close()
}