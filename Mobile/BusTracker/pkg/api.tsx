import {Bus} from "./models";

let address: string = "192.168.1.2:8080";

export async function getBusLocation(LicenseNo: string): Promise<Bus | undefined>{
    try{
        //let url: string = "http://"+address+"/getBusLocation/"+LicenseNo
        let res = await fetch(`http://${address}/getBusLocation/${LicenseNo}`);
        let body  = await res.json();
        let bus: Bus= {
            Location:{
                Latitude: body.Latitude,
                Longitude: body.Longitude
            }
                
            }
            return bus;
    }catch(e){
        console.log("error getting bus location");
        return;
    }
}
    


export async function getBusAttributes(LicenseNo: string):Promise<Bus | undefined>{
    try{
        let res = await fetch(`http://${address}/getBusAttributes/${LicenseNo}`);
        let body = await res.json();            
        let bus: Bus = {
                Attributes:{
                    PathNo: body.PathNo,
                    AC: body.AC,
                    Availability: body.Availability
                }
            }
        return bus;
    }catch(err){
        console.log(err)
        return;
    }
}

export async function getBusAvailability(bus: Bus):Promise<Bus | undefined>{
    try{
        let res = await fetch(`http://${address}/getBusAttributes/${bus.LicenseNo}`);
        let body = await res.json();            
        bus = {
                Attributes:{
                    PathNo: bus.Attributes?.PathNo!,
                    AC: bus.Attributes?.AC!,
                    Availability: body.Availability
                }
            };
        return bus;

    }catch(err){
        console.log(err)
        return;
    }
}


console.log(getBusLocation("NC-7783"));