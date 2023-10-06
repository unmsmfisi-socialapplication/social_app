//TODO: Create function reutilization file for all the constant values
import { passwordMismatchError } from "./Constant";

export function validatePassword(password1: string, password2: string): boolean {
  return password1 === password2;
}