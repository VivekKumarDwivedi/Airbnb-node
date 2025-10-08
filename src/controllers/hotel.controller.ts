import {Request,Response,NextFunction} from "express";
import { createHotelService, getAllHotelsService, getHotelByIdService } from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";
export async function createHotelHandler(req:Request, res:Response,next:NextFunction){
    //1. call service layer

    const hotelResponse= await createHotelService(req.body);
    //2. send the response

    res.status(StatusCodes.CREATED).json({
        message:"Hotel Created Successfully",
        data:hotelResponse,
        success:true,
    })

}

export async function getHotelByIdHandler(req:Request, res:Response,next:NextFunction){
    //1. call service layer
    const hotelResponse= await getHotelByIdService(Number(req.params.id));
    //2. send the response
    res.status(StatusCodes.OK).json({
        message:"Hotel Retrieved Successfully",
        data:hotelResponse,
        success:true,
    })
}

export async function getAllHotelsHandler(req:Request, res:Response,next:NextFunction){
    //1.call service layer
    const hotelResponse= await getAllHotelsService();
    //2.send the response
    res.status(StatusCodes.OK).json({
        message:"All Hotels Retrieved Successfully",
        data:hotelResponse,
        success:true,
    });
}

