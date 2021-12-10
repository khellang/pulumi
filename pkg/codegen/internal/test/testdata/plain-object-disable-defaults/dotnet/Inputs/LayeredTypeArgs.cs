// *** WARNING: this file was generated by test. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Example.Inputs
{

    /// <summary>
    /// Make sure that defaults propagate through types
    /// </summary>
    public sealed class LayeredTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The answer to the question
        /// </summary>
        [Input("answer")]
        public Input<double>? Answer { get; set; }

        [Input("other", required: true)]
        public Input<Inputs.HelmReleaseSettingsArgs> Other { get; set; } = null!;

        /// <summary>
        /// Test how plain types interact
        /// </summary>
        [Input("plainOther")]
        public Inputs.HelmReleaseSettingsArgs? PlainOther { get; set; }

        /// <summary>
        /// The question already answered
        /// </summary>
        [Input("question")]
        public Input<string>? Question { get; set; }

        [Input("recursive")]
        public Input<Inputs.LayeredTypeArgs>? Recursive { get; set; }

        /// <summary>
        /// To ask and answer
        /// </summary>
        [Input("thinker", required: true)]
        public Input<string> Thinker { get; set; } = null!;

        public LayeredTypeArgs()
        {
            Answer = 42;
            Question = Utilities.GetEnv("PULUMI_THE_QUESTION") ?? "<unknown>";
            Thinker = "not a good interaction";
        }
    }
}
