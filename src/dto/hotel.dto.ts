export type createHotelDTO = {
    name: string;
    address: string;
    location:string;
    rating:number;
    rating_count:number;
}

export interface updateHotelDTO {
  location?: string;
  name?: string;
  address?: string ;
  rating?: number;
  rating_count?: number ;
}