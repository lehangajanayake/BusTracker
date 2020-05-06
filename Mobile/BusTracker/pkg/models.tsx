
export interface Bus{
    LicenseNo?: string;
    Location?: Location;
    Attributes? : BusAttributes;
}
export interface Location{
    Latitude: number;
    Longitude: number;
}
export interface BusAttributes{
    PathNo: number;
    AC: boolean;
    Availability: number;
}

interface Search{
    PathNo?: number;
    Destination: string;
    CurrentLocation: Location
}