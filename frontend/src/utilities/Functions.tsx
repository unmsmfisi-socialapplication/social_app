//TODO: Create function reutilization file for all the constant values
import { passwordCompare } from "./Constant";

export function validarContraseñas(contraseña1: string, contraseña2: string): string | null {
    if (contraseña1 === contraseña2) {
      return null; // Las contraseñas coinciden
    } else {
      return passwordCompare;
    }
  }