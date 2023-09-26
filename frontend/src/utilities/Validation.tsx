import { nameRegex } from "./Constant";

export const validateName = (inputString: string): boolean => {
    return nameRegex.test(inputString);
  };
  