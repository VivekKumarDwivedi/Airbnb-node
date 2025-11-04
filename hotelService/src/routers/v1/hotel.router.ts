import express from "express";
import { createHotelHandler, getAllHotelsHandler, getHotelByIdHandler,deleteHotelHandler, updateHotelHandler } from "../../controllers/hotel.controller";
import { validateRequestBody } from "../../validators";
import { hotelSchema, hotelUpdateSchema } from "../../validators/hotel.vaildator";
const hotelRouter=express.Router();

hotelRouter.post('/',
    validateRequestBody(hotelSchema),
    createHotelHandler);

hotelRouter.get('/:id',getHotelByIdHandler);
hotelRouter.get('/',getAllHotelsHandler);
hotelRouter.delete('/:id',deleteHotelHandler);
hotelRouter.put('/:id',
    validateRequestBody(hotelUpdateSchema),
    updateHotelHandler);


export default hotelRouter;