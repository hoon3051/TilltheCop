import axios from "axios";

const API_URL = "http://localhost:8080"; // TODO: 환경변수로 변경

const ApiManager = axios.create({
  baseURL: API_URL,
  responseType: "json",
  withCredentials: true,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});


// 요청 전에 실행되는 인터셉터
ApiManager.interceptors.request.use(
  (config) => {
    // 로컬 스토리지에서 access_token을 가져옵니다.
    const accessToken = localStorage.getItem("accessToken");

    // access_token이 존재하면 authorization 헤더에 포함시킵니다.
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`;
    }

    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

export default ApiManager;