package mutation

import (
	"Go_Video/app/global/structs"
	"Go_Video/app/graphql/schema"
	"github.com/graphql-go/graphql"
)

//test mutation 方法
func TestMutation() (*graphql.Field) {
	var Mutation = graphql.Field{
		Name: "TestMutation",
		Type: schema.HelloType,
		Args: schema.HelloArgs,
		Resolve: func(p graphql.ResolveParams) (i interface{}, e error) {
			foo, _ := p.Args["foo"].(string)
			bar, _ := p.Args["bar"].(string)
			var hello structs.HelloResponse
			hello.Foo = foo
			hello.Bar = bar
			return hello, nil
		},
	}
	return &Mutation
}
