import { googleLogin } from '../api/oauth'; // Google OAuth 로그인을 처리하는 API 함수
import { Box, Button, Typography } from '@mui/material';
import GoogleIcon from '@mui/icons-material/Google';

function LoginPage() {

  const handleGoogleLogin = async () => {
    try {
      const callbackURL = 'http://localhost:5173/oauth/google/callback';
      const registerURL = 'http://localhost:5173/oauth/register';
      const response = await googleLogin(callbackURL, registerURL);
      window.location.href = response.url;
      // 처리 로직 추가
    } catch (error) {
      console.error('Error:', error);
      // 오류 처리 로직 추가
    }
  };

  return (
    <Box
      sx={{
        display: 'flex',
        flexDirection: 'column',
        justifyContent: 'center',
        alignItems: 'center',
        minHeight: '100vh',
        textAlign: 'center',
        // 아래 코드를 추가하여 화면 전체에 대해 중앙 정렬을 합니다.
        width: '100vw',
        maxWidth: '100%', // 최대 너비를 100%로 설정하여 오버플로우 방지
        margin: '0 auto', // 좌우 마진을 auto로 설정하여 중앙 정렬
      }}
    >
      <Typography variant="h4" gutterBottom component="div" sx={{ display: 'flex', alignItems: 'center', gap: 1 }}>
        Till the Cop <span role="img" aria-label="police car">🚔</span>
      </Typography>
      <Button
        variant="contained"
        color="primary"
        startIcon={<GoogleIcon />}
        onClick={handleGoogleLogin}
        sx={{ mt: 2, px: 5, py: 1 }}
      >
        Google로 로그인하기
      </Button>
    </Box>
  );
}

export default LoginPage;
