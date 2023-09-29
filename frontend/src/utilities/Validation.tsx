import { emailRegex, nameRegex, passwordRegex } from "./Constant";

export const validateName = (inputString: string): boolean => {
    return nameRegex.test(inputString);
  };

export const validateEmail = (inputString: string): boolean => {
  return emailRegex.test(inputString);
};


export const validatePassword = (inputString: string): boolean => {
  return passwordRegex.test(inputString);
};
  