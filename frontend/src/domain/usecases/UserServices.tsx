import { jwtDecode } from "jwt-decode";

export default class UserServices {
    static async decodeToken(token: string) {
      const decoded = jwtDecode(token);
      /**
       * token format:
        {
         "exp": 1699729258,
         "iat": 1699642858,
         "jti": "95211d8f-4b9f-43c0-a9fd-2fcd278d003c",
         "rol": "user",
         "sub": "myuser12"
       }
       */
      return decoded
    }
}
