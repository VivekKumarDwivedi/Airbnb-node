import {z} from "zod";

export const hotelSchema = z.object({
    name: z.string().min(1),
    address: z.string().min(1),
    location: z.string().min(1),
    rating: z.number().optional(),
    rating_count: z.number().optional(),
})

export const hotelUpdateSchema = hotelSchema.partial({
    location: true,
    name: true,
    address: true,
    rating: true,
    rating_count: true,
})