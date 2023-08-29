import { User, UserApiResponse } from "../components/User/interfaces";
import { instance } from "../services/api";
import { globalConstants, identityConstants } from "../utils/constants";

  export async function register(user: User): Promise<UserApiResponse> {
    var formData = new FormData();
    for (var key in user) {
        formData.append(key, (user as any)[key]);
    }
  
    try {
      if (
        !user.name
      ) {
        throw new Error(identityConstants.FILL_REQUIRED_FIELDS);
      }

      const result = await instance.post(`${globalConstants.BASE_URL}/register`, formData);
      return result.data as UserApiResponse;
    } catch (error: any) {
      throw new Error(error.message);
    }
  }
