import axios from "axios";

export const docAPI = () => {
    return axios.create({
      baseURL: import.meta.env.VITE_SERVER_API_URL,
    });
  };