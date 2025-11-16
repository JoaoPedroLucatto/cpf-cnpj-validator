import axios from "axios";

const token = import.meta.env.VITE_API_KEY;

const api = axios.create({
  baseURL: import.meta.env.VITE_API_HOST,
  timeout: 8000,
});

api.interceptors.request.use((config) => {
  if (token) {
    if (config.headers instanceof axios.AxiosHeaders) {
      config.headers.set("Authorization", `Bearer ${token}`);
    } else {
      config.headers = new axios.AxiosHeaders(config.headers);
      config.headers.set("Authorization", `Bearer ${token}`);
    }
  }
  return config;
});

api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.error("API Error:", error);
    return Promise.reject(error);
  }
);

export default api;
