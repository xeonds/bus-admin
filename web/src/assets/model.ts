declare interface Company {
    ID: number;
    Name: string;
    CreatedAt: string;
}

declare interface Team {
    ID: number;
    Name: string;
    Company: Company;
    ManagerName: string;
}

declare interface Route {
    ID: number;
    Name: string;
    Team: Team;
}

declare interface Driver {
    ID: number;
    Name: string;
    Route: Route;
}

declare interface RoadManager {
    ID: number;
    Name: string;
    Route: Route;
}

declare interface Violation {
    ID: number;
    Driver: Driver;
    Vehicle: Vehicle;
    Team: Team;
    Route: Route;
    OccurredAt: string;
    ViolationType: string;
}

declare interface Vehicle {
    ID: number;
    VIN: string;
}