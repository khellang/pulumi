package dotnet

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pulumi/pulumi/pkg/v3/codegen/internal/test"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/pkg/v3/testing/integration"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/executable"
)

func TestGeneratePackage(t *testing.T) {
	test.TestSDKCodegen(t, &test.TestSDKCodegenOptions{
		Language:   "dotnet",
		GenPackage: generatePackageWithVersion,
		Checks: map[string]test.CodegenCheck{
			"dotnet/compile": typeCheckGeneratedPackage,
		},
	})
}

// TODO replace this with GeneratePackage when https://github.com/pulumi/pulumi/pull/7938 lands.
func generatePackageWithVersion(tool string, pkg *schema.Package, extraFiles map[string][]byte) (map[string][]byte, error) {
	if extraFiles == nil {
		extraFiles = make(map[string][]byte)
	}
	extraFiles["version.txt"] = []byte("0.0.0\n")
	return GeneratePackage(tool, pkg, extraFiles)
}

func typeCheckGeneratedPackage(t *testing.T, pwd string) {
	var err error
	var dotnet string
	dotnet, err = executable.FindExecutable("dotnet")
	require.NoError(t, err)
	cmdOptions := integration.ProgramTestOptions{}
	err = integration.RunCommand(t, "dotnet build", []string{dotnet, "build"}, pwd, &cmdOptions)
	require.NoError(t, err)
}

func TestGenerateType(t *testing.T) {
	cases := []struct {
		typ      schema.Type
		expected string
	}{
		{
			&schema.InputType{
				ElementType: &schema.ArrayType{
					ElementType: &schema.InputType{
						ElementType: &schema.ArrayType{
							ElementType: &schema.InputType{
								ElementType: schema.NumberType,
							},
						},
					},
				},
			},
			"InputList<ImmutableArray<double>>",
		},
		{
			&schema.InputType{
				ElementType: &schema.MapType{
					ElementType: &schema.InputType{
						ElementType: &schema.ArrayType{
							ElementType: &schema.InputType{
								ElementType: schema.NumberType,
							},
						},
					},
				},
			},
			"InputMap<ImmutableArray<double>>",
		},
	}

	mod := &modContext{mod: "main"}
	for _, c := range cases {
		t.Run(c.typ.String(), func(t *testing.T) {
			typeString := mod.typeString(c.typ, "", true, false, false)
			assert.Equal(t, c.expected, typeString)
		})
	}
}

func TestGenerateTypeNames(t *testing.T) {
	test.TestTypeNameCodegen(t, "dotnet", func(pkg *schema.Package) test.TypeNameGeneratorFunc {
		modules, _, err := generateModuleContextMap("test", pkg)
		require.NoError(t, err)

		root, ok := modules[""]
		require.True(t, ok)

		return func(t schema.Type) string {
			return root.typeString(t, "", false, false, false)
		}
	})
}
