//TODO: Create function reutilization file for all the constant values
import { passwordCompare } from "./Constant";

export function validatePassword(password1: string, password2: string): string | null {
    if (password1 === password2) {
      return null; // The passwords match
    } else {
      return passwordCompare;
    }
  }