import { emailRegex, nameRegex } from "./Constant";

export const validateName = (inputString: string): boolean => {
    return nameRegex.test(inputString);
  };

export const validateEmail = (inputString: string): boolean => {
  return emailRegex.test(inputString);
};
  