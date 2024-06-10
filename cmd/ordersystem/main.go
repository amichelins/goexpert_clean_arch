package main

import (
    "database/sql"
    "fmt"
    "net"
    "net/http"

    graphql_handler "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/amichelins/goexpert_clean_arch/configs"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/database"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/graph"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/grpc/pb"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/grpc/service"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/web"
    "github.com/amichelins/goexpert_clean_arch/internal/infra/web/webserver"
    "github.com/amichelins/goexpert_clean_arch/internal/usecase"
    _ "github.com/go-sql-driver/mysql"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
)

func main() {
    configs, err := configs.LoadConfig(".")

    if err != nil {
        panic(err)
    }

    db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))

    if err != nil {
        panic(err)
    }
    defer db.Close()

    webserver := webserver.NewWebServer(configs.WebServerPort)
    webOrderHandler := web.NewWebOrderHandler(database.NewOrderRepository(db))
    webserver.AddHandler("/order", webOrderHandler.Create)
    webserver.AddHandler("/list", webOrderHandler.List)

    fmt.Println("Starting web server on port", configs.WebServerPort)
    go webserver.Start()
    //webserver.Start()

    grpcServer := grpc.NewServer()
    createOrderService := service.NewOrderService(*usecase.NewListOrderUseCase(database.NewOrderRepository(db)))
    pb.RegisterOrderServiceServer(grpcServer, createOrderService)
    reflection.Register(grpcServer)

    fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
    lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
    if err != nil {
        panic(err)
    }
    go grpcServer.Serve(lis)

    ListOrderUseCase := usecase.NewListOrderUseCase(database.NewOrderRepository(db))
    srv := graphql_handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
        ListOrderUseCase: *ListOrderUseCase,
    }}))
    http.Handle("/", playground.Handler("GraphQL playground", "/query"))
    http.Handle("/query", srv)

    fmt.Println("Starting GraphQL server on port", configs.GraphQLServerPort)
    _ = http.ListenAndServe(":"+configs.GraphQLServerPort, nil)
}
