import axios from "axios";
import { globalConstants } from "../utils/constants";

export const instance = axios.create({
  baseURL: globalConstants.BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  withCredentials: true,
});

instance.interceptors.request.use(
  (config) => {
    const auth = localStorage.getItem("auth");

    if (!!auth) {
      const { token } = JSON.parse(auth);
      config.headers["Authorization"] = `ApiKey ${token}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

instance.interceptors.response.use(
  (res) => {
    return res;
  },
  async (err) => {
    return Promise.reject(err);
  }
);
