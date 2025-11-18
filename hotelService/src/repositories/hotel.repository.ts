import Hotel from "../db/models/hotel";
import logger from "../config/logger.config";
import { NotFoundError } from "../utils/errors/app.error";
import BaseRepository from "./base.repository";

export class HotelRepository extends BaseRepository<Hotel> {
    constructor(){
        super(Hotel);
    }

    async findAll(){
        const hotels = await this.model.findAll({
            where:{
                deletedAt:null
            }
        });
        if(!hotels){
            logger.error(`No Hotels found`);
            throw new NotFoundError(`No Hotels found`);
        }
    
    logger.info(`Hotels Found : ${hotels.length}`);
    return hotels;
    }

    async softDelete(id:number) {
            const hotel = await Hotel.findByPk(id);

        if(!hotel){
            logger.error(`hotel note found with id ${id}`);
            throw new NotFoundError(`Hotel with id ${id} not found`);
        }

        hotel.deletedAt = new Date();
        await hotel.save();
        logger.info(`Hotel soft deleted with id ${id}`);
        return true;

    }
}