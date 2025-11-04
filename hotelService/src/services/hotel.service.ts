import { createHotelDTO,updateHotelDTO } from "../dto/hotel.dto";
import { createHotel, getAllHotels, getHotelById, softDeleteHotel, updateHotelById }from "../repositories/hotel.repository";

export async function createHotelService(hotelData: createHotelDTO) {

    const hotel = await createHotel(hotelData);
    return hotel;

}

export async function getHotelByIdService(id:number){
    const hotel = await getHotelById(id);
    return hotel;
}

export async function getAllHotelsService(){
    const hotels=await getAllHotels();
    return hotels;
}

export async function deleteHotelService(id:number){
    const response= await softDeleteHotel(id);
    return response;
}

export async function updateHotelService(id:number,updateData: updateHotelDTO){
    const hotel = await updateHotelById(id,updateData);
    return hotel;
}