import axios from "axios";

const API_URL = "http://localhost:8080"; // TODO: 환경변수로 변경

const ApiManager = axios.create({
  baseURL: API_URL,
  responseType: "json",
  withCredentials: true,
  timeout: 10000,
  headers: { "Content-Type": "application/json" },
});



// 요청 전에 실행되는 인터셉터 (access_token이 있는 경우 헤더에 포함)
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

// 재시도 횟수를 추적하는 변수
let retryCount = 0;
const maxRetryLimit = 3; // 최대 재시도 횟수

// 토큰 만료 시 재발급하는 인터셉터
ApiManager.interceptors.response.use(response => response, async error => {
  const originalRequest = error.config;
  if (error.response.status === 401 && !originalRequest._retry && retryCount < maxRetryLimit) {
    originalRequest._retry = true;
    retryCount++; // 재시도 횟수 증가

    try {
      const refreshToken = localStorage.getItem('refreshToken');
      // 리프레시 토큰으로 새 액세스 토큰 요청
      const response = await ApiManager.post("/refresh", { refreshToken });
      const {access_token} = response.data;
      localStorage.setItem('accessToken', access_token);
      // 새 액세스 토큰으로 원래 요청 재시도
      originalRequest.headers.Authorization = `Bearer ${access_token}`;
      return ApiManager(originalRequest);
    } catch (refreshError) {
      // 리프레시 토큰 요청 실패 시 알림 1초 후 로그인 페이지로 리다이렉트
      return Promise.reject(refreshError);
    }
  } else if (retryCount >= maxRetryLimit) {
    // 최대 재시도 횟수를 초과한 경우, 로그인 페이지로 리다이렉트

    alert("세션이 만료되었습니다. 다시 로그인해주세요.");
    window.location.href = "http://localhost:5173/oauth/google/login";
  }
  return Promise.reject(error);
});

export default ApiManager;