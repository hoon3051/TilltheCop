// callback 페이지 클라이언트 코드
import { useEffect } from "react";

const CallbackPage = () => {

  useEffect(() => {
    const hash = window.location.hash;
    const params = new URLSearchParams(hash.substring(1)); // 해시 제거
    const tokenString = params.get('token'); // 'token'은 해시에 포함된 토큰 정보의 키
    if (tokenString) {
      try {
        // decodeURIComponent를 사용하여 URL 인코딩을 해제하고, JSON.parse로 파싱합니다.
        const token = JSON.parse(decodeURIComponent(tokenString));

        // accessToken과 refreshToken을 localStorage에 저장
        localStorage.setItem("accessToken", token.access_token);
        localStorage.setItem("refreshToken", token.refresh_token);

        // 메인 페이지로 이동
        window.location.href = "/";
      } catch (error) {
        console.error("Error parsing token from URL:", error);
        // 여기서 사용자에게 오류가 발생했다는 것을 알리고, 적절한 조치를 취할 수 있습니다.
        // 예: 오류 페이지로 리디렉션하거나, 로그인 페이지로 돌아가기 등
      }
    }
  }, []);

  return (
    <div>
      <h1>Logging in...</h1>
      {/* 로그인 중임을 사용자에게 표시할 수 있는 컴포넌트 */}
    </div>
  );
};

export default CallbackPage;
