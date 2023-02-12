# CSE5234 Project

## Dependencies

- MySQL 8.0
- RabbitMQ
- Go 1.19
- npm

## Directories
```
cse5234/
|  inventory/     inventory-management microservice
|  |  migration/  SQL migrations that set up the database tables
|  |  pkg/        sub packages of our Go project
|  |  |  config/  config package helps reading config files
|  |  |  domain/  high level definitions of our business domain
|  |  |  mysql/   database-specific implementations of our business domain
|  |  |  routes/  HTTP routes and their handlers
|  |  Makefile    simplifies the compilation process with make
|  |  config.json server configurations
|  |  main.go     main package of inventory-management microservice
|  order/         order-management microservice
|  payment/       payment processing microservice
|  shipment/      shipment microservice
|  storefront/    presentation tier client
|  |  public/     static public assets
|  |  src/
|  |  |  app/       layout and routes of the app
|  |  |  common/    common components that may be reused
|  |  |  features/  Redux reducers
|  |  |  routes/    implementation of routes
|  |  |  index.js   entry point of ReactJS application
|  |  |  index.css  global styles
```

