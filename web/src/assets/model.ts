declare interface Company {
  ID: number
  Name: string
  CreatedAt: string
}

declare interface Team {
  ID: number
  Name: string
  CompanyID: number
  Company: Company
  ManagerName: string
}

declare interface Route {
  ID: number
  Name: string
  TeamID: number
  Team: Team
}

declare interface Driver {
  ID: number
  Name: string
  RouteID: number
  Route: Route
}

declare interface RoadManager {
  ID: number
  Name: string
  RouteID: number
  Route: Route
}

declare interface Violation {
  ID: number
  DriverID: number
  VehicleID: number
  TeamID: number
  RouteID: number
  OccurredAt: string
  ViolationType: string
  Driver: Driver
  Vehicle: Vehicle
  Team: Team
  Route: Route
}

declare interface Vehicle {
  ID: number
  VIN: string
}
