// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * An Azure Cosmos DB container.
 * API Version: 2021-03-15.
 *
 * ## Example Usage
 * ### CosmosDBSqlContainerCreateUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azure_native from "@pulumi/azure-native";
 *
 * const sqlResourceSqlContainer = new azure_native.documentdb.SqlResourceSqlContainer("sqlResourceSqlContainer", {
 *     accountName: "ddb1",
 *     containerName: "containerName",
 *     databaseName: "databaseName",
 *     location: "West US",
 *     options: {},
 *     resource: {
 *         conflictResolutionPolicy: {
 *             conflictResolutionPath: "/path",
 *             mode: "LastWriterWins",
 *         },
 *         defaultTtl: 100,
 *         id: "containerName",
 *         indexingPolicy: {
 *             automatic: true,
 *             excludedPaths: [],
 *             includedPaths: [{
 *                 indexes: [
 *                     {
 *                         dataType: "String",
 *                         kind: "Range",
 *                         precision: -1,
 *                     },
 *                     {
 *                         dataType: "Number",
 *                         kind: "Range",
 *                         precision: -1,
 *                     },
 *                 ],
 *                 path: "/*",
 *             }],
 *             indexingMode: "consistent",
 *         },
 *         partitionKey: {
 *             kind: "Hash",
 *             paths: ["/AccountNumber"],
 *         },
 *         uniqueKeyPolicy: {
 *             uniqueKeys: [{
 *                 paths: ["/testPath"],
 *             }],
 *         },
 *     },
 *     resourceGroupName: "rg1",
 *     tags: {},
 * });
 *
 * ```
 *
 * ## Import
 *
 * An existing resource can be imported using its type token, name, and identifier, e.g.
 *
 * ```sh
 * $ pulumi import azure-native:documentdb:SqlResourceSqlContainer containerName /subscriptions/subid/resourceGroups/rg1/providers/Microsoft.DocumentDB/databaseAccounts/ddb1/sqlDatabases/databaseName/sqlContainers/containerName 
 * ```
 */
export class SqlResourceSqlContainer extends pulumi.CustomResource {
    /**
     * Get an existing SqlResourceSqlContainer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): SqlResourceSqlContainer {
        return new SqlResourceSqlContainer(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-native:documentdb:SqlResourceSqlContainer';

    /**
     * Returns true if the given object is an instance of SqlResourceSqlContainer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SqlResourceSqlContainer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SqlResourceSqlContainer.__pulumiType;
    }

    public /*out*/ readonly resource!: pulumi.Output<outputs.documentdb.SqlContainerGetPropertiesResponseResource | undefined>;

    /**
     * Create a SqlResourceSqlContainer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: SqlResourceSqlContainerArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["resource"] = undefined /*out*/;
        } else {
            resourceInputs["resource"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SqlResourceSqlContainer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a SqlResourceSqlContainer resource.
 */
export interface SqlResourceSqlContainerArgs {
}
