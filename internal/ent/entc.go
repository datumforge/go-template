//go:build ignore

// See Upstream docs for more details: https://entgo.io/docs/code-gen/#use-entc-as-a-package

package main

import (
	"log"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/datumforge/entx/genhooks"
	"github.com/datumforge/fgax"
	"github.com/datumforge/fgax/entfga"
	"go.uber.org/zap"
	"gocloud.dev/secrets"

	"github.com/datumforge/entx"
)

var (
	graphSchemaDir = "./schema/"
	// graphQueryDir  = "./query/" // uncomment to enable query generation
)

func main() {
	xExt, err := entx.NewExtension(
		entx.WithJSONScalar(),
	)
	if err != nil {
		log.Fatalf("creating entx extension: %v", err)
	}

	// Ensure the schema directory exists before running entc.
	_ = os.Mkdir("schema", 0755)

	gqlExt, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("schema/ent.graphql"),
		entgql.WithConfigPath("gqlgen.yml"),
		entgql.WithWhereInputs(true),
		entgql.WithSchemaHook(xExt.GQLSchemaHooks()...),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	if err := entc.Generate("./internal/ent/schema", &gen.Config{
		Target:    "./internal/ent/generated",
		Templates: entgql.AllTemplates,
		Hooks: []gen.Hook{
			genhooks.GenSchema(graphSchemaDir),
			// genhooks.GenQuery(graphQueryDir), // uncomment to enable query generation
		},
		Package: "github.com/datumforge/go-template/internal/ent/generated",
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeaturePrivacy,
			gen.FeatureSnapshot,
			gen.FeatureEntQL,
			gen.FeatureNamedEdges,
			gen.FeatureSchemaConfig,
			gen.FeatureIntercept,
		},
	},
		entc.Dependency(
			entc.DependencyType(&secrets.Keeper{}),
		),
		entc.Dependency(
			entc.DependencyName("Authz"),
			entc.DependencyType(fgax.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Logger"),
			entc.DependencyType(zap.SugaredLogger{}),
		),
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.Extensions(
			gqlExt,
			entfga.New(
				entfga.WithSoftDeletes(),
			),
		)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
