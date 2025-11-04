import {Request,Response,NextFunction} from "express";
import { createHotelService, getAllHotelsService, getHotelByIdService ,deleteHotelService,updateHotelService} from "../services/hotel.service";
import { StatusCodes } from "http-status-codes";
import { hotelUpdateSchema } from "../validators/hotel.vaildator";
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


export async function deleteHotelHandler(req:Request, res:Response,next:NextFunction){

    //1.call service layer
    const hotelResponse= await deleteHotelService(Number(req.params.id));
    //2.send the response

    res.status(StatusCodes.OK).json({
        message:"Hotel Deleted Successfully",
        data:hotelResponse,
        success:true,
    })

}
export async function updateHotelHandler(req:Request, res:Response,next:NextFunction){
const hotelId = Number(req.params.id);
  const parsed = hotelUpdateSchema.safeParse(req.body);

  if (!parsed.success) {
    return res.status(StatusCodes.BAD_REQUEST).json({
      message: "Validation Failed",
      errors: parsed.error.format(),
      success: false,
    });
  }

  try {
    const hotelResponse = await updateHotelService(hotelId, parsed.data);
    res.status(StatusCodes.OK).json({
      message: "Hotel Updated Successfully",
      data: hotelResponse,
      success: true,
    });
  } catch (error) {
    const errorMessage = (error instanceof Error) ? error.message : "Hotel not found";
    res.status(StatusCodes.NOT_FOUND).json({
      message: errorMessage,
      success: false,
    });
  }

}
