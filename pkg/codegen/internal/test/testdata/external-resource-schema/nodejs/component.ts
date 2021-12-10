// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

import * as pulumiAws from "@pulumi/aws";
import * as pulumiKubernetes from "@pulumi/kubernetes";

export class Component extends pulumi.CustomResource {
    /**
     * Get an existing Component resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Component {
        return new Component(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'example::Component';

    /**
     * Returns true if the given object is an instance of Component.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Component {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Component.__pulumiType;
    }

    public /*out*/ readonly provider!: pulumi.Output<pulumiKubernetes.Provider | undefined>;
    public /*out*/ readonly securityGroup!: pulumi.Output<pulumiAws.ec2.SecurityGroup>;
    public /*out*/ readonly storageClasses!: pulumi.Output<{[key: string]: pulumiKubernetes.storage.v1.StorageClass} | undefined>;

    /**
     * Create a Component resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ComponentArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.requiredMetadata === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredMetadata'");
            }
            if ((!args || args.requiredMetadataArray === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredMetadataArray'");
            }
            if ((!args || args.requiredMetadataMap === undefined) && !opts.urn) {
                throw new Error("Missing required property 'requiredMetadataMap'");
            }
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataArray"] = args ? args.metadataArray : undefined;
            resourceInputs["metadataMap"] = args ? args.metadataMap : undefined;
            resourceInputs["requiredMetadata"] = args ? args.requiredMetadata : undefined;
            resourceInputs["requiredMetadataArray"] = args ? args.requiredMetadataArray : undefined;
            resourceInputs["requiredMetadataMap"] = args ? args.requiredMetadataMap : undefined;
            resourceInputs["provider"] = undefined /*out*/;
            resourceInputs["securityGroup"] = undefined /*out*/;
            resourceInputs["storageClasses"] = undefined /*out*/;
        } else {
            resourceInputs["provider"] = undefined /*out*/;
            resourceInputs["securityGroup"] = undefined /*out*/;
            resourceInputs["storageClasses"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Component.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Component resource.
 */
export interface ComponentArgs {
    metadata?: pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>;
    metadataArray?: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>[]>;
    metadataMap?: pulumi.Input<{[key: string]: pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>}>;
    requiredMetadata: pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>;
    requiredMetadataArray: pulumi.Input<pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>[]>;
    requiredMetadataMap: pulumi.Input<{[key: string]: pulumi.Input<pulumiKubernetes.types.input.meta.v1.ObjectMetaArgs>}>;
}
