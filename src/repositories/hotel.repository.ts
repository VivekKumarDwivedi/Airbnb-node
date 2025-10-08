import Hotel from "../db/models/hotel";
import { createHotelDTO } from "../dto/hotel.dto";
import logger from "../config/logger.config";
import { NotFoundError } from "../utils/errors/app.error";

export async function createHotel(hotelData: createHotelDTO) {

    const hotel = await Hotel.create({
        name: hotelData.name,
        address: hotelData.address,
        location: hotelData.location,
        rating: hotelData.rating ,
        rating_count: hotelData.rating_count,
    });
    logger.info(`Hotel created: ${hotel.id}`);
    return hotel;
}

export async function getHotelById(id: number) {
    const hotel = await Hotel.findByPk(id);
    if (!hotel) {
        logger.warn(`Hotel not found: ${id}`);
        throw new NotFoundError(`Hotel with id ${id} not found`);
    }
    logger.info(`Hotel retrieved: ${hotel.id}`);
    return hotel;
}

export async function getAllHotels() {
    const hotels = await Hotel.findAll();
    if(!hotels){
        logger.error('No hotel Found');
        throw new NotFoundError('No Hotel Found');
    }
    logger.info(`Hotels Found: ${hotels.length}`);
    return hotels;
}