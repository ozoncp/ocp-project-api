// create a package
package main

// import some pandora stuff
// and stuff you need for your scenario
// and protobuf contracts for your grpc service

import (
	"context"
	"log"

	pb "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"

	"github.com/spf13/afero"
	"github.com/yandex/pandora/cli"
	phttp "github.com/yandex/pandora/components/phttp/import"
	"github.com/yandex/pandora/core"
	"github.com/yandex/pandora/core/aggregator/netsample"
	coreimport "github.com/yandex/pandora/core/import"
	"github.com/yandex/pandora/core/register"
	"google.golang.org/grpc"
)

type Ammo struct {
	Tag        string
	ReposCount uint
	RepoId     uint64
	ProjectId  uint64
	UserId     uint64
	Link       string
}

type GunConfig struct {
	Target string `validate:"required"` // Configuration will fail, without target defined
}

type Gun struct {
	// Configured on construction.
	client *grpc.ClientConn
	conf   GunConfig
	// Configured on Bind, before shooting
	aggr core.Aggregator // May be your custom Aggregator.
	core.GunDeps
}

func NewGun(conf GunConfig) *Gun {
	return &Gun{conf: conf}
}

func (g *Gun) Bind(aggr core.Aggregator, deps core.GunDeps) error {
	// create gRPC stub at gun initialization
	conn, err := grpc.DialContext(
		context.TODO(),
		g.conf.Target,
		grpc.WithInsecure(),
		grpc.WithUserAgent("load test, pandora custom shooter"))
	if err != nil {
		log.Fatalf("FATAL: %s", err)
	}
	g.client = conn
	g.aggr = aggr
	g.GunDeps = deps
	return nil
}

func (g *Gun) Shoot(ammo core.Ammo) {
	customAmmo := ammo.(*Ammo)
	g.shoot(customAmmo)
}

func (g *Gun) create_method(client pb.OcpRepoApiClient, ammo *Ammo) int {
	code := 0
	out, err := client.CreateRepo(
		context.TODO(),
		&pb.CreateRepoRequest{ProjectId: ammo.ProjectId, UserId: ammo.UserId, Link: ammo.Link},
	)

	if err != nil {
		log.Printf("FATAL: %s", err)
		code = 500
	}

	if out != nil {
		code = 200
	}
	return code
}

func (g *Gun) update_method(client pb.OcpRepoApiClient, ammo *Ammo) int {
	code := 0
	var repo = pb.Repo{Id: ammo.RepoId, ProjectId: ammo.ProjectId, UserId: ammo.UserId, Link: ammo.Link}
	out, err := client.UpdateRepo(
		context.TODO(),
		&pb.UpdateRepoRequest{Repo: &repo},
	)

	if err != nil {
		log.Printf("FATAL: %s", err)
		code = 500
	}

	if out != nil {
		code = 200
	}
	return code
}

func (g *Gun) multi_create_method(client pb.OcpRepoApiClient, ammo *Ammo) int {
	code := 0
	repos := make([]*pb.NewRepo, 0, ammo.ReposCount)
	for i := 1; i <= int(ammo.ReposCount); i++ {
		var rep = pb.NewRepo{ProjectId: uint64(i), UserId: uint64(i), Link: "1"}
		repos = append(repos, &rep)
	}

	out, err := client.MultiCreateRepo(
		context.TODO(),
		&pb.MultiCreateRepoRequest{Repos: repos},
	)

	if err != nil {
		log.Printf("FATAL: %s", err)
		code = 500
	}

	if out != nil {
		code = 200
	}
	return code
}

func (g *Gun) shoot(ammo *Ammo) {
	code := 0
	sample := netsample.Acquire(ammo.Tag)

	conn := g.client
	client := pb.NewOcpRepoApiClient(conn)

	switch ammo.Tag {
	case "create":
		code = g.create_method(client, ammo)
	case "update":
		code = g.update_method(client, ammo)
	case "multiCreate":
		code = g.multi_create_method(client, ammo)
	default:
		code = 404
	}

	defer func() {
		sample.SetProtoCode(code)
		g.aggr.Report(sample)
	}()
}

func main() {
	//debug.SetGCPercent(-1)
	// Standard imports.
	fs := afero.NewOsFs()
	coreimport.Import(fs)
	// May not be imported, if you don't need http guns and etc.
	phttp.Import(fs)

	// Custom imports. Integrate your custom types into configuration system.
	coreimport.RegisterCustomJSONProvider("custom_provider", func() core.Ammo { return &Ammo{} })

	register.Gun("try_bfg", NewGun, func() GunConfig {
		return GunConfig{
			Target: "default target",
		}
	})

	cli.Run()
}