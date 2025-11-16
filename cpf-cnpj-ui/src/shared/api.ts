import axios from "axios";

const token = "437b89a7eb8b3ef111f448069eb55923";

const api = axios.create({
  baseURL: "http://localhost:3000",
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
