import { googleLogin } from '../api/oauth'; // Google OAuth 로그인을 처리하는 API 함수

function LoginPage() {

  const handleGoogleLogin = async () => {
    try {
      const response = await googleLogin();

      window.location.href = response.url;
      
      // 처리 로직 추가
    } catch (error) {
      console.error('Error:', error);
      // 오류 처리 로직 추가
    }
  };


  return (
    <div>
      <h2>Login Page</h2>
      <button onClick={handleGoogleLogin}>Login with Google</button>
    </div>
  );
}

export default LoginPage;
