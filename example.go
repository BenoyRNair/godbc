// Copyright (c) 2009 Benoy R Nair. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
 * 22-Nov-09 Benoy R Nair	First draft
 * 23-Nov-09 Benoy R Nair	For SQLDriverConnect()
 */
package main

import "godbc"
import "fmt"

var env, dbc godbc.GS_HANDLE;

func init()
{
	env.GS_AllocHandle ( godbc.GS_HANDLE_ENV, godbc.NULL_HANDLE );
	env.GS_SetEnvAttr ( godbc.GS_ATTR_ODBC_VERSION, godbc.GS_OV_ODBC3, 0 );
	dbc.GS_AllocHandle ( godbc.GS_HANDLE_DBC, env );
}

func main()
{
	listDataSources();
	listDrivers();
	connect();
}

func listDataSources()
{
	fmt.Printf ( "Listing Data Sources (odbcinst -q -s) >>\n" );

	var x int;
	var dsn, desc string;

	direction := godbc.GS_FETCH_FIRST;

	for
	{
		x, dsn, desc = env.GS_DataSources ( direction );

		if ( ! godbc.GS_Succeeded ( x ) )
		{
			break;
		}

		direction = godbc.GS_FETCH_NEXT;

		fmt.Printf ( "DSN: %s Desc: %s\n", dsn, desc );

		if x == godbc.GS_SUCCESS_WITH_INFO
		{
			fmt.Printf ( "\tdata truncation\n" );
		}
	}
}

func listDrivers()
{
	fmt.Printf ( "Listing Drivers (odbcinst -q -d) >>\n" );

	var x int;
	var driver, attr string;

	direction := godbc.GS_FETCH_FIRST;

	for
	{
		x, driver, attr = env.GS_Drivers ( direction );

		if ( ! godbc.GS_Succeeded ( x ) )
		{
			break;
		}

		direction = godbc.GS_FETCH_NEXT;

		fmt.Printf ( "Driver: %s Attr: %s\n", driver, attr );

		if x == godbc.GS_SUCCESS_WITH_INFO
		{
			fmt.Printf ( "\tdata truncation\n" );
		}
	}
}

func connect()
{
	x, str := dbc.GS_DriverConnect ( 0, "DSN=dsn1mysql;UID=root;PWD=password", godbc.GS_DRIVER_COMPLETE );

	if ( godbc.GS_Succeeded ( x ) )
	{
		fmt.Printf ( "Connected. Details: %s\n", str );
	}
	else
	{
		fmt.Printf ( "Unable to connect. Check the login credentials provided in the code.\n" );
	}
}
