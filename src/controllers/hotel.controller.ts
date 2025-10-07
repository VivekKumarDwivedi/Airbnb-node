import {Request,Response,NextFunction} from "express";
import { createHotelService, getHotelByIdService } from "../services/hotel.service";
export async function createHotelHandler(req:Request, res:Response,next:NextFunction){
    //1. call service layer

    const hotelResponse= await createHotelService(req.body);
    //2. send the response

    res.status(201).json({
        message:"Hotel Created Successfully",
        data:hotelResponse,
        success:true,
    })

}

export async function getHotelByIdHandler(req:Request, res:Response,next:NextFunction){
    //1. call service layer
    const hotelResponse= await getHotelByIdService(req.params.id);
    //2. send the response
    res.status(200).json({
        message:"Hotel Retrieved Successfully",
        data:hotelResponse,
        success:true,
    })
}