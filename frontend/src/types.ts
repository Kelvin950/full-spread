
import { FieldError, UseFormRegister } from "react-hook-form";
 import { z, ZodType } from "zod";
export interface User {
  id: number;
  email: string;
  
 
}
export interface AutData {
  user: User;
  token: string;
}
export  interface formData {
  email?: string;
  password?: string;
  amount?: number;
  currency?:string
  baseCurrency?:string
}


export type FormFieldProps = {
  type: string;
  placeholder: string;
  name: ValidFieldNames;
  register: UseFormRegister<formData>;
  error: FieldError | undefined;
  valueAsNumber?: boolean;
  className:string ;
};

export interface LoginRequest {
  email: string;
  password: string;
}

export interface convertRequest{
 
  from:string 
  to:string
  amount:number
  
}

export interface  ErrorType{

    statusCode:number 
    message:string | string[]


}

export interface Transaction {
  id: number;
  amount: number;
  toCurrency: string;
  result: number;
  rate: number;
  fromCurrency: string;
  createdAt: string;
}




export interface Rates{
  rates:Rate
}


type  Rate = Record<string , number>


  export type ValidFieldNames =
    | "email"
    | "password"
    | "amount"
    | "currency"
    


 export const LoginSchema: ZodType<formData> = z
   .object({
     email: z.string().email(),
   
  
     password: z
       .string()
       .nonempty("Password cannot be empty")
       .min(8, { message: "Password is too short" })
       .max(20, { message: "Password is too long" }),

   })


    export const ConvertSchema: ZodType<formData> = z.object({
      amount: z
        .number({message:"Field is required"})
        .nonnegative()
        .gt(0),
      baseCurrency: z.string(),
      currency: z.string(),
    });
   