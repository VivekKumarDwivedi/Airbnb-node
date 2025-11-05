import { v4 as uuidv4} from "uuid";

export function generateIdempotencyKey(): string {
    // Generate a unique idempotency key using UUID v4
  return uuidv4();
}